// +build integration

/*
2021 © Postgres.ai
*/

package snapshot

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	initialScript96 = `
-- SCHEMA
begin;
create table timezones
(
  id         serial PRIMARY KEY,
  created    timestamptz DEFAULT now() NOT NULL,
  modified   timestamptz DEFAULT now() NOT NULL,
  name       text                      NOT NULL,
  timeoffset smallint                  NOT NULL,
  identifier text                      NOT NULL
);
commit;
select pg_switch_xlog();

-- SEED
begin;
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (1, 'eastern', '-5', 'est');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (2, 'central', '-6', 'cst');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (3, 'mountain', '-7', 'mst');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (4, 'pacific', '-8', 'pst');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (5, 'alaska', '-9', 'ast');
alter sequence timezones_id_seq restart with 6;
commit;
select pg_switch_xlog();
`

	initialScript = `
-- SCHEMA
begin;
create table timezones
(
  id         serial PRIMARY KEY,
  created    timestamptz DEFAULT now() NOT NULL,
  modified   timestamptz DEFAULT now() NOT NULL,
  name       text                      NOT NULL,
  timeoffset smallint                  NOT NULL,
  identifier text                      NOT NULL
);
commit;
select pg_switch_wal();

-- SEED
begin;
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (1, 'eastern', '-5', 'est');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (2, 'central', '-6', 'cst');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (3, 'mountain', '-7', 'mst');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (4, 'pacific', '-8', 'pst');
INSERT INTO timezones  (id, name, timeoffset, identifier) VALUES  (5, 'alaska', '-9', 'ast');
alter sequence timezones_id_seq restart with 6;
commit;
select pg_switch_wal();
`
)

const (
	port         = "5432/tcp"
	dbname       = "postgres"
	user         = "postgres"
	testPassword = "password"
)

func TestParsingWAL96(t *testing.T) {
	dockerCLI, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatal("Failed to create a Docker client:", err)
	}

	testWALParsing(t, dockerCLI, 9.6, initialScript96)
}

func TestParsingWAL(t *testing.T) {
	dockerCLI, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatal("Failed to create a Docker client:", err)
	}

	postgresVersions := []float64{10, 11, 12, 13, 14}

	for _, pgVersion := range postgresVersions {
		testWALParsing(t, dockerCLI, pgVersion, initialScript)
	}
}

func testWALParsing(t *testing.T, dockerCLI *client.Client, pgVersion float64, initialSQL string) {
	ctx := context.Background()

	pgVersionString := fmt.Sprintf("%g", pgVersion)

	// Create a temporary directory to store PGDATA.
	dir, err := os.MkdirTemp("", "pg_test_"+pgVersionString+"_")
	require.Nil(t, err)

	defer os.Remove(dir)

	// Run a test container.
	logStrategyForAcceptingConnections := wait.NewLogStrategy("database system is ready to accept connections")
	logStrategyForAcceptingConnections.Occurrence = 2

	dbURL := func(port nat.Port) string {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			"localhost", port.Port(), user, testPassword, dbname)
	}

	req := testcontainers.ContainerRequest{
		Name:         "pg_test_" + pgVersionString,
		Image:        "postgres:" + pgVersionString,
		ExposedPorts: []string{port},
		WaitingFor: wait.ForAll(
			logStrategyForAcceptingConnections,
			wait.ForLog("PostgreSQL init process complete; ready for start up."),
			wait.ForSQL(nat.Port(port), "postgres", dbURL).Timeout(30*time.Second),
		),
		BindMounts: map[string]string{
			"/tmp": "/tmp", // To provide local access to the container temporary directory.
		},
		Env: map[string]string{
			"POSTGRES_PASSWORD": testPassword,
		},

	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.Nil(t, err)

	defer func() { _ = postgresContainer.Terminate(ctx) }()

	// Prepare test data.
	code, err := postgresContainer.Exec(ctx, []string{"psql", "-U", user, "-d", dbname, "-XAtc", initialSQL})

	require.Nil(t, err)
	assert.Equal(t, 0, code)

	p := &PhysicalInitial{
		dockerClient: dockerCLI,
	}

	// Prepare local copies of WAL files
	// since it's impossible to have access in the original PGDATA because permissions denied.
	tmpWaldir := walDir(dir, pgVersion)

	code, err = postgresContainer.Exec(ctx, []string{"mkdir", "-p", tmpWaldir})
	require.Nil(t, err)
	assert.Equal(t, 0, code)

	originalPGData := "/var/lib/postgresql/data/"
	code, err = postgresContainer.Exec(ctx, []string{"cp", "-R", walDir(originalPGData, pgVersion), dir})
	require.Nil(t, err)
	assert.Equal(t, 0, code)

	code, err = postgresContainer.Exec(ctx, []string{"chmod", "777", "-R", tmpWaldir})
	require.Nil(t, err)
	assert.Equal(t, 0, code)

	// Check WAL parsing.
	dsa, err := p.getDSAFromWAL(ctx, pgVersion, postgresContainer.GetContainerID(), dir)
	assert.Nil(t, err)
	assert.NotEmpty(t, dsa)

	t.Log("DSA: ", dsa)
}

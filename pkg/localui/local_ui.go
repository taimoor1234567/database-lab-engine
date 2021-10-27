/*
2021 © Postgres.ai
*/

// Package localui manages local UI container.
package localui

import (
	"context"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"

	"gitlab.com/postgres-ai/database-lab/v2/pkg/retrieval/engine/postgres/tools/cont"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/services/provision/docker"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/services/provision/runners"
)

// Config defines configs for a local UI container.
type Config struct {
	Enabled     bool   `yaml:"enabled"`
	DockerImage string `yaml:"dockerImage"`
	Port        int    `yaml:"port"`
}

// UIManager manages local UI container.
type UIManager struct {
	runner runners.Runner
	docker *client.Client
	cfg    *Config
}

// New creates a new UI Manager.
func New(cfg *Config, runner runners.Runner, docker *client.Client) *UIManager {
	return &UIManager{runner: runner, docker: docker, cfg: cfg}
}

// RunUI runs local UI container.
func (ui *UIManager) RunUI(ctx context.Context, instanceID string) error {
	if err := PrepareImage(ui.runner, ui.cfg.DockerImage); err != nil {
		return fmt.Errorf("failed to prepare Docker image: %w", err)
	}

	localUI, err := ui.docker.ContainerCreate(ctx,
		&container.Config{
			ExposedPorts: nat.PortSet{nat.Port(strconv.Itoa(ui.cfg.Port)) + "/tcp": struct{}{}},
			Labels: map[string]string{
				cont.DBLabControlLabel:    cont.DBLabLocalUILabel,
				cont.DBLabInstanceIDLabel: instanceID,
			},
			Image: ui.cfg.DockerImage,
		},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		"dblab_local_ui")

	if err != nil {
		return fmt.Errorf("failed to prepare Docker image for LocalUI: %w", err)
	}

	if err := ui.docker.ContainerStart(ctx, localUI.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container %q: %w", localUI.ID, err)
	}

	return nil
}

// PrepareImage prepares a Docker image to use.
func PrepareImage(runner runners.Runner, dockerImage string) error {
	imageExists, err := docker.ImageExists(runner, dockerImage)
	if err != nil {
		return fmt.Errorf("cannot check docker image existence: %w", err)
	}

	if imageExists {
		return nil
	}

	if err := docker.PullImage(runner, dockerImage); err != nil {
		return fmt.Errorf("cannot pull docker image: %w", err)
	}

	return nil
}

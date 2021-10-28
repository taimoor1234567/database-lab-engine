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

	"gitlab.com/postgres-ai/database-lab/v2/pkg/retrieval/engine/postgres/tools"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/retrieval/engine/postgres/tools/cont"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/services/provision/docker"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/services/provision/runners"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/util/networks"
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
	cfg    Config
}

// New creates a new UI Manager.
func New(cfg Config, runner runners.Runner, docker *client.Client) *UIManager {
	return &UIManager{runner: runner, docker: docker, cfg: cfg}
}

// Reload reloads configuration of UI manager and adjusts a UI container according to it.
func (ui *UIManager) Reload(ctx context.Context, cfg Config, instanceID string) error {
	originalConfig := ui.cfg
	ui.cfg = cfg

	if !ui.isConfigChanged(originalConfig) {
		return nil
	}

	if !cfg.Enabled {
		ui.Stop(ctx, instanceID)
		return nil
	}

	if !originalConfig.Enabled {
		return ui.Run(ctx, instanceID)
	}

	return ui.Restart(ctx, instanceID)
}

func (ui *UIManager) isConfigChanged(cfg Config) bool {
	return ui.cfg.Enabled != cfg.Enabled ||
		ui.cfg.DockerImage != cfg.DockerImage ||
		ui.cfg.Port != cfg.Port
}

// Run creates a new local UI container.
func (ui *UIManager) Run(ctx context.Context, instanceID string) error {
	if err := docker.PrepareImage(ui.runner, ui.cfg.DockerImage); err != nil {
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
		getLocalUIName(instanceID),
	)

	if err != nil {
		return fmt.Errorf("failed to prepare Docker image for LocalUI: %w", err)
	}

	if err := ui.docker.ContainerStart(ctx, localUI.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container %q: %w", localUI.ID, err)
	}

	if err := networks.Connect(ctx, ui.docker, instanceID, localUI.ID); err != nil {
		return fmt.Errorf("failed to connect UI container to the internal Docker network: %w", err)
	}

	return nil
}

// Restart destroys and creates a new local UI container.
func (ui *UIManager) Restart(ctx context.Context, instanceID string) error {
	ui.Stop(ctx, instanceID)

	if err := ui.Run(ctx, instanceID); err != nil {
		return fmt.Errorf("failed to start UI container: %w", err)
	}

	return nil
}

// Stop removes a local UI container.
func (ui *UIManager) Stop(ctx context.Context, instanceID string) {
	tools.RemoveContainer(ctx, ui.docker, getLocalUIName(instanceID), cont.StopTimeout)
}

func getLocalUIName(instanceID string) string {
	return cont.DBLabLocalUILabel + "_" + instanceID
}

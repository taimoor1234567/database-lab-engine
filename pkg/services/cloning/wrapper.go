/*
2019 © Postgres.ai
*/

// Package cloning provides a cloning service.
package cloning

import (
	"time"

	"gitlab.com/postgres-ai/database-lab/v2/pkg/models"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/services/provision/resources"
)

// CloneWrapper represents a cloning service wrapper.
type CloneWrapper struct {
	Clone   *models.Clone
	Session *resources.Session

	TimeCreatedAt time.Time
	TimeStartedAt time.Time
}

// NewCloneWrapper constructs a new CloneWrapper.
func NewCloneWrapper(clone *models.Clone, createdAt time.Time) *CloneWrapper {
	w := &CloneWrapper{
		Clone:         clone,
		TimeCreatedAt: createdAt,
	}

	return w
}

// IsProtected checks if clone is protected.
func (cw CloneWrapper) IsProtected() bool {
	return cw.Clone != nil && cw.Clone.Protected
}

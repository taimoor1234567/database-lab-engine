/*
2019 © Postgres.ai
*/

package models

import (
	"fmt"
	"math/big"

	"github.com/dustin/go-humanize"
)

// Clone defines a clone model.
type Clone struct {
	ID        string        `json:"id"`
	Snapshot  *Snapshot     `json:"snapshot"`
	Protected bool          `json:"protected"`
	DeleteAt  string        `json:"deleteAt"`
	CreatedAt string        `json:"createdAt"`
	Status    Status        `json:"status"`
	DB        Database      `json:"db"`
	Metadata  CloneMetadata `json:"metadata"`
}

// CloneMetadata contains fields describing a clone model.
type CloneMetadata struct {
	CloneDiffSize  uint64  `json:"cloneDiffSize"`
	LogicalSize    uint64  `json:"logicalSize"`
	CloningTime    float64 `json:"cloningTime"`
	MaxIdleMinutes uint    `json:"maxIdleMinutes"`
}

// Size describes amount of disk space.
type Size uint64

// MarshalJSON marshals the Size struct.
func (s Size) MarshalJSON() ([]byte, error) {
	humanReadableSize := humanize.BigIBytes(big.NewInt(int64(s)))
	return []byte(fmt.Sprintf("%q", humanReadableSize)), nil
}

// CloneView represents a view of clone model.
type CloneView struct {
	*Clone
	Snapshot *SnapshotView     `json:"snapshot"`
	Metadata CloneMetadataView `json:"metadata"`
}

// CloneMetadataView represents a view of clone metadata.
type CloneMetadataView struct {
	*CloneMetadata
	CloneDiffSize Size `json:"cloneDiffSize"`
	LogicalSize   Size `json:"logicalSize"`
}

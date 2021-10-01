/*
2021 © Postgres.ai
*/

package cloning

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/postgres-ai/database-lab/v2/pkg/log"
	"gitlab.com/postgres-ai/database-lab/v2/pkg/util"
)

const sessionsFilename = "sessions.json"

// LoadSessionState reads sessions data from disk.
func (c *Base) LoadSessionState() error {
	c.cloneMutex.Lock()
	defer c.cloneMutex.Unlock()

	c.clones = make(map[string]*CloneWrapper)

	sessionsPath, err := util.GetMetaPath(sessionsFilename)
	if err != nil {
		return fmt.Errorf("failed to get path of a sessions file: %w", err)
	}

	data, err := ioutil.ReadFile(sessionsPath)
	if err != nil {
		if os.IsNotExist(err) {
			// no sessions data, ignore
			return nil
		}

		return fmt.Errorf("failed to read sessions data: %w", err)
	}

	return json.Unmarshal(data, &c.clones)
}

func (c *Base) filterRunningClones() {
	c.cloneMutex.Lock()
	defer c.cloneMutex.Unlock()

	clones := make(map[string]*CloneWrapper, len(c.clones))
	snapshotCache := make(map[string]struct{})

	for _, wrapper := range c.clones {
		if wrapper.Clone == nil {
			continue
		}

		if _, ok := snapshotCache[wrapper.Clone.Snapshot.ID]; !ok {
			snapshot, err := c.getSnapshotByID(wrapper.Clone.Snapshot.ID)
			if err != nil {
				continue
			}

			snapshotCache[snapshot.ID] = struct{}{}
		}

		clones[wrapper.Clone.ID] = wrapper
	}

	c.clones = clones
}

// SaveClonesState writes clones state to disk.
func (c *Base) SaveClonesState() {
	if err := c.saveClonesState(); err != nil {
		log.Err("Failed to save the state of running clones", err)
	}
}

// saveClonesState tries to write clones state to disk and returns an error on failure.
func (c *Base) saveClonesState() error {
	c.cloneMutex.Lock()
	defer c.cloneMutex.Unlock()

	sessionsPath, err := util.GetMetaPath(sessionsFilename)
	if err != nil {
		return fmt.Errorf("failed to get path of a sessions file: %w", err)
	}

	data, err := json.Marshal(c.clones)
	if err != nil {
		return fmt.Errorf("failed to encode session data: %w", err)
	}

	return ioutil.WriteFile(sessionsPath, data, 0600)
}

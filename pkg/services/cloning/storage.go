/*
2021 © Postgres.ai
*/

package cloning

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.com/postgres-ai/database-lab/v2/pkg/util"
)

const sessionsFilename = "sessions.json"

// Load reads sessions data from disk.
func (c *Base) Load() error {
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

// Save writes sessions data to disk.
func (c *Base) Save() error {
	c.cloneMutex.Lock()
	defer c.cloneMutex.Unlock()

	if len(c.clones) == 0 {
		return nil
	}

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

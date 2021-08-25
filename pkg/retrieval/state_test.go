/*
2021 © Postgres.ai
*/

package retrieval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelByAlertType(t *testing.T) {
	testCases := []struct {
		alertType string
		level     string
	}{
		{
			alertType: "refresh_failed",
			level:     errorLevel,
		},
		{
			alertType: "unschedulable_refresh_ahead",
			level:     warningLevel,
		},
		{
			alertType: "unknown_fail",
			level:     unknownLevel,
		},
	}

	for _, tc := range testCases {
		level := getLevelByAlertType(tc.alertType)
		assert.Equal(t, tc.level, level)
	}
}

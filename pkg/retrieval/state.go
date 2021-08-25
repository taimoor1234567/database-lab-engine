/*
2021 © Postgres.ai
*/

package retrieval

import (
	"sync"
	"time"
)

const (
	// Refresh statuses.
	inactiveStatus   = "inactive"
	refreshingStatus = "refreshing"
	finishedStatus   = "finished"

	// Alert types.
	refreshFailed             = "refresh_failed"
	unschedulableRefreshAhead = "unschedulable_refresh_ahead"

	// Alert levels.
	errorLevel   = "error"
	warningLevel = "warning"
	unknownLevel = "unknown"
)

// State contains state of retrieval service.
type State struct {
	Status string
	mu     sync.Mutex
	Alerts map[string]Alert
}

// Alert defines retrieval subsystem alert.
type Alert struct {
	Level    string
	Message  string
	LastSeen time.Time
	Count    int
}

func (s *State) addAlert(alertType, message string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	alert, ok := s.Alerts[alertType]
	if ok {
		alert.Count++
		alert.LastSeen = time.Now()
		alert.Message = message

		return
	}

	alert = Alert{
		Level:    getLevelByAlertType(alertType),
		Message:  message,
		LastSeen: time.Now(),
		Count:    1,
	}

	s.Alerts[alertType] = alert
}

func (s *State) cleanAlerts() {
	s.mu.Lock()
	s.Alerts = make(map[string]Alert)
	s.mu.Unlock()
}

func getLevelByAlertType(alertType string) string {
	switch alertType {
	case refreshFailed:
		return errorLevel

	case unschedulableRefreshAhead:
		return warningLevel

	default:
		return unknownLevel
	}
}

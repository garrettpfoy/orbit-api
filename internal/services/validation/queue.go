package validation

import (
	"fmt"
	"garrettpfoy/orbit-api/internal/models"
)

// ValidateQueue validates a queue item, if it is valid, it returns nil,
// otherwise it returns an error
func ValidateQueue(queue models.Queue) error {
	if queue.TrackURI == "" {
		return fmt.Errorf("track URI is required")
	}

	if queue.SessionID == 0 {
		return fmt.Errorf("session ID is required")
	}

	if queue.UserID == 0 {
		return fmt.Errorf("user ID is required")
	}

	return nil
}

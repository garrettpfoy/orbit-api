package validation

import (
	"fmt"
	"garrettpfoy/orbit-api/internal/models"
)

func ValidateSession(session models.Session) error {
	if session.Slug == "" {
		return fmt.Errorf("slug is empty")
	}

	if session.HostID == 0 {
		return fmt.Errorf("host ID is empty")
	}

	return nil
}

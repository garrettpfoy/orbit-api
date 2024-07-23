package validation

import (
	"fmt"
	"garrettpfoy/orbit-api/internal/models"
)

func ValidateUser(user models.User) error {
	// Verify the user has either a spotify ID or a google email attached to their account,
	// which is used in OAuth2 flows to verify the user's identity
	if (user.SpotifyUserID == nil || *user.SpotifyUserID == "") && (user.Email == nil || *user.Email == "") {
		return fmt.Errorf("a valid oauth2 ID (google email or spotify ID) is required and not provided")
	}

	return nil
}

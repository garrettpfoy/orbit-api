package validation

import (
	"fmt"
	"garrettpfoy/orbit-api/internal/models"
)

// ValidateAccessToken validates the access token, if it is valid, it returns no error (nil)
// otherwise, it returns an error with a message.
func ValidateAccessToken(accessToken models.AccessToken) error {
	// Verify the user ID is given (valid user is out of the scope of this function)
	if accessToken.UserID == 0 {
		return fmt.Errorf("user ID is empty")
	}

	// Verify the Session ID is given (valid session is out of the scope of this function)
	if accessToken.SessionID == 0 {
		return fmt.Errorf("session ID is empty")
	}

	// Verify the Access Token exists
	if accessToken.AccessToken == "" {
		return fmt.Errorf("access token is empty")
	}

	// Verify the Refresh Token exists
	if accessToken.RefreshToken == "" {
		return fmt.Errorf("refresh token is empty")
	}

	// Verify the Expiry Time exists and is not "zero"
	if accessToken.ExpiryTime.IsZero() {
		return fmt.Errorf("expiry time is empty")
	}

	// If all checks pass, return true
	return nil
}

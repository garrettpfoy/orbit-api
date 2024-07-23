package validation_test

import (
	"fmt"
	"garrettpfoy/orbit-api/internal/models"
	"garrettpfoy/orbit-api/internal/services/validation"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newString(s string) *string {
	return &s
}

func TestValidateAccessToken(t *testing.T) {
	tests := []struct {
		name        string
		accessToken models.AccessToken
		expectedErr error
	}{
		{
			name: "Valid AccessToken",
			accessToken: models.AccessToken{
				UserID:       1,
				SessionID:    1,
				AccessToken:  "valid_access_token",
				RefreshToken: "valid_refresh_token",
				ExpiryTime:   time.Now().Add(time.Hour),
			},
			expectedErr: nil,
		},
		{
			name: "Empty UserID",
			accessToken: models.AccessToken{
				UserID:       0,
				SessionID:    1,
				AccessToken:  "valid_access_token",
				RefreshToken: "valid_refresh_token",
				ExpiryTime:   time.Now().Add(time.Hour),
			},
			expectedErr: fmt.Errorf("user ID is empty"),
		},
		{
			name: "Empty SessionID",
			accessToken: models.AccessToken{
				UserID:       1,
				SessionID:    0,
				AccessToken:  "valid_access_token",
				RefreshToken: "valid_refresh_token",
				ExpiryTime:   time.Now().Add(time.Hour),
			},
			expectedErr: fmt.Errorf("session ID is empty"),
		},
		{
			name: "Empty AccessToken",
			accessToken: models.AccessToken{
				UserID:       1,
				SessionID:    1,
				AccessToken:  "",
				RefreshToken: "valid_refresh_token",
				ExpiryTime:   time.Now().Add(time.Hour),
			},
			expectedErr: fmt.Errorf("access token is empty"),
		},
		{
			name: "Empty RefreshToken",
			accessToken: models.AccessToken{
				UserID:       1,
				SessionID:    1,
				AccessToken:  "valid_access_token",
				RefreshToken: "",
				ExpiryTime:   time.Now().Add(time.Hour),
			},
			expectedErr: fmt.Errorf("refresh token is empty"),
		},
		{
			name: "Zero ExpiryTime",
			accessToken: models.AccessToken{
				UserID:       1,
				SessionID:    1,
				AccessToken:  "valid_access_token",
				RefreshToken: "valid_refresh_token",
				ExpiryTime:   time.Time{},
			},
			expectedErr: fmt.Errorf("expiry time is empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateAccessToken(tt.accessToken)
			if tt.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestValidateSession(t *testing.T) {
	tests := []struct {
		name        string
		session     models.Session
		expectedErr error
	}{
		{
			name: "Valid Session",
			session: models.Session{
				Slug:   "valid_slug",
				HostID: 1,
			},
			expectedErr: nil,
		},
		{
			name: "Empty Slug",
			session: models.Session{
				Slug:   "",
				HostID: 1,
			},
			expectedErr: fmt.Errorf("slug is empty"),
		},
		{
			name: "Empty HostID",
			session: models.Session{
				Slug:   "valid_slug",
				HostID: 0,
			},
			expectedErr: fmt.Errorf("host ID is empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateSession(tt.session)
			if tt.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestValidateQueue(t *testing.T) {
	tests := []struct {
		name        string
		queue       models.Queue
		expectedErr error
	}{
		{
			name: "Valid Queue",
			queue: models.Queue{
				TrackURI:  "spotify:track:123",
				SessionID: 1,
				UserID:    1,
			},
			expectedErr: nil,
		},
		{
			name: "Empty TrackURI",
			queue: models.Queue{
				TrackURI:  "",
				SessionID: 1,
				UserID:    1,
			},
			expectedErr: fmt.Errorf("track URI is required"),
		},
		{
			name: "Empty SessionID",
			queue: models.Queue{
				TrackURI:  "spotify:track:123",
				SessionID: 0,
				UserID:    1,
			},
			expectedErr: fmt.Errorf("session ID is required"),
		},
		{
			name: "Empty UserID",
			queue: models.Queue{
				TrackURI:  "spotify:track:123",
				SessionID: 1,
				UserID:    0,
			},
			expectedErr: fmt.Errorf("user ID is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateQueue(tt.queue)
			if tt.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name        string
		user        models.User
		expectedErr error
	}{
		{
			name: "Valid Spotify User",
			user: models.User{
				SpotifyUserID: newString("spotify123"),
				Email:         nil,
			},
			expectedErr: nil,
		},
		{
			name: "Valid Google User",
			user: models.User{
				SpotifyUserID: nil,
				Email:         newString("user@example.com"),
			},
			expectedErr: nil,
		},
		{
			name: "Valid Spotify and Google User",
			user: models.User{
				SpotifyUserID: newString("spotify123"),
				Email:         newString("user@example.com"),
			},
			expectedErr: nil,
		},
		{
			name: "Invalid User with no OAuth2 ID",
			user: models.User{
				SpotifyUserID: nil,
				Email:         nil,
			},
			expectedErr: fmt.Errorf("a valid oauth2 ID (google email or spotify ID) is required and not provided"),
		},
		{
			name: "Invalid User with empty Spotify ID",
			user: models.User{
				SpotifyUserID: newString(""),
				Email:         nil,
			},
			expectedErr: fmt.Errorf("a valid oauth2 ID (google email or spotify ID) is required and not provided"),
		},
		{
			name: "Invalid User with empty Email",
			user: models.User{
				SpotifyUserID: nil,
				Email:         newString(""),
			},
			expectedErr: fmt.Errorf("a valid oauth2 ID (google email or spotify ID) is required and not provided"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateUser(tt.user)
			if tt.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

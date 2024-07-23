package oauth2

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"

	"golang.org/x/oauth2"
)

// This package is in charge of verifying a user's identity via the Spotify OAuth2 flow.
// It is responsible for creating a new user in the database if the user does not exist,
// and returns a signed JWT token to the client if the user is successfully authenticated.

var Config *oauth2.Config

type StateData struct {
	State       string `json:"state"`
	RedirectURL string `json:"redirect_url"`
	ClientState string `json:"client_state"`
}

// GenerateRandomState generates a random state string of the specified length.
// The state string is generated using a combination of lowercase and uppercase letters,
// as well as digits from 0 to 9.
func GenerateRandomState(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	state := make([]rune, length)
	for i := range state {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		state[i] = chars[index.Int64()]
	}

	return string(state)
}

// EncodeState encodes the given random state and redirect URL into a string and returns it along with the state data.
// It takes a randomState string and a redirectURL string as input parameters.
// It returns a string representing the encoded state and a pointer to the StateData struct.
// If there is an error during the marshaling process, it logs the error and returns an empty string and nil.
func EncodeState(randomState string, clientState string) (string, *StateData) {
	state := StateData{
		State:       randomState,
		ClientState: clientState,
	}

	stateJSON, err := json.Marshal(state)
	if err != nil {
		return "", nil
	}

	return base64.URLEncoding.EncodeToString(stateJSON), &state
}

// Initialize initializes the OAuth2 configuration.
func Initialize(clientID, clientSecret, redirectURL, state string, scopes []string, authURL string, tokenURL string) {
	Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:   authURL,
			TokenURL:  tokenURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
}

// AuthCodeURL returns the URL to redirect the user to for authorization.
func AuthCodeURL(state string) string {
	return Config.AuthCodeURL(state)
}

// Exchange converts an authorization code into a token.
func Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return Config.Exchange(ctx, code)
}

// Client returns an HTTP client using the provided token.
func Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return Config.Client(ctx, token)
}

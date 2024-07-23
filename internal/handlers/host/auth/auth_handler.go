package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"garrettpfoy/orbit-api/internal/environment"
	"garrettpfoy/orbit-api/internal/services/oauth2"
	"net/http"
	"sync"
)

var (
	// stateStore is a map of state strings to StateData structs.
	stateStore = make(map[string]oauth2.StateData)

	// mu is a mutex to protect stateStore.
	mu sync.Mutex
)

// This package handles all interactions with the SpotifyAPI OAuth2 service. It is charged with
// verifying a user's identity via the Spotify OAuth2 flow. It is responsible for creating a new
// user in the database if the user does not exist, and returns a signed JWT token to the client
// if the user is successfully authenticated.

// handleLogin handles the login request and redirects the user to the appropriate URL.
// It validates the redirect URL and generates a random state for OAuth authentication.
// If any error occurs during the process, it redirects the user to the failure page with an error message.
//
// Parameters:
// - env: The OrbitEnvironment containing the necessary authentication environment variables.
// - w: The http.ResponseWriter used to write the response back to the client.
// - r: The http.Request representing the incoming request.
//
// Returns: None
func handleLogin(env *environment.OrbitEnvironment, w http.ResponseWriter, r *http.Request) {
	randomState := oauth2.GenerateRandomState(32)
	if randomState == "" {
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}

	var clientState string = ""
	if clientStateParam := r.URL.Query().Get("client_state"); clientStateParam != "" {
		clientState = clientStateParam
	}

	oAuthStateString, oAuthStateStruct := oauth2.EncodeState(randomState, clientState)

	mu.Lock()
	stateStore[randomState] = *oAuthStateStruct
	mu.Unlock()

	url := oauth2.AuthCodeURL(oAuthStateString)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// handleCallback handles the callback from the authentication provider.
// It exchanges the authorization code for an access token, retrieves user information,
// generates a JWT token, and redirects the user to the specified redirect URL.
// If any error occurs during the process, it redirects the user to the failure page.
//
// Parameters:
// - env: The authentication environment containing necessary configurations.
// - w: The http.ResponseWriter used to send the HTTP response.
// - r: The *http.Request representing the incoming HTTP request.
//
// Returns: None
func handleCallback(env *environment.OrbitEnvironment, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	stateEncoded := r.URL.Query().Get("state")

	code := r.URL.Query().Get("code")
	if code == "" {
		fmt.Println("No code provided.")
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}

	stateJSON, err := base64.URLEncoding.DecodeString(stateEncoded)
	if err != nil {
		fmt.Println("Failed to decode state.")
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}

	var state oauth2.StateData
	if err := json.Unmarshal(stateJSON, &state); err != nil {
		fmt.Println("Failed to unmarshal state data.")
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}

	mu.Lock()
	storedState, exists := stateStore[state.State]
	mu.Unlock()

	if !exists || storedState.RedirectURL != state.RedirectURL {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	mu.Lock()
	delete(stateStore, state.State)
	mu.Unlock()

	token, err := oauth2.Exchange(ctx, code)
	if err != nil {
		fmt.Println("Failed to exchange code for token.")
		http.Redirect(w, r, "https://google.com", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("Token recieved: ", token)

	// jwt, err := auth.CreateJWT(user.ID, env.JWTKey, (time.Minute * time.Duration(env.JWTLifespan)), []byte(env.JWTSecret))

	// if err != nil || jwt == "" {
	// 	http.Redirect(w, r, "/failure?error=Internal%20server%20error%20surrounding%20JWT%20generation", http.StatusTemporaryRedirect)
	// 	return
	// }

	// auth.ReturnJWT(env.IsProduction, env.JWTKey, jwt, env.Domain, time.Hour, w)

	// log.Log(log.InfoLevel, "User authenticated and JWT set.")

	// if state.Native {
	// 	// Add the JWT to the Authorization header for Native app
	// 	state.RedirectURL += "?token=" + jwt

	// 	log.Log(log.InfoLevel, "Redirecting to native app with redirect URL: "+state.RedirectURL)
	// }

	// http.Redirect(w, r, state.RedirectURL, http.StatusTemporaryRedirect)
}

package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type OrbitEnvironment struct {
	IS_PRODUCTION         bool   // Whether or not we are in production, used for various security measures
	ENCRYPTION_SECRET     string // Secret of size 32 that is used for salting/encrypting access tokens in the database
	JWT_SECRET            string // Secret that is used to sign JWT tokens
	SPOTIFY_CLIENT_ID     string // Spotify client ID
	SPOTIFY_CLIENT_SECRET string // Spotify client secret
	SPOTIFY_REDIRECT_URL  string // Spotify redirect URL
}

func LoadOrbitEnvironment(IS_PRODUCTION bool) (*OrbitEnvironment, error) {
	err := godotenv.Load()
	if err != nil && !IS_PRODUCTION {
		return nil, fmt.Errorf("there was an issue loading the .env file whilst in development mode %s", err.Error())
	}

	var orbitEnvironment OrbitEnvironment

	orbitEnvironment.IS_PRODUCTION = IS_PRODUCTION

	if encryptionSecret := os.Getenv("ENCRYPTION_SECRET"); encryptionSecret != "" && len(encryptionSecret) == 32 {
		// Valid encryption secret
		orbitEnvironment.ENCRYPTION_SECRET = encryptionSecret
	} else {
		return nil, fmt.Errorf("the required secret ENCRYPTION_SECRET is not valid or not supplied. Length: %d", len(encryptionSecret))
	}

	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		// Valid JWT secret
		orbitEnvironment.JWT_SECRET = jwtSecret
	} else {
		return nil, fmt.Errorf("the required secret JWT_SECRET is not valid or not supplied")
	}

	if spotifyClientID := os.Getenv("SPOTIFY_CLIENT_ID"); spotifyClientID != "" {
		// Valid Spotify client ID
		orbitEnvironment.SPOTIFY_CLIENT_ID = spotifyClientID
	} else {
		return nil, fmt.Errorf("the required secret SPOTIFY_CLIENT_ID is not valid or not supplied")
	}

	if spotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET"); spotifyClientSecret != "" {
		// Valid Spotify client secret
		orbitEnvironment.SPOTIFY_CLIENT_SECRET = spotifyClientSecret
	} else {
		return nil, fmt.Errorf("the required secret SPOTIFY_CLIENT_SECRET is not valid or not supplied")
	}

	if spotifyRedirectURL := os.Getenv("SPOTIFY_REDIRECT_URL"); spotifyRedirectURL != "" {
		// Valid Spotify redirect URL
		orbitEnvironment.SPOTIFY_REDIRECT_URL = spotifyRedirectURL
	} else {
		return nil, fmt.Errorf("the required secret SPOTIFY_REDIRECT_URL is not valid or not supplied")
	}

	return &orbitEnvironment, nil
}

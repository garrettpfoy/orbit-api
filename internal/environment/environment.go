package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type OrbitEnvironment struct {
	IS_PRODUCTION     bool   // Whether or not we are in production, used for various security measures
	ENCRYPTION_SECRET string // Secret of size 32 that is used for salting/encrypting access tokens in the database
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

	return &orbitEnvironment, nil
}

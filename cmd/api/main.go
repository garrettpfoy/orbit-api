package main

import (
	"log"
	"time"

	"garrettpfoy/orbit-api/internal/repositories/access_token"
	// "garrettpfoy/orbit-api/internal/repositories/queue"
	// "garrettpfoy/orbit-api/internal/repositories/session"
	// "garrettpfoy/orbit-api/internal/repositories/user"
	// "garrettpfoy/orbit-api/internal/repositories/vote"

	"garrettpfoy/orbit-api/internal/services/encryption"

	"garrettpfoy/orbit-api/internal/environment"

	"garrettpfoy/orbit-api/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	environment, err := environment.LoadOrbitEnvironment(false)
	if err != nil {
		log.Fatal("failed to establish the orbit environment: ", err)
	}

	models.SetEncryptionService(encryption.NewEncryptionService(environment.ENCRYPTION_SECRET))

	// Auto migrate the schema
	db.AutoMigrate(&models.Session{}, &models.AccessToken{}, &models.Queue{}, &models.User{})

	accessTokenRepo := access_token.NewGormAccessTokenRepository(db)

	// Example: Creating a new access token
	newToken := models.AccessToken{
		UserID:       1, // assuming user ID 1 exists
		SessionID:    1,
		AccessToken:  "example_access_token",
		RefreshToken: "example_refresh_token",
		ExpiryTime:   time.Now().Add(time.Hour * 1),
	}
	err = accessTokenRepo.CreateAccessToken(&newToken)
	if err != nil {
		log.Fatalf("Failed to create access token: %v", err)
	}

	// Fetch and decrypt the access token
	token, err := accessTokenRepo.GetAccessToken(newToken.ID)
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}
	log.Printf("Decrypted Access Token: %s", token.AccessToken)
	log.Printf("Decrypted Refresh Token: %s", token.RefreshToken)
}

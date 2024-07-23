package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	auth "garrettpfoy/orbit-api/internal/services/jwt"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("my-secret-key")
	userID := "123"
	duration := time.Hour

	token, err := auth.CreateJWT(userID, duration, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token claims
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secret, nil
	})
	assert.NoError(t, err)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, userID, claims["id"])
}

func TestVerifyJWT(t *testing.T) {
	secret := []byte("my-secret-key")
	userID := "123"
	duration := time.Hour

	token, err := auth.CreateJWT(userID, duration, secret)
	assert.NoError(t, err)

	err = auth.VerifyJWT(token, secret)
	assert.NoError(t, err)

	// Test with an invalid token
	invalidToken := "invalid-token"
	err = auth.VerifyJWT(invalidToken, secret)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse JWT token")
}

func TestReturnJWT(t *testing.T) {
	secret := []byte("my-secret-key")
	userID := "123"
	duration := time.Hour
	domain := "example.com"
	isProd := true
	key := "savanna-jwt"

	token, err := auth.CreateJWT(userID, duration, secret)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	auth.ReturnJWT(isProd, key, token, domain, duration, w)

	cookie := w.Result().Cookies()[0]
	assert.Equal(t, key, cookie.Name)
	assert.Equal(t, token, cookie.Value)
	assert.Equal(t, domain, cookie.Domain)
	assert.Equal(t, true, cookie.HttpOnly)
	assert.Equal(t, isProd, cookie.Secure)
	assert.Equal(t, http.SameSiteStrictMode, cookie.SameSite)
}

func TestSetCrossOriginResourceSharing(t *testing.T) {
	domain := "example.com"
	w := httptest.NewRecorder()

	auth.SetCrossOriginResourceSharing(domain, w)

	headers := w.Result().Header
	assert.Equal(t, "https://*.example.com", headers.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET, POST, PATCH, PUT, DELETE", headers.Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type, Authorization", headers.Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "true", headers.Get("Access-Control-Allow-Credentials"))
}

func TestSetContentSecurityPolicy(t *testing.T) {
	w := httptest.NewRecorder()

	auth.SetContentSecurityPolicy(nil, w)
	headers := w.Result().Header
	assert.Equal(t, "default-src 'self'; script-src 'self'; style-src 'self'; object-src 'none'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'; block-all-mixed-content; upgrade-insecure-requests;", headers.Get("Content-Security-Policy"))

	w2 := httptest.NewRecorder()
	customPolicy := "default-src 'none';"
	auth.SetContentSecurityPolicy(&customPolicy, w2)
	headers = w2.Result().Header
	assert.Equal(t, customPolicy, headers.Get("Content-Security-Policy"))
}

func TestSetCookie(t *testing.T) {
	isProd := true
	domain := "example.com"
	jwtToken := "my-jwt-token"
	key := "orbit"
	duration := time.Hour

	w := httptest.NewRecorder()
	auth.SetCookie(isProd, domain, jwtToken, key, duration, w)

	cookie := w.Result().Cookies()[0]
	assert.Equal(t, key, cookie.Name)
	assert.Equal(t, jwtToken, cookie.Value)
	assert.Equal(t, domain, cookie.Domain)
	assert.Equal(t, true, cookie.HttpOnly)
	assert.Equal(t, isProd, cookie.Secure)
	assert.Equal(t, http.SameSiteStrictMode, cookie.SameSite)
	assert.WithinDuration(t, time.Now().Add(duration), cookie.Expires, time.Second)
}

package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateJWT creates a new JWT token with the provided claims and signs it with the given secret.
// The function returns the generated token as a string.
//
// Parameters:
//   - userID: The ID of the user.
//   - duration: The duration of the token's validity.
//   - secret: The secret key used for signing the token.
//
// Returns:
//   - string: The generated JWT token.
//
// Example:
//
//	secret := []byte("my-secret-key")
//	token := CreateJWT("123", "John Doe", "john@example.com", time.Hour, secret)
func CreateJWT(userID string, duration time.Duration, secret []byte) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS512)

	// Set standard claims
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "orbit"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(duration).Unix()

	// Set custom claims
	claims["id"] = userID

	// Sign the token with the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}

	return tokenString, nil
}

// VerifyJWT verifies the provided JWT token using the given secret key.
// The function returns an error if the JWT is not valid, otherwise,
// it returns nil.
func VerifyJWT(tokenString string, secret []byte) error {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token's signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return fmt.Errorf("failed to parse JWT token: %w", err)
	}

	// Validate the token and its claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check token expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return fmt.Errorf("token has expired")
			}
		} else {
			return fmt.Errorf("token does not contain an exp claim")
		}
		return nil
	} else {
		return fmt.Errorf("invalid token")
	}
}

// ReturnJWT sets the JWT cookie and additional headers based on the provided parameters.
// If the isProd flag is set to true, it sets the cookie with the JWT and sets the necessary
// headers for production environment. If the flag is set to false, it only sets the cookie
// and does not set the additional headers.
//
// Parameters:
//   - isProd: A boolean flag indicating whether the code is running in production environment.
//   - jwt: The JWT token to be set as a cookie.
//   - domain: The domain for which the cookie is valid.
//   - duration: The duration for which the cookie is valid.
//   - w: The http.ResponseWriter to write the cookie and headers to.
//
// Example usage:
//
//	ReturnJWT(true, "my-jwt-token", "example.com", time.Hour, w)
func ReturnJWT(isProd bool, key, jwt, domain string, duration time.Duration, w http.ResponseWriter) {
	// First, we need to set the cookie with the JWT
	// see the Readme for specifics and reasoning behind each option
	SetCookie(isProd, domain, jwt, key, duration, w)

	// Cookie has been set, if we want to be fully secure (e.g. in
	// production), we should also set the following headers
	if !isProd {
		// Next, setup the CSP headers
		SetContentSecurityPolicy(nil, w)

		// Now we can set CORS headers
		SetCrossOriginResourceSharing(domain, w)
	}
}

// setCrossOriginResourceSharing sets the necessary headers for Cross-Origin Resource Sharing (CORS).
// It takes a domain string and a http.ResponseWriter as parameters.
func SetCrossOriginResourceSharing(domain string, w http.ResponseWriter) {
	// Format the domain into a Allow-Origin header
	allowedOrigin := fmt.Sprintf("https://*.%s", domain)

	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

// setCookie sets a JWT cookie in the HTTP response.
// The cookie is named "savanna-jwt" and its value is set to the provided JWT.
// The cookie is set to expire after the specified duration.
// If isProd is true, the cookie is set as secure.
// The cookie is set with HttpOnly flag to prevent client-side access.
// The cookie is set with SameSiteLaxMode to allow cross-site requests with safe HTTP methods.
func SetCookie(isProd bool, domain, jwt, key string, duration time.Duration, w http.ResponseWriter) {
	var siteMode http.SameSite = http.SameSiteStrictMode
	if !isProd {
		siteMode = http.SameSiteLaxMode
	}

	http.SetCookie(w, &http.Cookie{
		Name:     key,
		Value:    jwt,
		Path:     "/",
		HttpOnly: true,
		Secure:   isProd,
		Expires:  time.Now().Add(duration),
		SameSite: siteMode,
		Domain:   domain,
	})
}

// setContentSecurityPolicy sets the Content-Security-Policy header for the HTTP response.
// If the policy parameter is nil, it sets a default CSP policy with strict settings.
// If the policy parameter is provided, it sets the CSP policy based on the provided value.
// The Content-Security-Policy header defines the content security policy for the web page,
// specifying which resources can be loaded and executed by the browser.
//
// Parameters:
// - policy: A pointer to a string representing the CSP policy. If nil, a default policy is used.
// - w: The http.ResponseWriter to set the header on.
//
// Example usage:
//
//	setContentSecurityPolicy(nil, w)
//	setContentSecurityPolicy(&customPolicy, w)
func SetContentSecurityPolicy(policy *string, w http.ResponseWriter) {
	// See README for reasoning behind CSP headers and settings
	if policy == nil {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; object-src 'none'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'; block-all-mixed-content; upgrade-insecure-requests;")
		return
	}
	w.Header().Set("Content-Security-Policy", *policy)
}

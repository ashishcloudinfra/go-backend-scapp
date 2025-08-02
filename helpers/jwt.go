package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"server.simplifycontrol.com/secrets"
)

type Claims struct {
	ID          string
	Username    string
	Role        string
	Permissions string
}

// GenerateJWT generates a new JWT for a given user ID
func GenerateJWT(obj Claims) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Define claims (you can add more custom claims here)
	claims := jwt.MapClaims{
		"id":          obj.ID,
		"username":    obj.Username,
		"role":        obj.Role,
		"permissions": obj.Permissions,
		"exp":         expirationTime.UnixMilli(),
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(secrets.SecretJSON.JWTConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates the JWT and returns the claims if valid
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secrets.SecretJSON.JWTConfig.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and return the claims if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

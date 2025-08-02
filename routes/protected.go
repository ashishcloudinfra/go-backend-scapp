package routes

import (
	"net/http"
	"strings"

	"server.simplifycontrol.com/helpers"
)

// ProtectedHandler requires a valid JWT to access
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		helpers.SendJSONError(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := helpers.ValidateJWT(tokenString)
	if err != nil {
		helpers.SendJSONError(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{
		"message": "Welcome to the protected route!",
		"user_id": claims["user_id"],
	})
}

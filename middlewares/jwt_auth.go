package middleware

import (
	"net/http"
	"strings"

	"server.simplifycontrol.com/helpers"
)

// JWTAuthMiddleware checks for a valid JWT token in the Authorization header,
// but bypasses validation for endpoints defined in jwtBypassRules.
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Loop through bypass rules to see if current request should bypass JWT check.
		for _, rule := range helpers.JwtBypassRules {
			if rule.Pattern.MatchString(r.URL.Path) {
				// If Methods slice is not empty, ensure the request's method is allowed.
				if len(rule.Methods) > 0 {
					methodAllowed := false
					for _, method := range rule.Methods {
						if strings.EqualFold(r.Method, method) {
							methodAllowed = true
							break
						}
					}
					if methodAllowed {
						next.ServeHTTP(w, r)
						return
					}
				} else {
					// If Methods is empty, bypass for any method.
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		// Proceed with JWT validation for all requests that aren't bypassed.
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		// Expect the Authorization header to be in the format "Bearer <token>".
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Unauthorized: invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Validate the token using your helper function.
		_, err := helpers.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		// If token is valid, proceed to the next handler.
		next.ServeHTTP(w, r)
	})
}

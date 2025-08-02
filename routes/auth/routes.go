package auth

import "github.com/gorilla/mux"

// RegisterAuthRoutes registers routes for authentication
func RegisterAuthRoutes(mux *mux.Router) {
	mux.HandleFunc("/login", LoginHandler).Methods("Post")
	mux.HandleFunc("/signup", SignupHandler).Methods("Post")
}

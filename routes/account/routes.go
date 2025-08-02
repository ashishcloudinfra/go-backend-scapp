package account

import "github.com/gorilla/mux"

func RegisterAccountRoutes(mux *mux.Router) {
	mux.HandleFunc("/account/user/{userId}", GetAccountDetailsHandler).Methods("GET")
}

package personalDetail

import "github.com/gorilla/mux"

func RegisterPersonalDetailsRoutes(mux *mux.Router) {
	mux.HandleFunc("/pd/c/{companyId}/role/{role}", GetAllUsersByCompanyIDAndRole).Methods("GET")
	mux.HandleFunc("/pd/{userId}", GetUserDetailByUserID).Methods("GET")
	mux.HandleFunc("/pd/c/{companyId}", RegisterUserDetails).Methods("POST")
	mux.HandleFunc("/register/c/{companyId}", RegisterNewUser).Methods("POST")
	mux.HandleFunc("/user-details/{userDetailId}", DeleteUserDetail).Methods("DELETE")
	mux.HandleFunc("/pd/{userId}", EditUserDetails).Methods("PUT")
}

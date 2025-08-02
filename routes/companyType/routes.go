package companyType

import "github.com/gorilla/mux"

func RegisterCompanyTypeRoutes(mux *mux.Router) {
	mux.HandleFunc("/companyTypes", GetCompanyTypesHandler).Methods("Get")
}

package company

import "github.com/gorilla/mux"

// RegisterCompanyRoutes registers routes for company-related operations
func RegisterCompanyRoutes(mux *mux.Router) {
	mux.HandleFunc("/company", CreateCompanyHandler).Methods("Post")
	mux.HandleFunc("/companies/{userId}", GetCompanyDetailsHandler).Methods("Get")
	mux.HandleFunc("/cd/{companyId}", FetchCompanyByCompanyId).Methods("Get")
	mux.HandleFunc("/company/{userId}", FetchAssociatedCompany).Methods("Get")
}

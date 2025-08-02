package membershipPlan

import "github.com/gorilla/mux"

func RegisterMembershipPlanRoutes(mux *mux.Router) {
	mux.HandleFunc("/membershipPlan/{companyId}", RegisterMembershipPlan).Methods("POST")
	mux.HandleFunc("/membershipPlan/{companyId}", GetAllMembershipPlan).Methods("GET")
	mux.HandleFunc("/c/{companyId}/mp/{membershipPlanId}", GetAllMembershipPlanById).Methods("GET")
	mux.HandleFunc("/membershipPlan/{planId}", DeleteMembershipPlan).Methods("DELETE")
	mux.HandleFunc("/membershipPlan/{planId}", EditMembershipPlan).Methods("PUT")
}

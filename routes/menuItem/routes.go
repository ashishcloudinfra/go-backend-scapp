package menuItem

import "github.com/gorilla/mux"

func RegisterMenuItemRoutes(mux *mux.Router) {
	mux.HandleFunc("/c/{companyId}/menuItem", RegisterMenuItem).Methods("POST")
	mux.HandleFunc("/c/{companyId}/menuItem/bulk", RegisterBulkMenuItem).Methods("POST")
	mux.HandleFunc("/c/{companyId}/scan", ScanMenu).Methods("POST")
	mux.HandleFunc("/c/{companyId}/menuItems", GetAllMenuItem).Methods("GET")
	mux.HandleFunc("/c/{companyId}/menuItem/{menuItemId}", GetMenuItemWithId).Methods("GET")
	mux.HandleFunc("/c/{companyId}/menuItem/{menuItemId}", UpdateMenuItem).Methods("PUT")
	mux.HandleFunc("/c/{companyId}/menuItem/{menuItemId}", DeleteMenuItem).Methods("DELETE")
}

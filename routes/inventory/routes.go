package inventory

import "github.com/gorilla/mux"

func RegisterInventoryRoutes(mux *mux.Router) {
	mux.HandleFunc("/c/{companyId}/itemDetails", RegisterItemDetails).Methods("POST")
	mux.HandleFunc("/c/{companyId}/addItems", AddEquipments).Methods("POST")
	mux.HandleFunc("/c/{companyId}/updateStatus", UpdateEquipmentsStatus).Methods("PUT")
	mux.HandleFunc("/c/{companyId}/itemTypes", GetItemTypes).Methods("GET")
	mux.HandleFunc("/equipments", GetEquipmentList).Methods("GET")
}

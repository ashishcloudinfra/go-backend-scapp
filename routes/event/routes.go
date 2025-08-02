package event

import "github.com/gorilla/mux"

func RegisterEventRoutes(mux *mux.Router) {
	mux.HandleFunc("/c/{companyId}/event", RegisterEvent).Methods("POST")
	// mux.HandleFunc("/eventRoom/{eventRoomId}", EditEventRoom).Methods("PUT")
	mux.HandleFunc("/c/{companyId}/events", getAllEvents).Methods("GET")
	// mux.HandleFunc("/c/{companyId}/eventRoom/{eventRoomId}", GetEventRoomById).Methods("GET")
}

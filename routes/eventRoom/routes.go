package eventRoom

import "github.com/gorilla/mux"

func RegisterEventRoomRoutes(mux *mux.Router) {
	mux.HandleFunc("/c/{companyId}/eventRoom", RegisterEventRoom).Methods("POST")
	mux.HandleFunc("/eventRoom/{eventRoomId}", EditEventRoom).Methods("PUT")
	mux.HandleFunc("/c/{companyId}/eventRooms", GetEventRoom).Methods("GET")
	mux.HandleFunc("/c/{companyId}/eventRoom/{eventRoomId}", GetEventRoomById).Methods("GET")
}

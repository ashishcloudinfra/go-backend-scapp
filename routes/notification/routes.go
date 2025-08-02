package notification

import "github.com/gorilla/mux"

func RegisterNotificationRoutes(mux *mux.Router) {
	mux.HandleFunc("/notification/sendEmail", SendEmailHandler).Methods("POST")
}

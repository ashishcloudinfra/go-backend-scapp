package payment

import "github.com/gorilla/mux"

func RegisterPaymentRoutes(mux *mux.Router) {
	mux.HandleFunc("/create-payment-order", CreateOrder).Methods("POST")
}

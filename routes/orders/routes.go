package orders

import "github.com/gorilla/mux"

func RegisterOrdersRoutes(mux *mux.Router) {
	mux.HandleFunc("/orders/order", RegisterOrder).Methods("POST")
	mux.HandleFunc("/orders/tn/{tableNumber}/c/{companyId}", GetOrderByTableNumberAndCompanyId).Methods("GET")
	mux.HandleFunc("/orders/admin/c/{companyId}", GetOrdersListForAdmin).Methods("GET")
	mux.HandleFunc("/orders/tn/{tableNumber}/c/{companyId}", UpdateOrder).Methods("PUT")
	mux.HandleFunc("/order/{orderId}/status/{status}", UpdateOrderStatus).Methods("PUT")
}

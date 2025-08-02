package personalfinance

import "github.com/gorilla/mux"

func RegisterPersonalFinanceRoutes(mux *mux.Router) {
	mux.HandleFunc("/{userId}/refreshPortfolio", RefreshPortfolio).Methods("POST")
	mux.HandleFunc("/{userId}/personalFinance/prompt", Promt).Methods("POST")
}

package scripts

import "github.com/gorilla/mux"

func RegisterScriptsRoutes(mux *mux.Router) {
	mux.HandleFunc("/price/{tickerSymbol}", GetStockPriceData).Methods("GET")
}

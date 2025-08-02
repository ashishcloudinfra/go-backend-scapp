package assets

import "github.com/gorilla/mux"

func RegisterAssetsRoutes(mux *mux.Router) {
	mux.HandleFunc("/assetTypes", GetAssetsType).Methods("GET")
	mux.HandleFunc("/{userId}/assets", GetAssets).Methods("GET")
	mux.HandleFunc("/{userId}/asset", AddAsset).Methods("POST")
	mux.HandleFunc("/assetType", AddAssetType).Methods("POST")
	mux.HandleFunc("/{userId}/asset/{assetId}", UpdateAsset).Methods("PUT")
	mux.HandleFunc("/{userId}/asset/{assetId}", DeleteAsset).Methods("DELETE")
}

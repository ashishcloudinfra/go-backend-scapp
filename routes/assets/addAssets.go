package assets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/types"
)

func AddAssetType(w http.ResponseWriter, r *http.Request) {
	var assetTypeReqBody types.AssetTypeReqBody
	if err := json.NewDecoder(r.Body).Decode(&assetTypeReqBody); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := helpers.InsertAssetType(assetTypeReqBody)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting asset type", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": "success"})
}

func AddAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}
	var assetItemReqBody types.AssetItemReqBody
	if err := json.NewDecoder(r.Body).Decode(&assetItemReqBody); err != nil {
		log.Println(err)
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := helpers.InsertAssetItem(userId, assetItemReqBody)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting asset type", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": "success"})
}

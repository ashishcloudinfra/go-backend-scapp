package assets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/types"
)

func UpdateAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	assetId := vars["assetId"]

	// Validate required path parameters
	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}
	if assetId == "" {
		helpers.SendJSONError(w, "assetId is required", http.StatusBadRequest)
		return
	}

	// Decode request body
	var assetItemReqBody types.AssetItemReqBody
	if err := json.NewDecoder(r.Body).Decode(&assetItemReqBody); err != nil {
		log.Println(err)
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call your helper to update the asset
	err := helpers.UpdateAssetItem(userId, assetId, assetItemReqBody)
	if err != nil {
		log.Printf("Error updating asset: %v\n", err)
		helpers.SendJSONError(w, "Error updating asset", http.StatusInternalServerError)
		return
	}

	// On successful update
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{
		"data": "update successful",
	})
}

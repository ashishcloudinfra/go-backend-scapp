package assets

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
)

func DeleteAsset(w http.ResponseWriter, r *http.Request) {
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

	// Delete the asset via the helper
	err := helpers.DeleteAssetItem(userId, assetId)
	if err != nil {
		log.Printf("Error deleting asset: %v\n", err)
		helpers.SendJSONError(w, "Error deleting asset", http.StatusBadRequest)
		return
	}

	// On successful deletion
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{
		"data": "delete successful",
	})
}

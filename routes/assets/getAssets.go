package assets

import (
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
)

func GetAssetsType(w http.ResponseWriter, r *http.Request) {
	data, err := helpers.GetAssetTypes()
	if err != nil {
		helpers.SendJSONError(w, "Error fetching asset types", http.StatusBadRequest)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": data})
}

func GetAssets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}
	data, err := helpers.GetAssetItems(userId)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching assets", http.StatusBadRequest)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": data})
}

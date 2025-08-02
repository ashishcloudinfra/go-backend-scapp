package inventory

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func AddEquipments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var reqBody types.AddItemRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (status, equipmentId, companyId) VALUES ($1, $2, $3)`, database.INVENTORY_ITEM_TABLE_NAME)
	_, err := database.Execute(query, reqBody.Status, reqBody.EquipmentId, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into items database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

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

func RegisterItemDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var reqBody types.InventoryItemDescriptionRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, img, type, instructions) VALUES ($1, $2, $3, $4)`, database.INVENTORY_EQUIPMENT_TABLE_NAME)
	_, err := database.Execute(query, reqBody.Name, reqBody.Img, reqBody.Type, reqBody.Instructions)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into inventory category database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

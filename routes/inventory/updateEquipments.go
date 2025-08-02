package inventory

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func UpdateEquipmentsStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var reqBody types.UpdateItemRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT count(*) FROM %s WHERE equipmentId = $1 and companyId = $2 and status = $3`, database.INVENTORY_ITEM_TABLE_NAME)
	rows, err := database.Query(query, reqBody.EquipmentId, companyId, reqBody.PrevStatus)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from inventory item database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result []int
	for rows.Next() {
		var mp int
		err := rows.Scan(&mp)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		result = append(result, mp)
	}

	// check if count is greater or equal to the count in the database
	if result[0] < reqBody.Count {
		helpers.SendJSONError(w, "Count is greater than the count in the database", http.StatusBadRequest)
		return
	}

	// update the status of the all count item
	for range reqBody.Count {
		query = fmt.Sprintf(`UPDATE %s SET status = $1 WHERE id in (SELECT id 
    FROM item 
    WHERE equipmentId = $2 and companyId = $3 and status = $4
    LIMIT 1)`, database.INVENTORY_ITEM_TABLE_NAME)
		_, err = database.Query(query, reqBody.NewStatus, reqBody.EquipmentId, companyId, reqBody.PrevStatus)
		if err != nil {
			helpers.SendJSONError(w, "Error updating item database", http.StatusInternalServerError)
			return
		}
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": "success"})
}

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

func GetEquipmentList(w http.ResponseWriter, r *http.Request) {
	query := fmt.Sprintf(`SELECT id, name, img, type, instructions FROM %s`, database.INVENTORY_EQUIPMENT_TABLE_NAME)
	rows, err := database.Query(query)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from inventory category database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result []types.Equipment
	for rows.Next() {
		var mp types.Equipment
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Img, &mp.Type, &mp.Instructions)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		result = append(result, mp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]types.Equipment{"data": result})
}

func GetItemTypes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	workingEquipments, err := helpers.GetInventoryByStatus(companyId, "available")
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from inventory database", http.StatusInternalServerError)
		return
	}

	maintenanceEquipments, err := helpers.GetInventoryByStatus(companyId, "maintenance")
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from inventory database", http.StatusInternalServerError)
		return
	}

	decommissionedEquipments, err := helpers.GetInventoryByStatus(companyId, "decommissioned")
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from inventory database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": map[string]interface{}{
		"available":      workingEquipments,
		"maintenance":    maintenanceEquipments,
		"decommissioned": decommissionedEquipments,
	}})
}

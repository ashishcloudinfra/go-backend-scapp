package eventRoom

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

func GetEventRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * from %s where companyId = $1`, database.EVENT_ROOMS_TABLE_NAME)
	rows, err := database.Query(query, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var eventRooms []types.EventRoom
	for rows.Next() {
		var mp types.EventRoom
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Capacity, &mp.Location, &mp.IsUnderMaintenance, &mp.StartTime, &mp.EndTime, &mp.CreatedAt, &mp.UpdatedAt, &mp.CompanyId)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		eventRooms = append(eventRooms, mp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]types.EventRoom{"data": eventRooms})
}

func GetEventRoomById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	eventRoomId := vars["eventRoomId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * from %s where companyId = $1 and id = $2`, database.EVENT_ROOMS_TABLE_NAME)
	rows, err := database.Query(query, companyId, eventRoomId)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var eventRooms []types.EventRoom
	for rows.Next() {
		var mp types.EventRoom
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Capacity, &mp.Location, &mp.IsUnderMaintenance, &mp.StartTime, &mp.EndTime, &mp.CreatedAt, &mp.UpdatedAt, &mp.CompanyId)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		eventRooms = append(eventRooms, mp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]types.EventRoom{"data": eventRooms[0]})
}

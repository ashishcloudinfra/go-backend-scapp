package eventRoom

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func RegisterEventRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var eventRoomReq types.EventRoomRequestBody
	if err := json.NewDecoder(r.Body).Decode(&eventRoomReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, capacity, location, isUnderMaintenance, startTime, endTime, companyId) VALUES ($1, $2, $3, $4, $5, $6, $7)`, database.EVENT_ROOMS_TABLE_NAME)
	_, err := database.Execute(query, eventRoomReq.Name, eventRoomReq.Capacity, eventRoomReq.Location, eventRoomReq.IsUnderMaintenance, eventRoomReq.StartTime, eventRoomReq.EndTime, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into identity database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

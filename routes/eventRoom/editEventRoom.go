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

func EditEventRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventRoomId := vars["eventRoomId"]

	if eventRoomId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var eventRoomReq types.EventRoomRequestBody
	if err := json.NewDecoder(r.Body).Decode(&eventRoomReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`UPDATE %s SET name = $1, capacity = $2, location = $3, isUnderMaintenance = $4, startTime = $5, endTime = $6 where id = $7`, database.EVENT_ROOMS_TABLE_NAME)
	_, err := database.Execute(query, eventRoomReq.Name, eventRoomReq.Capacity, eventRoomReq.Location, eventRoomReq.IsUnderMaintenance, eventRoomReq.StartTime, eventRoomReq.EndTime, eventRoomId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into identity database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

package event

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

func RegisterEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var eventReq types.EventFormValues
	if err := json.NewDecoder(r.Body).Decode(&eventReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	metadata, marsErr := json.Marshal(eventReq.Metadata)
	if marsErr != nil {
		helpers.SendJSONError(w, "Error marshalling metadata", http.StatusInternalServerError)
		return
	}

	eventRoomId := eventReq.EventRoomId

	if eventRoomId == "" {
		eventRoomId = helpers.DEFAULT_UUID
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, description, organiserId, eventRoomId, startDate, endDate, startTime, endTime, isAllDayEvent, isRecurring, recurrenceType, status, metadata, companyId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`, database.EVENTS_TABLE_NAME)
	_, err := database.Execute(query, eventReq.Name, eventReq.Description, eventReq.OrganiserId, eventRoomId, eventReq.StartDate, eventReq.EndDate, eventReq.StartTime, eventReq.EndTime, eventReq.IsAllDayEvent, eventReq.IsRecurring, eventReq.RecurrenceType, eventReq.Status, metadata, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into events database", http.StatusInternalServerError)
		return
	}

	query = fmt.Sprintf(`select id from %s where name = $1 and description = $2`, database.EVENTS_TABLE_NAME)
	rows, err := database.Query(query, eventReq.Name, eventReq.Description)
	if err != nil {
		helpers.SendJSONError(w, "Error querying from events database", http.StatusInternalServerError)
		return
	}

	var eventId string
	var events []string
	for rows.Next() {
		var event string
		err := rows.Scan(&event)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		events = append(events, event)
	}

	eventId = events[0]

	// Create Scheduled Events
	if err := helpers.CreateScheduledEvents(eventId, eventReq); err != nil {
		helpers.SendJSONError(w, "Error creating scheduled events", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

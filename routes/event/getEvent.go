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

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`select f.*, er.name as eventRoomName from (select event.*, pd.firstName || ' ' || pd.lastName as organiser from %s
		join %s as pd
		on pd.iamId = event.organiserId
		where event.companyId = $1) as f
		left join %s as er
		on f.eventRoomId = er.id`, database.EVENTS_TABLE_NAME, database.PERSONAL_DETAIL_TABLE_NAME, database.EVENT_ROOMS_TABLE_NAME)
	rows, err := database.Query(query, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into events database", http.StatusInternalServerError)
		return
	}

	var events []types.EventWithOrganiserAndRoomName
	for rows.Next() {
		var event types.EventWithOrganiserAndRoomName
		var metadata []byte
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.OrganiserID,
			&event.EventRoomID,
			&event.StartDate,
			&event.EndDate,
			&event.StartTime,
			&event.EndTime,
			&event.IsAllDayEvent,
			&event.IsRecurring,
			&event.RecurrenceType,
			&event.Status,
			&metadata,
			&event.CreatedAt,
			&event.UpdatedAt,
			&event.CompanyID,
			&event.Organiser,
			&event.EventRoomName,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(metadata, &event.Metadata); err != nil {
			helpers.SendJSONError(w, "Error unmarshalling data", http.StatusInternalServerError)
			return
		}

		events = append(events, event)
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": events})
}

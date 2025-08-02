package membershipPlan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
)

func DeleteMembershipPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	planId := vars["planId"]

	if planId == "" {
		helpers.SendJSONError(w, "planId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT count(*) as count from %s where plans = $1 and status = 'authorized'`, database.ACCOUNT_DETAIL_TABLE_NAME)
	rows, err := database.Query(query, planId)
	if err != nil {
		helpers.SendJSONError(w, "Error deleting from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var counts []int
	for rows.Next() {
		var cnt int
		err := rows.Scan(&cnt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		counts = append(counts, cnt)
	}

	if counts[0] > 0 {
		helpers.SendJSONError(w, "Cannot delete membership plan with existing accounts", http.StatusBadRequest)
		return
	}

	query = fmt.Sprintf(`DELETE from %s where id = $1`, database.MEMBERSHIP_PLAN_TABLE_NAME)
	_, err = database.Execute(query, planId)
	if err != nil {
		helpers.SendJSONError(w, "Error deleting from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"success": "true"})
}

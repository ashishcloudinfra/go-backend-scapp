package membershipPlan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func GetAllMembershipPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * from %s where companyId = $1`, database.MEMBERSHIP_PLAN_TABLE_NAME)
	rows, err := database.Query(query, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var membershipPlans []types.MembershipPlan
	for rows.Next() {
		var mp types.MembershipPlan
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Price, &mp.Duration, &mp.Discount, pq.Array(&mp.Features), &mp.Status, pq.Array(&mp.CancellationPolicy), &mp.CompanyId, &mp.CreatedAt, &mp.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		membershipPlans = append(membershipPlans, mp)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]types.MembershipPlan{"data": membershipPlans})
}

func GetAllMembershipPlanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	membershipPlanId := vars["membershipPlanId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * from %s where companyId = $1 and id = $2`, database.MEMBERSHIP_PLAN_TABLE_NAME)
	rows, err := database.Query(query, companyId, membershipPlanId)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var membershipPlans []types.MembershipPlan
	for rows.Next() {
		var mp types.MembershipPlan
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Price, &mp.Duration, &mp.Discount, pq.Array(&mp.Features), &mp.Status, pq.Array(&mp.CancellationPolicy), &mp.CompanyId, &mp.CreatedAt, &mp.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		membershipPlans = append(membershipPlans, mp)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]types.MembershipPlan{"data": membershipPlans[0]})
}

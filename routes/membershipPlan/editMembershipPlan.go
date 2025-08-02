package membershipPlan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func EditMembershipPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	planId := vars["planId"]

	if planId == "" {
		helpers.SendJSONError(w, "planId is required", http.StatusBadRequest)
		return
	}

	var membershipPlanReq types.MembershipPlanFormValues
	if err := json.NewDecoder(r.Body).Decode(&membershipPlanReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`UPDATE %s SET name = $1, price = $2, duration = $3, discount = $4, features = $5, status = $6, cancellation_policy = $7 WHERE id = $8`, database.MEMBERSHIP_PLAN_TABLE_NAME)
	_, err := database.Execute(query, membershipPlanReq.Name, membershipPlanReq.Price, membershipPlanReq.Duration, membershipPlanReq.Discount, pq.Array(membershipPlanReq.Features), membershipPlanReq.Status, pq.Array(membershipPlanReq.CancellationPolicy), planId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"success": "true"})
}

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

func RegisterMembershipPlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyId"]

	if companyID == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var membershipPlanReq types.MembershipPlanFormValues
	if err := json.NewDecoder(r.Body).Decode(&membershipPlanReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, price, duration, discount, features, status, cancellation_policy, companyId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, database.MEMBERSHIP_PLAN_TABLE_NAME)
	_, err := database.Execute(query, membershipPlanReq.Name, membershipPlanReq.Price, membershipPlanReq.Duration, membershipPlanReq.Discount, pq.Array(membershipPlanReq.Features), membershipPlanReq.Status, pq.Array(membershipPlanReq.CancellationPolicy), companyID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

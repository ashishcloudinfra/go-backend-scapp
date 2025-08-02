package personalDetail

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func EditUserDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId are required", http.StatusBadRequest)
		return
	}

	var memberDetail types.UserDetailPostRequestBody
	err := json.NewDecoder(r.Body).Decode(&memberDetail)
	if err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`UPDATE %s SET status = $1, role = $2, permissions = $3, plans = $4 WHERE iamId = $5`, database.ACCOUNT_DETAIL_TABLE_NAME)
	_, err = database.Execute(query, memberDetail.Status, memberDetail.Role, memberDetail.Permissions, memberDetail.Plans, userId)
	if err != nil {
		http.Error(w, "Error updating into account details database", http.StatusInternalServerError)
		return
	}

	err = helpers.UpdateUserDetailTable(w, types.UserDetailReq{
		FirstName: memberDetail.FirstName,
		LastName:  memberDetail.LastName,
		Email:     memberDetail.Email,
		Phone:     memberDetail.Phone,
		Address:   memberDetail.Address,
		City:      memberDetail.City,
		State:     memberDetail.State,
		Zip:       memberDetail.Zip,
		Country:   memberDetail.Country,
		Dob:       memberDetail.Dob,
		Gender:    memberDetail.Gender,
		Metadata:  memberDetail.Metadata,
	}, userId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into user detail table", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"success": "true"})
}

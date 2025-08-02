package personalDetail

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anandvarma/namegen"
	"github.com/gorilla/mux"
	"github.com/sethvargo/go-password/password"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func RegisterUserDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyId"]

	if companyID == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var userReq types.UserDetailPostRequestBody
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := namegen.New().Get()
	passwd, err := password.Generate(20, 5, 5, false, false)
	if err != nil {
		log.Println("Error generating password: ", err)
		helpers.SendJSONError(w, "Error generating password", http.StatusUnauthorized)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (username, password) VALUES ($1, $2)`, database.IDENTITY_TABLE_NAME)
	row, tbErr := database.Query(query, username, passwd)
	if tbErr != nil {
		log.Println("Error inserting into identity database: ", tbErr.Error())
		helpers.SendJSONError(w, "Error inserting into identity database", http.StatusUnauthorized)
		return
	}
	defer row.Close()

	getIAMQuery := fmt.Sprintf(`SELECT * from %s WHERE username = $1`, database.IDENTITY_TABLE_NAME)
	rows, getIAMErr := database.Query(getIAMQuery, username)
	if getIAMErr != nil {
		helpers.SendJSONError(w, "Error fetching from identity database", http.StatusInternalServerError)
		return
	}

	var users []types.IAM
	for rows.Next() {
		var a types.IAM
		err := rows.Scan(&a.ID, &a.Username, &a.Password, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning row from identity database", http.StatusInternalServerError)
			return
		}
		users = append(users, a)
	}
	err = rows.Err()
	if err != nil {
		helpers.SendJSONError(w, "Error scanning rows from identity database", http.StatusInternalServerError)
		return
	}

	if len(users) != 1 {
		helpers.SendJSONError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// inserting into user detail
	err = helpers.InsertIntoUserDetailTable(w, types.UserDetailReq{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Email:     userReq.Email,
		Phone:     userReq.Phone,
		Address:   userReq.Address,
		City:      userReq.City,
		State:     userReq.State,
		Zip:       userReq.Zip,
		Country:   userReq.Country,
		Dob:       userReq.Dob,
		Gender:    userReq.Gender,
		Metadata:  userReq.Metadata,
	}, users[0].ID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into user detail table", http.StatusInternalServerError)
		return
	}
	// insert into account detail
	err = helpers.InsertAccountDetailsWithCompany(w, userReq.Role, userReq.Permissions, userReq.Plans, userReq.Status, companyID, users[0].ID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into account detail table", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyId"]

	if companyID == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var userReq types.UserSignUpRequestBody
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (username, password) VALUES ($1, $2)`, database.IDENTITY_TABLE_NAME)
	row, tbErr := database.Query(query, userReq.UserName, userReq.Password)
	if tbErr != nil {
		log.Println("Error inserting into identity database: ", tbErr.Error())
		helpers.SendJSONError(w, "Error inserting into identity database", http.StatusUnauthorized)
		return
	}
	defer row.Close()

	getIAMQuery := fmt.Sprintf(`SELECT * from %s WHERE username = $1`, database.IDENTITY_TABLE_NAME)
	rows, getIAMErr := database.Query(getIAMQuery, userReq.UserName)
	if getIAMErr != nil {
		helpers.SendJSONError(w, "Error fetching from identity database", http.StatusInternalServerError)
		return
	}

	var users []types.IAM
	for rows.Next() {
		var a types.IAM
		err := rows.Scan(&a.ID, &a.Username, &a.Password, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning row from identity database", http.StatusInternalServerError)
			return
		}
		users = append(users, a)
	}
	err := rows.Err()
	if err != nil {
		helpers.SendJSONError(w, "Error scanning rows from identity database", http.StatusInternalServerError)
		return
	}

	if len(users) != 1 {
		helpers.SendJSONError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// inserting into user detail
	err = helpers.InsertIntoUserDetailTable(w, types.UserDetailReq{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Email:     userReq.Email,
		Phone:     userReq.Phone,
		Address:   userReq.Address,
		City:      userReq.City,
		State:     userReq.State,
		Zip:       userReq.Zip,
		Country:   userReq.Country,
		Dob:       userReq.Dob,
		Gender:    userReq.Gender,
		Metadata:  userReq.Metadata,
	}, users[0].ID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into user detail table", http.StatusInternalServerError)
		return
	}
	// insert into account detail
	err = helpers.InsertAccountDetailsWithCompany(w, userReq.Role, userReq.Permissions, userReq.Plans, userReq.Status, companyID, users[0].ID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into account detail table", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

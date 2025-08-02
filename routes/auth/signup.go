package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/secrets"
	"server.simplifycontrol.com/types"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var signupReq types.SignupRequestBody
	err := json.NewDecoder(r.Body).Decode(&signupReq)
	if err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the email and password
	if signupReq.FirstName == "" || signupReq.LastName == "" || signupReq.Email == "" || signupReq.Password == "" || signupReq.Role == "" || signupReq.Contact == "" {
		helpers.SendJSONError(w, "All fields are required", http.StatusBadRequest)
		return
	}

	encryptPass, err := secrets.EncryptAESGCM(signupReq.Password)
	if err != nil {
		log.Println("Error encrypting password: ", err)
		helpers.SendJSONError(w, "Error encrypting password", http.StatusUnauthorized)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (username, password) VALUES ($1, $2)`, database.IDENTITY_TABLE_NAME)
	// Insert data into IAM
	row, tbErr := database.Query(query, signupReq.Username, encryptPass)
	if tbErr != nil {
		log.Println("Error inserting into identity database: ", tbErr.Error())
		helpers.SendJSONError(w, "Error inserting into identity database", http.StatusUnauthorized)
		return
	}
	defer row.Close()

	// Query the database again
	query = fmt.Sprintf(`SELECT * from %s WHERE username = $1`, database.IDENTITY_TABLE_NAME)
	rows, tbErr := database.Query(query, signupReq.Username)
	if tbErr != nil {
		log.Println("Error querying database: ", tbErr)
		helpers.SendJSONError(w, "Error querying database", http.StatusUnauthorized)
		return
	}
	defer rows.Close()

	var users []types.IAM
	for rows.Next() {
		var a types.IAM
		err := rows.Scan(&a.ID, &a.Username, &a.Password, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			log.Println("Error scanning rows from identity database: ", tbErr)
			helpers.SendJSONError(w, "Error querying database", http.StatusUnauthorized)
			return
		}
		users = append(users, a)
	}
	err = rows.Err()
	if err != nil {
		helpers.SendJSONError(w, "Error querying rows database", http.StatusUnauthorized)
		return
	}

	if len(users) != 1 {
		helpers.SendJSONError(w, "Something went wrong signing up user", http.StatusUnauthorized)
		return
	}

	helpers.InsertAccountDetails(w, signupReq.Role, signupReq.Permissions, signupReq.Plans, "", users[0].ID)

	err = helpers.InsertIntoUserDetailTable(w, types.UserDetailReq{
		FirstName: signupReq.FirstName,
		LastName:  signupReq.LastName,
		Email:     signupReq.Email,
		Phone:     signupReq.Contact,
		Address:   "",
		City:      "",
		State:     "",
		Zip:       "",
		Country:   "",
		Dob:       "",
		Gender:    "",
		Metadata:  "",
	}, users[0].ID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into user detail table", http.StatusInternalServerError)
		return
	}

	// Generate a JWT token
	token, err := helpers.GenerateJWT(helpers.Claims{
		ID:       users[0].ID,
		Username: signupReq.Username,
		Role:     signupReq.Role,
	})
	if err != nil {
		helpers.SendJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"token": token})
}

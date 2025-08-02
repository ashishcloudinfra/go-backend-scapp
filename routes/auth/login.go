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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var loginReq types.LoginRequestBody
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the email and password
	if loginReq.Username == "" || loginReq.Password == "" {
		helpers.SendJSONError(w, "username and password are required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * FROM %s WHERE username = $1`, database.IDENTITY_TABLE_NAME)
	rows, tbErr := database.Query(query, loginReq.Username)
	if tbErr != nil {
		log.Println("Error querying database: ", tbErr)
		helpers.SendJSONError(w, "Error fetching user from database", http.StatusUnauthorized)
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
		log.Println("Error scanning all rows from identity database: ", tbErr)
		helpers.SendJSONError(w, "Error querying database", http.StatusUnauthorized)
		return
	}

	if len(users) != 1 {
		helpers.SendJSONError(w, "Invalid email", http.StatusUnauthorized)
		return
	}

	decryptedPass, err := secrets.DecryptAESGCM(users[0].Password)
	if err != nil {
		helpers.SendJSONError(w, "Error decrypting password", http.StatusUnauthorized)
		return
	}

	if decryptedPass != loginReq.Password {
		helpers.SendJSONError(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	query = fmt.Sprintf(`SELECT * FROM %s WHERE iamId = $1`, database.ACCOUNT_DETAIL_TABLE_NAME)
	rows, tbErr = database.Query(query, users[0].ID)
	if tbErr != nil {
		log.Println("Error querying database: ", tbErr)
		helpers.SendJSONError(w, "Error fetching user from database", http.StatusUnauthorized)
		return
	}
	defer rows.Close()

	var ads []types.AccountDetail
	for rows.Next() {
		var a types.AccountDetail
		err := rows.Scan(&a.ID, &a.Role, &a.Permissions, &a.Plans, &a.Status, &a.CreatedAt, &a.UpdatedAt, &a.CompanyID, &a.IamID)
		if err != nil {
			log.Println("Error scanning account detail: ", err)
			helpers.SendJSONError(w, "Error scanning account detail:", http.StatusUnauthorized)
			return
		}
		ads = append(ads, a)
	}

	// Generate a JWT token
	token, err := helpers.GenerateJWT(helpers.Claims{
		ID:       users[0].ID,
		Username: users[0].Username,
		Role:     ads[0].Role,
	})
	if err != nil {
		helpers.SendJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"token": token})
}

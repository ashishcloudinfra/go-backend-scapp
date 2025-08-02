package company

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var companyReq types.PostCompanyRequestBody
	if err := json.NewDecoder(r.Body).Decode(&companyReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (name, email, contact, address, city, state, country, pincode, type, isDefault, iamId)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`, database.COMPANY_TABLE_NAME)
	_, err := database.Execute(query, companyReq.Name, companyReq.Email, companyReq.Contact, companyReq.Address, companyReq.City, companyReq.State, companyReq.Country, companyReq.Pincode, companyReq.Type, companyReq.IsDefault, companyReq.OwnerID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	query = fmt.Sprintf(`
		SELECT id from %s where iamId = $1 and name = $2
	`, database.COMPANY_TABLE_NAME)
	rows, err := database.Query(query, companyReq.OwnerID, companyReq.Name)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from company database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var companies []types.Company
	for rows.Next() {
		var company types.Company
		err := rows.Scan(&company.ID)
		if err != nil {
			log.Println("Error scanning data:", err)
			helpers.SendJSONError(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		companies = append(companies, company)
	}

	query = fmt.Sprintf(`
		UPDATE %s SET companyId = $1 WHERE iamId = $2
	`, database.ACCOUNT_DETAIL_TABLE_NAME)
	_, err = database.Query(query, companies[0].ID, companyReq.OwnerID)
	if err != nil {
		helpers.SendJSONError(w, "Error updating the account table", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

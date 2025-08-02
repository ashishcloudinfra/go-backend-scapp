package company

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

// GetCompanyDetailsHandler handles fetching companies by owner ID
func GetCompanyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * FROM %s WHERE iamId = $1`, database.COMPANY_TABLE_NAME)
	rows, err := database.Query(query, userId)
	if err != nil {
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var companies []types.Company
	for rows.Next() {
		var company types.Company
		err := rows.Scan(&company.ID, &company.Name, &company.Email, &company.Contact, &company.Address, &company.City, &company.State, &company.Country, &company.Pincode, &company.Type, &company.IsDefault, &company.IamId, &company.CreatedAt, &company.UpdatedAt)
		if err != nil {
			log.Println("Error scanning data:", err)
			helpers.SendJSONError(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		helpers.SendJSONError(w, "Error reading rows", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": companies})
}

func FetchDefaultCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * FROM %s WHERE iamId = $1 and isDefault = true`, database.COMPANY_TABLE_NAME)
	rows, err := database.Query(query, userId)
	if err != nil {
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var companies []types.Company
	for rows.Next() {
		var company types.Company
		err := rows.Scan(&company.ID, &company.Name, &company.Email, &company.Contact, &company.Address, &company.City, &company.State, &company.Country, &company.Pincode, &company.Type, &company.IsDefault, &company.IamId, &company.CreatedAt, &company.UpdatedAt)
		if err != nil {
			log.Println("Error scanning data:", err)
			helpers.SendJSONError(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		helpers.SendJSONError(w, "Error reading rows", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": companies[0]})
}

func FetchCompanyByCompanyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, database.COMPANY_TABLE_NAME)
	rows, err := database.Query(query, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var companies []types.Company
	for rows.Next() {
		var company types.Company
		err := rows.Scan(&company.ID, &company.Name, &company.Email, &company.Contact, &company.Address, &company.City, &company.State, &company.Country, &company.Pincode, &company.Type, &company.IsDefault, &company.IamId, &company.CreatedAt, &company.UpdatedAt)
		if err != nil {
			log.Println("Error scanning data:", err)
			helpers.SendJSONError(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		helpers.SendJSONError(w, "Error reading rows", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": companies[0]})
}

func FetchAssociatedCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`select * from %s where id = (select companyId from %s where iamid = $1)`, database.COMPANY_TABLE_NAME, database.ACCOUNT_DETAIL_TABLE_NAME)
	rows, err := database.Query(query, userId)
	if err != nil {
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var companies []types.Company
	for rows.Next() {
		var company types.Company
		err := rows.Scan(&company.ID, &company.Name, &company.Email, &company.Contact, &company.Address, &company.City, &company.State, &company.Country, &company.Pincode, &company.Type, &company.IsDefault, &company.IamId, &company.CreatedAt, &company.UpdatedAt)
		if err != nil {
			log.Println("Error scanning data:", err)
			helpers.SendJSONError(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		helpers.SendJSONError(w, "Error reading rows", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": companies[0]})
}

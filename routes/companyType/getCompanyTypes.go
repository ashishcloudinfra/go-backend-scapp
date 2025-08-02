package companyType

import (
	"fmt"
	"log"
	"net/http"

	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
)

func GetCompanyTypesHandler(w http.ResponseWriter, r *http.Request) {
	query := fmt.Sprintf(`SELECT name FROM %s`, database.COMPANY_TYPE_TABLE_NAME)
	rows, err := database.Query(query)
	if err != nil {
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var types []string
	for rows.Next() {
		var companyType string
		err := rows.Scan(&companyType)
		if err != nil {
			log.Println("Error scanning company type data: ", err)
			helpers.SendJSONError(w, "Error scanning company type data: ", http.StatusInternalServerError)
			return
		}
		types = append(types, companyType)
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": types})
}

package helpers

import (
	"fmt"
	"log"
	"net/http"

	"server.simplifycontrol.com/helpers/database"
)

func InsertAccountDetails(w http.ResponseWriter, role string, permissions string, plans string, status string, userId string) error {
	query := fmt.Sprintf(`INSERT INTO %s (role, permissions, plans, status, iamId) VALUES ($1, $2, $3, $4, $5)`, database.ACCOUNT_DETAIL_TABLE_NAME)
	_, err := database.Query(query, role, permissions, plans, status, userId)
	if err != nil {
		log.Println("Error inserting into account detail database: ", err.Error())
		return err
	}
	return nil
}

func InsertAccountDetailsWithCompany(w http.ResponseWriter, role string, permissions string, plans string, status string, companyId string, userId string) error {
	query := fmt.Sprintf(`INSERT INTO %s (role, permissions, plans, status, companyId, iamId) VALUES ($1, $2, $3, $4, $5, $6)`, database.ACCOUNT_DETAIL_TABLE_NAME)
	_, err := database.Query(query, role, permissions, plans, status, companyId, userId)
	if err != nil {
		log.Println("Error inserting into account detail database: ", err.Error())
		return err
	}
	return nil
}

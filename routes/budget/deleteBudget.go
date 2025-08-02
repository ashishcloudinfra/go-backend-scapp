package budget

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
)

func DeleteBudgetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"] // capturing categoryId from route

	if categoryId == "" {
		helpers.SendJSONError(w, "categoryId is required", http.StatusBadRequest)
		return
	}

	// Build the DELETE SQL statement
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, database.BUDGET_CATEGORY_TABLE_NAME)

	// Execute the DELETE query
	_, err := database.Execute(query, categoryId)
	if err != nil {
		log.Println("Error deleting category from database:", err)
		helpers.SendJSONError(w, "Error deleting category record", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func DeleteBudgetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"] // capturing itemId from route

	if itemId == "" {
		helpers.SendJSONError(w, "itemId is required", http.StatusBadRequest)
		return
	}

	// Build the DELETE SQL statement
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, database.BUDGET_ITEM_TABLE_NAME)

	// Execute the DELETE query
	_, err := database.Execute(query, itemId)
	if err != nil {
		log.Println("Error deleting from database:", err)
		helpers.SendJSONError(w, "Error deleting record", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

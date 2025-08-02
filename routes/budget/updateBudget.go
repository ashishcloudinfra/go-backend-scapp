package budget

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/secrets"
	"server.simplifycontrol.com/types"
)

func UpdateBudgetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"] // assuming you capture the item ID from the route

	if itemId == "" {
		helpers.SendJSONError(w, "itemId is required", http.StatusBadRequest)
		return
	}

	// Parse the request body into your BudgetItemReqBody struct
	var budgetItemReq types.BudgetItemReqBody
	if err := json.NewDecoder(r.Body).Decode(&budgetItemReq); err != nil {
		log.Println(err)
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`
		UPDATE %s
		SET
			categoryId = $1,
			itemName = $2,
			description = $3,
			allocatedAmount = $4,
			actualAmount = $5,
			currencyCode = $6,
			status = $7,
			month = $8,
			year = $9
		WHERE
			id = $10
	`, database.BUDGET_ITEM_TABLE_NAME)

	encyptedAllocatedAmount, err := secrets.EncryptAESGCM(helpers.FloatToString(budgetItemReq.AllocatedAmount))
	if err != nil {
		log.Println("Error encrypting allocated amount: ", err)
		helpers.SendJSONError(w, "Error encrypting allocated amount: ", http.StatusInternalServerError)
		return
	}
	encyptedActualAmount, err := secrets.EncryptAESGCM(helpers.FloatToString(budgetItemReq.ActualAmount))
	if err != nil {
		log.Println("Error encrypting actual amount: ", err)
		helpers.SendJSONError(w, "Error encrypting actual amount: ", http.StatusInternalServerError)
		return
	}

	_, err = database.Execute(
		query,
		budgetItemReq.CategoryID,
		budgetItemReq.ItemName,
		budgetItemReq.Description,
		encyptedAllocatedAmount,
		encyptedActualAmount,
		budgetItemReq.CurrencyCode,
		budgetItemReq.Status,
		budgetItemReq.Month,
		budgetItemReq.Year,
		itemId,
	)
	if err != nil {
		log.Println("Error updating database:", err)
		helpers.SendJSONError(w, "Error updating database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func UpdateBudgetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"] // capturing categoryId from route

	if categoryId == "" {
		helpers.SendJSONError(w, "categoryId is required", http.StatusBadRequest)
		return
	}

	var budgetCategoryReq types.BudgetCategoryReqBody
	if err := json.NewDecoder(r.Body).Decode(&budgetCategoryReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Handle possible null parentId
	var parentId interface{}
	if budgetCategoryReq.ParentId == "" {
		parentId = nil
	} else {
		parentId = budgetCategoryReq.ParentId
	}

	// Build the UPDATE query
	query := fmt.Sprintf(
		`UPDATE %s
			 SET categoryName = $1,
				categoryDescription = $2,
				parentId = $3,
				categoryTypeId = $4,
				month = $5,
				year = $6
			 WHERE id = $7`,
		database.BUDGET_CATEGORY_TABLE_NAME,
	)

	// Execute the UPDATE query
	_, err := database.Execute(
		query,
		budgetCategoryReq.CategoryName,
		budgetCategoryReq.CategoryDescription,
		parentId,
		budgetCategoryReq.CategoryTypeId,
		budgetCategoryReq.Month,
		budgetCategoryReq.Year,
		categoryId,
	)
	if err != nil {
		fmt.Println(err)
		helpers.SendJSONError(w, "Error updating category from database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

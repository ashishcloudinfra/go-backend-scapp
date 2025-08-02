package budget

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func AddBudgetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}

	var budgetItemReq types.BudgetItemReqBody
	if err := json.NewDecoder(r.Body).Decode(&budgetItemReq); err != nil {
		fmt.Println(err)
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := helpers.InsertBudgetItem(budgetItemReq, userId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func AddBudgetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}

	var budgetCategoryReq types.BudgetCategoryReqBody
	if err := json.NewDecoder(r.Body).Decode(&budgetCategoryReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := helpers.InsertBudgetCategory(budgetCategoryReq, userId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func AddBudgetCategoryType(w http.ResponseWriter, r *http.Request) {
	var budgetCategoryReq types.BudgetCategoryTypeReqBody
	if err := json.NewDecoder(r.Body).Decode(&budgetCategoryReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (type, bgColor, textColor) VALUES ($1, $2, $3)`, database.BUDGET_CATEGORY_TYPE_TABLE_NAME)
	_, err := database.Execute(query, budgetCategoryReq.Type, budgetCategoryReq.BgColor, budgetCategoryReq.TextColor)
	if err != nil {
		fmt.Println(err)
		helpers.SendJSONError(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func CopyBudgetOfGivenMonthAndYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId and month and year is required", http.StatusBadRequest)
		return
	}

	var copyBudgetReq types.CopyBudgetReqBody
	if err := json.NewDecoder(r.Body).Decode(&copyBudgetReq); err != nil {
		fmt.Println(err)
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := helpers.DeleteBudgetCategoriesOfGivenMonthAndYear(strconv.Itoa(copyBudgetReq.CurrentMonth), strconv.Itoa(copyBudgetReq.CurrentYear))
	if err != nil {
		helpers.SendJSONError(w, "Error deleting budget categories", http.StatusInternalServerError)
		return
	}

	err = helpers.CopyBudgetCategoriesAndItems(userId, copyBudgetReq)
	if err != nil {
		helpers.SendJSONError(w, "Error copying budget categories and items", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

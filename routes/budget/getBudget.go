package budget

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/types"
)

func GetBudgetItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	month := vars["month"]
	year := vars["year"]

	if userId == "" || month == "" || year == "" {
		helpers.SendJSONError(w, "userId and month and year is required", http.StatusBadRequest)
		return
	}

	items, err := helpers.GetBudgetItems(userId, month, year)
	if err != nil {
		helpers.SendJSONError(w, "Error retrieving budget items", http.StatusInternalServerError)
		return
	}

	// Return the result as JSON
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": items})
}

func GetBudgetCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	month := vars["month"]
	year := vars["year"]

	if userId == "" || month == "" || year == "" {
		helpers.SendJSONError(w, "userId and month and year is required", http.StatusBadRequest)
		return
	}

	categories, err := helpers.GetBudgetCategories(userId, month, year)
	if err != nil {
		helpers.SendJSONError(w, "Error querying budget categories database", http.StatusInternalServerError)
		return
	}

	if categories != nil {
		helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": categories})
		return
	}

	reqYear, err := strconv.Atoi(year)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		return
	}

	reqMonth, err := strconv.Atoi(month)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		return
	}

	var prevMonth int
	var prevYear int
	if reqMonth == 0 {
		prevMonth = 11
		prevYear = reqYear - 1
	} else {
		prevMonth = reqMonth - 1
		prevYear = reqYear
	}

	prevMonthCategories, err := helpers.GetBudgetCategories(userId, strconv.Itoa(prevMonth), strconv.Itoa(prevYear))
	if err != nil {
		helpers.SendJSONError(w, "Error querying previous budget categories database", http.StatusInternalServerError)
		return
	}

	if prevMonthCategories != nil {
		err := helpers.CopyBudgetCategoriesAndItems(userId, types.CopyBudgetReqBody{
			OldMonth:     prevMonth,
			OldYear:      prevYear,
			CurrentMonth: reqMonth,
			CurrentYear:  reqYear,
		})
		if err != nil {
			helpers.SendJSONError(w, "Error copying budget categories and items", http.StatusInternalServerError)
			return
		}
		categories, err = helpers.GetBudgetCategories(userId, month, year)
		if err != nil {
			helpers.SendJSONError(w, "Error querying current month categories database", http.StatusInternalServerError)
			return
		}
		helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": categories})
		return
	}

	categoryTypes, err := helpers.GetBudgetCategoryTypes()
	if err != nil {
		log.Printf("GetBudgetCategoryTypes error: %v", err)
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}

	if categoryTypes == nil {
		helpers.SendJSONError(w, "No data in category type database", http.StatusInternalServerError)
		return
	}

	// loop in category types
	for _, categoryType := range categoryTypes {
		_, err := helpers.InsertBudgetCategory(types.BudgetCategoryReqBody{
			CategoryName:        categoryType.Type + "s",
			CategoryDescription: "",
			Month:               reqMonth,
			Year:                reqYear,
			ParentId:            "",
			CategoryTypeId:      categoryType.Id,
		}, userId)
		if err != nil {
			helpers.SendJSONError(w, "Error inserting budget category into database", http.StatusInternalServerError)
			return
		}
	}

	categories, err = helpers.GetBudgetCategories(userId, month, year)
	if err != nil {
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": categories})
}

func GetBudgetCategoryTypes(w http.ResponseWriter, r *http.Request) {
	categoryTypes, err := helpers.GetBudgetCategoryTypes()
	if err != nil {
		log.Printf("GetBudgetCategoryTypes error: %v", err)
		helpers.SendJSONError(w, "Error querying database", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": categoryTypes})
}

func GetBudgetStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId and month and year is required", http.StatusBadRequest)
		return
	}

	stats, err := helpers.GetBudgetStatsByMonthAndYear(userId)
	if err != nil {
		log.Printf("GetBudgetCategoryTypes error: %v", err)
		helpers.SendJSONError(w, "Error fetching stats", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": stats})
}

func GetRawStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId and year is required", http.StatusBadRequest)
		return
	}

	stats, err := helpers.GetRawUserStats(userId)
	if err != nil {
		log.Printf("GetBudgetCategoryTypes error: %v", err)
		helpers.SendJSONError(w, "Error fetching stats", http.StatusInternalServerError)
		return
	}

	// Return the result
	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": stats})
}

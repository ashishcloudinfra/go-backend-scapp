package menuItem

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	menuItemId := vars["menuItemId"]

	if companyId == "" || menuItemId == "" {
		helpers.SendJSONError(w, "companyId and menuItemId is required", http.StatusBadRequest)
		return
	}

	var menuItemReq types.MenuItemRequestBody
	if err := json.NewDecoder(r.Body).Decode(&menuItemReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT id from %s where name = $1 and companyId = $2`, database.MENU_ITEM_CATEGORY_TABLE_NAME)
	rows, err := database.Query(query, menuItemReq.Category, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error scanning into menu item category database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categoryIds []string
	for rows.Next() {
		var categoryId string
		err := rows.Scan(&categoryId)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning menu item category", http.StatusBadRequest)
			return
		}
		categoryIds = append(categoryIds, categoryId)
	}

	var categoryId string
	if len(categoryIds) == 0 {
		query = fmt.Sprintf(`INSERT INTO %s (name, companyId) VALUES ($1, $2) RETURNING id`, database.MENU_ITEM_CATEGORY_TABLE_NAME)
		rows, err := database.Query(query, menuItemReq.Category, companyId)
		if err != nil {
			helpers.SendJSONError(w, "Error inserting into menu item category database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&categoryId); err != nil {
				helpers.SendJSONError(w, "Error scanning into orders database", http.StatusInternalServerError)
				return
			}
		}
	} else {
		categoryId = categoryIds[0]
	}

	query = fmt.Sprintf(`UPDATE %s SET name = $1, description = $2, cookingTime = $3, photo = $4, isVeg = $5, categoryId = $6 WHERE id = $7`, database.MENU_ITEM_TABLE_NAME)
	rows, err = database.Query(query, menuItemReq.Name, menuItemReq.Description, menuItemReq.CookingTime, menuItemReq.Photo, menuItemReq.IsVeg, categoryId, menuItemId)
	if err != nil {
		helpers.SendJSONError(w, "Error updating menu item database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Insert into menu pricing
	for _, value := range menuItemReq.Varieties {
		query = fmt.Sprintf(`UPDATE %s SET varietyType = $1, price = $2 WHERE id = $3`, database.MENU_ITEM_PRICING_TABLE_NAME)
		_, err = database.Execute(query, value.Name, value.Price, value.Id)
		if err != nil {
			helpers.SendJSONError(w, "Error updating menu item pricing database", http.StatusInternalServerError)
			return
		}
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

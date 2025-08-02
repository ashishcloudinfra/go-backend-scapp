package menuItem

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func GetAllMenuItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT
				mi.id AS menuItemId,
				mi.name AS menuItemName,
				mi.description AS menuItemDescription,
				mi.cookingTime AS cookingTime,
				mi.photo AS menuItemPhoto,
				mi.isVeg AS isVeg,
			  mic.name AS categoryName,
				mip.id AS pricingId,
				mip.varietyType AS varietyType,
				mip.price AS price
		FROM %s mic
		JOIN %s mi
		ON mic.id = mi.categoryId
		JOIN %s mip
		ON mi.id = mip.menuItemId
		WHERE mi.companyid = $1`, database.MENU_ITEM_CATEGORY_TABLE_NAME, database.MENU_ITEM_TABLE_NAME, database.MENU_ITEM_PRICING_TABLE_NAME)
	rows, err := database.Query(query, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from menu item database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var menuItems []types.MenuItem
	for rows.Next() {
		var menuItem types.MenuItem
		err := rows.Scan(
			&menuItem.MenuItemID,
			&menuItem.MenuItemName,
			&menuItem.MenuItemDescription,
			&menuItem.CookingTime,
			&menuItem.MenuItemPhoto,
			&menuItem.IsVeg,
			&menuItem.CategoryName,
			&menuItem.PricingID,
			&menuItem.VarietyType,
			&menuItem.Price,
		)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning menu items", http.StatusBadRequest)
			return
		}
		menuItems = append(menuItems, menuItem)
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": menuItems})
}

func GetMenuItemWithId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	menuItemId := vars["menuItemId"]

	if companyId == "" || menuItemId == "" {
		helpers.SendJSONError(w, "companyId and menuItemId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT
				mi.id AS menuItemId,
				mi.name AS menuItemName,
				mi.description AS menuItemDescription,
				mi.cookingTime AS cookingTime,
				mi.photo AS menuItemPhoto,
				mi.isVeg AS isVeg,
			  mic.name AS categoryName,
				mip.id AS pricingId,
				mip.varietyType AS varietyType,
				mip.price AS price
		FROM %s mic
		JOIN %s mi
		ON mic.id = mi.categoryId
		JOIN %s mip
		ON mi.id = mip.menuItemId
		WHERE mi.companyid = $1 and mi.id = $2`, database.MENU_ITEM_CATEGORY_TABLE_NAME, database.MENU_ITEM_TABLE_NAME, database.MENU_ITEM_PRICING_TABLE_NAME)
	rows, err := database.Query(query, companyId, menuItemId)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from menu item database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var menuItems []types.MenuItem
	for rows.Next() {
		var menuItem types.MenuItem
		err := rows.Scan(
			&menuItem.MenuItemID,
			&menuItem.MenuItemName,
			&menuItem.MenuItemDescription,
			&menuItem.CookingTime,
			&menuItem.MenuItemPhoto,
			&menuItem.IsVeg,
			&menuItem.CategoryName,
			&menuItem.PricingID,
			&menuItem.VarietyType,
			&menuItem.Price,
		)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning menu items", http.StatusBadRequest)
			return
		}
		menuItems = append(menuItems, menuItem)
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": menuItems})
}

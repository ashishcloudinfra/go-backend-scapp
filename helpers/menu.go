package helpers

import (
	"fmt"

	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func InsertMenuItem(companyId string, menuItemReq types.MenuItemRequestBody) error {
	query := fmt.Sprintf(`SELECT id from %s where name = $1 and companyId = $2`, database.MENU_ITEM_CATEGORY_TABLE_NAME)
	rows, err := database.Query(query, menuItemReq.Category, companyId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var categoryIds []string
	for rows.Next() {
		var categoryId string
		err := rows.Scan(&categoryId)
		if err != nil {
			return err
		}
		categoryIds = append(categoryIds, categoryId)
	}

	var categoryId string
	if len(categoryIds) == 0 {
		query = fmt.Sprintf(`INSERT INTO %s (name, companyId) VALUES ($1, $2) RETURNING id`, database.MENU_ITEM_CATEGORY_TABLE_NAME)
		rows, err := database.Query(query, menuItemReq.Category, companyId)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&categoryId); err != nil {
				return err
			}
		}
	} else {
		categoryId = categoryIds[0]
	}

	query = fmt.Sprintf(`INSERT INTO %s (name, description, cookingTime, photo, isVeg, categoryId, companyId) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, database.MENU_ITEM_TABLE_NAME)
	rows, err = database.Query(query, menuItemReq.Name, menuItemReq.Description, menuItemReq.CookingTime, menuItemReq.Photo, menuItemReq.IsVeg, categoryId, companyId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var menuItemId string
	for rows.Next() {
		if err := rows.Scan(&menuItemId); err != nil {
			return err
		}
	}

	// Insert into menu pricing
	for _, value := range menuItemReq.Varieties {
		query = fmt.Sprintf(`INSERT INTO %s (varietyType, price, menuItemId) VALUES ($1, $2, $3)`, database.MENU_ITEM_PRICING_TABLE_NAME)
		_, err = database.Execute(query, value.Name, value.Price, menuItemId)
		if err != nil {
			return err
		}
	}

	return nil
}

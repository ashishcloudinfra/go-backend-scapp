package orders

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func GetOrderByTableNumberAndCompanyId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	tableNumber := vars["tableNumber"]

	if companyId == "" || tableNumber == "" {
		helpers.SendJSONError(w, "companyId and tableNumber is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT 
			mi.name AS menuitemName,
			mip.varietyType,
			mip.price,
			mi.photo,
			mi.isVeg,
			SUM(oi.quantity) AS quantity
	FROM %s oi
	JOIN %s mip ON oi.menuItemPricingId = mip.id
	JOIN %s mi ON mip.menuItemId = mi.id
	JOIN %s o ON oi.orderId = o.id
	WHERE o.tableNumber = $1
	AND o.companyId = $2
	GROUP BY mi.name, mip.varietyType, mip.price, mi.photo, mi.isVeg;`, database.ORDER_ITEM_TABLE_NAME, database.MENU_ITEM_PRICING_TABLE_NAME, database.MENU_ITEM_TABLE_NAME, database.ORDERS_TABLE_NAME)
	rows, err := database.Query(query, tableNumber, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into orders database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var menuItems []types.MenuWithQuantity
	for rows.Next() {
		var item types.MenuWithQuantity
		err := rows.Scan(
			&item.MenuitemName,
			&item.VarietyType,
			&item.Price,
			&item.Photo,
			&item.IsVeg,
			&item.Quantity,
		)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning menu items", http.StatusBadRequest)
			return
		}
		menuItems = append(menuItems, item)
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": menuItems})
}

func GetOrdersListForAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId and tableNumber is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`SELECT 
			o.id AS orderId,
			oi.id AS orderItemId,
			o.orderNumber,
			o.status,
			mi.name AS menuitemName,
			mip.varietyType,
			mi.isVeg,
			oi.quantity,
			o.tableNumber
	FROM %s oi
	JOIN %s o ON oi.orderId = o.id
	JOIN %s mip ON oi.menuItemPricingId = mip.id
	JOIN %s mi ON mip.menuItemId = mi.id
	WHERE o.companyId = $1`, database.ORDER_ITEM_TABLE_NAME, database.ORDERS_TABLE_NAME, database.MENU_ITEM_PRICING_TABLE_NAME, database.MENU_ITEM_TABLE_NAME)
	rows, err := database.Query(query, companyId)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into orders database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []types.AdminOrder
	for rows.Next() {
		var order types.AdminOrder
		err := rows.Scan(
			&order.OrderId,
			&order.OrderItemId,
			&order.OrderNumber,
			&order.Status,
			&order.MenuitemName,
			&order.VarietyType,
			&order.IsVeg,
			&order.Quantity,
			&order.TableNumber,
		)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning order items", http.StatusBadRequest)
			return
		}
		orders = append(orders, order)
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": orders})
}

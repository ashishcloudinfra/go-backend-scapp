package orders

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/types"
)

func RegisterOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq types.OrderRequestBody
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`INSERT INTO %s (tableNumber, isDineIn, status, phone, email, companyId) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`, database.ORDERS_TABLE_NAME)
	rows, err := database.Query(query, orderReq.TableNumber, orderReq.IsDineIn, orderReq.Status, orderReq.Phone, orderReq.Email, orderReq.CompanyID)
	if err != nil {
		helpers.SendJSONError(w, "Error inserting into orders database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var lastInsertID string
	for rows.Next() {
		if err := rows.Scan(&lastInsertID); err != nil {
			helpers.SendJSONError(w, "Error scanning into orders database", http.StatusInternalServerError)
			return
		}
	}
	if err := rows.Err(); err != nil {
		helpers.SendJSONError(w, "Error scanning into orders rows database", http.StatusInternalServerError)
	}

	for index, value := range orderReq.Items {
		fmt.Println(index, value)
		query = fmt.Sprintf(`INSERT INTO %s (quantity, menuItemPricingId, orderId) VALUES ($1, $2, $3)`, database.ORDER_ITEM_TABLE_NAME)
		_, err = database.Execute(query, value.Quantity, value.MenuItemPricingId, lastInsertID)
		if err != nil {
			helpers.SendJSONError(w, "Error inserting into order item database", http.StatusInternalServerError)
			return
		}
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"orderId": lastInsertID})
}

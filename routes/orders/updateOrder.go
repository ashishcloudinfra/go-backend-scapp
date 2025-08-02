package orders

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	tableNumber := vars["tableNumber"]

	if companyId == "" || tableNumber == "" {
		helpers.SendJSONError(w, "companyId and tableNumber is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`select id from %s where companyid = $1 and tablenumber = $2 and status = 'inprogress'`, database.ORDERS_TABLE_NAME)
	rows, err := database.Query(query, companyId, tableNumber)
	if err != nil {
		helpers.SendJSONError(w, "Error fetching from orders database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orderIds []string
	for rows.Next() {
		var orderId string
		err := rows.Scan(&orderId)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning order items", http.StatusBadRequest)
			return
		}
		orderIds = append(orderIds, orderId)
	}

	for _, orderId := range orderIds {
		query := fmt.Sprintf(`update %s set status = 'completed' where id = $1`, database.ORDERS_TABLE_NAME)
		_, err := database.Query(query, orderId)
		if err != nil {
			helpers.SendJSONError(w, "Error updating order status", http.StatusInternalServerError)
			return
		}
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	status := vars["status"]

	if orderId == "" {
		helpers.SendJSONError(w, "orderId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`UPDATE %s SET status = $1 WHERE id = $2`, database.ORDERS_TABLE_NAME)
	_, err := database.Query(query, status, orderId)
	if err != nil {
		helpers.SendJSONError(w, "Error updating order status", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

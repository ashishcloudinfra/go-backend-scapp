package menuItem

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
)

func DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]
	menuItemId := vars["menuItemId"]

	if companyId == "" || menuItemId == "" {
		helpers.SendJSONError(w, "companyId and menuItemId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`DELETE from %s where companyId = $1 and id = $2`, database.MENU_ITEM_TABLE_NAME)
	rows, err := database.Query(query, companyId, menuItemId)
	if err != nil {
		helpers.SendJSONError(w, "Error delete from menu item database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

package menuItem

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/types"
)

func RegisterMenuItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var menuItemReq types.MenuItemRequestBody
	if err := json.NewDecoder(r.Body).Decode(&menuItemReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := helpers.InsertMenuItem(companyId, menuItemReq)
	if err != nil {
		fmt.Println("Error inserting menu item:", err)
		helpers.SendJSONError(w, "Error inserting menu item", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

func RegisterBulkMenuItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	if companyId == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	var menuItemsReq []types.MenuItemRequestBody
	if err := json.NewDecoder(r.Body).Decode(&menuItemsReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for _, menuItem := range menuItemsReq {
		err := helpers.InsertMenuItem(companyId, menuItem)
		if err != nil {
			fmt.Println("Error inserting menu item:", err)
			helpers.SendJSONError(w, "Error inserting menu item", http.StatusInternalServerError)
			return
		}
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

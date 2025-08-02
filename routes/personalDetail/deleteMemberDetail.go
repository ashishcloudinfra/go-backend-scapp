package personalDetail

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
)

func DeleteUserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userDetailId := vars["userDetailId"]

	if userDetailId == "" {
		helpers.SendJSONError(w, "userDetailId are required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`DELETE from %s WHERE id = $2`, database.USER_DETAIL_TABLE_NAME)
	_, err := database.Execute(query, userDetailId)
	if err != nil {
		helpers.SendJSONError(w, "Error delete member detail", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"success": "true"})
}

package account

import (
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
)

func GetAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "userId is required", http.StatusBadRequest)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": ""})
}

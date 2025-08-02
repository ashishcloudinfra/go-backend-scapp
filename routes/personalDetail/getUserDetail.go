package personalDetail

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/helpers/database"
	"server.simplifycontrol.com/secrets"
	"server.simplifycontrol.com/types"
)

func GetAllUsersByCompanyIDAndRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyId"]
	role := vars["role"]

	if companyID == "" {
		helpers.SendJSONError(w, "companyId is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`select id, firstname, lastname, gender, status, phone, email, pd.iamId from %s as pd
		join (select iamId, status from %s
		where companyId = $1 and role = $2) as tmp
		on tmp.iamId = pd.iamId`, database.PERSONAL_DETAIL_TABLE_NAME, database.ACCOUNT_DETAIL_TABLE_NAME)

	rows, err := database.Query(query, companyID, role)
	if err != nil {
		log.Println("Error fetching form DB", err)
		helpers.SendJSONError(w, "Error fetching data from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allMemberDetail []types.MemberShortDetail
	for rows.Next() {
		var md types.MemberShortDetail
		err := rows.Scan(&md.ID, &md.FirstName, &md.LastName, &md.Gender, &md.Status, &md.Phone, &md.Email, &md.IamID)
		if err != nil {
			helpers.SendJSONError(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		allMemberDetail = append(allMemberDetail, md)
	}

	if err := rows.Err(); err != nil {
		helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": allMemberDetail})
}

func GetUserDetailByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	if userId == "" {
		helpers.SendJSONError(w, "UserID is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`select pd.*, ad.status, ad.role, ad.permissions, ad.plans from %s as pd
		join %s as ad
		on pd.iamId = ad.iamId
		where pd.iamId = $1`, database.PERSONAL_DETAIL_TABLE_NAME, database.ACCOUNT_DETAIL_TABLE_NAME)
	rows, err := database.Query(query, userId)
	if err != nil {
		log.Println("Error fetching from DB", err)
		helpers.SendJSONError(w, "Error fetching data from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var personalDetails []types.PersonalDetailWithStatus
	for rows.Next() {
		var md types.PersonalDetailWithStatus
		var firstName, lastName, email, phone, address, city, state, zip, country, dob, gender, metadata string
		err := rows.Scan(&md.ID, &firstName, &lastName, &email, &phone, &address, &city, &state, &zip, &country, &dob, &gender, &metadata, &md.CreatedAt, &md.UpdatedAt, &md.IamID, &md.Status, &md.Role, &md.Permissions, &md.Plans)
		if err != nil {
			log.Println("Error scanning data", err)
			helpers.SendJSONError(w, "Error scanning data", http.StatusInternalServerError)
			return
		}

		md.FirstName, err = secrets.DecryptAESGCM(firstName)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.LastName, err = secrets.DecryptAESGCM(lastName)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Email, err = secrets.DecryptAESGCM(email)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Phone, err = secrets.DecryptAESGCM(phone)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Address, err = secrets.DecryptAESGCM(address)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.City, err = secrets.DecryptAESGCM(city)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.State, err = secrets.DecryptAESGCM(state)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Zip, err = secrets.DecryptAESGCM(zip)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Country, err = secrets.DecryptAESGCM(country)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Dob, err = secrets.DecryptAESGCM(dob)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Gender, err = secrets.DecryptAESGCM(gender)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		md.Metadata, err = secrets.DecryptAESGCM(metadata)
		if err != nil {
			helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
			return
		}

		personalDetails = append(personalDetails, md)
	}

	if err := rows.Err(); err != nil {
		helpers.SendJSONError(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": personalDetails[0]})
}

package notification

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"server.simplifycontrol.com/helpers"
	"server.simplifycontrol.com/secrets"
	"server.simplifycontrol.com/types"
)

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming JSON request
	var reqData types.RequestPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	name := reqData.Name
	email := reqData.Email
	phone := reqData.Phone
	message := reqData.Message

	mailjetClient := mailjet.NewMailjetClient(secrets.SecretJSON.Mailjet.PublicApiKey, secrets.SecretJSON.Mailjet.PrivateApiKey)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "shivamkumarrajput56789@gmail.com",
				Name:  "Shivam the great",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  name,
				},
			},
			Subject:  "Your email flight plan!",
			TextPart: "Dear passenger 1, welcome to Mailjet! May the delivery force be with you!",
			HTMLPart: fmt.Sprintf(`<h3>Dear passenger 1, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!</h3><br />May the delivery force be with you! %s and %s`, phone, message),
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Println(err)
		helpers.SendJSONError(w, "Error sending mail", http.StatusBadRequest)
		return
	}
	fmt.Printf("Data: %+v\n", res)
}

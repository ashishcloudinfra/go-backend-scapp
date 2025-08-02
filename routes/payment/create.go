package payment

import (
	"encoding/json"
	"log"
	"net/http"

	"server.simplifycontrol.com/helpers"
	paymentHelpers "server.simplifycontrol.com/payment"
	"server.simplifycontrol.com/types"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var paymentReq types.PaymentRequestBody
	if err := json.NewDecoder(r.Body).Decode(&paymentReq); err != nil {
		helpers.SendJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if paymentReq.Amount < 100 {
		helpers.SendJSONError(w, "Amount must be at least ₹1 (100 paise)", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"amount":          paymentReq.Amount,
		"currency":        paymentReq.Currency,
		"receipt":         paymentReq.Receipt,
		"payment_capture": 1,
	}

	order, err := paymentHelpers.RazorpayClient.Order.Create(data, nil)
	if err != nil {
		log.Println(err)
		helpers.SendJSONError(w, "Error creating payment order", http.StatusInternalServerError)
		return
	}

	helpers.SendJSONSuccessResponse(w, map[string]interface{}{"data": order})
}

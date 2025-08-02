package types

type PaymentRequestBody struct {
	Amount   int    `json:"amount"`   // in paise
	Currency string `json:"currency"` // "INR"
	Receipt  string `json:"receipt"`  // any string ID
}

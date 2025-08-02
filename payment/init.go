package paymentHelpers

import (
	"github.com/razorpay/razorpay-go"
	"server.simplifycontrol.com/secrets"
)

var RazorpayClient *razorpay.Client

func InitPayment() {
	RazorpayClient = razorpay.NewClient(secrets.SecretJSON.Razorpay.KeyId, secrets.SecretJSON.Razorpay.KeySecret)
}

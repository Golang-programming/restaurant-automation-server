package service

import (
	"fmt"

	billService "github.co/golang-programming/restaurant/api/bill/service"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
)

// var stripeProvider payment.PaymentProvider

func CreatePaymentIntent(customerID uint) (string, error) {

	bill, err := billService.GetBillByTableID(customerID)
	if err != nil {
		return "", err
	}

	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(int64(bill.Total * 100)),
		Description: stripe.String("Payment for Order #12345"),
		Currency:    stripe.String(string(stripe.CurrencyPKR)),
	}

	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}

	fmt.Println("paymentIntent", paymentIntent.ClientSecret)

	return paymentIntent.ClientSecret, nil
}

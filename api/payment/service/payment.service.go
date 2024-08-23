package service

import (
	"fmt"
	"log"

	pp "github.com/golang-programming/pp"
)

var stripeProvider pp.PaymentProvider

func init() {
	fmt.Println("================================")

	// Get the payment provider (returns an interface, not a pointer to an interface)
	var err error
	stripeProvider, err = pp.GetPaymentProvider("stripe")
	if err != nil {
		log.Fatalf("Error getting payment provider: %v", err)
	}

	err = stripeProvider.Initialize("sk_test_51Pqfl5AJqNngFtp6jgEGMLENxAFOi9KslDus5drjHxqUjNrZUXAWnVsVFAuc9TL1UIx2oYQpYmlLl9XP2s3orsWE00kK1t8y2p", nil)
	if err != nil {
		log.Fatalf("Error initializing provider: %v", err)
	}
}

func Charge() {
	transactionID, err := stripeProvider.Charge(100.0, "USD", "source-id")
	if err != nil {
		log.Fatalf("Error charging payment: %v", err)
	}

	fmt.Println("transactionID", transactionID)
}

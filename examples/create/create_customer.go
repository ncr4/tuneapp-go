package main

import (
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	apiEmail := os.Getenv("API_EMAIL")
	apiKey := os.Getenv("API_KEY")

	// Create a customer passing in all required params
	customer := tuneuptechnology.CreateCustomer(
		&tuneuptechnology.Customer{
			Auth:       apiEmail,
			APIKey:     apiKey,
			Firstname:  "Jake",
			Lastname:   "Peralta",
			Email:      "jake@example.com",
			Phone:      "8015551234",
			UserID:     1,
			Notes:      "Believes he is a good detective.",
			LocationID: 1,
		},
	)

	tuneuptechnology.PrettyPrint(customer)
}

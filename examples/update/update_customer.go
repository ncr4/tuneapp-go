package main

import (
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	apiEmail := os.Getenv("API_EMAIL")
	apiKey := os.Getenv("API_KEY")

	// Update a customer passing in all the params you want to update
	customer := tuneuptechnology.UpdateCustomer(
		&tuneuptechnology.Customer{
			Auth:       apiEmail,
			APIKey:     apiKey,
			ID:         1, // the ID of the customer you are updating
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

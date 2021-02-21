package main

import (
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	apiEmail := os.Getenv("API_EMAIL")
	apiKey := os.Getenv("API_KEY")

	// Delete a customer
	customer := tuneuptechnology.DeleteCustomer(
		&tuneuptechnology.Customer{
			Auth:   apiEmail,
			APIKey: apiKey,
			ID:     1, // the ID of the customer you are deleting
		},
	)

	tuneuptechnology.PrettyPrint(customer)
}

package main

import (
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	apiEmail := os.Getenv("API_EMAIL")
	apiKey := os.Getenv("API_KEY")

	// Retrieve all customers
	customers := tuneuptechnology.AllCustomers(
		&tuneuptechnology.Customer{
			Auth:   apiEmail,
			APIKey: apiKey,
		},
	)

	tuneuptechnology.PrettyPrint(customers)
}

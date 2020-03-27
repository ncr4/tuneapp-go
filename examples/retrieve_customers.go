package main

import (
	"os"
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	api_email := os.Getenv("API_EMAIL")
	api_key := os.Getenv("API_KEY")

	// Retrieve all customers
	customers := tuneuptechnology.AllCustomers(
		&tuneuptechnology.Customer{
			Auth: api_email,
			APIKey: api_key,
		},
	)

	tuneuptechnology.PrettyPrint(customers)
}

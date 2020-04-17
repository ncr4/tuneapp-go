package main

import (
	"os"
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	api_email := os.Getenv("API_EMAIL")
	api_key := os.Getenv("API_KEY")

	// Delete a customer
	customer := tuneuptechnology.DeleteCustomer(
		&tuneuptechnology.Customer{
			Auth: api_email,
			APIKey: api_key,
			Id: 1, // the ID of the customer you are deleting
		},
	)

	tuneuptechnology.PrettyPrint(customer)
}

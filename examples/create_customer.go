package main

import (
	"os"
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	api_email := os.Getenv("API_EMAIL")
	api_key := os.Getenv("API_KEY")

	// Create a customer passing in all required params
	customer := tuneuptechnology.CreateCustomer(
		&tuneuptechnology.Customer{
			Auth: api_email,
			APIKey: api_key,
			Firstname: "Jake",
			Lastname: "Peralta",
			Email: "jake@example.com",
			Phone: "8015551234",
			UserId: 1,
			Notes: "Believes he is a good detective.",
			LocationId: 1,
		},
	)

	tuneuptechnology.PrettyPrint(customer)
}

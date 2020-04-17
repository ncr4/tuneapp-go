package main

import (
	"os"
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	api_email := os.Getenv("API_EMAIL")
	api_key := os.Getenv("API_KEY")

	// Update a customer passing in all the params you want to update
	customer := tuneuptechnology.UpdateCustomer(
		&tuneuptechnology.Customer{
			Auth: api_email,
			APIKey: api_key,
			Id: 1, // the ID of the customer you are updating
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

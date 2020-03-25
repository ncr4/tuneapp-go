package main

import (
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	tuneuptechnology.CreateCustomer(
		// Pass in your email and API key
		&tuneuptechnology.Client{
			Auth: "",
			APIKey: "",
		},

		// Create a customer passing in all required params
		&tuneuptechnology.Customer{
			Firstname: "Go",
			Lastname: "Test",
			Email: "go-test@test.com",
			Phone: "8018981234",
			UserId: 1,
			Notes: "Testing some notes here",
			LocationId: 1,
		},
	)
}

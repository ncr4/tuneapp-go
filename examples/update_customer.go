package main

import (
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	tuneuptechnology.UpdateCustomer(
		// Pass in your email and API key
		&tuneuptechnology.Client{
			Auth: "",
			APIKey: "",
		},
		
		// Update a customer passing in all the desired params to be updated
		&tuneuptechnology.Customer{
			Id: 1, // the ID of the customer you want to update
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

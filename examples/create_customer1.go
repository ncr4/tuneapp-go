package main

import (
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Define your credentials
	var auth string = "YOUR_EMAIL_HERE"
	var api_key string = "YOUR_API_KEY_HERE"
	
	// Call an action
	tuneuptechnology.CreateCustomer(auth, api_key, 
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

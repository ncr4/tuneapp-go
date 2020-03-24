package main

import (
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Define your credentials
	var auth string = "YOUR_EMAIL_HERE"
	var api_key string = "YOUR_API_KEY_HERE"
	
	// Call an action
	tuneuptechnology.RetrieveCustomers(auth, api_key)
}

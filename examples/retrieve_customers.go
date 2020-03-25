package main

import (
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	tuneuptechnology.RetrieveCustomers(
		// Pass in your email and API key
		&tuneuptechnology.Client{
			Auth: "",
			APIKey: "",
		},

		// Retrieves all customers
	)
}

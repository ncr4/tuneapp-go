package main

import (
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	tuneuptechnology.DeleteCustomer(		
		// Pass in your email and API key
		&tuneuptechnology.Client{
			Auth: "",
			APIKey: "",
		},

		// deletes customer with ID #23
		&tuneuptechnology.Customer{
			Id: 23,
		},
	)
}

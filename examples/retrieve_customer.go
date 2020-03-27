package main

import (
	"os"
	"github.com/ncr4/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	api_email := os.Getenv("API_EMAIL")
	api_key := os.Getenv("API_KEY")

	// Retrieve a single customer
	customer := tuneuptechnology.RetrieveCustomer(
		&tuneuptechnology.Customer{
			Auth: api_email,
			APIKey: api_key,
			Id: 1, // the ID of the customer you are retrieving
		},
	)

	tuneuptechnology.PrettyPrint(customer)

	// Alternatively, print individual items from the response:
	// fmt.Println(customer["data"].(map[string]interface{})["firstname"], customer["data"].(map[string]interface{})["lastname"])
}

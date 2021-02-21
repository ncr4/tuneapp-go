package main

import (
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func main() {
	// Setup your email and API key
	apiEmail := os.Getenv("API_EMAIL")
	apiKey := os.Getenv("API_KEY")

	// Retrieve a single customer
	customer := tuneuptechnology.RetrieveCustomer(
		&tuneuptechnology.Customer{
			Auth:   apiEmail,
			APIKey: apiKey,
			ID:     1, // the ID of the customer you are retrieving
		},
	)

	tuneuptechnology.PrettyPrint(customer)

	// Alternatively, print individual items from the response:
	// fmt.Println(customer["data"].(map[string]interface{})["firstname"], customer["data"].(map[string]interface{})["lastname"])
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func main() {
	client := tuneuptechnology.New(os.Getenv("API_EMAIL"), os.Getenv("API_KEY"))

	customer := client.UpdateCustomer(
		1,
		&tuneuptechnology.Customer{
			Firstname:  "Jake",
			Lastname:   "Peralta",
			Email:      "jake@example.com",
			Phone:      "8015551234",
			UserID:     1,
			Notes:      "Believes he is a good detective.",
			LocationID: 1,
		},
	)

	prettyJSON, err := json.MarshalIndent(customer, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

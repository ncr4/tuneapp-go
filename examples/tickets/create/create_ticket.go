package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go/v3"
)

func main() {
	client := tuneuptechnology.New(os.Getenv("API_EMAIL"), os.Getenv("API_KEY"))

	ticket := client.CreateTicket(
		&tuneuptechnology.Ticket{
			CustomerID:   1,
			TicketTypeID: 1,
			Serial:       "10000",
			UserID:       1,
			Notes:        "here are some notes",
			Title:        "Fancy Title",
			Status:       1,
			Device:       "2",
			IMEI:         10000,
			LocationID:   1,
		},
	)

	prettyJSON, err := json.MarshalIndent(ticket, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

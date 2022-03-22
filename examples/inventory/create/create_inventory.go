package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go/v3"
)

func main() {
	client := tuneuptechnology.New(os.Getenv("API_EMAIL"), os.Getenv("API_KEY"))

	inventory := client.CreateInventory(
		&tuneuptechnology.Inventory{
			Name:            "Inventory Item",
			InventoryTypeID: 1,
			PartNumber:      "1234",
			SKU:             "1234",
			Notes:           "here are some notes",
			PartPrice:       "19.99",
			LocationID:      1,
			Quantity:        1,
		},
	)

	prettyJSON, err := json.MarshalIndent(inventory, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

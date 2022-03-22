package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go/v3"
)

func main() {
	client := tuneuptechnology.New(os.Getenv("API_EMAIL"), os.Getenv("API_KEY"))

	inventory := client.DeleteInventory(23)

	prettyJSON, err := json.MarshalIndent(inventory, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

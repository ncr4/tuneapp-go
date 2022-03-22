package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuneuptechnology/tuneuptechnology-go/v3"
)

func main() {
	client := tuneuptechnology.New(os.Getenv("API_EMAIL"), os.Getenv("API_KEY"))

	location := client.CreateLocation(
		&tuneuptechnology.Location{
			Name:   "Location Name",
			Street: "123 California Ave",
			City:   "Salt Lake",
			State:  "UT",
			Zip:    84043,
		},
	)

	prettyJSON, err := json.MarshalIndent(location, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

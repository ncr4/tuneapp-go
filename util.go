package tuneuptechnology

import (
    "fmt"
	"os"
	"encoding/json"
)

// Helper util to pretty print JSON
func PrettyPrint(data map[string]interface{}) {
    prettyJSON, err := json.MarshalIndent(data, "", "    ")
    if err != nil {
    	fmt.Fprintln(os.Stderr, "error creating JSON:", err)
    }
    fmt.Printf("%s\n", string(prettyJSON))
}

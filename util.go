package tuneuptechnology

import (
	"encoding/json"
	"fmt"
	"os"
)

// PrettyPrint is a helper util to pretty print JSON
func PrettyPrint(data map[string]interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating JSON:", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

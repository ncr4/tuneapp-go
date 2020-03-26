package tuneuptechnology

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"log"
	"bytes"
    "encoding/json"
)

// Setup the HTTP client and response functionality
const APIBaseUrl = "https://app.tuneuptechnology.com/api/"
// const UserAgent = "Tuneup_Technology_App/GoClient/" + Version // TODO: Implement this

type Client struct {
	Auth 		string `json:"auth"`
	APIKey		string `json:"api_key"`
}

func Response(values map[string]interface{}, endpoint string) {
	jsonData, _ := json.Marshal(values)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatal(err)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	
    fmt.Println(string(responseData))
}

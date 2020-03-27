package tuneuptechnology

import (
    "io/ioutil"
    "net/http"
	"log"
	"bytes"
    "encoding/json"
)

// Setup the HTTP client and response functionality
const APIBaseUrl = "https://app.tuneuptechnology.com/api/"

// Request a response from the API with supplied data
func Response(data interface{}, endpoint string) map[string]interface{} {
    // Make a request to the API
    jsonData, _ := json.Marshal(data)
	request, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatal(err)
    }

    // Read the response data from the API
    responseData, err := ioutil.ReadAll(request.Body)
    if err != nil {
        log.Fatal(err)
    }

    // Convert the response to a map
    var response map[string]interface{}
    err = json.Unmarshal(responseData, &response)
    if err != nil {
        log.Fatal(err)
    }

    return response
}

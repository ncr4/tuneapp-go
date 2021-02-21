package tuneuptechnology

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// APIBaseURL sets up the HTTP client and response functionality
const APIBaseURL = "https://app.tuneuptechnology.com/api/"

// UserAgent sets the user-agent for requests
const UserAgent = "TuneupTechnologyApp/GoClient/" + Version

// makeHTTPRequest requests a response from the API with supplied data
func makeHTTPRequest(data interface{}, endpoint string) map[string]interface{} {
	jsonData, _ := json.Marshal(data)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	request, err := client.Post(
		endpoint,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", UserAgent)

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

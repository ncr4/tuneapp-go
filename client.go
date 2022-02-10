package tuneuptechnology

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	// APIEmail is the email of the user and is required on every request
	APIEmail string
	// APIKey is the API key of the user and is required on every request
	APIKey string
	// BaseURL is the base url where the API resides
	BaseURL string
	// Timeout is the timeout of each request made to the API
	Timeout int
	// Version is the version of this client library
	Version string
	// Client is an HTTP client used to make API requests
	Client *http.Client
}

// Version is the client library version
const Version = "3.0.1"

// Timeout is the number of seconds before a request will timeout
const Timeout = 10

// UserAgent sets the user-agent for requests
const UserAgent = "TuneupTechnologyApp/GoClient/" + Version

// baseURL sets the base url of a request, if not specified, a default will be used
func (client *Client) baseURL() string {
	if client.BaseURL != "" {
		return client.BaseURL
	}
	return "https://app.tuneuptechnology.com/api"
}

func (client *Client) httpClient() *http.Client {
	if client.Client != nil {
		return client.Client
	}
	return &http.Client{
		Timeout: time.Second * time.Duration(client.Timeout),
	}
}

// New returns a new Client with the given API email and key
func New(apiEmail string, apiKey string) *Client {
	if apiEmail == "" || apiKey == "" {
		log.Fatal("apiEmail and apiKey are required to create a Client.")
	}

	return &Client{
		APIEmail: apiEmail,
		APIKey:   apiKey,
	}
}

// makeHTTPRequest requests a response from the API with supplied data
func (client *Client) makeHTTPRequest(method string, endpoint string, data interface{}) map[string]interface{} {
	// Based on https://christiangiacomi.com/posts/simple-put-patch-request-go/
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest(strings.ToUpper(method), endpoint, bytes.NewBuffer(jsonData))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", UserAgent)
	request.Header.Set("Email", client.APIEmail)
	request.Header.Set("Api-Key", client.APIKey)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.httpClient().Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseJson map[string]interface{}
	err = json.Unmarshal(responseData, &responseJson)
	if err != nil {
		log.Fatal(err)
	}

	return responseJson
}

package tuneuptechnology

import (
    "fmt" // format i/o
    "io/ioutil" // read JSON
    "log" // log errors
    "net/http" // client for accessing the API
	"os" // allows us to exit on error
	"bytes" // needed for json
    "encoding/json" // json output
    "strconv" // converts int's to strings
)

// Location Struct
type Location struct {    
    Id     int     `json:"id"`
    Name   string  `json:"name"`
    Street string  `json:"street"`
    City   string  `json:"city"`
    State  string  `json:"state"`
    Zip    int     `json:"zip"`
}

// Create a location
func CreateLocation(client *Client, data *Location) {
	endpoint := APIBaseUrl + "locations/create"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "name": data.Name,
        "street": data.Street,
        "city": data.City,
        "state": data.State,
        "zip": data.Zip,
	}
	jsonValue, _ := json.Marshal(values)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}

// Retrieve a list of locations
func RetrieveLocations(client *Client) {
	endpoint := APIBaseUrl + "locations"

	values := map[string]string{
		"auth": client.Auth,
        "api_key": client.APIKey,
	}
	jsonValue, _ := json.Marshal(values)

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}

// Retrieve a single location
func RetrieveLocation(client *Client, data *Location) {
    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id)

    values := map[string]string{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }
    jsonValue, _ := json.Marshal(values)

    response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}

// Update a location
func UpdateLocation(client *Client, data *Location) {
    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id) + "/update"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "name": data.Name,
        "street": data.Street,
        "city": data.City,
        "state": data.State,
        "zip": data.Zip,
	}
    jsonValue, _ := json.Marshal(values)

    response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}

// Delete a location
func DeleteLocation(client *Client, data *Location) {
    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id) + "/delete"

    values := map[string]string{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }
    jsonValue, _ := json.Marshal(values)

    response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}

package tuneuptechnology

import (
    "strconv"
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
	values := map[string]interface{}{
		"auth":     client.Auth,
        "api_key":  client.APIKey,
        "name":     data.Name,
        "street":   data.Street,
        "city":     data.City,
        "state":    data.State,
        "zip":      data.Zip,
	}

    endpoint := APIBaseUrl + "locations/create"

    Response(values, endpoint)
}

// Retrieve a list of locations
func RetrieveLocations(client *Client) {
	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
	}

    endpoint := APIBaseUrl + "locations"

    Response(values, endpoint)
}

// Retrieve a single location
func RetrieveLocation(client *Client, data *Location) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id)

    Response(values, endpoint)
}

// Update a location
func UpdateLocation(client *Client, data *Location) {
	values := map[string]interface{}{
		"auth":     client.Auth,
        "api_key":  client.APIKey,
        "name":     data.Name,
        "street":   data.Street,
        "city":     data.City,
        "state":    data.State,
        "zip":      data.Zip,
    }
    
    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id) + "/update"

    Response(values, endpoint)
}

// Delete a location
func DeleteLocation(client *Client, data *Location) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id) + "/delete"

    Response(values, endpoint)
}

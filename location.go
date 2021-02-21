package tuneuptechnology

import (
	"strconv"
)

// Location lists all properties of a location
type Location struct {
	Auth   string `json:"auth"`
	APIKey string `json:"api_key"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    int    `json:"zip"`
}

// CreateLocation creates a location
func CreateLocation(data *Location) map[string]interface{} {
	endpoint := APIBaseURL + "locations/create"
	return makeHTTPRequest(data, endpoint)
}

// AllLocations retrieves a list of locations
func AllLocations(data *Location) map[string]interface{} {
	endpoint := APIBaseURL + "locations"
	return makeHTTPRequest(data, endpoint)
}

// RetrieveLocation retrieves a single location
func RetrieveLocation(data *Location) map[string]interface{} {
	endpoint := APIBaseURL + "locations/" + strconv.Itoa(data.ID)
	return makeHTTPRequest(data, endpoint)
}

// UpdateLocation updates a location
func UpdateLocation(data *Location) map[string]interface{} {
	endpoint := APIBaseURL + "locations/" + strconv.Itoa(data.ID) + "/update"
	return makeHTTPRequest(data, endpoint)
}

// DeleteLocation deletes a location
func DeleteLocation(data *Location) map[string]interface{} {
	endpoint := APIBaseURL + "locations/" + strconv.Itoa(data.ID) + "/delete"
	return makeHTTPRequest(data, endpoint)
}

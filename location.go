package tuneuptechnology

import (
    "strconv"
)

type Location struct {
    Auth    string  `json:"auth"`
    APIKey  string  `json:"api_key"`
    Id      int     `json:"id"`
    Name    string  `json:"name"`
    Street  string  `json:"street"`
    City    string  `json:"city"`
    State   string  `json:"state"`
    Zip     int     `json:"zip"`
}

// Create a location
func CreateLocation(data *Location) map[string]interface{} {
    endpoint := APIBaseUrl + "locations/create"
    return Response(data, endpoint)
}

// Retrieve a list of locations
func AllLocations(data *Location) map[string]interface{} {
    endpoint := APIBaseUrl + "locations"
    return Response(data, endpoint)
}

// Retrieve a single location
func RetrieveLocation(data *Location) map[string]interface{} {
    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id)
    return Response(data, endpoint)
}

// Update a location
func UpdateLocation(data *Location) map[string]interface{} {
    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id) + "/update"
    return Response(data, endpoint)
}

// Delete a location
func DeleteLocation(data *Location) map[string]interface{} {
    endpoint := APIBaseUrl + "locations/" + strconv.Itoa(data.Id) + "/delete"
    return Response(data, endpoint)
}

package tuneuptechnology

import (
	"strconv"
	"time"
)

// Location lists all properties of a location
type Location struct {
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Street    string     `json:"street,omitempty"`
	City      string     `json:"city,omitempty"`
	State     string     `json:"state,omitempty"`
	Zip       int        `json:"zip,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// CreateLocation creates a location
func (client *Client) CreateLocation(data *Location) map[string]interface{} {
	endpoint := client.baseURL() + "/locations"
	return client.makeHTTPRequest("post", endpoint, data)
}

// AllLocations retrieves a list of locations
func (client *Client) AllLocations() map[string]interface{} {
	endpoint := client.baseURL() + "/locations"
	return client.makeHTTPRequest("get", endpoint, nil)
}

// RetrieveLocation retrieves a single location
func (client *Client) RetrieveLocation(id int) map[string]interface{} {
	endpoint := client.baseURL() + "/locations/" + strconv.Itoa(id)
	return client.makeHTTPRequest("get", endpoint, nil)
}

// UpdateLocation updates a location
func (client *Client) UpdateLocation(id int, data *Location) map[string]interface{} {
	endpoint := client.baseURL() + "/locations/" + strconv.Itoa(id)
	return client.makeHTTPRequest("patch", endpoint, data)
}

// DeleteLocation deletes a location
func (client *Client) DeleteLocation(id int) map[string]interface{} {
	endpoint := client.baseURL() + "/locations/" + strconv.Itoa(id)
	return client.makeHTTPRequest("delete", endpoint, nil)
}

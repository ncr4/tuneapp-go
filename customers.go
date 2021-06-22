package tuneuptechnology

import (
	"strconv"
	"time"
)

// Customer lists all properties of a customer
type Customer struct {
	ID         int        `json:"id,omitempty"`
	Firstname  string     `json:"firstname,omitempty"`
	Lastname   string     `json:"lastname,omitempty"`
	Email      string     `json:"email,omitempty"`
	Phone      string     `json:"phone,omitempty"`
	UserID     int        `json:"user_id,omitempty"`
	Notes      string     `json:"notes,omitempty"`
	LocationID int        `json:"location_id,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

// CreateCustomer creates a customer
func (client *Client) CreateCustomer(data *Customer) map[string]interface{} {
	endpoint := client.baseURL() + "/customers"
	return client.makeHTTPRequest("post", endpoint, data)
}

// AllCustomers retrieves a list of customers
func (client *Client) AllCustomers() map[string]interface{} {
	endpoint := client.baseURL() + "/customers"
	return client.makeHTTPRequest("get", endpoint, nil)
}

// RetrieveCustomer retrieves a single customer
func (client *Client) RetrieveCustomer(id int) map[string]interface{} {
	endpoint := client.baseURL() + "/customers/" + strconv.Itoa(id)
	return client.makeHTTPRequest("get", endpoint, nil)
}

// UpdateCustomer updates a customer
func (client *Client) UpdateCustomer(id int, data *Customer) map[string]interface{} {
	endpoint := client.baseURL() + "/customers/" + strconv.Itoa(id)
	return client.makeHTTPRequest("patch", endpoint, data)
}

// DeleteCustomer deletes a customer
func (client *Client) DeleteCustomer(id int) map[string]interface{} {
	endpoint := client.baseURL() + "/customers/" + strconv.Itoa(id)
	return client.makeHTTPRequest("delete", endpoint, nil)
}

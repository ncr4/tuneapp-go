package tuneuptechnology

import (
	"strconv"
	"time"
)

// Ticket lists all properties of a ticket
type Ticket struct {
	ID           int        `json:"id,omitempty"`
	CustomerID   int        `json:"customer_id,omitempty"`
	TicketTypeID int        `json:"ticket_type_id,omitempty"`
	Serial       string     `json:"serial,omitempty"`
	UserID       int        `json:"user_id,omitempty"`
	Notes        string     `json:"notes,omitempty"`
	Title        string     `json:"title,omitempty"`
	Status       int        `json:"status,omitempty"`
	Device       string     `json:"device,omitempty"`
	IMEI         int        `json:"imei,omitempty"`
	LocationID   int        `json:"location_id,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// CreateTicket creates a ticket
func (client *Client) CreateTicket(data *Ticket) map[string]interface{} {
	endpoint := client.baseURL() + "/tickets"
	return client.makeHTTPRequest("post", endpoint, data)
}

// AllTickets retrieves a list of tickets
func (client *Client) AllTickets() map[string]interface{} {
	endpoint := client.baseURL() + "/tickets"
	return client.makeHTTPRequest("get", endpoint, nil)
}

// RetrieveTicket retrieves a single ticket
func (client *Client) RetrieveTicket(id int) map[string]interface{} {
	endpoint := client.baseURL() + "/tickets/" + strconv.Itoa(id)
	return client.makeHTTPRequest("get", endpoint, nil)
}

// UpdateTicket updates a ticket
func (client *Client) UpdateTicket(id int, data *Ticket) map[string]interface{} {
	endpoint := client.baseURL() + "/tickets/" + strconv.Itoa(id)
	return client.makeHTTPRequest("patch", endpoint, data)
}

// DeleteTicket deletes a ticket
func (client *Client) DeleteTicket(id int) map[string]interface{} {
	endpoint := client.baseURL() + "/tickets/" + strconv.Itoa(id)
	return client.makeHTTPRequest("delete", endpoint, nil)
}

package tuneuptechnology

import (
	"strconv"
)

// Ticket lists all properties of a ticket
type Ticket struct {
	Auth         string `json:"auth"`
	APIKey       string `json:"api_key"`
	ID           int    `json:"id"`
	CustomerID   int    `json:"customer_id"`
	TicketTypeID int    `json:"ticket_type_id"`
	Serial       string `json:"serial"`
	UserID       int    `json:"user_id"`
	Notes        string `json:"notes"`
	Title        string `json:"title"`
	Status       string `json:"status"`
	Device       string `json:"device"`
	IMEI         string `json:"imei"`
	LocationID   int    `json:"location_id"`
}

// CreateTicket creates a ticket
func CreateTicket(data *Ticket) map[string]interface{} {
	endpoint := APIBaseURL + "tickets/create"
	return Response(data, endpoint)
}

// AllTickets retrieves a list of tickets
func AllTickets(data *Ticket) map[string]interface{} {
	endpoint := APIBaseURL + "tickets"
	return Response(data, endpoint)
}

// RetrieveTicket retrieves a single ticket
func RetrieveTicket(data *Ticket) map[string]interface{} {
	endpoint := APIBaseURL + "tickets/" + strconv.Itoa(data.ID)
	return Response(data, endpoint)
}

// UpdateTicket updates a ticket
func UpdateTicket(data *Ticket) map[string]interface{} {
	endpoint := APIBaseURL + "tickets/" + strconv.Itoa(data.ID) + "/update"
	return Response(data, endpoint)
}

// DeleteTicket deletes a ticket
func DeleteTicket(data *Ticket) map[string]interface{} {
	endpoint := APIBaseURL + "tickets/" + strconv.Itoa(data.ID) + "/delete"
	return Response(data, endpoint)
}

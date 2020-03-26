package tuneuptechnology

import (
    "strconv"
)

// Ticket Struct
type Ticket struct {    
    Id              int     `json:"id"`
    CustomerId      int     `json:"customer_id"`
    TicketTypeId    int     `json:"ticket_type_id"`
    Serial          string  `json:"serial"`
    UserId          int     `json:"user_id"`
    Notes           string  `json:"notes"`
    Title           string  `json:"title"`
    Status          string  `json:"status"`
    Device          string  `json:"device"`
    IMEI            string  `json:"imei"`
    LocationId      int     `json:"location_id"`
}

// Create a ticket
func CreateTicket(client *Client, data *Ticket) {
	values := map[string]interface{}{
		"auth":             client.Auth,
        "api_key":          client.APIKey,
        "customer_id":      data.CustomerId,
        "ticket_type_id":   data.TicketTypeId,
        "serial":           data.Serial,
        "user_id":          data.UserId,
        "notes":            data.Notes,
        "title":            data.Title,
        "status":           data.Status,
        "device":           data.Device,
        "imei":             data.IMEI,
        "location_id":      data.LocationId,
    }

    endpoint := APIBaseUrl + "tickets/create"

    Response(values, endpoint)
}

// Retrieve a list of tickets
func RetrieveTickets(client *Client) {
	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }
    
    endpoint := APIBaseUrl + "tickets"

    Response(values, endpoint)
}

// Retrieve a single ticket
func RetrieveTicket(client *Client, data *Ticket) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id)

    Response(values, endpoint)
}

// Update a ticket
func UpdateTicket(client *Client, data *Ticket) {
	values := map[string]interface{}{
		"auth":             client.Auth,
        "api_key":          client.APIKey,
        "customer_id":      data.CustomerId,
        "ticket_type_id":   data.TicketTypeId,
        "serial":           data.Serial,
        "user_id":          data.UserId,
        "notes":            data.Notes,
        "title":            data.Title,
        "status":           data.Status,
        "device":           data.Device,
        "imei":             data.IMEI,
        "location_id":      data.LocationId,
    }
    
    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id) + "/update"

    Response(values, endpoint)
}

// Delete a ticket
func DeleteTicket(client *Client, data *Ticket) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id) + "/delete"

    Response(values, endpoint)
}

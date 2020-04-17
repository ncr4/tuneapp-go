package tuneuptechnology

import (
    "strconv"
)

type Ticket struct {
    Auth            string  `json:"auth"`
    APIKey          string  `json:"api_key"`
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
func CreateTicket(data *Ticket) map[string]interface{} {
    endpoint := APIBaseUrl + "tickets/create"
    return Response(data, endpoint)
}

// Retrieve a list of tickets
func AllTickets(data *Ticket) map[string]interface{} {
    endpoint := APIBaseUrl + "tickets"
    return Response(data, endpoint)
}

// Retrieve a single ticket
func RetrieveTicket(data *Ticket) map[string]interface{} {
    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id)
    return Response(data, endpoint)
}

// Update a ticket
func UpdateTicket(data *Ticket) map[string]interface{} {
    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id) + "/update"
    return Response(data, endpoint)
}

// Delete a ticket
func DeleteTicket(data *Ticket) map[string]interface{} {
    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id) + "/delete"
    return Response(data, endpoint)
}

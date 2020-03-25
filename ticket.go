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
	endpoint := APIBaseUrl + "tickets/create"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "customer_id": data.CustomerId,
        "ticket_type_id": data.TicketTypeId,
        "serial": data.Serial,
        "user_id": data.UserId,
        "notes": data.Notes,
        "title": data.Title,
        "status": data.Status,
        "device": data.Device,
        "imei": data.IMEI,
        "location_id": data.LocationId,
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

// Retrieve a list of tickets
func RetrieveTickets(client *Client) {
	endpoint := APIBaseUrl + "tickets"

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

// Retrieve a single ticket
func RetrieveTicket(client *Client, data *Ticket) {
    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id)

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

// Update a ticket
func UpdateTicket(client *Client, data *Ticket) {
    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id) + "/update"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "customer_id": data.CustomerId,
        "ticket_type_id": data.TicketTypeId,
        "serial": data.Serial,
        "user_id": data.UserId,
        "notes": data.Notes,
        "title": data.Title,
        "status": data.Status,
        "device": data.Device,
        "imei": data.IMEI,
        "location_id": data.LocationId,
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

// Delete a ticket
func DeleteTicket(client *Client, data *Ticket) {
    endpoint := APIBaseUrl + "tickets/" + strconv.Itoa(data.Id) + "/delete"

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

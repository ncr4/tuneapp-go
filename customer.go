package tuneuptechnology

import (
    "strconv"
)

// Customer Struct
type Customer struct {    
    Id          int     `json:"id"`
    Firstname   string  `json:"firstname"`
    Lastname    string  `json:"lastname"`
    Email       string  `json:"email"`
    Phone       string  `json:"phone"`
    UserId      int     `json:"user_id"`
    Notes       string  `json:"notes"`
    LocationId  int     `json:"location_id"`
}

// Create a customer
func CreateCustomer(client *Client, data *Customer) {
	values := map[string]interface{}{
		"auth":         client.Auth,
        "api_key":      client.APIKey,
        "firstname":    data.Firstname,
        "lastname":     data.Lastname,
        "email":        data.Email,
        "phone":        data.Phone,
        "user_id":      data.UserId,
        "notes":        data.Notes,
        "location_id":  data.LocationId,
	}

    endpoint := APIBaseUrl + "customers/create"

    Response(values, endpoint)
}

// Retrieve a list of customers
func RetrieveCustomers(client *Client) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "customers"

    Response(values, endpoint)
}

// Retrieve a single customer
func RetrieveCustomer(client *Client, data *Customer) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id)

    Response(values, endpoint)
}

// Update a customer
func UpdateCustomer(client *Client, data *Customer) {
	values := map[string]interface{}{
		"auth":         client.Auth,
        "api_key":      client.APIKey,
        "firstname":    data.Firstname,
        "lastname":     data.Lastname,
        "email":        data.Email,
        "phone":        data.Phone,
        "user_id":      data.UserId,
        "notes":        data.Notes,
        "location_id":  data.LocationId,
	}

    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id) + "/update"

    Response(values, endpoint)
}

// Delete a customer
func DeleteCustomer(client *Client, data *Customer) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id) + "/delete"

    Response(values, endpoint)
}

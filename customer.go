package tuneuptechnology

import (
    "strconv"
)

type Customer struct {
    Auth        string  `json:"auth"`
    APIKey      string  `json:"api_key"`
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
func CreateCustomer(data *Customer) map[string]interface{} {
    endpoint := APIBaseUrl + "customers/create"
    return Response(data, endpoint)
}

// Retrieve a list of customers
func AllCustomers(data *Customer) map[string]interface{} {
    endpoint := APIBaseUrl + "customers"
    return Response(data, endpoint)
}

// Retrieve a single customer
func RetrieveCustomer(data *Customer) map[string]interface{} {
    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id)
    return Response(data, endpoint)
}

// Update a customer
func UpdateCustomer(data *Customer) map[string]interface{} {
    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id) + "/update"
    return Response(data, endpoint)
}

// Delete a customer
func DeleteCustomer(data *Customer) map[string]interface{} {
    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id) + "/delete"
    return Response(data, endpoint)
}

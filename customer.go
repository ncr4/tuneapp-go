package tuneuptechnology

import (
	"strconv"
)

// Customer lists all properties of a customer
type Customer struct {
	Auth       string `json:"auth"`
	APIKey     string `json:"api_key"`
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	UserID     int    `json:"user_id"`
	Notes      string `json:"notes"`
	LocationID int    `json:"location_id"`
}

// CreateCustomer creates a customer
func CreateCustomer(data *Customer) map[string]interface{} {
	endpoint := APIBaseURL + "customers/create"
	return Response(data, endpoint)
}

// AllCustomers retrieves a list of customers
func AllCustomers(data *Customer) map[string]interface{} {
	endpoint := APIBaseURL + "customers"
	return Response(data, endpoint)
}

// RetrieveCustomer retrieves a single customer
func RetrieveCustomer(data *Customer) map[string]interface{} {
	endpoint := APIBaseURL + "customers/" + strconv.Itoa(data.ID)
	return Response(data, endpoint)
}

// UpdateCustomer updates a customer
func UpdateCustomer(data *Customer) map[string]interface{} {
	endpoint := APIBaseURL + "customers/" + strconv.Itoa(data.ID) + "/update"
	return Response(data, endpoint)
}

// DeleteCustomer deletes a customer
func DeleteCustomer(data *Customer) map[string]interface{} {
	endpoint := APIBaseURL + "customers/" + strconv.Itoa(data.ID) + "/delete"
	return Response(data, endpoint)
}

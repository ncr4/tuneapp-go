package tuneuptechnology_test

import (
	"reflect"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func (c *ClientTests) TestCreateCustomer() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	customer := client.CreateCustomer(
		&tuneuptechnology.Customer{
			Firstname:  "Jake",
			Lastname:   "Peralta",
			Email:      "jake@example.com",
			Phone:      "8015551234",
			UserID:     1,
			Notes:      "Believes he is a good detective.",
			LocationID: 1,
		},
	)

	assert.Equal(customer["message"], "Customer created successfully.")
}

func (c *ClientTests) TestRetrieveCustomer() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	customer := client.RetrieveCustomer(1)

	assert.Equal(customer["message"], "Customer retrieved successfully.")
}

func (c *ClientTests) TestAllCustomers() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	customer := client.AllCustomers()

	assert.Equal(customer["message"], "Customers retrieved successfully.")
	assert.Greater(reflect.ValueOf(customer["data"]).Len(), 1)
}

func (c *ClientTests) TestUpdateCustomer() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	customer := client.UpdateCustomer(
		1,
		&tuneuptechnology.Customer{
			Firstname:  "Jake",
			Lastname:   "Peralta",
			Email:      "jake@example.com",
			Phone:      "8015551234",
			UserID:     1,
			Notes:      "Believes he is a good detective.",
			LocationID: 1,
		},
	)

	assert.Equal(customer["message"], "Customer updated successfully.")
}

func (c *ClientTests) TestDeleteCustomer() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	customer := client.DeleteCustomer(1)

	assert.Equal(customer["message"], "Customer deleted successfully.")
}

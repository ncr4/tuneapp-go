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
			LocationID: 2,
		},
	)

	assert.Equal(customer["firstname"], "Jake")
}

func (c *ClientTests) TestRetrieveCustomer() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	customer := client.RetrieveCustomer(1)

	require.NotNil(customer["firstname"])
}

func (c *ClientTests) TestAllCustomers() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	customers := client.AllCustomers()

	assert.Greater(reflect.ValueOf(customers["data"]).Len(), 1)
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
			LocationID: 2,
		},
	)

	assert.Equal(customer["firstname"], "Jake")
}

func (c *ClientTests) TestDeleteCustomer() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	customer := client.DeleteCustomer(1)

	require.NotNil(customer["deleted_at"])
}

package tuneuptechnology_test

import (
	"reflect"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func (c *ClientTests) TestCreateInventory() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	inventory := client.CreateInventory(
		&tuneuptechnology.Inventory{
			Name:            "Inventory Item",
			InventoryTypeID: 1,
			PartNumber:      "1234",
			SKU:             "1234",
			Notes:           "here are some notes",
			PartPrice:       "19.99",
			LocationID:      2,
			Quantity:        1,
		},
	)

	assert.Equal(inventory["name"], "Inventory Item")
}

func (c *ClientTests) TestRetrieveInventory() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	inventory := client.RetrieveInventory(1)

	require.NotNil(inventory["name"])
}

func (c *ClientTests) TestAllInventory() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	inventory := client.AllInventory()

	assert.Greater(reflect.ValueOf(inventory["data"]).Len(), 1)
}

func (c *ClientTests) TestUpdateInventory() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	inventory := client.UpdateInventory(
		1,
		&tuneuptechnology.Inventory{
			Name:            "Inventory Item",
			InventoryTypeID: 1,
			PartNumber:      "1234",
			SKU:             "1234",
			Notes:           "here are some notes",
			PartPrice:       "19.99",
			LocationID:      2,
			Quantity:        1,
		},
	)

	assert.Equal(inventory["name"], "Inventory Item")
}

func (c *ClientTests) TestDeleteInventory() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	inventory := client.DeleteInventory(1)

	require.NotNil(inventory["deleted_at"])
}

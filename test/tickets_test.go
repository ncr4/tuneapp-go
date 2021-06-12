package tuneuptechnology_test

import (
	"reflect"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func (c *ClientTests) TestCreateTicket() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	ticket := client.CreateTicket(
		&tuneuptechnology.Ticket{
			CustomerID:   2,
			TicketTypeID: 1,
			Serial:       "10000",
			UserID:       1,
			Notes:        "here are some notes",
			Title:        "Fancy Title",
			Status:       1,
			Device:       "iPhone",
			IMEI:         10000,
			LocationID:   2,
		},
	)

	assert.Equal(ticket["title"], "Fancy Title")
}

func (c *ClientTests) TestRetrieveTicket() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	ticket := client.RetrieveTicket(1)

	require.NotNil(ticket["title"])
}

func (c *ClientTests) TestAllTickets() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	tickets := client.AllTickets()

	assert.Greater(reflect.ValueOf(tickets["data"]).Len(), 1)
}

func (c *ClientTests) TestUpdateTicket() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	ticket := client.UpdateTicket(
		1,
		&tuneuptechnology.Ticket{
			CustomerID:   2,
			TicketTypeID: 1,
			Serial:       "10000",
			UserID:       1,
			Notes:        "here are some notes",
			Title:        "Fancy Title",
			Status:       1,
			Device:       "iPhone",
			IMEI:         10000,
			LocationID:   2,
		},
	)

	assert.Equal(ticket["title"], "Fancy Title")
}

func (c *ClientTests) TestDeleteTicket() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	ticket := client.DeleteTicket(1)

	require.NotNil(ticket["deleted_at"])
}

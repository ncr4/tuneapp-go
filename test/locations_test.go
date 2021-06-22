package tuneuptechnology_test

import (
	"reflect"

	"github.com/tuneuptechnology/tuneuptechnology-go"
)

func (c *ClientTests) TestCreateLocation() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	location := client.CreateLocation(
		&tuneuptechnology.Location{
			Name:   "Location Name",
			Street: "123 California Ave",
			City:   "Salt Lake",
			State:  "UT",
			Zip:    84043,
		},
	)

	assert.Equal(location["name"], "Location Name")
}

func (c *ClientTests) TestRetrieveLocation() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	location := client.RetrieveLocation(1)

	require.NotNil(location["name"])
}

func (c *ClientTests) TestAllLocations() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	locations := client.AllLocations()

	assert.Greater(reflect.ValueOf(locations["data"]).Len(), 1)
}

func (c *ClientTests) TestUpdateLocation() {
	client := c.TestClient()
	assert := c.Assert()

	client.BaseURL = "http://tuneapp.localhost/api"

	location := client.UpdateLocation(
		1,
		&tuneuptechnology.Location{
			Name:   "Location Name",
			Street: "123 California Ave",
			City:   "Salt Lake",
			State:  "UT",
			Zip:    84043,
		},
	)

	assert.Equal(location["name"], "Location Name")
}

func (c *ClientTests) TestDeleteLocation() {
	client := c.TestClient()
	require := c.Require()

	client.BaseURL = "http://tuneapp.localhost/api"

	location := client.DeleteLocation(1)

	require.NotNil(location["deleted_at"])
}

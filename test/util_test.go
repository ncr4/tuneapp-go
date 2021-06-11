package tuneuptechnology_test

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/suite"
	"github.com/tuneuptechnology/tuneuptechnology-go"
)

type ClientTests struct {
	suite.Suite
	recorder *recorder.Recorder
}

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientTests))
}

func (c *ClientTests) TestClient() *tuneuptechnology.Client {
	return &tuneuptechnology.Client{
		APIEmail: os.Getenv("API_EMAIL"),
		APIKey:   os.Getenv("API_KEY"),
		Client:   &http.Client{Transport: c.recorder},
	}
}

func (c *ClientTests) SetupTest() {
	pathComponents := append(
		[]string{"cassettes"}, strings.Split(c.T().Name(), "/")...,
	)
	r, err := recorder.New(filepath.Join(pathComponents...))
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Email")
		delete(i.Request.Headers, "Api-Key")
		return nil
	})
	c.recorder = r
}

func (c *ClientTests) TearDownTest() {
	_ = c.recorder.Stop()
}

package tuneuptechnology

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"bytes"
	"encoding/json"
)

// Create a customer
func CreateCustomer() {
	auth := ""
	api_key := ""
	endpoint := "https://app.tuneuptechnology.com/api/create-customer"


	values := map[string]string{
		"auth": auth,
		"api_key": api_key,
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

// Retrieve a list of customers
func RetrieveCustomers() {
	auth := ""
	api_key := ""
	endpoint := "https://app.tuneuptechnology.com/api/read-customers"


	values := map[string]string{
		"auth": auth,
		"api_key": api_key,
	}
	jsonValue, _ := json.Marshal(values)

	response, err := http.Get(endpoint)

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

// Retrieve a single customer
func RetrieveCustomer() {
	auth := ""
	api_key := ""
	endpoint := "https://app.tuneuptechnology.com/api/read-customer/{id}" // TODO: Change ID to variable


	values := map[string]string{
		"auth": auth,
		"api_key": api_key,
	}
	jsonValue, _ := json.Marshal(values)

	response, err := http.Get(endpoint)

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

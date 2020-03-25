package tuneuptechnology

import (
    "fmt" // format i/o
    "io/ioutil" // read JSON
    "log" // log errors
    "net/http" // client for accessing the API
	"os" // allows us to exit on error
	"bytes" // needed for json
    "encoding/json" // json output
    "strconv" // converts int's to strings
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
	endpoint := APIBaseUrl + "customers/create"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "firstname": data.Firstname,
        "lastname": data.Lastname,
        "email": data.Email,
        "phone": data.Phone,
        "user_id": data.UserId,
        "notes": data.Notes,
        "location_id": data.LocationId,
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
func RetrieveCustomers(client *Client) {
	endpoint := APIBaseUrl + "customers"

	values := map[string]string{
		"auth": client.Auth,
        "api_key": client.APIKey,
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

// Retrieve a single customer
func RetrieveCustomer(client *Client, data *Customer) {
    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id)

    values := map[string]string{
		"auth": client.Auth,
        "api_key": client.APIKey,
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

// Update a customer
func UpdateCustomer(client *Client, data *Customer) {
    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id) + "/update"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "firstname": data.Firstname,
        "lastname": data.Lastname,
        "email": data.Email,
        "phone": data.Phone,
        "user_id": data.UserId,
        "notes": data.Notes,
        "location_id": data.LocationId,
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

// Delete a customer
func DeleteCustomer(client *Client, data *Customer) {
    endpoint := APIBaseUrl + "customers/" + strconv.Itoa(data.Id) + "/delete"

    values := map[string]string{
		"auth": client.Auth,
        "api_key": client.APIKey,
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

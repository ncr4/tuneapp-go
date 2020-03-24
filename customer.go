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
func CreateCustomer(auth string, api_key string) {
	endpoint := apiBaseURL + "customers/create"

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
func RetrieveCustomers(auth string, api_key string) {
	endpoint := apiBaseURL + "customers"

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

// Retrieve a single customer
func RetrieveCustomer(auth string, api_key string, Id int) {
    endpoint := apiBaseURL + "customers/" + strconv.Itoa(Id)

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

// Update a customer
func UpdateCustomer(auth string, api_key string, Id int) {
    endpoint := apiBaseURL + "customers/" + strconv.Itoa(Id) + "/update"

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

// Delete a customer
func DeleteCustomer(auth string, api_key string, Id int) {
    endpoint := apiBaseURL + "customers/" + strconv.Itoa(Id) + "/delete"

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

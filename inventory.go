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

// Inventory Struct
type Inventory struct {    
    Id              int     `json:"id"`
    Name            string  `json:"name"`
    Quantity        int     `json:"quantity"`
    InventoryTypeId int     `json:"inventory_type_id"`
    SKU             string  `json:"sku"`
    PartPrice       string  `json:"part_price"`
    LocationId      int     `json:"location_id"`
}

// Create an inventory item
func CreateInventory(client *Client, data *Inventory) {
	endpoint := APIBaseUrl + "inventory/create"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "name": data.Name,
        "quantity": data.Quantity,
        "inventory_type_id": data.InventoryTypeId,
        "sku": data.SKU,
        "part_price": data.PartPrice,
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

// Retrieve a list of inventory items
func RetrieveInventorys(client *Client) {
	endpoint := APIBaseUrl + "inventory"

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

// Retrieve a single inventory item
func RetrieveInventory(client *Client, data *Inventory) {
    endpoint := APIBaseUrl + "inventory/" + strconv.Itoa(data.Id)

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

// Update an inventory item
func UpdateInventory(client *Client, data *Inventory) {
    endpoint := APIBaseUrl + "inventory/" + strconv.Itoa(data.Id) + "/update"

	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
        "name": data.Name,
        "quantity": data.Quantity,
        "inventory_type_id": data.InventoryTypeId,
        "sku": data.SKU,
        "part_price": data.PartPrice,
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

// Delete an inventory item
func DeleteInventory(client *Client, data *Inventory) {
    endpoint := APIBaseUrl + "inventorys/" + strconv.Itoa(data.Id) + "/delete"

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

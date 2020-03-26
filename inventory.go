package tuneuptechnology

import (
    "strconv"
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
	values := map[string]interface{}{
		"auth":                 client.Auth,
        "api_key":              client.APIKey,
        "name":                 data.Name,
        "quantity":             data.Quantity,
        "inventory_type_id":    data.InventoryTypeId,
        "sku":                  data.SKU,
        "part_price":           data.PartPrice,
        "location_id":          data.LocationId,
    }
    
    endpoint := APIBaseUrl + "inventory/create"

    Response(values, endpoint)
}

// Retrieve a list of inventory items
func RetrieveInventorys(client *Client) {
	values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
	}

    endpoint := APIBaseUrl + "inventory"

    Response(values, endpoint)
}

// Retrieve a single inventory item
func RetrieveInventory(client *Client, data *Inventory) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "inventory/" + strconv.Itoa(data.Id)

    Response(values, endpoint)
}

// Update an inventory item
func UpdateInventory(client *Client, data *Inventory) {
	values := map[string]interface{}{
		"auth":                 client.Auth,
        "api_key":              client.APIKey,
        "name":                 data.Name,
        "quantity":             data.Quantity,
        "inventory_type_id":    data.InventoryTypeId,
        "sku":                  data.SKU,
        "part_price":           data.PartPrice,
        "location_id":          data.LocationId,
	}

    endpoint := APIBaseUrl + "inventory/" + strconv.Itoa(data.Id) + "/update"

    Response(values, endpoint)
}

// Delete an inventory item
func DeleteInventory(client *Client, data *Inventory) {
    values := map[string]interface{}{
		"auth": client.Auth,
        "api_key": client.APIKey,
    }

    endpoint := APIBaseUrl + "inventorys/" + strconv.Itoa(data.Id) + "/delete"

    Response(values, endpoint)
}

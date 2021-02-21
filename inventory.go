package tuneuptechnology

import (
	"strconv"
)

// Inventory lists all properties of inventory
type Inventory struct {
	Auth            string `json:"auth"`
	APIKey          string `json:"api_key"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Quantity        int    `json:"quantity"`
	InventoryTypeID int    `json:"inventory_type_id"`
	SKU             string `json:"sku"`
	PartPrice       string `json:"part_price"`
	LocationID      int    `json:"location_id"`
}

// CreateInventory creates an inventory item
func CreateInventory(data *Inventory) map[string]interface{} {
	endpoint := APIBaseURL + "inventory/create"
	return makeHTTPRequest(data, endpoint)
}

// AllInventory retrieves a list of inventory items
func AllInventory(data *Inventory) map[string]interface{} {
	endpoint := APIBaseURL + "inventory"
	return makeHTTPRequest(data, endpoint)
}

// RetrieveInventory retrieves a single inventory item
func RetrieveInventory(data *Inventory) map[string]interface{} {
	endpoint := APIBaseURL + "inventory/" + strconv.Itoa(data.ID)
	return makeHTTPRequest(data, endpoint)
}

// UpdateInventory updates an inventory item
func UpdateInventory(data *Inventory) map[string]interface{} {
	endpoint := APIBaseURL + "inventory/" + strconv.Itoa(data.ID) + "/update"
	return makeHTTPRequest(data, endpoint)
}

// DeleteInventory deletes an inventory item
func DeleteInventory(data *Inventory) map[string]interface{} {
	endpoint := APIBaseURL + "inventorys/" + strconv.Itoa(data.ID) + "/delete"
	return makeHTTPRequest(data, endpoint)
}

package tuneuptechnology

import (
	"strconv"
)

// Inventory lists all properties of inventory
type Inventory struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	InventoryTypeID int    `json:"inventory_type_id,omitempty"`
	SKU             string `json:"sku,omitempty"`
	PartPrice       string `json:"part_price,omitempty"`
	LocationID      int    `json:"location_id,omitempty"`
}

// CreateInventory creates an inventory item
func (client *Client) CreateInventory(data *Inventory) map[string]interface{} {
	endpoint := client.baseURL() + "inventory"
	return client.makeHTTPRequest("post", endpoint, data)
}

// AllInventory retrieves a list of inventory items
func (client *Client) AllInventory() map[string]interface{} {
	endpoint := client.baseURL() + "inventory"
	return client.makeHTTPRequest("get", endpoint, nil)
}

// RetrieveInventory retrieves a single inventory item
func (client *Client) RetrieveInventory(id int) map[string]interface{} {
	endpoint := client.baseURL() + "inventory/" + strconv.Itoa(id)
	return client.makeHTTPRequest("get", endpoint, nil)
}

// UpdateInventory updates an inventory item
func (client *Client) UpdateInventory(id int, data *Inventory) map[string]interface{} {
	endpoint := client.baseURL() + "inventory/" + strconv.Itoa(id)
	return client.makeHTTPRequest("patch", endpoint, data)
}

// DeleteInventory deletes an inventory item
func (client *Client) DeleteInventory(id int) map[string]interface{} {
	endpoint := client.baseURL() + "inventory/" + strconv.Itoa(id)
	return client.makeHTTPRequest("delete", endpoint, nil)
}

package tuneuptechnology

import (
    "strconv"
)

type Inventory struct {
    Auth            string  `json:"auth"`
    APIKey          string  `json:"api_key"`
    Id              int     `json:"id"`
    Name            string  `json:"name"`
    Quantity        int     `json:"quantity"`
    InventoryTypeId int     `json:"inventory_type_id"`
    SKU             string  `json:"sku"`
    PartPrice       string  `json:"part_price"`
    LocationId      int     `json:"location_id"`
}

// Create an inventory item
func CreateInventory(data *Inventory) map[string]interface{} {
    endpoint := APIBaseUrl + "inventory/create"
    return Response(data, endpoint)
}

// Retrieve a list of inventory items
func AllInventory(data *Inventory) map[string]interface{} {
    endpoint := APIBaseUrl + "inventory"
    return Response(data, endpoint)
}

// Retrieve a single inventory item
func RetrieveInventory(data *Inventory) map[string]interface{} {
    endpoint := APIBaseUrl + "inventory/" + strconv.Itoa(data.Id)
    return Response(data, endpoint)
}

// Update an inventory item
func UpdateInventory(data *Inventory) map[string]interface{} {
    endpoint := APIBaseUrl + "inventory/" + strconv.Itoa(data.Id) + "/update"
    return Response(data, endpoint)
}

// Delete an inventory item
func DeleteInventory(data *Inventory) map[string]interface{} {
    endpoint := APIBaseUrl + "inventorys/" + strconv.Itoa(data.Id) + "/delete"
    return Response(data, endpoint)
}

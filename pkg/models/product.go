package models

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"nombre"`
	Type            string  `json:"tipo"`
	Count           int     `json:"cantidad"`
	Price           float64 `json:"precio"`
	Warehouse       string  `json:"warehouse"`
	WarehouseAdress string  `json:"warehouseadress"`
}

package inventory

import (
	"github.com/jdotw/go-utils/model"
)

type Vendor struct {
	model.Defaults

	Name string

	ScrapeType string

	Products []Product
}

type Product struct {
	model.Defaults

	Name       string
	stockLevel int

	InventoryHistory []InventorySnapshot

	VendorID string
}

type InventorySnapshot struct {
	model.Defaults

	ProductID string

	StockLevel int
}

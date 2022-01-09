// DO NOT EDIT
// This file was code-generated by github.com/jdotw/oapi-gokit-codegen version (devel)
// It is expected that this file will be re-generated and overwitten to
// adapt to changes in the OpenAPI spec that was used to generate it

package vendor

import (
	_ "embed"
)

// Create Syrup Vendor
type CreateVendorRequest struct {
	Name       *string `json:"name,omitempty"`
	ScrapeType *string `json:"scrape_type,omitempty"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Message *string `json:"message,omitempty"`
}

// Point-in-Time Inventory Data
type InventorySnapshot struct {
	CreatedAt  *string `json:"created_at,omitempty"`
	ID         *string `json:"id,omitempty"`
	ProductID  *string `json:"product_id,omitempty"`
	StockLevel *int    `json:"stock_level,omitempty"`
}

// Syrup
type Product struct {
	ID                 *string              `json:"id,omitempty"`
	InventorySnapshots *[]InventorySnapshot `json:"inventory_snapshots,omitempty"`
	Name               *string              `json:"name,omitempty"`
	StockLevel         *int                 `json:"stock_level,omitempty"`
	VendorID           *string              `json:"vendor_id,omitempty"`
}

// Syrup Vendor
type Vendor struct {
	ID         *string    `gorm:"primaryKey;unique;type:uuid;default:uuid_generate_v4();" json:"id,omitempty"`
	Name       *string    `gorm:"not null" json:"name,omitempty"`
	Products   *[]Product `json:"products,omitempty"`
	ScrapeType *string    `gorm:"not null" json:"scrape_type,omitempty"`
}

// BadRequestError defines model for BadRequestError.
type BadRequestError HTTPError

// ForbiddenError defines model for ForbiddenError.
type ForbiddenError HTTPError

// InternalServerError defines model for InternalServerError.
type InternalServerError HTTPError

// NotFoundError defines model for NotFoundError.
type NotFoundError HTTPError

// UnauthorizedError defines model for UnauthorizedError.
type UnauthorizedError HTTPError

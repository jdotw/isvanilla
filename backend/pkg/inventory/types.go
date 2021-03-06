// DO NOT EDIT
// This file was code-generated by github.com/jdotw/oapi-gokit-codegen version (devel)
// It is expected that this file will be re-generated and overwitten to
// adapt to changes in the OpenAPI spec that was used to generate it

package inventory

import (
	_ "embed"
	"time"
)

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Message *string `json:"message,omitempty"`
}

// Point-in-Time Inventory Data
type InventorySnapshot struct {
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ID         *string    `gorm:"primaryKey;unique;type:uuid;default:uuid_generate_v4();" json:"id,omitempty"`
	ProductID  *string    `gorm:"not null" json:"product_id,omitempty"`
	StockLevel *int       `json:"stock_level,omitempty"`
}

// Point-in-Time Inventory Data
type MutateInventorySnapshot struct {
	StockLevel *int `json:"stock_level,omitempty"`
}

// BadRequestError defines model for BadRequestError.
type BadRequestError HTTPError

// ForbiddenError defines model for ForbiddenError.
type ForbiddenError HTTPError

// InternalServerError defines model for InternalServerError.
type InternalServerError HTTPError

// UnauthorizedError defines model for UnauthorizedError.
type UnauthorizedError HTTPError

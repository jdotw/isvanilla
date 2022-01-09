package inventory

import (
	"context"
	_ "embed"
)

type Repository interface {
	GetInventorySnapshots(ctx context.Context, vendorID string, productID string) (*[]InventorySnapshot, error)

	CreateInventorySnapshot(ctx context.Context, vendorID string, productID string, inventorySnapshot *InventorySnapshot) (*InventorySnapshot, error)
}

package vendor

import (
	"context"
	_ "embed"
)

type Repository interface {
	CreateVendor(ctx context.Context, vendor *Vendor) (*Vendor, error)

	DeleteVendor(ctx context.Context, vendorID string) error

	GetVendor(ctx context.Context, vendorID string) (*Vendor, error)

	UpdateVendor(ctx context.Context, vendorID string, vendor *Vendor) (*Vendor, error)

	GetVendors(ctx context.Context) (*[]Vendor, error)
}

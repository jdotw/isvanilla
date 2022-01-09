package product

import (
	"context"
	_ "embed"
)

type Repository interface {
	CreateProduct(ctx context.Context, vendorID string, product *Product) (*Product, error)

	DeleteProduct(ctx context.Context, vendorID string, productID string) error

	GetProduct(ctx context.Context, vendorID string, productID string) (*Product, error)

	UpdateProduct(ctx context.Context, vendorID string, productID string, product *Product) (*Product, error)

	GetProducts(ctx context.Context, vendorID string) (*[]Product, error)
}

package inventory

import (
	"context"
)

type Repository interface {
	GetProduct(ctx context.Context, ProductID string) (*Product, error)
	UpdateProduct(ctx context.Context, ProductID string, Product *Product) (*Product, error)
	CreateProduct(ctx context.Context, profileID string, Product *Product) (*Product, error)
}

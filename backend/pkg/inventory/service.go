package inventory

import (
	"context"
	_ "embed"

	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/syrupstock/pkg/product"
	"github.com/opentracing/opentracing-go"
)

type Service interface {
	GetInventorySnapshots(ctx context.Context, vendorID string, productID string) (*[]InventorySnapshot, error)

	CreateInventorySnapshot(ctx context.Context, vendorID string, productID string, inventorySnapshot *InventorySnapshot) (*InventorySnapshot, error)
}

type service struct {
	repository     Repository
	productService *product.Service
}

func NewService(repository Repository, productService *product.Service, logger log.Factory, tracer opentracing.Tracer) Service {
	var svc Service
	{
		svc = &service{
			repository:     repository,
			productService: productService,
		}
	}
	return svc
}

func (f *service) GetInventorySnapshots(ctx context.Context, vendorID string, productID string) (*[]InventorySnapshot, error) {
	v, err := f.repository.GetInventorySnapshots(ctx, vendorID, productID)
	return v, err
}

func (f *service) CreateInventorySnapshot(ctx context.Context, vendorID string, productID string, inventorySnapshot *InventorySnapshot) (*InventorySnapshot, error) {
	v, err := f.repository.CreateInventorySnapshot(ctx, vendorID, productID, inventorySnapshot)
	if err != nil {
		return v, err
	}

	if f.productService != nil {
		p := product.Product{
			StockLevel: inventorySnapshot.StockLevel,
		}
		ps := *f.productService
		_, err = ps.UpdateProduct(ctx, vendorID, productID, &p)
	}

	return v, err
}

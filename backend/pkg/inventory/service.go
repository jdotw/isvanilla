package inventory

import (
	"context"
	_ "embed"

	"github.com/jdotw/go-utils/log"
	"github.com/opentracing/opentracing-go"
)

type Service interface {
	GetInventorySnapshots(ctx context.Context, vendorID string, productID string) (*[]InventorySnapshot, error)

	CreateInventorySnapshot(ctx context.Context, vendorID string, productID string, inventorySnapshot *InventorySnapshot) (*InventorySnapshot, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository, logger log.Factory, tracer opentracing.Tracer) Service {
	var svc Service
	{
		svc = &service{
			repository: repository,
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
	return v, err
}

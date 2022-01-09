package product

import (
	"context"
	_ "embed"

	"github.com/jdotw/go-utils/log"
	"github.com/opentracing/opentracing-go"
)

type Service interface {
	CreateProduct(ctx context.Context, vendorID string, product *Product) (*Product, error)

	DeleteProduct(ctx context.Context, vendorID string, productID string) error

	GetProduct(ctx context.Context, vendorID string, productID string) (*Product, error)

	UpdateProduct(ctx context.Context, vendorID string, productID string, product *Product) (*Product, error)

	GetProducts(ctx context.Context, vendorID string) (*[]Product, error)
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

func (f *service) CreateProduct(ctx context.Context, vendorID string, product *Product) (*Product, error) {
	v, err := f.repository.CreateProduct(ctx, vendorID, product)
	return v, err
}

func (f *service) DeleteProduct(ctx context.Context, vendorID string, productID string) error {
	v, err := f.repository.DeleteProduct(ctx, vendorID, productID)
	return v, err
}

func (f *service) GetProduct(ctx context.Context, vendorID string, productID string) (*Product, error) {
	v, err := f.repository.GetProduct(ctx, vendorID, productID)
	return v, err
}

func (f *service) UpdateProduct(ctx context.Context, vendorID string, productID string, product *Product) (*Product, error) {
	v, err := f.repository.UpdateProduct(ctx, vendorID, productID, product)
	return v, err
}

func (f *service) GetProducts(ctx context.Context, vendorID string) (*[]Product, error) {
	v, err := f.repository.GetProducts(ctx, vendorID)
	return v, err
}

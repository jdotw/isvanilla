package inventory

import (
	"context"

	"github.com/jdotw/go-utils/log"
	"github.com/opentracing/opentracing-go"
)

type Service interface {
	GetProduct(ctx context.Context, productID string) (*Product, error)
	UpdateProduct(ctx context.Context, productID string, product *Product) (*Product, error)
	CreateProduct(ctx context.Context, productID string, product *Product) (*Product, error)
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

func (f *service) GetProduct(ctx context.Context, productID string) (*Product, error) {
	v, err := f.repository.GetProduct(ctx, productID)
	return v, err
}

func (f *service) UpdateProduct(ctx context.Context, productID string, product *Product) (*Product, error) {
	v, err := f.repository.UpdateProduct(ctx, productID, product)
	return v, err
}

func (f *service) CreateProduct(ctx context.Context, productID string, product *Product) (*Product, error) {
	v, err := f.repository.CreateProduct(ctx, productID, product)
	return v, err
}

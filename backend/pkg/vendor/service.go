package vendor

import (
	"context"
	_ "embed"

	"github.com/jdotw/go-utils/log"
	"github.com/opentracing/opentracing-go"
)

type Service interface {
	CreateVendor(ctx context.Context, vendor *Vendor) (*Vendor, error)

	DeleteVendor(ctx context.Context, vendorID string) error

	GetVendor(ctx context.Context, vendorID string) (*Vendor, error)

	UpdateVendor(ctx context.Context, vendorID string, vendor *Vendor) (*Vendor, error)

	GetVendors(ctx context.Context) (*[]Vendor, error)
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

func (f *service) CreateVendor(ctx context.Context, vendor *Vendor) (*Vendor, error) {
	v, err := f.repository.CreateVendor(ctx, vendor)
	return v, err
}

func (f *service) DeleteVendor(ctx context.Context, vendorID string) error {
	err := f.repository.DeleteVendor(ctx, vendorID)
	return err
}

func (f *service) GetVendor(ctx context.Context, vendorID string) (*Vendor, error) {
	v, err := f.repository.GetVendor(ctx, vendorID)
	return v, err
}

func (f *service) UpdateVendor(ctx context.Context, vendorID string, vendor *Vendor) (*Vendor, error) {
	v, err := f.repository.UpdateVendor(ctx, vendorID, vendor)
	return v, err
}

func (f *service) GetVendors(ctx context.Context) (*[]Vendor, error) {
	v, err := f.repository.GetVendors(ctx)
	return v, err
}

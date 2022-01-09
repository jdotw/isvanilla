package vendor

import (
	"context"
	_ "embed"
	"errors"

	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/go-utils/recorderrors"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

type repository struct {
	db *gorm.DB
}

func NewGormRepository(ctx context.Context, connString string, logger log.Factory, tracer opentracing.Tracer) (Repository, error) {
	var r Repository
	{
		db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
		if err != nil {
			logger.For(ctx).Fatal("Failed to open db", zap.Error(err))
		}

		db.Use(gormopentracing.New(gormopentracing.WithTracer(tracer)))

		// TODO: Ensure these migrations are correct
		// The OpenAPI Spec used to generate this code often uses
		// results in AutoMigrate statements being generated for
		// request/response body objects instead of actual data models

		err = db.AutoMigrate(&Vendor{})
		if err != nil {
			logger.For(ctx).Fatal("Failed to migrate db for type Vendor", zap.Error(err))
		}

		r = &repository{db: db}
	}

	return r, nil
}

func (p *repository) CreateVendor(ctx context.Context, vendor *Vendor) (*Vendor, error) {

	var tx *gorm.DB
	var v Vendor

	tx = p.db.WithContext(ctx).Create(&vendor)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v, nil

}

func (p *repository) DeleteVendor(ctx context.Context, vendorID string) error {

	// TODO: Unable to generate code for this Operation
	return errors.New("Not Implemented")

}

func (p *repository) GetVendor(ctx context.Context, vendorID string) (*Vendor, error) {

	// TODO: Check the .First query as codegen is not able
	// to elegantly deal with multiple request parameters
	var v Vendor
	tx := p.db.WithContext(ctx).Model(&Vendor{}).First(&v, "vendorID = ? ", vendorID)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, recorderrors.ErrNotFound
	}
	return &v, tx.Error

}

func (p *repository) UpdateVendor(ctx context.Context, vendorID string, vendor *Vendor) (*Vendor, error) {

	// TODO: Check the .Where queries as codegen is not able
	// to elegantly deal with multiple request parameters
	var v Vendor

	tx := p.db.WithContext(ctx).Model(&Vendor{}).Where("vendorID = ?", vendorID).UpdateColumns(vendor)
	if tx.RowsAffected == 0 {
		return nil, recorderrors.ErrNotFound
	}

	return &v, tx.Error

}

func (p *repository) GetVendors(ctx context.Context) (*[]Vendor, error) {
	var v []Vendor
	tx := p.db.WithContext(ctx).Find(&v)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, recorderrors.ErrNotFound
	}
	return &v, tx.Error

}

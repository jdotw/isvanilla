package product

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

		err = db.AutoMigrate(&Product{})
		if err != nil {
			logger.For(ctx).Fatal("Failed to migrate db for type Product", zap.Error(err))
		}

		r = &repository{db: db}
	}

	return r, nil
}

func (p *repository) CreateProduct(ctx context.Context, vendorID string, product *Product) (*Product, error) {

	var tx *gorm.DB
	var v Product

	tx = p.db.WithContext(ctx).Create(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v, nil

}

func (p *repository) DeleteProduct(ctx context.Context, vendorID string, productID string) error {

	// TODO: Unable to generate code for this Operation
	return errors.New("Not Implemented")

}

func (p *repository) GetProduct(ctx context.Context, vendorID string, productID string) (*Product, error) {

	// TODO: Check the .First query as codegen is not able
	// to elegantly deal with multiple request parameters
	var v Product
	tx := p.db.WithContext(ctx).Model(&Product{}).First(&v, "vendorID = ? productID = ? ", vendorID, productID)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, recorderrors.ErrNotFound
	}
	return &v, tx.Error

}

func (p *repository) UpdateProduct(ctx context.Context, vendorID string, productID string, product *Product) (*Product, error) {

	// TODO: Check the .Where queries as codegen is not able
	// to elegantly deal with multiple request parameters
	var v Product

	tx := p.db.WithContext(ctx).Model(&Product{}).Where("vendorID = ?", vendorID).Where("productID = ?", productID).UpdateColumns(product)
	if tx.RowsAffected == 0 {
		return nil, recorderrors.ErrNotFound
	}

	return &v, tx.Error

}

func (p *repository) GetProducts(ctx context.Context, vendorID string) (*[]Product, error) {

	// TODO: Check the .First query as codegen is not able
	// to elegantly deal with multiple request parameters
	var v []Product
	tx := p.db.WithContext(ctx).Model(&[]Product{}).First(&v, "vendorID = ? ", vendorID)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, recorderrors.ErrNotFound
	}
	return &v, tx.Error

}

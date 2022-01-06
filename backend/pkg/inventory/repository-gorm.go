package inventory

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"

	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/go-utils/recorderrors"
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

		err = db.AutoMigrate(&Vendor{})
		if err != nil {
			logger.For(ctx).Fatal("Failed to migrate db for type Vendor", zap.Error(err))
		}
		err = db.AutoMigrate(&Product{})
		if err != nil {
			logger.For(ctx).Fatal("Failed to migrate db for type Product", zap.Error(err))
		}
		err = db.AutoMigrate(&InventorySnapshot{})
		if err != nil {
			logger.For(ctx).Fatal("Failed to migrate db for type InventorySnapshot", zap.Error(err))
		}

		r = &repository{db: db}
	}

	return r, nil
}

func (p *repository) GetProduct(ctx context.Context, productID string) (*Product, error) {
	var v Product
	tx := p.db.WithContext(ctx).Model(&Product{}).First(&v, "id = ?", productID)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, recorderrors.ErrNotFound
	}
	return &v, tx.Error
}

func (p *repository) UpdateProduct(ctx context.Context, productID string, product *Product) (*Product, error) {
	var v Product

	tx := p.db.WithContext(ctx).Model(&Product{}).Where("id = ?", productID).UpdateColumns(product)
	if tx.RowsAffected == 0 {
		return nil, recorderrors.ErrNotFound
	}

	return &v, tx.Error

}

func (p *repository) CreateProduct(ctx context.Context, profileID string, product *Product) (*Product, error) {
	var tx *gorm.DB

	tx = p.db.WithContext(ctx).Create(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}

package inventory

import (
	"context"
	_ "embed"

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

		err = db.AutoMigrate(&InventorySnapshot{})
		if err != nil {
			logger.For(ctx).Fatal("Failed to migrate db for type InventorySnapshot", zap.Error(err))
		}

		r = &repository{db: db}
	}

	return r, nil
}

func (p *repository) GetInventorySnapshots(ctx context.Context, vendorID string, productID string) (*[]InventorySnapshot, error) {

	// TODO: Check the .First query as codegen is not able
	// to elegantly deal with multiple request parameters
	var v []InventorySnapshot
	tx := p.db.WithContext(ctx).Model(&[]InventorySnapshot{}).First(&v, "vendorID = ? productID = ? ", vendorID, productID)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, recorderrors.ErrNotFound
	}
	return &v, tx.Error

}

func (p *repository) CreateInventorySnapshot(ctx context.Context, vendorID string, productID string, inventorySnapshot *InventorySnapshot) (*InventorySnapshot, error) {

	var tx *gorm.DB
	var v InventorySnapshot

	tx = p.db.WithContext(ctx).Create(&inventorySnapshot)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &v, nil

}

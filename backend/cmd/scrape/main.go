package main

import (
	"context"
	"os"
	"time"

	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/go-utils/tracing"
	"github.com/jdotw/syrupstock/pkg/inventory"
	"github.com/jdotw/syrupstock/pkg/product"
	"github.com/jdotw/syrupstock/pkg/scraper"
	"github.com/jdotw/syrupstock/pkg/vendor"
	"go.uber.org/zap"
)

func main() {
	serviceName := "scrape"

	// Logging and Tracing
	logger, metricsFactory := log.Init(serviceName)
	tracer := tracing.Init(serviceName, metricsFactory, logger)

	inventoryURL := os.Getenv("INVENTORY_URL")
	if len(inventoryURL) == 0 {
		inventoryURL = "http://localhost:8080"
	}

	vc, err := vendor.NewClientWithResponses(inventoryURL)
	if err != nil {
		logger.Bg().Fatal("failed to create vendor client", zap.Error(err))
	}

	pc, err := product.NewClientWithResponses(inventoryURL)
	if err != nil {
		logger.Bg().Fatal("failed to create product client", zap.Error(err))
	}

	ic, err := inventory.NewClientWithResponses(inventoryURL)
	if err != nil {
		logger.Bg().Fatal("failed to create inventory client", zap.Error(err))
	}

	ctx := context.Background()

	var loop bool
	if os.Getenv("LOOP") == "1" {
		loop = true
	} else {
		loop = false
	}

	for {
		start := time.Now()

		vs, err := vc.GetVendorsWithResponse(ctx)
		if err != nil {
			logger.Bg().Fatal("failed to get vendors", zap.Error(err))
		}

		for i := 0; i < len(*vs.JSON200); i++ {
			v := (*vs.JSON200)[i]
			logger.Bg().Info("scraping vendor", zap.String("name", *v.Name))

			ps, err := pc.GetProductsWithResponse(ctx, *v.ID)
			if err != nil {
				logger.Bg().Fatal("failed to get products", zap.Error(err))
			}

			s := scraper.NewScraper(logger, tracer)

			for i := 0; i < len(*ps.JSON200); i++ {
				p := (*ps.JSON200)[i]
				logger.Bg().Info("product", zap.String("name", *p.Name))

				l := s.ScrapeStockLevel(ic, p)
				logger.Bg().Info("stock", zap.String("name", *p.Name), zap.Int("level", l))
			}
		}

		end := time.Now()
		duration := end.Sub(start)
		logger.Bg().Info("scraped", zap.Duration("duration", duration))

		if !loop {
			logger.Bg().Info("exit", zap.Bool("loop", loop))
			os.Exit(0)
		}

		d := os.Getenv("LOOP_DELAY")
		if len(d) > 0 {
			duration, err := time.ParseDuration(d)
			if err != nil {
				logger.Bg().Fatal("failed to parse loop delay duration", zap.Error(err))
			}
			logger.Bg().Info("pause", zap.Duration("duration", duration))
			time.Sleep(duration)
		}
	}
}

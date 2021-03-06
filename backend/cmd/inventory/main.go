package main

import (
	"context"
	_ "embed"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/go-utils/tracing"
	"github.com/jdotw/syrupstock/pkg/inventory"
	"github.com/jdotw/syrupstock/pkg/product"
	"github.com/jdotw/syrupstock/pkg/vendor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	serviceName := "inventory"

	// Logging and Tracing
	logger, metricsFactory := log.Init(serviceName)
	tracer := tracing.Init(serviceName, metricsFactory, logger)

	// HTTP Router
	r := mux.NewRouter()

	// Craft PostgreSQL DSN
	dsn := os.Getenv("POSTGRES_DSN")
	if len(dsn) == 0 {
		dsn = "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USERNAME") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=isvanilla" + " port=" + os.Getenv("POSTGRES_PORT")
	}

	logger.Bg().Info("DB", zap.String("POSTGRES_HOST", os.Getenv("POSTGRES_HOST")))

	// Vendor Service
	{
		repo, err := vendor.NewGormRepository(context.Background(), dsn, logger, tracer)
		if err != nil {
			logger.Bg().Fatal("Failed to create vendor repository", zap.Error(err))
		}
		service := vendor.NewService(repo, logger, tracer)
		endPoints := vendor.NewEndpointSet(service, logger, tracer)
		vendor.AddHTTPRoutes(r, endPoints, logger, tracer)
	}

	// Product Service
	var productService *product.Service
	{
		repo, err := product.NewGormRepository(context.Background(), dsn, logger, tracer)
		if err != nil {
			logger.Bg().Fatal("Failed to create product repository", zap.Error(err))
		}
		service := product.NewService(repo, logger, tracer)
		endPoints := product.NewEndpointSet(service, logger, tracer)
		product.AddHTTPRoutes(r, endPoints, logger, tracer)
		productService = &service
	}

	// Inventory Service
	{
		repo, err := inventory.NewGormRepository(context.Background(), dsn, logger, tracer)
		if err != nil {
			logger.Bg().Fatal("Failed to create inventory repository", zap.Error(err))
		}
		service := inventory.NewService(repo, productService, logger, tracer)
		endPoints := inventory.NewEndpointSet(service, logger, tracer)
		inventory.AddHTTPRoutes(r, endPoints, logger, tracer)
	}

	// HTTP Mux
	m := tracing.NewServeMux(tracer)
	m.Handle("/metrics", promhttp.Handler()) // Prometheus
	m.Handle("/", r)

	// Start Transports
	go func() error {
		// HTTP
		httpHost := os.Getenv("HTTP_LISTEN_HOST")
		httpPort := os.Getenv("HTTP_LISTEN_PORT")
		if len(httpPort) == 0 {
			httpPort = "8080"
		}
		httpAddr := httpHost + ":" + httpPort
		logger.Bg().Info("Listening", zap.String("transport", "http"), zap.String("host", httpHost), zap.String("port", httpPort), zap.String("addr", httpAddr))
		err := http.ListenAndServe(httpAddr, m)
		logger.Bg().Fatal("Exit", zap.Error(err))
		return err
	}()

	// Select Loop
	select {}
}

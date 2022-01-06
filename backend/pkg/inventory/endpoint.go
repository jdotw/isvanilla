package inventory

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kittracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	// _ "embed"

	"github.com/jdotw/go-utils/log"
)

type EndpointSet struct {
	GetProductEndpoint    endpoint.Endpoint
	UpdateProductEndpoint endpoint.Endpoint
	CreateProductEndpoint endpoint.Endpoint
}

// //go:embed policies/endpoint.rego
// var endpointPolicy string

func NewEndpointSet(s Service, logger log.Factory, tracer opentracing.Tracer) EndpointSet {
	// authn := jwt.NewAuthenticator(logger, tracer)
	// authz := opa.NewAuthorizor(logger, tracer)

	var getProductEndpoint endpoint.Endpoint
	{
		getProductEndpoint = makeGetProductEndpoint(s, logger, tracer)
		// getProductEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.get_product")(getProductEndpoint)
		// getProductEndpoint = authn.NewMiddleware()(getProductEndpoint)
		getProductEndpoint = kittracing.TraceServer(tracer, "GetProductEndpoint")(getProductEndpoint)
	}
	var updateProductEndpoint endpoint.Endpoint
	{
		updateProductEndpoint = makeUpdateProductEndpoint(s, logger, tracer)
		// updateProductEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.update_product")(updateProductEndpoint)
		// updateProductEndpoint = authn.NewMiddleware()(updateProductEndpoint)
		updateProductEndpoint = kittracing.TraceServer(tracer, "UpdateProductEndpoint")(updateProductEndpoint)
	}
	var createProductEndpoint endpoint.Endpoint
	{
		createProductEndpoint = makeCreateProductEndpoint(s, logger, tracer)
		// createProductEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.create_product")(createProductEndpoint)
		// createProductEndpoint = authn.NewMiddleware()(createProductEndpoint)
		createProductEndpoint = kittracing.TraceServer(tracer, "CreateProductEndpoint")(createProductEndpoint)
	}
	return EndpointSet{
		GetProductEndpoint:    getProductEndpoint,
		UpdateProductEndpoint: updateProductEndpoint,
		CreateProductEndpoint: createProductEndpoint,
	}
}

// GetProduct

type GetProductEndpointRequest struct {
	ProductID string
}

func makeGetProductEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.GetProductEndpoint received request")
		req := request.(GetProductEndpointRequest)
		v, err := s.GetProduct(ctx, req.ProductID)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// UpdateProduct

type UpdateProductEndpointRequest struct {
	ProductID string

	Product *Product
}

func makeUpdateProductEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.UpdateProductEndpoint received request")
		req := request.(UpdateProductEndpointRequest)
		v, err := s.UpdateProduct(ctx, req.ProductID, req.Product)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// CreateProduct

type CreateProductEndpointRequest struct {
	ProductID string

	Product *Product
}

func makeCreateProductEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.CreateProductEndpoint received request")
		req := request.(CreateProductEndpointRequest)
		v, err := s.CreateProduct(ctx, req.ProductID, req.Product)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

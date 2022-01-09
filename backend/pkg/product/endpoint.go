package product

import (
	"context"
	_ "embed"

	"github.com/go-kit/kit/endpoint"
	kittracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/jdotw/go-utils/authn/jwt"
	"github.com/jdotw/go-utils/authz/opa"
	"github.com/jdotw/go-utils/log"
	"github.com/opentracing/opentracing-go"
)

type EndpointSet struct {
	CreateProductEndpoint endpoint.Endpoint
	DeleteProductEndpoint endpoint.Endpoint
	GetProductEndpoint    endpoint.Endpoint
	UpdateProductEndpoint endpoint.Endpoint
	GetProductsEndpoint   endpoint.Endpoint
}

//go:embed policies/endpoint.rego
var endpointPolicy string

func NewEndpointSet(s Service, logger log.Factory, tracer opentracing.Tracer) EndpointSet {
	authn := jwt.NewAuthenticator(logger, tracer)
	authz := opa.NewAuthorizor(logger, tracer)

	var createProductEndpoint endpoint.Endpoint
	{
		createProductEndpoint = makeCreateProductEndpoint(s, logger, tracer)
		createProductEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.create_product")(createProductEndpoint)
		createProductEndpoint = authn.NewMiddleware()(createProductEndpoint)
		createProductEndpoint = kittracing.TraceServer(tracer, "CreateProductEndpoint")(createProductEndpoint)
	}
	var deleteProductEndpoint endpoint.Endpoint
	{
		deleteProductEndpoint = makeDeleteProductEndpoint(s, logger, tracer)
		deleteProductEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.delete_product")(deleteProductEndpoint)
		deleteProductEndpoint = authn.NewMiddleware()(deleteProductEndpoint)
		deleteProductEndpoint = kittracing.TraceServer(tracer, "DeleteProductEndpoint")(deleteProductEndpoint)
	}
	var getProductEndpoint endpoint.Endpoint
	{
		getProductEndpoint = makeGetProductEndpoint(s, logger, tracer)
		getProductEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.get_product")(getProductEndpoint)
		getProductEndpoint = authn.NewMiddleware()(getProductEndpoint)
		getProductEndpoint = kittracing.TraceServer(tracer, "GetProductEndpoint")(getProductEndpoint)
	}
	var updateProductEndpoint endpoint.Endpoint
	{
		updateProductEndpoint = makeUpdateProductEndpoint(s, logger, tracer)
		updateProductEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.update_product")(updateProductEndpoint)
		updateProductEndpoint = authn.NewMiddleware()(updateProductEndpoint)
		updateProductEndpoint = kittracing.TraceServer(tracer, "UpdateProductEndpoint")(updateProductEndpoint)
	}
	var getProductsEndpoint endpoint.Endpoint
	{
		getProductsEndpoint = makeGetProductsEndpoint(s, logger, tracer)
		getProductsEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.product.endpoint.authz.get_products")(getProductsEndpoint)
		getProductsEndpoint = authn.NewMiddleware()(getProductsEndpoint)
		getProductsEndpoint = kittracing.TraceServer(tracer, "GetProductsEndpoint")(getProductsEndpoint)
	}
	return EndpointSet{
		CreateProductEndpoint: createProductEndpoint,
		DeleteProductEndpoint: deleteProductEndpoint,
		GetProductEndpoint:    getProductEndpoint,
		UpdateProductEndpoint: updateProductEndpoint,
		GetProductsEndpoint:   getProductsEndpoint,
	}
}

// CreateProduct

type CreateProductEndpointRequest struct {
	VendorID             string
	MutateProductRequest *MutateProductRequest
}

func makeCreateProductEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.CreateProductEndpoint received request")
		req := request.(CreateProductEndpointRequest)
		c := Product{
			Name:       req.MutateProductRequest.Name,
			StockLevel: req.MutateProductRequest.StockLevel,
		}
		v, err := s.CreateProduct(ctx, req.VendorID, &c)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// DeleteProduct

type DeleteProductEndpointRequest struct {
	VendorID  string
	ProductID string
}

func makeDeleteProductEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.DeleteProductEndpoint received request")
		req := request.(DeleteProductEndpointRequest)
		err := s.DeleteProduct(ctx, req.VendorID, req.ProductID)
		return nil, err
	}
}

// GetProduct

type GetProductEndpointRequest struct {
	VendorID  string
	ProductID string
}

func makeGetProductEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.GetProductEndpoint received request")
		req := request.(GetProductEndpointRequest)
		v, err := s.GetProduct(ctx, req.VendorID, req.ProductID)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// UpdateProduct

type UpdateProductEndpointRequest struct {
	VendorID  string
	ProductID string

	Product *Product
}

func makeUpdateProductEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.UpdateProductEndpoint received request")
		req := request.(UpdateProductEndpointRequest)
		v, err := s.UpdateProduct(ctx, req.VendorID, req.ProductID, req.Product)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// GetProducts

type GetProductsEndpointRequest struct {
	VendorID string
}

func makeGetProductsEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Product.GetProductsEndpoint received request")
		req := request.(GetProductsEndpointRequest)
		v, err := s.GetProducts(ctx, req.VendorID)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

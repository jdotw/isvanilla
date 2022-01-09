package vendor

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
	CreateVendorEndpoint endpoint.Endpoint
	DeleteVendorEndpoint endpoint.Endpoint
	GetVendorEndpoint    endpoint.Endpoint
	UpdateVendorEndpoint endpoint.Endpoint
	GetVendorsEndpoint   endpoint.Endpoint
}

//go:embed policies/endpoint.rego
var endpointPolicy string

func NewEndpointSet(s Service, logger log.Factory, tracer opentracing.Tracer) EndpointSet {
	authn := jwt.NewAuthenticator(logger, tracer)
	authz := opa.NewAuthorizor(logger, tracer)

	var createVendorEndpoint endpoint.Endpoint
	{
		createVendorEndpoint = makeCreateVendorEndpoint(s, logger, tracer)
		createVendorEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.vendor.endpoint.authz.create_vendor")(createVendorEndpoint)
		createVendorEndpoint = authn.NewMiddleware()(createVendorEndpoint)
		createVendorEndpoint = kittracing.TraceServer(tracer, "CreateVendorEndpoint")(createVendorEndpoint)
	}
	var deleteVendorEndpoint endpoint.Endpoint
	{
		deleteVendorEndpoint = makeDeleteVendorEndpoint(s, logger, tracer)
		deleteVendorEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.vendor.endpoint.authz.delete_vendor")(deleteVendorEndpoint)
		deleteVendorEndpoint = authn.NewMiddleware()(deleteVendorEndpoint)
		deleteVendorEndpoint = kittracing.TraceServer(tracer, "DeleteVendorEndpoint")(deleteVendorEndpoint)
	}
	var getVendorEndpoint endpoint.Endpoint
	{
		getVendorEndpoint = makeGetVendorEndpoint(s, logger, tracer)
		getVendorEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.vendor.endpoint.authz.get_vendor")(getVendorEndpoint)
		// getVendorEndpoint = authn.NewMiddleware()(getVendorEndpoint)
		getVendorEndpoint = kittracing.TraceServer(tracer, "GetVendorEndpoint")(getVendorEndpoint)
	}
	var updateVendorEndpoint endpoint.Endpoint
	{
		updateVendorEndpoint = makeUpdateVendorEndpoint(s, logger, tracer)
		updateVendorEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.vendor.endpoint.authz.update_vendor")(updateVendorEndpoint)
		updateVendorEndpoint = authn.NewMiddleware()(updateVendorEndpoint)
		updateVendorEndpoint = kittracing.TraceServer(tracer, "UpdateVendorEndpoint")(updateVendorEndpoint)
	}
	var getVendorsEndpoint endpoint.Endpoint
	{
		getVendorsEndpoint = makeGetVendorsEndpoint(s, logger, tracer)
		getVendorsEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.vendor.endpoint.authz.get_vendors")(getVendorsEndpoint)
		// getVendorsEndpoint = authn.NewMiddleware()(getVendorsEndpoint)
		getVendorsEndpoint = kittracing.TraceServer(tracer, "GetVendorsEndpoint")(getVendorsEndpoint)
	}
	return EndpointSet{
		CreateVendorEndpoint: createVendorEndpoint,
		DeleteVendorEndpoint: deleteVendorEndpoint,
		GetVendorEndpoint:    getVendorEndpoint,
		UpdateVendorEndpoint: updateVendorEndpoint,
		GetVendorsEndpoint:   getVendorsEndpoint,
	}
}

// CreateVendor

type CreateVendorEndpointRequest struct {
	Vendor *Vendor
}

func makeCreateVendorEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Vendor.CreateVendorEndpoint received request")
		req := request.(CreateVendorEndpointRequest)
		v, err := s.CreateVendor(ctx, req.Vendor)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// DeleteVendor

type DeleteVendorEndpointRequest struct {
	VendorID string
}

func makeDeleteVendorEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Vendor.DeleteVendorEndpoint received request")
		req := request.(DeleteVendorEndpointRequest)
		err := s.DeleteVendor(ctx, req.VendorID)
		return nil, err
	}
}

// GetVendor

type GetVendorEndpointRequest struct {
	VendorID string
}

func makeGetVendorEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Vendor.GetVendorEndpoint received request")
		req := request.(GetVendorEndpointRequest)
		v, err := s.GetVendor(ctx, req.VendorID)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// UpdateVendor

type UpdateVendorEndpointRequest struct {
	VendorID string

	Vendor *Vendor
}

func makeUpdateVendorEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Vendor.UpdateVendorEndpoint received request")
		req := request.(UpdateVendorEndpointRequest)
		v, err := s.UpdateVendor(ctx, req.VendorID, req.Vendor)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// GetVendors

func makeGetVendorsEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Vendor.GetVendorsEndpoint received request")
		v, err := s.GetVendors(ctx)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

package inventory

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
	GetInventorySnapshotsEndpoint   endpoint.Endpoint
	CreateInventorySnapshotEndpoint endpoint.Endpoint
}

//go:embed policies/endpoint.rego
var endpointPolicy string

func NewEndpointSet(s Service, logger log.Factory, tracer opentracing.Tracer) EndpointSet {
	authn := jwt.NewAuthenticator(logger, tracer)
	authz := opa.NewAuthorizor(logger, tracer)

	var getInventorySnapshotsEndpoint endpoint.Endpoint
	{
		getInventorySnapshotsEndpoint = makeGetInventorySnapshotsEndpoint(s, logger, tracer)
		getInventorySnapshotsEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.inventory.endpoint.authz.get_inventory_snapshots")(getInventorySnapshotsEndpoint)
		getInventorySnapshotsEndpoint = authn.NewMiddleware()(getInventorySnapshotsEndpoint)
		getInventorySnapshotsEndpoint = kittracing.TraceServer(tracer, "GetInventorySnapshotsEndpoint")(getInventorySnapshotsEndpoint)
	}
	var createInventorySnapshotEndpoint endpoint.Endpoint
	{
		createInventorySnapshotEndpoint = makeCreateInventorySnapshotEndpoint(s, logger, tracer)
		createInventorySnapshotEndpoint = authz.NewInProcessMiddleware(endpointPolicy, "data.inventory.endpoint.authz.create_inventory_snapshot")(createInventorySnapshotEndpoint)
		createInventorySnapshotEndpoint = authn.NewMiddleware()(createInventorySnapshotEndpoint)
		createInventorySnapshotEndpoint = kittracing.TraceServer(tracer, "CreateInventorySnapshotEndpoint")(createInventorySnapshotEndpoint)
	}
	return EndpointSet{
		GetInventorySnapshotsEndpoint:   getInventorySnapshotsEndpoint,
		CreateInventorySnapshotEndpoint: createInventorySnapshotEndpoint,
	}
}

// GetInventorySnapshots

type GetInventorySnapshotsEndpointRequest struct {
	VendorID  string
	ProductID string
}

func makeGetInventorySnapshotsEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Inventory.GetInventorySnapshotsEndpoint received request")
		req := request.(GetInventorySnapshotsEndpointRequest)
		v, err := s.GetInventorySnapshots(ctx, req.VendorID, req.ProductID)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

// CreateInventorySnapshot

type CreateInventorySnapshotEndpointRequest struct {
	VendorID  string
	ProductID string

	InventorySnapshot *InventorySnapshot
}

func makeCreateInventorySnapshotEndpoint(s Service, logger log.Factory, tracer opentracing.Tracer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.For(ctx).Info("Inventory.CreateInventorySnapshotEndpoint received request")
		req := request.(CreateInventorySnapshotEndpointRequest)
		v, err := s.CreateInventorySnapshot(ctx, req.VendorID, req.ProductID, req.InventorySnapshot)
		if err != nil {
			return &v, err
		}
		return &v, nil
	}
}

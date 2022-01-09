package inventory

import (
	"context"
	_ "embed"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/jdotw/go-utils/authn/jwt"
	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/go-utils/transport"
	"github.com/opentracing/opentracing-go"
)

func AddHTTPRoutes(r *mux.Router, endpoints EndpointSet, logger log.Factory, tracer opentracing.Tracer) {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(transport.HTTPErrorEncoder),
		httptransport.ServerBefore(jwt.HTTPAuthorizationToContext()),
	}

	getInventorySnapshotsHandler := httptransport.NewServer(
		endpoints.GetInventorySnapshotsEndpoint,
		decodeGetInventorySnapshotsEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}/product/{product_id}/inventory", getInventorySnapshotsHandler).Methods("GET")

	createInventorySnapshotHandler := httptransport.NewServer(
		endpoints.CreateInventorySnapshotEndpoint,
		decodeCreateInventorySnapshotEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}/product/{product_id}/inventory", createInventorySnapshotHandler).Methods("POST")

}

// GetInventorySnapshots

func decodeGetInventorySnapshotsEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)

	request := GetInventorySnapshotsEndpointRequest{
		VendorID:  vars["vendor_id"],
		ProductID: vars["product_id"],
	}
	return request, nil
}

// CreateInventorySnapshot

func decodeCreateInventorySnapshotEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var inventorySnapshot MutateInventorySnapshot

	if err := json.NewDecoder(r.Body).Decode(&inventorySnapshot); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)

	request := CreateInventorySnapshotEndpointRequest{
		VendorID:          vars["vendor_id"],
		ProductID:         vars["product_id"],
		InventorySnapshot: &inventorySnapshot,
	}
	return request, nil
}

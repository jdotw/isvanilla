package vendor

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

	createVendorHandler := httptransport.NewServer(
		endpoints.CreateVendorEndpoint,
		decodeCreateVendorEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor", createVendorHandler).Methods("POST")

	deleteVendorHandler := httptransport.NewServer(
		endpoints.DeleteVendorEndpoint,
		decodeDeleteVendorEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}", deleteVendorHandler).Methods("DELETE")

	getVendorHandler := httptransport.NewServer(
		endpoints.GetVendorEndpoint,
		decodeGetVendorEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}", getVendorHandler).Methods("GET")

	updateVendorHandler := httptransport.NewServer(
		endpoints.UpdateVendorEndpoint,
		decodeUpdateVendorEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}", updateVendorHandler).Methods("PATCH")

	getVendorsHandler := httptransport.NewServer(
		endpoints.GetVendorsEndpoint,
		decodeGetVendorsEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendors", getVendorsHandler).Methods("GET")

}

// CreateVendor

func decodeCreateVendorEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var vendor Vendor

	if err := json.NewDecoder(r.Body).Decode(&vendor); err != nil {
		return nil, err
	}

	request := CreateVendorEndpointRequest{
		Vendor: &vendor,
	}
	return request, nil
}

// DeleteVendor

func decodeDeleteVendorEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)

	request := DeleteVendorEndpointRequest{
		VendorID: vars["vendor_id"],
	}
	return request, nil
}

// GetVendor

func decodeGetVendorEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)

	request := GetVendorEndpointRequest{
		VendorID: vars["vendor_id"],
	}
	return request, nil
}

// UpdateVendor

func decodeUpdateVendorEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var vendor Vendor

	if err := json.NewDecoder(r.Body).Decode(&vendor); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)

	request := UpdateVendorEndpointRequest{
		VendorID: vars["vendor_id"],
		Vendor:   &vendor,
	}
	return request, nil
}

// GetVendors

func decodeGetVendorsEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

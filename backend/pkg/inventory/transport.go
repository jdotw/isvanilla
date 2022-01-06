package inventory

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	"github.com/jdotw/go-utils/log"
	"github.com/jdotw/go-utils/transport"
)

func AddHTTPRoutes(r *mux.Router, endpoints EndpointSet, logger log.Factory, tracer opentracing.Tracer) {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(transport.HTTPErrorEncoder),
		// httptransport.ServerBefore(jwt.HTTPAuthorizationToContext()),
	}

	getProductHandler := httptransport.NewServer(
		endpoints.GetProductEndpoint,
		decodeGetProductEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/product/{product_id}", getProductHandler).Methods("GET")

	updateProductHandler := httptransport.NewServer(
		endpoints.UpdateProductEndpoint,
		decodeUpdateProductEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/product/{product_id}", updateProductHandler).Methods("PATCH")

	createProductHandler := httptransport.NewServer(
		endpoints.CreateProductEndpoint,
		decodeCreateProductEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/profile/{profile_id}/product", createProductHandler).Methods("POST")

}

// GetProduct

func decodeGetProductEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	request := GetProductEndpointRequest{
		ProductID: vars["product_id"],
	}
	return request, nil
}

// UpdateProduct

func decodeUpdateProductEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var product Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)

	request := UpdateProductEndpointRequest{
		ProductID: vars["product_id"],
		Product:   &product,
	}

	return request, nil
}

// CreateProduct

func decodeCreateProductEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var product Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)

	request := CreateProductEndpointRequest{
		ProductID: vars["product_id"],
		Product:   &product,
	}
	return request, nil
}

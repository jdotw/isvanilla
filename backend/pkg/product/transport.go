package product

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

	createProductHandler := httptransport.NewServer(
		endpoints.CreateProductEndpoint,
		decodeCreateProductEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}/product", createProductHandler).Methods("POST")

	deleteProductHandler := httptransport.NewServer(
		endpoints.DeleteProductEndpoint,
		decodeDeleteProductEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}/product/{product_id}", deleteProductHandler).Methods("DELETE")

	getProductHandler := httptransport.NewServer(
		endpoints.GetProductEndpoint,
		decodeGetProductEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}/product/{product_id}", getProductHandler).Methods("GET")

	updateProductHandler := httptransport.NewServer(
		endpoints.UpdateProductEndpoint,
		decodeUpdateProductEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}/product/{product_id}", updateProductHandler).Methods("PATCH")

	getProductsHandler := httptransport.NewServer(
		endpoints.GetProductsEndpoint,
		decodeGetProductsEndpointRequest,
		transport.HTTPEncodeResponse,
		options...,
	)
	r.Handle("/vendor/{vendor_id}/products", getProductsHandler).Methods("GET")

}

// CreateProduct

func decodeCreateProductEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var product MutateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	request := CreateProductEndpointRequest{
		VendorID:             vars["vendor_id"],
		MutateProductRequest: &product,
	}
	return request, nil
}

// DeleteProduct

func decodeDeleteProductEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)

	request := DeleteProductEndpointRequest{
		VendorID:  vars["vendor_id"],
		ProductID: vars["product_id"],
	}
	return request, nil
}

// GetProduct

func decodeGetProductEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)

	request := GetProductEndpointRequest{
		VendorID:  vars["vendor_id"],
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
		VendorID:  vars["vendor_id"],
		ProductID: vars["product_id"],
		Product:   &product,
	}
	return request, nil
}

// GetProducts

func decodeGetProductsEndpointRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)

	request := GetProductsEndpointRequest{
		VendorID: vars["vendor_id"],
	}
	return request, nil
}

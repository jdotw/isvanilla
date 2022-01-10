// DO NOT EDIT
// This file was code-generated by github.com/jdotw/oapi-gokit-codegen version (devel)
// It is expected that this file will be re-generated and overwitten to
// adapt to changes in the OpenAPI spec that was used to generate it

package vendor

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jdotw/oapi-gokit-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateVendor request with any body
	CreateVendorWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateVendor(ctx context.Context, createVendorRequest CreateVendorRequest, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteVendor request
	DeleteVendor(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVendor request
	GetVendor(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateVendor request with any body
	UpdateVendorWithBody(ctx context.Context, vendorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateVendor(ctx context.Context, vendorID string, vendor Vendor, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetVendors request
	GetVendors(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateVendorWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateVendorRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateVendor(ctx context.Context, createVendorRequest CreateVendorRequest, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateVendorRequest(c.Server, createVendorRequest)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteVendor(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteVendorRequest(c.Server, vendorID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetVendor(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVendorRequest(c.Server, vendorID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateVendorWithBody(ctx context.Context, vendorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateVendorRequestWithBody(c.Server, vendorID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateVendor(ctx context.Context, vendorID string, vendor Vendor, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateVendorRequest(c.Server, vendorID, vendor)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetVendors(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetVendorsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateVendorRequest calls the generic CreateVendor builder with application/json body
func NewCreateVendorRequest(server string, createVendorRequest CreateVendorRequest) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(createVendorRequest)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateVendorRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateVendorRequestWithBody generates requests for CreateVendor with any type of body
func NewCreateVendorRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vendor")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteVendorRequest generates requests for DeleteVendor
func NewDeleteVendorRequest(server string, vendorID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "vendor_id", runtime.ParamLocationPath, vendorID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vendor/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVendorRequest generates requests for GetVendor
func NewGetVendorRequest(server string, vendorID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "vendor_id", runtime.ParamLocationPath, vendorID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vendor/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateVendorRequest calls the generic UpdateVendor builder with application/json body
func NewUpdateVendorRequest(server string, vendorID string, vendor Vendor) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(vendor)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateVendorRequestWithBody(server, vendorID, "application/json", bodyReader)
}

// NewUpdateVendorRequestWithBody generates requests for UpdateVendor with any type of body
func NewUpdateVendorRequestWithBody(server string, vendorID string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "vendor_id", runtime.ParamLocationPath, vendorID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vendor/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetVendorsRequest generates requests for GetVendors
func NewGetVendorsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/vendors")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateVendor request with any body
	CreateVendorWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateVendorResponse, error)

	CreateVendorWithResponse(ctx context.Context, createVendorRequest CreateVendorRequest, reqEditors ...RequestEditorFn) (*CreateVendorResponse, error)

	// DeleteVendor request
	DeleteVendorWithResponse(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*DeleteVendorResponse, error)

	// GetVendor request
	GetVendorWithResponse(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*GetVendorResponse, error)

	// UpdateVendor request with any body
	UpdateVendorWithBodyWithResponse(ctx context.Context, vendorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateVendorResponse, error)

	UpdateVendorWithResponse(ctx context.Context, vendorID string, vendor Vendor, reqEditors ...RequestEditorFn) (*UpdateVendorResponse, error)

	// GetVendors request
	GetVendorsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetVendorsResponse, error)
}

type CreateVendorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Vendor
	JSON400      *HTTPError
	JSON401      *HTTPError
	JSON403      *HTTPError
	JSON500      *HTTPError
}

// Status returns HTTPResponse.Status
func (r CreateVendorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateVendorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteVendorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *HTTPError
	JSON401      *HTTPError
	JSON403      *HTTPError
	JSON404      *HTTPError
	JSON500      *HTTPError
}

// Status returns HTTPResponse.Status
func (r DeleteVendorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteVendorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetVendorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Vendor
	JSON400      *HTTPError
	JSON404      *HTTPError
	JSON500      *HTTPError
}

// Status returns HTTPResponse.Status
func (r GetVendorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVendorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateVendorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Vendor
	JSON400      *HTTPError
	JSON401      *HTTPError
	JSON403      *HTTPError
	JSON404      *HTTPError
	JSON500      *HTTPError
}

// Status returns HTTPResponse.Status
func (r UpdateVendorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateVendorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetVendorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Vendor
	JSON400      *HTTPError
	JSON500      *HTTPError
}

// Status returns HTTPResponse.Status
func (r GetVendorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVendorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateVendorWithBodyWithResponse request with arbitrary body returning *CreateVendorResponse
func (c *ClientWithResponses) CreateVendorWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateVendorResponse, error) {
	rsp, err := c.CreateVendorWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateVendorResponse(rsp)
}

func (c *ClientWithResponses) CreateVendorWithResponse(ctx context.Context, createVendorRequest CreateVendorRequest, reqEditors ...RequestEditorFn) (*CreateVendorResponse, error) {
	rsp, err := c.CreateVendor(ctx, createVendorRequest, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateVendorResponse(rsp)
}

// DeleteVendorWithResponse request returning *DeleteVendorResponse
func (c *ClientWithResponses) DeleteVendorWithResponse(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*DeleteVendorResponse, error) {
	rsp, err := c.DeleteVendor(ctx, vendorID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteVendorResponse(rsp)
}

// GetVendorWithResponse request returning *GetVendorResponse
func (c *ClientWithResponses) GetVendorWithResponse(ctx context.Context, vendorID string, reqEditors ...RequestEditorFn) (*GetVendorResponse, error) {
	rsp, err := c.GetVendor(ctx, vendorID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVendorResponse(rsp)
}

// UpdateVendorWithBodyWithResponse request with arbitrary body returning *UpdateVendorResponse
func (c *ClientWithResponses) UpdateVendorWithBodyWithResponse(ctx context.Context, vendorID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateVendorResponse, error) {
	rsp, err := c.UpdateVendorWithBody(ctx, vendorID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateVendorResponse(rsp)
}

func (c *ClientWithResponses) UpdateVendorWithResponse(ctx context.Context, vendorID string, vendor Vendor, reqEditors ...RequestEditorFn) (*UpdateVendorResponse, error) {
	rsp, err := c.UpdateVendor(ctx, vendorID, vendor, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateVendorResponse(rsp)
}

// GetVendorsWithResponse request returning *GetVendorsResponse
func (c *ClientWithResponses) GetVendorsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetVendorsResponse, error) {
	rsp, err := c.GetVendors(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetVendorsResponse(rsp)
}

// ParseCreateVendorResponse parses an HTTP response from a CreateVendorWithResponse call
func ParseCreateVendorResponse(rsp *http.Response) (*CreateVendorResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateVendorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Vendor
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteVendorResponse parses an HTTP response from a DeleteVendorWithResponse call
func ParseDeleteVendorResponse(rsp *http.Response) (*DeleteVendorResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteVendorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetVendorResponse parses an HTTP response from a GetVendorWithResponse call
func ParseGetVendorResponse(rsp *http.Response) (*GetVendorResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVendorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Vendor
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseUpdateVendorResponse parses an HTTP response from a UpdateVendorWithResponse call
func ParseUpdateVendorResponse(rsp *http.Response) (*UpdateVendorResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateVendorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Vendor
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetVendorsResponse parses an HTTP response from a GetVendorsWithResponse call
func ParseGetVendorsResponse(rsp *http.Response) (*GetVendorsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetVendorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	println("Content-type: ", rsp.Header.Get("Content-Type"))

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Vendor
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HTTPError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

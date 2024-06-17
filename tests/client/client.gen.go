// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// ClerkWebhookEmailAddress defines model for ClerkWebhookEmailAddress.
type ClerkWebhookEmailAddress struct {
	// EmailAddress User's email address
	EmailAddress string `json:"email_address"`

	// Id Unique identifier for the email address
	Id string `json:"id"`

	// LinkedTo (Array is empty for this event)
	LinkedTo *[]map[string]interface{} `json:"linked_to"`

	// Object Object type (always "email_address" for this event)
	Object       *string `json:"object,omitempty"`
	Verification *struct {
		// Status Verification status (e.g., "verified", "unverified")
		Status *string `json:"status,omitempty"`

		// Strategy Verification strategy (e.g., "ticket", "link")
		Strategy *string `json:"strategy,omitempty"`
	} `json:"verification,omitempty"`
}

// ClerkWebhookUserCreatedData defines model for ClerkWebhookUserCreatedData.
type ClerkWebhookUserCreatedData struct {
	// Birthday User's birthday (empty string if not set)
	Birthday *string `json:"birthday,omitempty"`

	// CreatedAt Timestamp (epoch milliseconds) representing user creation time
	CreatedAt      int                        `json:"created_at"`
	EmailAddresses []ClerkWebhookEmailAddress `json:"email_addresses"`

	// ExternalAccounts (Array is empty for this event)
	ExternalAccounts *[]map[string]interface{} `json:"external_accounts,omitempty"`

	// ExternalId User's external identifier
	ExternalId *string `json:"external_id"`

	// FirstName User's first name
	FirstName *string `json:"first_name"`

	// Gender User's gender (empty string if not set)
	Gender *string `json:"gender,omitempty"`

	// Id Unique identifier for the user
	Id string `json:"id"`

	// ImageUrl User's image URL (may be redacted)
	ImageUrl *string `json:"image_url,omitempty"`

	// LastName User's last name
	LastName *string `json:"last_name"`

	// LastSignInAt Timestamp (epoch milliseconds) representing last sign-in time
	LastSignInAt *int `json:"last_sign_in_at"`

	// Object Object type (always "user" for this event)
	Object *string `json:"object,omitempty"`

	// PasswordEnabled Whether the user has password authentication enabled
	PasswordEnabled bool `json:"password_enabled"`

	// PhoneNumbers (Array is empty for this event)
	PhoneNumbers *[]map[string]interface{} `json:"phone_numbers,omitempty"`

	// PrimaryEmailAddressId Unique identifier for the primary email address
	PrimaryEmailAddressId *string `json:"primary_email_address_id"`

	// PrimaryPhoneNumberId Unique identifier for the primary phone number (null if not set)
	PrimaryPhoneNumberId *string `json:"primary_phone_number_id"`

	// PrimaryWeb3WalletId Unique identifier for the primary web3 wallet (null if not set)
	PrimaryWeb3WalletId *string `json:"primary_web3_wallet_id"`

	// PrivateMetadata User's private metadata (empty object for this event)
	PrivateMetadata *map[string]interface{} `json:"private_metadata,omitempty"`

	// ProfileImageUrl User's profile image URL (may be redacted)
	ProfileImageUrl *string `json:"profile_image_url,omitempty"`

	// PublicMetadata User's public metadata (empty object for this event)
	PublicMetadata *map[string]interface{} `json:"public_metadata,omitempty"`

	// TwoFactorEnabled Whether two-factor authentication is enabled
	TwoFactorEnabled bool `json:"two_factor_enabled"`

	// UnsafeMetadata User's unsafe metadata (empty object for this event)
	UnsafeMetadata *map[string]interface{} `json:"unsafe_metadata,omitempty"`

	// UpdatedAt Timestamp (epoch milliseconds) representing user update time
	UpdatedAt *int `json:"updated_at,omitempty"`

	// Username Username (null if not set)
	Username *string `json:"username"`

	// Web3Wallets (Array is empty for this event)
	Web3Wallets *[]map[string]interface{} `json:"web3_wallets,omitempty"`
}

// CreateAccountRequest defines model for CreateAccountRequest.
type CreateAccountRequest struct {
	Data ClerkWebhookUserCreatedData `json:"data"`

	// Object Event type (always "user.created" for this event)
	Object string `json:"object"`

	// Type Event type (always "user.created" for this event)
	Type string `json:"type"`
}

// CreateApplicationRequest defines model for CreateApplicationRequest.
type CreateApplicationRequest struct {
	Name string `json:"name"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error custom error code such as 'email_in_use'
	Error string `json:"error"`

	// Message A description about the error
	Message string `json:"message"`
}

// DefaultError defines model for DefaultError.
type DefaultError = ErrorResponse

// CreateAccountParams defines parameters for CreateAccount.
type CreateAccountParams struct {
	SvixId        string `json:"svix-id"`
	SvixTimestamp string `json:"svix-timestamp"`
	SvixSignature string `json:"svix-signature"`
}

// CreateApplicationJSONRequestBody defines body for CreateApplication for application/json ContentType.
type CreateApplicationJSONRequestBody = CreateApplicationRequest

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody = CreateAccountRequest

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
	// CreateApplicationWithBody request with any body
	CreateApplicationWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateApplication(ctx context.Context, body CreateApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// HealthCheck request
	HealthCheck(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateAccountWithBody request with any body
	CreateAccountWithBody(ctx context.Context, params *CreateAccountParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateAccount(ctx context.Context, params *CreateAccountParams, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateApplicationWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateApplicationRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateApplication(ctx context.Context, body CreateApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateApplicationRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) HealthCheck(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewHealthCheckRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAccountWithBody(ctx context.Context, params *CreateAccountParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAccountRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateAccount(ctx context.Context, params *CreateAccountParams, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateAccountRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateApplicationRequest calls the generic CreateApplication builder with application/json body
func NewCreateApplicationRequest(server string, body CreateApplicationJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateApplicationRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateApplicationRequestWithBody generates requests for CreateApplication with any type of body
func NewCreateApplicationRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/applications")
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

// NewHealthCheckRequest generates requests for HealthCheck
func NewHealthCheckRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/health")
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

// NewCreateAccountRequest calls the generic CreateAccount builder with application/json body
func NewCreateAccountRequest(server string, params *CreateAccountParams, body CreateAccountJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateAccountRequestWithBody(server, params, "application/json", bodyReader)
}

// NewCreateAccountRequestWithBody generates requests for CreateAccount with any type of body
func NewCreateAccountRequestWithBody(server string, params *CreateAccountParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/webhooks/account")
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

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "svix-id", runtime.ParamLocationHeader, params.SvixId)
		if err != nil {
			return nil, err
		}

		req.Header.Set("svix-id", headerParam0)

		var headerParam1 string

		headerParam1, err = runtime.StyleParamWithLocation("simple", false, "svix-timestamp", runtime.ParamLocationHeader, params.SvixTimestamp)
		if err != nil {
			return nil, err
		}

		req.Header.Set("svix-timestamp", headerParam1)

		var headerParam2 string

		headerParam2, err = runtime.StyleParamWithLocation("simple", false, "svix-signature", runtime.ParamLocationHeader, params.SvixSignature)
		if err != nil {
			return nil, err
		}

		req.Header.Set("svix-signature", headerParam2)

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
	// CreateApplicationWithBodyWithResponse request with any body
	CreateApplicationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateApplicationResponse, error)

	CreateApplicationWithResponse(ctx context.Context, body CreateApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateApplicationResponse, error)

	// HealthCheckWithResponse request
	HealthCheckWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*HealthCheckResponse, error)

	// CreateAccountWithBodyWithResponse request with any body
	CreateAccountWithBodyWithResponse(ctx context.Context, params *CreateAccountParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error)

	CreateAccountWithResponse(ctx context.Context, params *CreateAccountParams, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error)
}

type CreateApplicationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *DefaultError
}

// Status returns HTTPResponse.Status
func (r CreateApplicationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateApplicationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type HealthCheckResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *DefaultError
}

// Status returns HTTPResponse.Status
func (r HealthCheckResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r HealthCheckResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateAccountResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *DefaultError
}

// Status returns HTTPResponse.Status
func (r CreateAccountResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAccountResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateApplicationWithBodyWithResponse request with arbitrary body returning *CreateApplicationResponse
func (c *ClientWithResponses) CreateApplicationWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateApplicationResponse, error) {
	rsp, err := c.CreateApplicationWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateApplicationResponse(rsp)
}

func (c *ClientWithResponses) CreateApplicationWithResponse(ctx context.Context, body CreateApplicationJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateApplicationResponse, error) {
	rsp, err := c.CreateApplication(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateApplicationResponse(rsp)
}

// HealthCheckWithResponse request returning *HealthCheckResponse
func (c *ClientWithResponses) HealthCheckWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*HealthCheckResponse, error) {
	rsp, err := c.HealthCheck(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseHealthCheckResponse(rsp)
}

// CreateAccountWithBodyWithResponse request with arbitrary body returning *CreateAccountResponse
func (c *ClientWithResponses) CreateAccountWithBodyWithResponse(ctx context.Context, params *CreateAccountParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error) {
	rsp, err := c.CreateAccountWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAccountResponse(rsp)
}

func (c *ClientWithResponses) CreateAccountWithResponse(ctx context.Context, params *CreateAccountParams, body CreateAccountJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateAccountResponse, error) {
	rsp, err := c.CreateAccount(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateAccountResponse(rsp)
}

// ParseCreateApplicationResponse parses an HTTP response from a CreateApplicationWithResponse call
func ParseCreateApplicationResponse(rsp *http.Response) (*CreateApplicationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateApplicationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest DefaultError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseHealthCheckResponse parses an HTTP response from a HealthCheckWithResponse call
func ParseHealthCheckResponse(rsp *http.Response) (*HealthCheckResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &HealthCheckResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest DefaultError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseCreateAccountResponse parses an HTTP response from a CreateAccountWithResponse call
func ParseCreateAccountResponse(rsp *http.Response) (*CreateAccountResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAccountResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest DefaultError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

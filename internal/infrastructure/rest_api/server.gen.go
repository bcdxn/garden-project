//go:build go1.22

// Package rest_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package rest_api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
)

// ListPermissionsResponse defines model for ListPermissionsResponse.
type ListPermissionsResponse = []map[string]RBACPermission

// ListRolesResponse defines model for ListRolesResponse.
type ListRolesResponse = []map[string]RBACRole

// RBACAction defines model for RBACAction.
type RBACAction struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// RBACPermission defines model for RBACPermission.
type RBACPermission struct {
	Action   *map[string]RBACAction   `json:"action,omitempty"`
	Resource *map[string]RBACResource `json:"resource,omitempty"`
}

// RBACResource defines model for RBACResource.
type RBACResource struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// RBACRole defines model for RBACRole.
type RBACRole struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/v1/roles)
	GetApiV1Roles(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/roles/{roleId}/permissions)
	GetApiV1RolesRoleIdPermissions(w http.ResponseWriter, r *http.Request, roleId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetApiV1Roles operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Roles(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Roles(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetApiV1RolesRoleIdPermissions operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RolesRoleIdPermissions(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "roleId" -------------
	var roleId string

	err = runtime.BindStyledParameterWithOptions("simple", "roleId", r.PathValue("roleId"), &roleId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "roleId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RolesRoleIdPermissions(w, r, roleId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/api/v1/roles", wrapper.GetApiV1Roles)
	m.HandleFunc("GET "+options.BaseURL+"/api/v1/roles/{roleId}/permissions", wrapper.GetApiV1RolesRoleIdPermissions)

	return m
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RUwW7bMAz9FYHb0ajc7VL4lu0wBNghyGGXoSg0mUnY2ZJGMQGCIP8+UE4Wp8nWYWtz",
	"iG2Keu/xidQOfOxTDBgkQ7OD7FfYu/L6mbLMkHvKmWLIc8wphoy6RIJ9yXFtS0IxuG7GMSELYYm/ZVxA",
	"A2/sCd0eoO38w+TjCRf2Fcg2ITQQvz2il1HAMbutfquUeezwZUUo4l/Ra/LEK4XCpjMSavX/sCMLU1jq",
	"luB6vLJwje2JHxcM7hfzvxZ60H6NnDHHNXv8Lx+PGJcEv6t3PqJ9DT/L0b44toYoLKImt5g9UxqOBh5m",
	"nQvGhdYIO//dbOOazdJxi+EBKhAS1QOfSsTMOCqgmcymUMEGeTh4uL2pb2pVFBMGlwgaeF9CFSQnq1KD",
	"dYns5tayjoMGliiXepYoxnWdUS9MSTUUjKzQ5G0W7KFwsNP0aavKUCaJvtyWKYPSF2XQCsW7utaHj0Ew",
	"FDaXUke+bLePeejOoSWea5jLUd4Pv+q8NLvTx7Td23S6g56td3RfmUVk40xO6GlBvrjw57LnhXAE8do+",
	"XLtdj2Ykx65HQc7QfN0BaZXaBHDsVxj8KRp/rImxhUZ4jdVIwNOOvlfkjLw5wp672EXvOtPiBruYegxi",
	"MGyIY9B3qGDNHTSwEkmNtSV5FbM0d/VdDfv7/c8AAAD//2QKVrxMBgAA",
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

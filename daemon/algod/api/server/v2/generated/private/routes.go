// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

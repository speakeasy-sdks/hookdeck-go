// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck/pkg/models/shared"
	"net/http"
)

type UpdateSourceRequestBody struct {
	// List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.
	AllowedHTTPMethods []shared.SourceAllowedHTTPMethod `json:"allowed_http_methods,omitempty"`
	// Custom response object
	CustomResponse *shared.SourceCustomResponse `json:"custom_response,omitempty"`
	// A unique name for the source
	Name *string `json:"name,omitempty"`
}

type UpdateSourceRequest struct {
	RequestBody UpdateSourceRequestBody `request:"mediaType=application/json"`
	ID          string                  `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateSourceResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single source
	Source      *shared.Source
	StatusCode  int
	RawResponse *http.Response
}
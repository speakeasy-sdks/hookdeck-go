// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type CreateSourceRequestBody struct {
	// List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.
	AllowedHTTPMethods []shared.SourceAllowedHTTPMethod `json:"allowed_http_methods,omitempty"`
	// Custom response object
	CustomResponse *shared.SourceCustomResponse `json:"custom_response,omitempty"`
	// A unique name for the source
	Name string `json:"name"`
}

func (o *CreateSourceRequestBody) GetAllowedHTTPMethods() []shared.SourceAllowedHTTPMethod {
	if o == nil {
		return nil
	}
	return o.AllowedHTTPMethods
}

func (o *CreateSourceRequestBody) GetCustomResponse() *shared.SourceCustomResponse {
	if o == nil {
		return nil
	}
	return o.CustomResponse
}

func (o *CreateSourceRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type CreateSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single source
	Source *shared.Source
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceResponse) GetSource() *shared.Source {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *CreateSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

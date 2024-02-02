// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"net/http"
)

type GetIntegrationRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetIntegrationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetIntegrationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single integration
	Integration *components.Integration
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIntegrationResponse) GetIntegration() *components.Integration {
	if o == nil {
		return nil
	}
	return o.Integration
}

func (o *GetIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

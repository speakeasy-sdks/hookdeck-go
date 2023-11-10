// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type GetIntegrationsRequest struct {
	Label *string `queryParam:"style=form,explode=true,name=label"`
	// The provider name
	Provider *components.IntegrationProvider `queryParam:"style=form,explode=true,name=provider"`
}

func (o *GetIntegrationsRequest) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *GetIntegrationsRequest) GetProvider() *components.IntegrationProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

type GetIntegrationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of integrations
	IntegrationPaginatedResult *components.IntegrationPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIntegrationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIntegrationsResponse) GetIntegrationPaginatedResult() *components.IntegrationPaginatedResult {
	if o == nil {
		return nil
	}
	return o.IntegrationPaginatedResult
}

func (o *GetIntegrationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIntegrationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
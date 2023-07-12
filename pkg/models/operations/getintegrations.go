// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetIntegrationsRequest struct {
	Label *string `queryParam:"style=form,explode=true,name=label"`
	// The provider name
	Provider *shared.IntegrationProvider `queryParam:"style=form,explode=true,name=provider"`
}

type GetIntegrationsResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// List of integrations
	IntegrationPaginatedResult *shared.IntegrationPaginatedResult
	StatusCode                 int
	RawResponse                *http.Response
}

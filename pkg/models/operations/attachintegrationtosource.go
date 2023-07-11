// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck/pkg/models/shared"
	"net/http"
)

type AttachIntegrationToSourceRequest struct {
	ID       string `pathParam:"style=simple,explode=false,name=id"`
	SourceID string `pathParam:"style=simple,explode=false,name=source_id"`
}

type AttachIntegrationToSourceResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	// Attach operation success status
	AttachedIntegrationToSource *shared.AttachedIntegrationToSource
	ContentType                 string
	StatusCode                  int
	RawResponse                 *http.Response
}
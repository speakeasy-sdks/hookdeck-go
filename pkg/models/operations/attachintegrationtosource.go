// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type AttachIntegrationToSourceRequest struct {
	ID       string `pathParam:"style=simple,explode=false,name=id"`
	SourceID string `pathParam:"style=simple,explode=false,name=source_id"`
}

func (o *AttachIntegrationToSourceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AttachIntegrationToSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type AttachIntegrationToSourceResponse struct {
	// Attach operation success status
	AttachedIntegrationToSource *shared.AttachedIntegrationToSource
	ContentType                 string
	StatusCode                  int
	RawResponse                 *http.Response
}

func (o *AttachIntegrationToSourceResponse) GetAttachedIntegrationToSource() *shared.AttachedIntegrationToSource {
	if o == nil {
		return nil
	}
	return o.AttachedIntegrationToSource
}

func (o *AttachIntegrationToSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AttachIntegrationToSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AttachIntegrationToSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

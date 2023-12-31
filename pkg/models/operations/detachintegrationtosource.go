// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type DetachIntegrationToSourceRequest struct {
	ID       string `pathParam:"style=simple,explode=false,name=id"`
	SourceID string `pathParam:"style=simple,explode=false,name=source_id"`
}

func (o *DetachIntegrationToSourceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DetachIntegrationToSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type DetachIntegrationToSourceResponse struct {
	ContentType string
	// Detach operation success status
	DetachedIntegrationFromSource *shared.DetachedIntegrationFromSource
	StatusCode                    int
	RawResponse                   *http.Response
}

func (o *DetachIntegrationToSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DetachIntegrationToSourceResponse) GetDetachedIntegrationFromSource() *shared.DetachedIntegrationFromSource {
	if o == nil {
		return nil
	}
	return o.DetachedIntegrationFromSource
}

func (o *DetachIntegrationToSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DetachIntegrationToSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

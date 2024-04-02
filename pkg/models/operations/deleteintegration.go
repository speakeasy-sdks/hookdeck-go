// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type DeleteIntegrationRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteIntegrationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteIntegrationResponse struct {
	ContentType string
	// An object with deleted integration id
	DeletedIntegration *shared.DeletedIntegration
	StatusCode         int
	RawResponse        *http.Response
}

func (o *DeleteIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteIntegrationResponse) GetDeletedIntegration() *shared.DeletedIntegration {
	if o == nil {
		return nil
	}
	return o.DeletedIntegration
}

func (o *DeleteIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
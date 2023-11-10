// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type UnarchiveConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UnarchiveConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UnarchiveConnectionResponse struct {
	// A single connection
	Connection *components.Connection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UnarchiveConnectionResponse) GetConnection() *components.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UnarchiveConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UnarchiveConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UnarchiveConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
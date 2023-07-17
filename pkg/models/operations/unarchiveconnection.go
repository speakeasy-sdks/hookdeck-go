// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UnarchiveConnectionResponse) GetConnection() *shared.Connection {
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

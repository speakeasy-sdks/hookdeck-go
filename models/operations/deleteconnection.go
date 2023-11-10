// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteConnectionResponseBody - A single connection
type DeleteConnectionResponseBody struct {
	// ID of the connection
	ID string `json:"id"`
}

func (o *DeleteConnectionResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A single connection
	Object *DeleteConnectionResponseBody
}

func (o *DeleteConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteConnectionResponse) GetObject() *DeleteConnectionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
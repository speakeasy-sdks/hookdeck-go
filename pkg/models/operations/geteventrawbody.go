// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetEventRawBodyRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetEventRawBodyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetEventRawBodyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A request raw body data
	RawBody *shared.RawBody
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEventRawBodyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEventRawBodyResponse) GetRawBody() *shared.RawBody {
	if o == nil {
		return nil
	}
	return o.RawBody
}

func (o *GetEventRawBodyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEventRawBodyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

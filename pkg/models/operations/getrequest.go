// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetRequestRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetRequestRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRequestResponse struct {
	ContentType string
	// A single request
	Request     *shared.Request
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetRequestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRequestResponse) GetRequest() *shared.Request {
	if o == nil {
		return nil
	}
	return o.Request
}

func (o *GetRequestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRequestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

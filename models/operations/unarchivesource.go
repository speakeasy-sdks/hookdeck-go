// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type UnarchiveSourceRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UnarchiveSourceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UnarchiveSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single source
	Source *components.Source
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UnarchiveSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UnarchiveSourceResponse) GetSource() *components.Source {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *UnarchiveSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UnarchiveSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
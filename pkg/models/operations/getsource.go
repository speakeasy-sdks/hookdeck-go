// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetSourceRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSourceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSourceResponse struct {
	ContentType string
	// A single source
	Source      *shared.Source
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSourceResponse) GetSource() *shared.Source {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *GetSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
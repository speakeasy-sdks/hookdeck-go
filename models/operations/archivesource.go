// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"net/http"
)

type ArchiveSourceRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ArchiveSourceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ArchiveSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single source
	Source *components.Source
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ArchiveSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ArchiveSourceResponse) GetSource() *components.Source {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *ArchiveSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ArchiveSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

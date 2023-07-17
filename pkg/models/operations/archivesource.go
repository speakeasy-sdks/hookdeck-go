// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
	ContentType string
	// A single source
	Source      *shared.Source
	StatusCode  int
	RawResponse *http.Response
}

func (o *ArchiveSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ArchiveSourceResponse) GetSource() *shared.Source {
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

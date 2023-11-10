// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type GetBookmarkRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetBookmarkRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetBookmarkResponse struct {
	// A single bookmark
	Bookmark *components.Bookmark
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBookmarkResponse) GetBookmark() *components.Bookmark {
	if o == nil {
		return nil
	}
	return o.Bookmark
}

func (o *GetBookmarkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBookmarkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBookmarkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
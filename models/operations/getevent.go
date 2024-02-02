// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"net/http"
)

type GetEventRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetEventRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetEventResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single event
	Event *components.Event
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEventResponse) GetEvent() *components.Event {
	if o == nil {
		return nil
	}
	return o.Event
}

func (o *GetEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

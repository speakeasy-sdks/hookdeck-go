// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"net/http"
)

type GetIssueTriggerRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetIssueTriggerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetIssueTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single issue trigger
	IssueTrigger *components.IssueTrigger
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIssueTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIssueTriggerResponse) GetIssueTrigger() *components.IssueTrigger {
	if o == nil {
		return nil
	}
	return o.IssueTrigger
}

func (o *GetIssueTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIssueTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

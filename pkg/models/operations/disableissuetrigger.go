// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type DisableIssueTriggerRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DisableIssueTriggerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DisableIssueTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single issue trigger
	IssueTrigger *shared.IssueTrigger
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DisableIssueTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DisableIssueTriggerResponse) GetIssueTrigger() *shared.IssueTrigger {
	if o == nil {
		return nil
	}
	return o.IssueTrigger
}

func (o *DisableIssueTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DisableIssueTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

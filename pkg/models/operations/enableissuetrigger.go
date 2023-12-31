// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type EnableIssueTriggerRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *EnableIssueTriggerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type EnableIssueTriggerResponse struct {
	ContentType string
	// A single issue trigger
	IssueTrigger *shared.IssueTrigger
	StatusCode   int
	RawResponse  *http.Response
}

func (o *EnableIssueTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EnableIssueTriggerResponse) GetIssueTrigger() *shared.IssueTrigger {
	if o == nil {
		return nil
	}
	return o.IssueTrigger
}

func (o *EnableIssueTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EnableIssueTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

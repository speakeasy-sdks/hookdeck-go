// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type DeleteIssueTriggerRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteIssueTriggerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteIssueTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// An object with deleted issue trigger's id
	DeletedIssueTriggerResponse *components.DeletedIssueTriggerResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteIssueTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteIssueTriggerResponse) GetDeletedIssueTriggerResponse() *components.DeletedIssueTriggerResponse {
	if o == nil {
		return nil
	}
	return o.DeletedIssueTriggerResponse
}

func (o *DeleteIssueTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteIssueTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
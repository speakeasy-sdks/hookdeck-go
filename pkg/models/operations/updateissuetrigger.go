// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck/pkg/models/shared"
	"net/http"
	"time"
)

type UpdateIssueTriggerRequestBody struct {
	// Notification channels object for the specific channel type
	Channels *shared.IssueTriggerChannels `json:"channels,omitempty"`
	// Configuration object for the specific issue type selected
	Configs interface{} `json:"configs,omitempty"`
	// Date when the issue trigger was disabled
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	// Optional unique name to use as reference when using the API
	Name *string `json:"name,omitempty"`
}

type UpdateIssueTriggerRequest struct {
	RequestBody UpdateIssueTriggerRequestBody `request:"mediaType=application/json"`
	ID          string                        `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateIssueTriggerResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single issue trigger
	IssueTrigger *shared.IssueTrigger
	StatusCode   int
	RawResponse  *http.Response
}
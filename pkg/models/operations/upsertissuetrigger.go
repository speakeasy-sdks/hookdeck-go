// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type UpsertIssueTriggerRequestBody struct {
	// Notification channels object for the specific channel type
	Channels shared.IssueTriggerChannels `json:"channels"`
	// Configuration object for the specific issue type selected
	Configs interface{} `json:"configs,omitempty"`
	// Required unique name to use as reference when using the API
	Name string `json:"name"`
	// Issue type
	Type shared.IssueType `json:"type"`
}

type UpsertIssueTriggerResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single issue trigger
	IssueTrigger *shared.IssueTrigger
	StatusCode   int
	RawResponse  *http.Response
}

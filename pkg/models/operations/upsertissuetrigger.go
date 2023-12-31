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

func (o *UpsertIssueTriggerRequestBody) GetChannels() shared.IssueTriggerChannels {
	if o == nil {
		return shared.IssueTriggerChannels{}
	}
	return o.Channels
}

func (o *UpsertIssueTriggerRequestBody) GetConfigs() interface{} {
	if o == nil {
		return nil
	}
	return o.Configs
}

func (o *UpsertIssueTriggerRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpsertIssueTriggerRequestBody) GetType() shared.IssueType {
	if o == nil {
		return shared.IssueType("")
	}
	return o.Type
}

type UpsertIssueTriggerResponse struct {
	ContentType string
	// A single issue trigger
	IssueTrigger *shared.IssueTrigger
	StatusCode   int
	RawResponse  *http.Response
}

func (o *UpsertIssueTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertIssueTriggerResponse) GetIssueTrigger() *shared.IssueTrigger {
	if o == nil {
		return nil
	}
	return o.IssueTrigger
}

func (o *UpsertIssueTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertIssueTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

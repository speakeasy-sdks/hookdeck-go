// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// IssueTrigger - A single issue trigger
type IssueTrigger struct {
	// Notification channels object for the specific channel type
	Channels *IssueTriggerChannels `json:"channels,omitempty"`
	// Configuration object for the specific issue type selected
	Configs interface{} `json:"configs"`
	// ISO timestamp for when the issue trigger was created
	CreatedAt time.Time `json:"created_at"`
	// ISO timestamp for when the issue trigger was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// ISO timestamp for when the issue trigger was disabled
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	// ID of the issue trigger
	ID string `json:"id"`
	// Optional unique name to use as reference when using the API
	Name *string `json:"name,omitempty"`
	// ID of the workspace
	TeamID *string `json:"team_id,omitempty"`
	// Issue type
	Type IssueType `json:"type"`
	// ISO timestamp for when the issue trigger was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

func (o *IssueTrigger) GetChannels() *IssueTriggerChannels {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *IssueTrigger) GetConfigs() interface{} {
	if o == nil {
		return nil
	}
	return o.Configs
}

func (o *IssueTrigger) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *IssueTrigger) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *IssueTrigger) GetDisabledAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DisabledAt
}

func (o *IssueTrigger) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *IssueTrigger) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IssueTrigger) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *IssueTrigger) GetType() IssueType {
	if o == nil {
		return IssueType("")
	}
	return o.Type
}

func (o *IssueTrigger) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

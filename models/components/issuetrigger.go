// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"time"
)

type IssueTrigger struct {
	// Notification channels object for the specific channel type
	Channels *IssueTriggerChannels `json:"channels,omitempty"`
	// Configuration object for the specific issue type selected
	Configs IssueTriggerReference `json:"configs"`
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

func (i IssueTrigger) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IssueTrigger) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IssueTrigger) GetChannels() *IssueTriggerChannels {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *IssueTrigger) GetConfigs() IssueTriggerReference {
	if o == nil {
		return IssueTriggerReference{}
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
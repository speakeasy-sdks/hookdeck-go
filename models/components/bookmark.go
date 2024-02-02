// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/hookdeck-go/v2/internal/utils"
	"time"
)

type Bookmark struct {
	// Alternate alias for the bookmark
	Alias *string `json:"alias,omitempty"`
	// Date the bookmark was created
	CreatedAt time.Time       `json:"created_at"`
	Data      *ShortEventData `json:"data,omitempty"`
	// ID of the bookmarked event data
	EventDataID string `json:"event_data_id"`
	// ID of the bookmark
	ID string `json:"id"`
	// Descriptive name of the bookmark
	Label string `json:"label"`
	// Date the bookmark was last manually triggered
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// ID of the workspace
	TeamID string `json:"team_id"`
	// Date the bookmark was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// ID of the associated connection
	WebhookID string `json:"webhook_id"`
}

func (b Bookmark) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *Bookmark) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Bookmark) GetAlias() *string {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *Bookmark) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Bookmark) GetData() *ShortEventData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Bookmark) GetEventDataID() string {
	if o == nil {
		return ""
	}
	return o.EventDataID
}

func (o *Bookmark) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Bookmark) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *Bookmark) GetLastUsedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUsedAt
}

func (o *Bookmark) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *Bookmark) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Bookmark) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

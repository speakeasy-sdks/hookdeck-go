// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type IgnoredEvent struct {
	Cause     IgnoredEventCause `json:"cause"`
	CreatedAt time.Time         `json:"created_at"`
	ID        string            `json:"id"`
	Meta      interface{}       `json:"meta,omitempty"`
	RequestID string            `json:"request_id"`
	TeamID    string            `json:"team_id"`
	UpdatedAt time.Time         `json:"updated_at"`
	WebhookID string            `json:"webhook_id"`
}

func (o *IgnoredEvent) GetCause() IgnoredEventCause {
	if o == nil {
		return IgnoredEventCause("")
	}
	return o.Cause
}

func (o *IgnoredEvent) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *IgnoredEvent) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *IgnoredEvent) GetMeta() interface{} {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *IgnoredEvent) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

func (o *IgnoredEvent) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *IgnoredEvent) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *IgnoredEvent) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type DeliveryIssueWithDataType string

const (
	DeliveryIssueWithDataTypeDelivery DeliveryIssueWithDataType = "delivery"
)

func (e DeliveryIssueWithDataType) ToPointer() *DeliveryIssueWithDataType {
	return &e
}

func (e *DeliveryIssueWithDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delivery":
		*e = DeliveryIssueWithDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeliveryIssueWithDataType: %v", v)
	}
}

// DeliveryIssueWithData - Delivery issue
type DeliveryIssueWithData struct {
	// Keys used as the aggregation keys a 'delivery' type issue
	AggregationKeys DeliveryIssueAggregationKeys `json:"aggregation_keys"`
	AutoResolvedAt  *time.Time                   `json:"auto_resolved_at,omitempty"`
	// ISO timestamp for when the issue was created
	CreatedAt string `json:"created_at"`
	// Delivery issue data
	Data *DeliveryIssueData `json:"data,omitempty"`
	// ISO timestamp for when the issue was dismissed
	DismissedAt *time.Time `json:"dismissed_at,omitempty"`
	// ISO timestamp for when the issue was first opened
	FirstSeenAt time.Time `json:"first_seen_at"`
	// Issue ID
	ID string `json:"id"`
	// ISO timestamp for when the issue last occured
	LastSeenAt time.Time `json:"last_seen_at"`
	// ID of the team member who last updated the issue status
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	MergedWith    *string `json:"merged_with,omitempty"`
	// ISO timestamp for when the issue was last opened
	OpenedAt time.Time `json:"opened_at"`
	// Reference to the event and attempt an issue is being created for.
	Reference DeliveryIssueReference `json:"reference"`
	// Issue status
	Status IssueStatus `json:"status"`
	// ID of the workspace
	TeamID string                    `json:"team_id"`
	Type   DeliveryIssueWithDataType `json:"type"`
	// ISO timestamp for when the issue was last updated
	UpdatedAt string `json:"updated_at"`
}

func (o *DeliveryIssueWithData) GetAggregationKeys() DeliveryIssueAggregationKeys {
	if o == nil {
		return DeliveryIssueAggregationKeys{}
	}
	return o.AggregationKeys
}

func (o *DeliveryIssueWithData) GetAutoResolvedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AutoResolvedAt
}

func (o *DeliveryIssueWithData) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *DeliveryIssueWithData) GetData() *DeliveryIssueData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DeliveryIssueWithData) GetDismissedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DismissedAt
}

func (o *DeliveryIssueWithData) GetFirstSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.FirstSeenAt
}

func (o *DeliveryIssueWithData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeliveryIssueWithData) GetLastSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastSeenAt
}

func (o *DeliveryIssueWithData) GetLastUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedBy
}

func (o *DeliveryIssueWithData) GetMergedWith() *string {
	if o == nil {
		return nil
	}
	return o.MergedWith
}

func (o *DeliveryIssueWithData) GetOpenedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.OpenedAt
}

func (o *DeliveryIssueWithData) GetReference() DeliveryIssueReference {
	if o == nil {
		return DeliveryIssueReference{}
	}
	return o.Reference
}

func (o *DeliveryIssueWithData) GetStatus() IssueStatus {
	if o == nil {
		return IssueStatus("")
	}
	return o.Status
}

func (o *DeliveryIssueWithData) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *DeliveryIssueWithData) GetType() DeliveryIssueWithDataType {
	if o == nil {
		return DeliveryIssueWithDataType("")
	}
	return o.Type
}

func (o *DeliveryIssueWithData) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

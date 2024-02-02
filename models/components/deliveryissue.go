// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/v2/internal/utils"
	"time"
)

type DeliveryIssueType string

const (
	DeliveryIssueTypeDelivery DeliveryIssueType = "delivery"
)

func (e DeliveryIssueType) ToPointer() *DeliveryIssueType {
	return &e
}

func (e *DeliveryIssueType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delivery":
		*e = DeliveryIssueType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeliveryIssueType: %v", v)
	}
}

// DeliveryIssue - Delivery issue
type DeliveryIssue struct {
	// Keys used as the aggregation keys a 'delivery' type issue
	AggregationKeys DeliveryIssueAggregationKeys `json:"aggregation_keys"`
	AutoResolvedAt  *time.Time                   `json:"auto_resolved_at,omitempty"`
	// ISO timestamp for when the issue was created
	CreatedAt string `json:"created_at"`
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
	TeamID string            `json:"team_id"`
	Type   DeliveryIssueType `json:"type"`
	// ISO timestamp for when the issue was last updated
	UpdatedAt string `json:"updated_at"`
}

func (d DeliveryIssue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeliveryIssue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DeliveryIssue) GetAggregationKeys() DeliveryIssueAggregationKeys {
	if o == nil {
		return DeliveryIssueAggregationKeys{}
	}
	return o.AggregationKeys
}

func (o *DeliveryIssue) GetAutoResolvedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AutoResolvedAt
}

func (o *DeliveryIssue) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *DeliveryIssue) GetDismissedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DismissedAt
}

func (o *DeliveryIssue) GetFirstSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.FirstSeenAt
}

func (o *DeliveryIssue) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeliveryIssue) GetLastSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastSeenAt
}

func (o *DeliveryIssue) GetLastUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedBy
}

func (o *DeliveryIssue) GetMergedWith() *string {
	if o == nil {
		return nil
	}
	return o.MergedWith
}

func (o *DeliveryIssue) GetOpenedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.OpenedAt
}

func (o *DeliveryIssue) GetReference() DeliveryIssueReference {
	if o == nil {
		return DeliveryIssueReference{}
	}
	return o.Reference
}

func (o *DeliveryIssue) GetStatus() IssueStatus {
	if o == nil {
		return IssueStatus("")
	}
	return o.Status
}

func (o *DeliveryIssue) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *DeliveryIssue) GetType() DeliveryIssueType {
	if o == nil {
		return DeliveryIssueType("")
	}
	return o.Type
}

func (o *DeliveryIssue) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

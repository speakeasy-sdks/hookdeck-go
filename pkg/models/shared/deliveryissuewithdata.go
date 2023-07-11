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

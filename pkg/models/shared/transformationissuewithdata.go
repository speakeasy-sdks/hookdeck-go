// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type TransformationIssueWithDataType string

const (
	TransformationIssueWithDataTypeTransformation TransformationIssueWithDataType = "transformation"
)

func (e TransformationIssueWithDataType) ToPointer() *TransformationIssueWithDataType {
	return &e
}

func (e *TransformationIssueWithDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "transformation":
		*e = TransformationIssueWithDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransformationIssueWithDataType: %v", v)
	}
}

// TransformationIssueWithData - Transformation issue
type TransformationIssueWithData struct {
	// Keys used as the aggregation keys a 'transformation' type issue
	AggregationKeys TransformationIssueAggregationKeys `json:"aggregation_keys"`
	AutoResolvedAt  *time.Time                         `json:"auto_resolved_at,omitempty"`
	// ISO timestamp for when the issue was created
	CreatedAt string `json:"created_at"`
	// Transformation issue data
	Data *TransformationIssueData `json:"data,omitempty"`
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
	// Reference to the event request transformation an issue is being created for.
	Reference TransformationIssueReference `json:"reference"`
	// Issue status
	Status IssueStatus `json:"status"`
	// ID of the workspace
	TeamID string                          `json:"team_id"`
	Type   TransformationIssueWithDataType `json:"type"`
	// ISO timestamp for when the issue was last updated
	UpdatedAt string `json:"updated_at"`
}

func (o *TransformationIssueWithData) GetAggregationKeys() TransformationIssueAggregationKeys {
	if o == nil {
		return TransformationIssueAggregationKeys{}
	}
	return o.AggregationKeys
}

func (o *TransformationIssueWithData) GetAutoResolvedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AutoResolvedAt
}

func (o *TransformationIssueWithData) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *TransformationIssueWithData) GetData() *TransformationIssueData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *TransformationIssueWithData) GetDismissedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DismissedAt
}

func (o *TransformationIssueWithData) GetFirstSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.FirstSeenAt
}

func (o *TransformationIssueWithData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TransformationIssueWithData) GetLastSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastSeenAt
}

func (o *TransformationIssueWithData) GetLastUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedBy
}

func (o *TransformationIssueWithData) GetMergedWith() *string {
	if o == nil {
		return nil
	}
	return o.MergedWith
}

func (o *TransformationIssueWithData) GetOpenedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.OpenedAt
}

func (o *TransformationIssueWithData) GetReference() TransformationIssueReference {
	if o == nil {
		return TransformationIssueReference{}
	}
	return o.Reference
}

func (o *TransformationIssueWithData) GetStatus() IssueStatus {
	if o == nil {
		return IssueStatus("")
	}
	return o.Status
}

func (o *TransformationIssueWithData) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *TransformationIssueWithData) GetType() TransformationIssueWithDataType {
	if o == nil {
		return TransformationIssueWithDataType("")
	}
	return o.Type
}

func (o *TransformationIssueWithData) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

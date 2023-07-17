// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type TransformationIssueType string

const (
	TransformationIssueTypeTransformation TransformationIssueType = "transformation"
)

func (e TransformationIssueType) ToPointer() *TransformationIssueType {
	return &e
}

func (e *TransformationIssueType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "transformation":
		*e = TransformationIssueType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransformationIssueType: %v", v)
	}
}

// TransformationIssue - Transformation issue
type TransformationIssue struct {
	// Keys used as the aggregation keys a 'transformation' type issue
	AggregationKeys TransformationIssueAggregationKeys `json:"aggregation_keys"`
	AutoResolvedAt  *time.Time                         `json:"auto_resolved_at,omitempty"`
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
	// Reference to the event request transformation an issue is being created for.
	Reference TransformationIssueReference `json:"reference"`
	// Issue status
	Status IssueStatus `json:"status"`
	// ID of the workspace
	TeamID string                  `json:"team_id"`
	Type   TransformationIssueType `json:"type"`
	// ISO timestamp for when the issue was last updated
	UpdatedAt string `json:"updated_at"`
}

func (o *TransformationIssue) GetAggregationKeys() TransformationIssueAggregationKeys {
	if o == nil {
		return TransformationIssueAggregationKeys{}
	}
	return o.AggregationKeys
}

func (o *TransformationIssue) GetAutoResolvedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AutoResolvedAt
}

func (o *TransformationIssue) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *TransformationIssue) GetDismissedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DismissedAt
}

func (o *TransformationIssue) GetFirstSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.FirstSeenAt
}

func (o *TransformationIssue) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TransformationIssue) GetLastSeenAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastSeenAt
}

func (o *TransformationIssue) GetLastUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedBy
}

func (o *TransformationIssue) GetMergedWith() *string {
	if o == nil {
		return nil
	}
	return o.MergedWith
}

func (o *TransformationIssue) GetOpenedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.OpenedAt
}

func (o *TransformationIssue) GetReference() TransformationIssueReference {
	if o == nil {
		return TransformationIssueReference{}
	}
	return o.Reference
}

func (o *TransformationIssue) GetStatus() IssueStatus {
	if o == nil {
		return IssueStatus("")
	}
	return o.Status
}

func (o *TransformationIssue) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *TransformationIssue) GetType() TransformationIssueType {
	if o == nil {
		return TransformationIssueType("")
	}
	return o.Type
}

func (o *TransformationIssue) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

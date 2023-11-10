// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// IssueStatus - Issue status
type IssueStatus string

const (
	IssueStatusOpened       IssueStatus = "OPENED"
	IssueStatusIgnored      IssueStatus = "IGNORED"
	IssueStatusAcknowledged IssueStatus = "ACKNOWLEDGED"
	IssueStatusResolved     IssueStatus = "RESOLVED"
)

func (e IssueStatus) ToPointer() *IssueStatus {
	return &e
}

func (e *IssueStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OPENED":
		fallthrough
	case "IGNORED":
		fallthrough
	case "ACKNOWLEDGED":
		fallthrough
	case "RESOLVED":
		*e = IssueStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IssueStatus: %v", v)
	}
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RetryRuleType - A retry rule must be of type `retry`
type RetryRuleType string

const (
	RetryRuleTypeRetry RetryRuleType = "retry"
)

func (e RetryRuleType) ToPointer() *RetryRuleType {
	return &e
}

func (e *RetryRuleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "retry":
		*e = RetryRuleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetryRuleType: %v", v)
	}
}

type RetryRule struct {
	// Maximum number of retries to attempt
	Count *int64 `json:"count,omitempty"`
	// Time in MS between each retry
	Interval *int64 `json:"interval,omitempty"`
	// Algorithm to use when calculating delay between retries
	Strategy RetryStrategy `json:"strategy"`
	// A retry rule must be of type `retry`
	Type RetryRuleType `json:"type"`
}
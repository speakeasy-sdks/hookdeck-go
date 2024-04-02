// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AlertStrategy - Alert strategy to use
type AlertStrategy string

const (
	AlertStrategyEachAttempt AlertStrategy = "each_attempt"
	AlertStrategyLastAttempt AlertStrategy = "last_attempt"
)

func (e AlertStrategy) ToPointer() *AlertStrategy {
	return &e
}

func (e *AlertStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "each_attempt":
		fallthrough
	case "last_attempt":
		*e = AlertStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AlertStrategy: %v", v)
	}
}
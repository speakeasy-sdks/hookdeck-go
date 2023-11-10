// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// RetryStrategy - Algorithm to use when calculating delay between retries
type RetryStrategy string

const (
	RetryStrategyLinear      RetryStrategy = "linear"
	RetryStrategyExponential RetryStrategy = "exponential"
)

func (e RetryStrategy) ToPointer() *RetryStrategy {
	return &e
}

func (e *RetryStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linear":
		fallthrough
	case "exponential":
		*e = RetryStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetryStrategy: %v", v)
	}
}

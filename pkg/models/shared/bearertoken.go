// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BearerTokenType - Type of auth method
type BearerTokenType string

const (
	BearerTokenTypeBearerToken BearerTokenType = "BEARER_TOKEN"
)

func (e BearerTokenType) ToPointer() *BearerTokenType {
	return &e
}

func (e *BearerTokenType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BEARER_TOKEN":
		*e = BearerTokenType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BearerTokenType: %v", v)
	}
}

// BearerToken - Bearer Token
type BearerToken struct {
	// Bearer token config for the destination's auth method
	Config *DestinationAuthMethodBearerTokenConfig `json:"config,omitempty"`
	// Type of auth method
	Type BearerTokenType `json:"type"`
}

func (o *BearerToken) GetConfig() *DestinationAuthMethodBearerTokenConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *BearerToken) GetType() BearerTokenType {
	if o == nil {
		return BearerTokenType("")
	}
	return o.Type
}

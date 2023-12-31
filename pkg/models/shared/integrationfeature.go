// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type IntegrationFeature string

const (
	IntegrationFeatureVerification IntegrationFeature = "VERIFICATION"
	IntegrationFeatureHandshake    IntegrationFeature = "HANDSHAKE"
	IntegrationFeaturePolling      IntegrationFeature = "POLLING"
)

func (e IntegrationFeature) ToPointer() *IntegrationFeature {
	return &e
}

func (e *IntegrationFeature) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VERIFICATION":
		fallthrough
	case "HANDSHAKE":
		fallthrough
	case "POLLING":
		*e = IntegrationFeature(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationFeature: %v", v)
	}
}

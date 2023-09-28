// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
)

// DestinationAuthMethodAPIKeyConfigTo - Whether the API key should be sent as a header or a query parameter
type DestinationAuthMethodAPIKeyConfigTo string

const (
	DestinationAuthMethodAPIKeyConfigToHeader DestinationAuthMethodAPIKeyConfigTo = "header"
	DestinationAuthMethodAPIKeyConfigToQuery  DestinationAuthMethodAPIKeyConfigTo = "query"
)

func (e DestinationAuthMethodAPIKeyConfigTo) ToPointer() *DestinationAuthMethodAPIKeyConfigTo {
	return &e
}

func (e *DestinationAuthMethodAPIKeyConfigTo) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "header":
		fallthrough
	case "query":
		*e = DestinationAuthMethodAPIKeyConfigTo(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAuthMethodAPIKeyConfigTo: %v", v)
	}
}

// DestinationAuthMethodAPIKeyConfig - API key config for the destination's auth method
type DestinationAuthMethodAPIKeyConfig struct {
	// API key for the API key auth
	APIKey string `json:"api_key"`
	// Key for the API key auth
	Key string `json:"key"`
	// Whether the API key should be sent as a header or a query parameter
	To *DestinationAuthMethodAPIKeyConfigTo `default:"header" json:"to"`
}

func (d DestinationAuthMethodAPIKeyConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAuthMethodAPIKeyConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAuthMethodAPIKeyConfig) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *DestinationAuthMethodAPIKeyConfig) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *DestinationAuthMethodAPIKeyConfig) GetTo() *DestinationAuthMethodAPIKeyConfigTo {
	if o == nil {
		return nil
	}
	return o.To
}

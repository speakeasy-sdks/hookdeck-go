// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
)

// APIKey - API Key
type APIKey struct {
	// API key config for the destination's auth method
	Config *DestinationAuthMethodAPIKeyConfig `json:"config,omitempty"`
	// Type of auth method
	type_ string `const:"API_KEY" json:"type"`
}

func (a APIKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *APIKey) GetConfig() *DestinationAuthMethodAPIKeyConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *APIKey) GetType() string {
	return "API_KEY"
}

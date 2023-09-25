// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
)

// BasicAuth - Basic Auth
type BasicAuth struct {
	// Basic auth config for the destination's auth method
	Config *DestinationAuthMethodBasicAuthConfig `json:"config,omitempty"`
	// Type of auth method
	type_ string `const:"BASIC_AUTH" json:"type"`
}

func (b BasicAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BasicAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *BasicAuth) GetConfig() *DestinationAuthMethodBasicAuthConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *BasicAuth) GetType() string {
	return "BASIC_AUTH"
}

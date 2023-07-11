// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ShopifyIntegrationConfigsRateLimitPeriod string

const (
	ShopifyIntegrationConfigsRateLimitPeriodLessThanNilGreaterThan ShopifyIntegrationConfigsRateLimitPeriod = "<nil>"
	ShopifyIntegrationConfigsRateLimitPeriodMinute                 ShopifyIntegrationConfigsRateLimitPeriod = "minute"
	ShopifyIntegrationConfigsRateLimitPeriodSecond                 ShopifyIntegrationConfigsRateLimitPeriod = "second"
)

func (e ShopifyIntegrationConfigsRateLimitPeriod) ToPointer() *ShopifyIntegrationConfigsRateLimitPeriod {
	return &e
}

func (e *ShopifyIntegrationConfigsRateLimitPeriod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "<nil>":
		fallthrough
	case "minute":
		fallthrough
	case "second":
		*e = ShopifyIntegrationConfigsRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ShopifyIntegrationConfigsRateLimitPeriod: %v", v)
	}
}

// ShopifyIntegrationConfigs - Decrypted Key/Value object of the associated configuration for that provider
type ShopifyIntegrationConfigs struct {
	APIKey           *string                                   `json:"api_key,omitempty"`
	APISecret        *string                                   `json:"api_secret,omitempty"`
	RateLimit        *float32                                  `json:"rate_limit,omitempty"`
	RateLimitPeriod  *ShopifyIntegrationConfigsRateLimitPeriod `json:"rate_limit_period,omitempty"`
	Shop             *string                                   `json:"shop,omitempty"`
	WebhookSecretKey string                                    `json:"webhook_secret_key"`
}

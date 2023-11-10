// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
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

// ShopifyIntegrationConfigsSchemas - Decrypted Key/Value object of the associated configuration for that provider
type ShopifyIntegrationConfigsSchemas struct {
	APIKey           *string                                   `json:"api_key,omitempty"`
	APISecret        *string                                   `json:"api_secret,omitempty"`
	RateLimit        *float32                                  `json:"rate_limit,omitempty"`
	RateLimitPeriod  *ShopifyIntegrationConfigsRateLimitPeriod `json:"rate_limit_period,omitempty"`
	Shop             *string                                   `json:"shop,omitempty"`
	WebhookSecretKey string                                    `json:"webhook_secret_key"`
}

func (o *ShopifyIntegrationConfigsSchemas) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *ShopifyIntegrationConfigsSchemas) GetAPISecret() *string {
	if o == nil {
		return nil
	}
	return o.APISecret
}

func (o *ShopifyIntegrationConfigsSchemas) GetRateLimit() *float32 {
	if o == nil {
		return nil
	}
	return o.RateLimit
}

func (o *ShopifyIntegrationConfigsSchemas) GetRateLimitPeriod() *ShopifyIntegrationConfigsRateLimitPeriod {
	if o == nil {
		return nil
	}
	return o.RateLimitPeriod
}

func (o *ShopifyIntegrationConfigsSchemas) GetShop() *string {
	if o == nil {
		return nil
	}
	return o.Shop
}

func (o *ShopifyIntegrationConfigsSchemas) GetWebhookSecretKey() string {
	if o == nil {
		return ""
	}
	return o.WebhookSecretKey
}

// BasicAuthIntegrationConfigsSchemas - Decrypted Key/Value object of the associated configuration for that provider
type BasicAuthIntegrationConfigsSchemas struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (o *BasicAuthIntegrationConfigsSchemas) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *BasicAuthIntegrationConfigsSchemas) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

// HandledHMACConfigsSchemas - Decrypted Key/Value object of the associated configuration for that provider
type HandledHMACConfigsSchemas struct {
	WebhookSecretKey string `json:"webhook_secret_key"`
}

func (o *HandledHMACConfigsSchemas) GetWebhookSecretKey() string {
	if o == nil {
		return ""
	}
	return o.WebhookSecretKey
}

// Schemas - Decrypted Key/Value object of the associated configuration for that provider
type Schemas struct {
	APIKey    string `json:"api_key"`
	HeaderKey string `json:"header_key"`
}

func (o *Schemas) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *Schemas) GetHeaderKey() string {
	if o == nil {
		return ""
	}
	return o.HeaderKey
}

type Encoding string

const (
	EncodingBase64 Encoding = "base64"
	EncodingHex    Encoding = "hex"
)

func (e Encoding) ToPointer() *Encoding {
	return &e
}

func (e *Encoding) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "base64":
		fallthrough
	case "hex":
		*e = Encoding(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Encoding: %v", v)
	}
}

// HMACIntegrationConfigsSchemas - Decrypted Key/Value object of the associated configuration for that provider
type HMACIntegrationConfigsSchemas struct {
	Algorithm        components.HMACAlgorithms `json:"algorithm"`
	Encoding         Encoding                  `json:"encoding"`
	HeaderKey        string                    `json:"header_key"`
	WebhookSecretKey string                    `json:"webhook_secret_key"`
}

func (o *HMACIntegrationConfigsSchemas) GetAlgorithm() components.HMACAlgorithms {
	if o == nil {
		return components.HMACAlgorithms("")
	}
	return o.Algorithm
}

func (o *HMACIntegrationConfigsSchemas) GetEncoding() Encoding {
	if o == nil {
		return Encoding("")
	}
	return o.Encoding
}

func (o *HMACIntegrationConfigsSchemas) GetHeaderKey() string {
	if o == nil {
		return ""
	}
	return o.HeaderKey
}

func (o *HMACIntegrationConfigsSchemas) GetWebhookSecretKey() string {
	if o == nil {
		return ""
	}
	return o.WebhookSecretKey
}

// One - Decrypted Key/Value object of the associated configuration for that provider
type One struct {
}

type CreateIntegrationRequestBody struct {
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs interface{} `json:"configs,omitempty"`
	// List of features to enable (see features list above)
	Features []components.IntegrationFeature `json:"features,omitempty"`
	// Label of the integration
	Label *string `json:"label,omitempty"`
	// The provider name
	Provider *components.IntegrationProvider `json:"provider,omitempty"`
}

func (o *CreateIntegrationRequestBody) GetConfigs() interface{} {
	if o == nil {
		return nil
	}
	return o.Configs
}

func (o *CreateIntegrationRequestBody) GetFeatures() []components.IntegrationFeature {
	if o == nil {
		return nil
	}
	return o.Features
}

func (o *CreateIntegrationRequestBody) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *CreateIntegrationRequestBody) GetProvider() *components.IntegrationProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

type CreateIntegrationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single integration
	Integration *components.Integration
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIntegrationResponse) GetIntegration() *components.Integration {
	if o == nil {
		return nil
	}
	return o.Integration
}

func (o *CreateIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// UpdateIntegration1 - Decrypted Key/Value object of the associated configuration for that provider
type UpdateIntegration1 struct {
}

type UpdateIntegrationRequestBody struct {
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs interface{} `json:"configs,omitempty"`
	// List of features to enable (see features list above)
	Features []components.IntegrationFeature `json:"features,omitempty"`
	// Label of the integration
	Label *string `json:"label,omitempty"`
	// The provider name
	Provider *components.IntegrationProvider `json:"provider,omitempty"`
}

func (o *UpdateIntegrationRequestBody) GetConfigs() interface{} {
	if o == nil {
		return nil
	}
	return o.Configs
}

func (o *UpdateIntegrationRequestBody) GetFeatures() []components.IntegrationFeature {
	if o == nil {
		return nil
	}
	return o.Features
}

func (o *UpdateIntegrationRequestBody) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *UpdateIntegrationRequestBody) GetProvider() *components.IntegrationProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

type UpdateIntegrationRequest struct {
	RequestBody UpdateIntegrationRequestBody `request:"mediaType=application/json"`
	ID          string                       `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateIntegrationRequest) GetRequestBody() UpdateIntegrationRequestBody {
	if o == nil {
		return UpdateIntegrationRequestBody{}
	}
	return o.RequestBody
}

func (o *UpdateIntegrationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateIntegrationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single integration
	Integration *components.Integration
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIntegrationResponse) GetIntegration() *components.Integration {
	if o == nil {
		return nil
	}
	return o.Integration
}

func (o *UpdateIntegrationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIntegrationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
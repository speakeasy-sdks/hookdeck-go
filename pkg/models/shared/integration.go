// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
	"time"
)

type IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod string

const (
	IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriodLessThanNilGreaterThan IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod = "<nil>"
	IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriodMinute                 IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod = "minute"
	IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriodSecond                 IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod = "second"
)

func (e IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod) ToPointer() *IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod {
	return &e
}

func (e *IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod) UnmarshalJSON(data []byte) error {
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
		*e = IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod: %v", v)
	}
}

// IntegrationConfigsShopifyIntegrationConfigs - Decrypted Key/Value object of the associated configuration for that provider
type IntegrationConfigsShopifyIntegrationConfigs struct {
	APIKey           *string                                                     `json:"api_key,omitempty"`
	APISecret        *string                                                     `json:"api_secret,omitempty"`
	RateLimit        *float32                                                    `json:"rate_limit,omitempty"`
	RateLimitPeriod  *IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod `json:"rate_limit_period,omitempty"`
	Shop             *string                                                     `json:"shop,omitempty"`
	WebhookSecretKey string                                                      `json:"webhook_secret_key"`
}

func (o *IntegrationConfigsShopifyIntegrationConfigs) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *IntegrationConfigsShopifyIntegrationConfigs) GetAPISecret() *string {
	if o == nil {
		return nil
	}
	return o.APISecret
}

func (o *IntegrationConfigsShopifyIntegrationConfigs) GetRateLimit() *float32 {
	if o == nil {
		return nil
	}
	return o.RateLimit
}

func (o *IntegrationConfigsShopifyIntegrationConfigs) GetRateLimitPeriod() *IntegrationConfigsShopifyIntegrationConfigsRateLimitPeriod {
	if o == nil {
		return nil
	}
	return o.RateLimitPeriod
}

func (o *IntegrationConfigsShopifyIntegrationConfigs) GetShop() *string {
	if o == nil {
		return nil
	}
	return o.Shop
}

func (o *IntegrationConfigsShopifyIntegrationConfigs) GetWebhookSecretKey() string {
	if o == nil {
		return ""
	}
	return o.WebhookSecretKey
}

// IntegrationConfigsBasicAuthIntegrationConfigs - Decrypted Key/Value object of the associated configuration for that provider
type IntegrationConfigsBasicAuthIntegrationConfigs struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (o *IntegrationConfigsBasicAuthIntegrationConfigs) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *IntegrationConfigsBasicAuthIntegrationConfigs) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

// IntegrationConfigsHandledHMACConfigs - Decrypted Key/Value object of the associated configuration for that provider
type IntegrationConfigsHandledHMACConfigs struct {
	WebhookSecretKey string `json:"webhook_secret_key"`
}

func (o *IntegrationConfigsHandledHMACConfigs) GetWebhookSecretKey() string {
	if o == nil {
		return ""
	}
	return o.WebhookSecretKey
}

// IntegrationConfigsAPIKeyIntegrationConfigs - Decrypted Key/Value object of the associated configuration for that provider
type IntegrationConfigsAPIKeyIntegrationConfigs struct {
	APIKey    string `json:"api_key"`
	HeaderKey string `json:"header_key"`
}

func (o *IntegrationConfigsAPIKeyIntegrationConfigs) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *IntegrationConfigsAPIKeyIntegrationConfigs) GetHeaderKey() string {
	if o == nil {
		return ""
	}
	return o.HeaderKey
}

type IntegrationConfigsHMACIntegrationConfigsEncoding string

const (
	IntegrationConfigsHMACIntegrationConfigsEncodingBase64 IntegrationConfigsHMACIntegrationConfigsEncoding = "base64"
	IntegrationConfigsHMACIntegrationConfigsEncodingHex    IntegrationConfigsHMACIntegrationConfigsEncoding = "hex"
)

func (e IntegrationConfigsHMACIntegrationConfigsEncoding) ToPointer() *IntegrationConfigsHMACIntegrationConfigsEncoding {
	return &e
}

func (e *IntegrationConfigsHMACIntegrationConfigsEncoding) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "base64":
		fallthrough
	case "hex":
		*e = IntegrationConfigsHMACIntegrationConfigsEncoding(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationConfigsHMACIntegrationConfigsEncoding: %v", v)
	}
}

// IntegrationConfigsHMACIntegrationConfigs - Decrypted Key/Value object of the associated configuration for that provider
type IntegrationConfigsHMACIntegrationConfigs struct {
	Algorithm        HMACAlgorithms                                   `json:"algorithm"`
	Encoding         IntegrationConfigsHMACIntegrationConfigsEncoding `json:"encoding"`
	HeaderKey        string                                           `json:"header_key"`
	WebhookSecretKey string                                           `json:"webhook_secret_key"`
}

func (o *IntegrationConfigsHMACIntegrationConfigs) GetAlgorithm() HMACAlgorithms {
	if o == nil {
		return HMACAlgorithms("")
	}
	return o.Algorithm
}

func (o *IntegrationConfigsHMACIntegrationConfigs) GetEncoding() IntegrationConfigsHMACIntegrationConfigsEncoding {
	if o == nil {
		return IntegrationConfigsHMACIntegrationConfigsEncoding("")
	}
	return o.Encoding
}

func (o *IntegrationConfigsHMACIntegrationConfigs) GetHeaderKey() string {
	if o == nil {
		return ""
	}
	return o.HeaderKey
}

func (o *IntegrationConfigsHMACIntegrationConfigs) GetWebhookSecretKey() string {
	if o == nil {
		return ""
	}
	return o.WebhookSecretKey
}

// IntegrationConfigs1 - Decrypted Key/Value object of the associated configuration for that provider
type IntegrationConfigs1 struct {
}

type IntegrationConfigsType string

const (
	IntegrationConfigsTypeIntegrationConfigs1                           IntegrationConfigsType = "Integration_configs_1"
	IntegrationConfigsTypeIntegrationConfigsHMACIntegrationConfigs      IntegrationConfigsType = "Integration_configs_HMACIntegrationConfigs"
	IntegrationConfigsTypeIntegrationConfigsAPIKeyIntegrationConfigs    IntegrationConfigsType = "Integration_configs_APIKeyIntegrationConfigs"
	IntegrationConfigsTypeIntegrationConfigsHandledHMACConfigs          IntegrationConfigsType = "Integration_configs_HandledHMACConfigs"
	IntegrationConfigsTypeIntegrationConfigsBasicAuthIntegrationConfigs IntegrationConfigsType = "Integration_configs_BasicAuthIntegrationConfigs"
	IntegrationConfigsTypeIntegrationConfigsShopifyIntegrationConfigs   IntegrationConfigsType = "Integration_configs_ShopifyIntegrationConfigs"
)

type IntegrationConfigs struct {
	IntegrationConfigs1                           *IntegrationConfigs1
	IntegrationConfigsHMACIntegrationConfigs      *IntegrationConfigsHMACIntegrationConfigs
	IntegrationConfigsAPIKeyIntegrationConfigs    *IntegrationConfigsAPIKeyIntegrationConfigs
	IntegrationConfigsHandledHMACConfigs          *IntegrationConfigsHandledHMACConfigs
	IntegrationConfigsBasicAuthIntegrationConfigs *IntegrationConfigsBasicAuthIntegrationConfigs
	IntegrationConfigsShopifyIntegrationConfigs   *IntegrationConfigsShopifyIntegrationConfigs

	Type IntegrationConfigsType
}

func CreateIntegrationConfigsIntegrationConfigs1(integrationConfigs1 IntegrationConfigs1) IntegrationConfigs {
	typ := IntegrationConfigsTypeIntegrationConfigs1

	return IntegrationConfigs{
		IntegrationConfigs1: &integrationConfigs1,
		Type:                typ,
	}
}

func CreateIntegrationConfigsIntegrationConfigsHMACIntegrationConfigs(integrationConfigsHMACIntegrationConfigs IntegrationConfigsHMACIntegrationConfigs) IntegrationConfigs {
	typ := IntegrationConfigsTypeIntegrationConfigsHMACIntegrationConfigs

	return IntegrationConfigs{
		IntegrationConfigsHMACIntegrationConfigs: &integrationConfigsHMACIntegrationConfigs,
		Type:                                     typ,
	}
}

func CreateIntegrationConfigsIntegrationConfigsAPIKeyIntegrationConfigs(integrationConfigsAPIKeyIntegrationConfigs IntegrationConfigsAPIKeyIntegrationConfigs) IntegrationConfigs {
	typ := IntegrationConfigsTypeIntegrationConfigsAPIKeyIntegrationConfigs

	return IntegrationConfigs{
		IntegrationConfigsAPIKeyIntegrationConfigs: &integrationConfigsAPIKeyIntegrationConfigs,
		Type: typ,
	}
}

func CreateIntegrationConfigsIntegrationConfigsHandledHMACConfigs(integrationConfigsHandledHMACConfigs IntegrationConfigsHandledHMACConfigs) IntegrationConfigs {
	typ := IntegrationConfigsTypeIntegrationConfigsHandledHMACConfigs

	return IntegrationConfigs{
		IntegrationConfigsHandledHMACConfigs: &integrationConfigsHandledHMACConfigs,
		Type:                                 typ,
	}
}

func CreateIntegrationConfigsIntegrationConfigsBasicAuthIntegrationConfigs(integrationConfigsBasicAuthIntegrationConfigs IntegrationConfigsBasicAuthIntegrationConfigs) IntegrationConfigs {
	typ := IntegrationConfigsTypeIntegrationConfigsBasicAuthIntegrationConfigs

	return IntegrationConfigs{
		IntegrationConfigsBasicAuthIntegrationConfigs: &integrationConfigsBasicAuthIntegrationConfigs,
		Type: typ,
	}
}

func CreateIntegrationConfigsIntegrationConfigsShopifyIntegrationConfigs(integrationConfigsShopifyIntegrationConfigs IntegrationConfigsShopifyIntegrationConfigs) IntegrationConfigs {
	typ := IntegrationConfigsTypeIntegrationConfigsShopifyIntegrationConfigs

	return IntegrationConfigs{
		IntegrationConfigsShopifyIntegrationConfigs: &integrationConfigsShopifyIntegrationConfigs,
		Type: typ,
	}
}

func (u *IntegrationConfigs) UnmarshalJSON(data []byte) error {

	integrationConfigs1 := IntegrationConfigs1{}
	if err := utils.UnmarshalJSON(data, &integrationConfigs1, "", true, true); err == nil {
		u.IntegrationConfigs1 = &integrationConfigs1
		u.Type = IntegrationConfigsTypeIntegrationConfigs1
		return nil
	}

	integrationConfigsHandledHMACConfigs := IntegrationConfigsHandledHMACConfigs{}
	if err := utils.UnmarshalJSON(data, &integrationConfigsHandledHMACConfigs, "", true, true); err == nil {
		u.IntegrationConfigsHandledHMACConfigs = &integrationConfigsHandledHMACConfigs
		u.Type = IntegrationConfigsTypeIntegrationConfigsHandledHMACConfigs
		return nil
	}

	integrationConfigsAPIKeyIntegrationConfigs := IntegrationConfigsAPIKeyIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &integrationConfigsAPIKeyIntegrationConfigs, "", true, true); err == nil {
		u.IntegrationConfigsAPIKeyIntegrationConfigs = &integrationConfigsAPIKeyIntegrationConfigs
		u.Type = IntegrationConfigsTypeIntegrationConfigsAPIKeyIntegrationConfigs
		return nil
	}

	integrationConfigsBasicAuthIntegrationConfigs := IntegrationConfigsBasicAuthIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &integrationConfigsBasicAuthIntegrationConfigs, "", true, true); err == nil {
		u.IntegrationConfigsBasicAuthIntegrationConfigs = &integrationConfigsBasicAuthIntegrationConfigs
		u.Type = IntegrationConfigsTypeIntegrationConfigsBasicAuthIntegrationConfigs
		return nil
	}

	integrationConfigsHMACIntegrationConfigs := IntegrationConfigsHMACIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &integrationConfigsHMACIntegrationConfigs, "", true, true); err == nil {
		u.IntegrationConfigsHMACIntegrationConfigs = &integrationConfigsHMACIntegrationConfigs
		u.Type = IntegrationConfigsTypeIntegrationConfigsHMACIntegrationConfigs
		return nil
	}

	integrationConfigsShopifyIntegrationConfigs := IntegrationConfigsShopifyIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &integrationConfigsShopifyIntegrationConfigs, "", true, true); err == nil {
		u.IntegrationConfigsShopifyIntegrationConfigs = &integrationConfigsShopifyIntegrationConfigs
		u.Type = IntegrationConfigsTypeIntegrationConfigsShopifyIntegrationConfigs
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u IntegrationConfigs) MarshalJSON() ([]byte, error) {
	if u.IntegrationConfigs1 != nil {
		return utils.MarshalJSON(u.IntegrationConfigs1, "", true)
	}

	if u.IntegrationConfigsHMACIntegrationConfigs != nil {
		return utils.MarshalJSON(u.IntegrationConfigsHMACIntegrationConfigs, "", true)
	}

	if u.IntegrationConfigsAPIKeyIntegrationConfigs != nil {
		return utils.MarshalJSON(u.IntegrationConfigsAPIKeyIntegrationConfigs, "", true)
	}

	if u.IntegrationConfigsHandledHMACConfigs != nil {
		return utils.MarshalJSON(u.IntegrationConfigsHandledHMACConfigs, "", true)
	}

	if u.IntegrationConfigsBasicAuthIntegrationConfigs != nil {
		return utils.MarshalJSON(u.IntegrationConfigsBasicAuthIntegrationConfigs, "", true)
	}

	if u.IntegrationConfigsShopifyIntegrationConfigs != nil {
		return utils.MarshalJSON(u.IntegrationConfigsShopifyIntegrationConfigs, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Integration struct {
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs IntegrationConfigs `json:"configs"`
	// Date the integration was created
	CreatedAt time.Time `json:"created_at"`
	// List of features to enable (see features list below)
	Features []IntegrationFeature `json:"features"`
	// ID of the integration
	ID string `json:"id"`
	// Label of the integration
	Label string `json:"label"`
	// The provider name
	Provider IntegrationProvider `json:"provider"`
	// List of source IDs the integration is attached to
	Sources []string `json:"sources"`
	// ID of the workspace
	TeamID string `json:"team_id"`
	// Date the integration was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

func (i Integration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Integration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Integration) GetConfigs() IntegrationConfigs {
	if o == nil {
		return IntegrationConfigs{}
	}
	return o.Configs
}

func (o *Integration) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Integration) GetFeatures() []IntegrationFeature {
	if o == nil {
		return []IntegrationFeature{}
	}
	return o.Features
}

func (o *Integration) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Integration) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *Integration) GetProvider() IntegrationProvider {
	if o == nil {
		return IntegrationProvider("")
	}
	return o.Provider
}

func (o *Integration) GetSources() []string {
	if o == nil {
		return []string{}
	}
	return o.Sources
}

func (o *Integration) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *Integration) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

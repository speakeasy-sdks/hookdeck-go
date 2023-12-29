// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"time"
)

type Integration1 struct {
}

type ConfigsType string

const (
	ConfigsTypeIntegration1                ConfigsType = "Integration_1"
	ConfigsTypeHMACIntegrationConfigs      ConfigsType = "HMACIntegrationConfigs"
	ConfigsTypeAPIKeyIntegrationConfigs    ConfigsType = "APIKeyIntegrationConfigs"
	ConfigsTypeHandledHMACConfigs          ConfigsType = "HandledHMACConfigs"
	ConfigsTypeBasicAuthIntegrationConfigs ConfigsType = "BasicAuthIntegrationConfigs"
	ConfigsTypeShopifyIntegrationConfigs   ConfigsType = "ShopifyIntegrationConfigs"
)

// Configs - Decrypted Key/Value object of the associated configuration for that provider
type Configs struct {
	Integration1                *Integration1
	HMACIntegrationConfigs      *HMACIntegrationConfigs
	APIKeyIntegrationConfigs    *APIKeyIntegrationConfigs
	HandledHMACConfigs          *HandledHMACConfigs
	BasicAuthIntegrationConfigs *BasicAuthIntegrationConfigs
	ShopifyIntegrationConfigs   *ShopifyIntegrationConfigs

	Type ConfigsType
}

func CreateConfigsIntegration1(integration1 Integration1) Configs {
	typ := ConfigsTypeIntegration1

	return Configs{
		Integration1: &integration1,
		Type:         typ,
	}
}

func CreateConfigsHMACIntegrationConfigs(hmacIntegrationConfigs HMACIntegrationConfigs) Configs {
	typ := ConfigsTypeHMACIntegrationConfigs

	return Configs{
		HMACIntegrationConfigs: &hmacIntegrationConfigs,
		Type:                   typ,
	}
}

func CreateConfigsAPIKeyIntegrationConfigs(apiKeyIntegrationConfigs APIKeyIntegrationConfigs) Configs {
	typ := ConfigsTypeAPIKeyIntegrationConfigs

	return Configs{
		APIKeyIntegrationConfigs: &apiKeyIntegrationConfigs,
		Type:                     typ,
	}
}

func CreateConfigsHandledHMACConfigs(handledHMACConfigs HandledHMACConfigs) Configs {
	typ := ConfigsTypeHandledHMACConfigs

	return Configs{
		HandledHMACConfigs: &handledHMACConfigs,
		Type:               typ,
	}
}

func CreateConfigsBasicAuthIntegrationConfigs(basicAuthIntegrationConfigs BasicAuthIntegrationConfigs) Configs {
	typ := ConfigsTypeBasicAuthIntegrationConfigs

	return Configs{
		BasicAuthIntegrationConfigs: &basicAuthIntegrationConfigs,
		Type:                        typ,
	}
}

func CreateConfigsShopifyIntegrationConfigs(shopifyIntegrationConfigs ShopifyIntegrationConfigs) Configs {
	typ := ConfigsTypeShopifyIntegrationConfigs

	return Configs{
		ShopifyIntegrationConfigs: &shopifyIntegrationConfigs,
		Type:                      typ,
	}
}

func (u *Configs) UnmarshalJSON(data []byte) error {

	integration1 := Integration1{}
	if err := utils.UnmarshalJSON(data, &integration1, "", true, true); err == nil {
		u.Integration1 = &integration1
		u.Type = ConfigsTypeIntegration1
		return nil
	}

	handledHMACConfigs := HandledHMACConfigs{}
	if err := utils.UnmarshalJSON(data, &handledHMACConfigs, "", true, true); err == nil {
		u.HandledHMACConfigs = &handledHMACConfigs
		u.Type = ConfigsTypeHandledHMACConfigs
		return nil
	}

	apiKeyIntegrationConfigs := APIKeyIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &apiKeyIntegrationConfigs, "", true, true); err == nil {
		u.APIKeyIntegrationConfigs = &apiKeyIntegrationConfigs
		u.Type = ConfigsTypeAPIKeyIntegrationConfigs
		return nil
	}

	basicAuthIntegrationConfigs := BasicAuthIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &basicAuthIntegrationConfigs, "", true, true); err == nil {
		u.BasicAuthIntegrationConfigs = &basicAuthIntegrationConfigs
		u.Type = ConfigsTypeBasicAuthIntegrationConfigs
		return nil
	}

	hmacIntegrationConfigs := HMACIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &hmacIntegrationConfigs, "", true, true); err == nil {
		u.HMACIntegrationConfigs = &hmacIntegrationConfigs
		u.Type = ConfigsTypeHMACIntegrationConfigs
		return nil
	}

	shopifyIntegrationConfigs := ShopifyIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &shopifyIntegrationConfigs, "", true, true); err == nil {
		u.ShopifyIntegrationConfigs = &shopifyIntegrationConfigs
		u.Type = ConfigsTypeShopifyIntegrationConfigs
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Configs) MarshalJSON() ([]byte, error) {
	if u.Integration1 != nil {
		return utils.MarshalJSON(u.Integration1, "", true)
	}

	if u.HMACIntegrationConfigs != nil {
		return utils.MarshalJSON(u.HMACIntegrationConfigs, "", true)
	}

	if u.APIKeyIntegrationConfigs != nil {
		return utils.MarshalJSON(u.APIKeyIntegrationConfigs, "", true)
	}

	if u.HandledHMACConfigs != nil {
		return utils.MarshalJSON(u.HandledHMACConfigs, "", true)
	}

	if u.BasicAuthIntegrationConfigs != nil {
		return utils.MarshalJSON(u.BasicAuthIntegrationConfigs, "", true)
	}

	if u.ShopifyIntegrationConfigs != nil {
		return utils.MarshalJSON(u.ShopifyIntegrationConfigs, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Integration struct {
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs Configs `json:"configs"`
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

func (o *Integration) GetConfigs() Configs {
	if o == nil {
		return Configs{}
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

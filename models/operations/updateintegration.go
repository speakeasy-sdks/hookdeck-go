// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type UpdateIntegration1 struct {
}

type UpdateIntegrationConfigsType string

const (
	UpdateIntegrationConfigsTypeUpdateIntegration1          UpdateIntegrationConfigsType = "updateIntegration_1"
	UpdateIntegrationConfigsTypeHMACIntegrationConfigs      UpdateIntegrationConfigsType = "HMACIntegrationConfigs"
	UpdateIntegrationConfigsTypeAPIKeyIntegrationConfigs    UpdateIntegrationConfigsType = "APIKeyIntegrationConfigs"
	UpdateIntegrationConfigsTypeHandledHMACConfigs          UpdateIntegrationConfigsType = "HandledHMACConfigs"
	UpdateIntegrationConfigsTypeBasicAuthIntegrationConfigs UpdateIntegrationConfigsType = "BasicAuthIntegrationConfigs"
	UpdateIntegrationConfigsTypeShopifyIntegrationConfigs   UpdateIntegrationConfigsType = "ShopifyIntegrationConfigs"
)

// UpdateIntegrationConfigs - Decrypted Key/Value object of the associated configuration for that provider
type UpdateIntegrationConfigs struct {
	UpdateIntegration1          *UpdateIntegration1
	HMACIntegrationConfigs      *components.HMACIntegrationConfigs
	APIKeyIntegrationConfigs    *components.APIKeyIntegrationConfigs
	HandledHMACConfigs          *components.HandledHMACConfigs
	BasicAuthIntegrationConfigs *components.BasicAuthIntegrationConfigs
	ShopifyIntegrationConfigs   *components.ShopifyIntegrationConfigs

	Type UpdateIntegrationConfigsType
}

func CreateUpdateIntegrationConfigsUpdateIntegration1(updateIntegration1 UpdateIntegration1) UpdateIntegrationConfigs {
	typ := UpdateIntegrationConfigsTypeUpdateIntegration1

	return UpdateIntegrationConfigs{
		UpdateIntegration1: &updateIntegration1,
		Type:               typ,
	}
}

func CreateUpdateIntegrationConfigsHMACIntegrationConfigs(hmacIntegrationConfigs components.HMACIntegrationConfigs) UpdateIntegrationConfigs {
	typ := UpdateIntegrationConfigsTypeHMACIntegrationConfigs

	return UpdateIntegrationConfigs{
		HMACIntegrationConfigs: &hmacIntegrationConfigs,
		Type:                   typ,
	}
}

func CreateUpdateIntegrationConfigsAPIKeyIntegrationConfigs(apiKeyIntegrationConfigs components.APIKeyIntegrationConfigs) UpdateIntegrationConfigs {
	typ := UpdateIntegrationConfigsTypeAPIKeyIntegrationConfigs

	return UpdateIntegrationConfigs{
		APIKeyIntegrationConfigs: &apiKeyIntegrationConfigs,
		Type:                     typ,
	}
}

func CreateUpdateIntegrationConfigsHandledHMACConfigs(handledHMACConfigs components.HandledHMACConfigs) UpdateIntegrationConfigs {
	typ := UpdateIntegrationConfigsTypeHandledHMACConfigs

	return UpdateIntegrationConfigs{
		HandledHMACConfigs: &handledHMACConfigs,
		Type:               typ,
	}
}

func CreateUpdateIntegrationConfigsBasicAuthIntegrationConfigs(basicAuthIntegrationConfigs components.BasicAuthIntegrationConfigs) UpdateIntegrationConfigs {
	typ := UpdateIntegrationConfigsTypeBasicAuthIntegrationConfigs

	return UpdateIntegrationConfigs{
		BasicAuthIntegrationConfigs: &basicAuthIntegrationConfigs,
		Type:                        typ,
	}
}

func CreateUpdateIntegrationConfigsShopifyIntegrationConfigs(shopifyIntegrationConfigs components.ShopifyIntegrationConfigs) UpdateIntegrationConfigs {
	typ := UpdateIntegrationConfigsTypeShopifyIntegrationConfigs

	return UpdateIntegrationConfigs{
		ShopifyIntegrationConfigs: &shopifyIntegrationConfigs,
		Type:                      typ,
	}
}

func (u *UpdateIntegrationConfigs) UnmarshalJSON(data []byte) error {

	updateIntegration1 := UpdateIntegration1{}
	if err := utils.UnmarshalJSON(data, &updateIntegration1, "", true, true); err == nil {
		u.UpdateIntegration1 = &updateIntegration1
		u.Type = UpdateIntegrationConfigsTypeUpdateIntegration1
		return nil
	}

	handledHMACConfigs := components.HandledHMACConfigs{}
	if err := utils.UnmarshalJSON(data, &handledHMACConfigs, "", true, true); err == nil {
		u.HandledHMACConfigs = &handledHMACConfigs
		u.Type = UpdateIntegrationConfigsTypeHandledHMACConfigs
		return nil
	}

	apiKeyIntegrationConfigs := components.APIKeyIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &apiKeyIntegrationConfigs, "", true, true); err == nil {
		u.APIKeyIntegrationConfigs = &apiKeyIntegrationConfigs
		u.Type = UpdateIntegrationConfigsTypeAPIKeyIntegrationConfigs
		return nil
	}

	basicAuthIntegrationConfigs := components.BasicAuthIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &basicAuthIntegrationConfigs, "", true, true); err == nil {
		u.BasicAuthIntegrationConfigs = &basicAuthIntegrationConfigs
		u.Type = UpdateIntegrationConfigsTypeBasicAuthIntegrationConfigs
		return nil
	}

	hmacIntegrationConfigs := components.HMACIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &hmacIntegrationConfigs, "", true, true); err == nil {
		u.HMACIntegrationConfigs = &hmacIntegrationConfigs
		u.Type = UpdateIntegrationConfigsTypeHMACIntegrationConfigs
		return nil
	}

	shopifyIntegrationConfigs := components.ShopifyIntegrationConfigs{}
	if err := utils.UnmarshalJSON(data, &shopifyIntegrationConfigs, "", true, true); err == nil {
		u.ShopifyIntegrationConfigs = &shopifyIntegrationConfigs
		u.Type = UpdateIntegrationConfigsTypeShopifyIntegrationConfigs
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateIntegrationConfigs) MarshalJSON() ([]byte, error) {
	if u.UpdateIntegration1 != nil {
		return utils.MarshalJSON(u.UpdateIntegration1, "", true)
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

type UpdateIntegrationRequestBody struct {
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs *UpdateIntegrationConfigs `json:"configs,omitempty"`
	// List of features to enable (see features list above)
	Features []components.IntegrationFeature `json:"features,omitempty"`
	// Label of the integration
	Label *string `json:"label,omitempty"`
	// The provider name
	Provider *components.IntegrationProvider `json:"provider,omitempty"`
}

func (o *UpdateIntegrationRequestBody) GetConfigs() *UpdateIntegrationConfigs {
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

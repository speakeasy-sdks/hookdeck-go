// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

// CreateIntegrationRequestBodyConfigs1 - Decrypted Key/Value object of the associated configuration for that provider
type CreateIntegrationRequestBodyConfigs1 struct {
}

type CreateIntegrationRequestBody struct {
	// Decrypted Key/Value object of the associated configuration for that provider
	Configs interface{} `json:"configs,omitempty"`
	// List of features to enable (see features list above)
	Features []shared.IntegrationFeature `json:"features,omitempty"`
	// Label of the integration
	Label *string `json:"label,omitempty"`
	// The provider name
	Provider *shared.IntegrationProvider `json:"provider,omitempty"`
}

func (o *CreateIntegrationRequestBody) GetConfigs() interface{} {
	if o == nil {
		return nil
	}
	return o.Configs
}

func (o *CreateIntegrationRequestBody) GetFeatures() []shared.IntegrationFeature {
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

func (o *CreateIntegrationRequestBody) GetProvider() *shared.IntegrationProvider {
	if o == nil {
		return nil
	}
	return o.Provider
}

type CreateIntegrationResponse struct {
	ContentType string
	// A single integration
	Integration *shared.Integration
	StatusCode  int
	RawResponse *http.Response
}

func (o *CreateIntegrationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIntegrationResponse) GetIntegration() *shared.Integration {
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

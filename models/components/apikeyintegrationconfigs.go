// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type APIKeyIntegrationConfigs struct {
	APIKey    string `json:"api_key"`
	HeaderKey string `json:"header_key"`
}

func (o *APIKeyIntegrationConfigs) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *APIKeyIntegrationConfigs) GetHeaderKey() string {
	if o == nil {
		return ""
	}
	return o.HeaderKey
}

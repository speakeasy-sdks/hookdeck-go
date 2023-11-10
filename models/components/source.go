// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"time"
)

// SourceIntegration - Integration object
type SourceIntegration struct {
	// List of enabled features
	Features []IntegrationFeature `json:"features"`
	// ID of the integration
	ID string `json:"id"`
	// Label of the integration
	Label string `json:"label"`
	// The provider name
	Provider IntegrationProvider `json:"provider"`
}

func (o *SourceIntegration) GetFeatures() []IntegrationFeature {
	if o == nil {
		return []IntegrationFeature{}
	}
	return o.Features
}

func (o *SourceIntegration) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SourceIntegration) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *SourceIntegration) GetProvider() IntegrationProvider {
	if o == nil {
		return IntegrationProvider("")
	}
	return o.Provider
}

// Source - Associated [Source](#source-object) object
type Source struct {
	// List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.
	AllowedHTTPMethods []SourceAllowedHTTPMethod `json:"allowed_http_methods,omitempty"`
	// Date the source was archived
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// Date the source was created
	CreatedAt time.Time `json:"created_at"`
	// Custom response object
	CustomResponse *SourceCustomResponse `json:"custom_response,omitempty"`
	// ID of the source
	ID string `json:"id"`
	// Integration object
	Integration *SourceIntegration `json:"integration,omitempty"`
	// ID of the integration
	IntegrationID *string `json:"integration_id,omitempty"`
	// Name for the source
	Name string `json:"name"`
	// ID of the workspace
	TeamID string `json:"team_id"`
	// Date the source was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// A unique URL that must be supplied to your webhook's provider
	URL string `json:"url"`
}

func (s Source) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Source) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Source) GetAllowedHTTPMethods() []SourceAllowedHTTPMethod {
	if o == nil {
		return nil
	}
	return o.AllowedHTTPMethods
}

func (o *Source) GetArchivedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *Source) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Source) GetCustomResponse() *SourceCustomResponse {
	if o == nil {
		return nil
	}
	return o.CustomResponse
}

func (o *Source) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Source) GetIntegration() *SourceIntegration {
	if o == nil {
		return nil
	}
	return o.Integration
}

func (o *Source) GetIntegrationID() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationID
}

func (o *Source) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Source) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *Source) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Source) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
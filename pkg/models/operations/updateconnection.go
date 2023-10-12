// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

// UpdateConnectionRequestBodyRuleset - Ruleset input object
type UpdateConnectionRequestBodyRuleset struct {
	IsTeamDefault *bool `json:"is_team_default,omitempty"`
	// Name for the ruleset
	Name string `json:"name"`
	// Array of rules to apply
	Rules []shared.Rule `json:"rules,omitempty"`
}

func (o *UpdateConnectionRequestBodyRuleset) GetIsTeamDefault() *bool {
	if o == nil {
		return nil
	}
	return o.IsTeamDefault
}

func (o *UpdateConnectionRequestBodyRuleset) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateConnectionRequestBodyRuleset) GetRules() []shared.Rule {
	if o == nil {
		return nil
	}
	return o.Rules
}

type UpdateConnectionRequestBody struct {
	// Unique name of the connection for the source
	Name *string `json:"name,omitempty"`
	// Array of rules to apply
	Rules []shared.Rule `json:"rules,omitempty"`
	// Ruleset input object
	Ruleset *UpdateConnectionRequestBodyRuleset `json:"ruleset,omitempty"`
	// ID of a rule to bind to the connection. Default to the Workspace default ruleset
	RulesetID *string `json:"ruleset_id,omitempty"`
}

func (o *UpdateConnectionRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateConnectionRequestBody) GetRules() []shared.Rule {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *UpdateConnectionRequestBody) GetRuleset() *UpdateConnectionRequestBodyRuleset {
	if o == nil {
		return nil
	}
	return o.Ruleset
}

func (o *UpdateConnectionRequestBody) GetRulesetID() *string {
	if o == nil {
		return nil
	}
	return o.RulesetID
}

type UpdateConnectionRequest struct {
	RequestBody UpdateConnectionRequestBody `request:"mediaType=application/json"`
	ID          string                      `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateConnectionRequest) GetRequestBody() UpdateConnectionRequestBody {
	if o == nil {
		return UpdateConnectionRequestBody{}
	}
	return o.RequestBody
}

func (o *UpdateConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateConnectionResponse struct {
	// A single connection
	Connection *shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UpdateConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

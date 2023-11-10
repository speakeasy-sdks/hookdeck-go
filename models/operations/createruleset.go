// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type CreateRulesetRequestBody struct {
	IsTeamDefault *bool `json:"is_team_default,omitempty"`
	// Name for the ruleset
	Name string `json:"name"`
	// Array of rules to apply
	Rules []components.Rule `json:"rules,omitempty"`
}

func (o *CreateRulesetRequestBody) GetIsTeamDefault() *bool {
	if o == nil {
		return nil
	}
	return o.IsTeamDefault
}

func (o *CreateRulesetRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateRulesetRequestBody) GetRules() []components.Rule {
	if o == nil {
		return nil
	}
	return o.Rules
}

type CreateRulesetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single ruleset
	Ruleset *components.Ruleset
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateRulesetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRulesetResponse) GetRuleset() *components.Ruleset {
	if o == nil {
		return nil
	}
	return o.Ruleset
}

func (o *CreateRulesetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRulesetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
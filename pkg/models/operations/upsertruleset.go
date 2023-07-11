// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck/pkg/models/shared"
	"net/http"
)

type UpsertRulesetRequestBody struct {
	IsTeamDefault *bool `json:"is_team_default,omitempty"`
	// Name for the ruleset
	Name string `json:"name"`
	// Array of rules to apply
	Rules []interface{} `json:"rules,omitempty"`
}

type UpsertRulesetResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single ruleset
	Ruleset     *shared.Ruleset
	StatusCode  int
	RawResponse *http.Response
}
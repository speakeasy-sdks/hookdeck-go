// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"hookdeck/pkg/models/shared"
	"net/http"
)

// CreateConnectionRequestBodyDestinationRateLimitPeriod - Period to rate limit attempts
type CreateConnectionRequestBodyDestinationRateLimitPeriod string

const (
	CreateConnectionRequestBodyDestinationRateLimitPeriodSecond CreateConnectionRequestBodyDestinationRateLimitPeriod = "second"
	CreateConnectionRequestBodyDestinationRateLimitPeriodMinute CreateConnectionRequestBodyDestinationRateLimitPeriod = "minute"
	CreateConnectionRequestBodyDestinationRateLimitPeriodHour   CreateConnectionRequestBodyDestinationRateLimitPeriod = "hour"
)

func (e CreateConnectionRequestBodyDestinationRateLimitPeriod) ToPointer() *CreateConnectionRequestBodyDestinationRateLimitPeriod {
	return &e
}

func (e *CreateConnectionRequestBodyDestinationRateLimitPeriod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "second":
		fallthrough
	case "minute":
		fallthrough
	case "hour":
		*e = CreateConnectionRequestBodyDestinationRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConnectionRequestBodyDestinationRateLimitPeriod: %v", v)
	}
}

// CreateConnectionRequestBodyDestination - Destination input object
type CreateConnectionRequestBodyDestination struct {
	// Config for the destination's auth method
	AuthMethod interface{} `json:"auth_method,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// HTTP method used on requests sent to the destination, overrides the method used on requests sent to the source.
	HTTPMethod *shared.DestinationHTTPMethod `json:"http_method,omitempty"`
	// Name for the destination
	Name                   string `json:"name"`
	PathForwardingDisabled *bool  `json:"path_forwarding_disabled,omitempty"`
	// Limit event attempts to receive per period
	RateLimit interface{} `json:"rate_limit,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *CreateConnectionRequestBodyDestinationRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Endpoint of the destination
	URL *string `json:"url,omitempty"`
}

// CreateConnectionRequestBodyRuleset - Ruleset input object
type CreateConnectionRequestBodyRuleset struct {
	IsTeamDefault *bool `json:"is_team_default,omitempty"`
	// Name for the ruleset
	Name string `json:"name"`
	// Array of rules to apply
	Rules []interface{} `json:"rules,omitempty"`
}

// CreateConnectionRequestBodySource - Source input object
type CreateConnectionRequestBodySource struct {
	// List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.
	AllowedHTTPMethods []shared.SourceAllowedHTTPMethod `json:"allowed_http_methods,omitempty"`
	// Custom response object
	CustomResponse *shared.SourceCustomResponse `json:"custom_response,omitempty"`
	// A unique name for the source
	Name string `json:"name"`
}

type CreateConnectionRequestBody struct {
	// Destination input object
	Destination *CreateConnectionRequestBodyDestination `json:"destination,omitempty"`
	// ID of a destination to bind to the connection
	DestinationID *string `json:"destination_id,omitempty"`
	// A unique name of the connection for the source
	Name string `json:"name"`
	// Array of rules to apply
	Rules []interface{} `json:"rules,omitempty"`
	// Ruleset input object
	Ruleset *CreateConnectionRequestBodyRuleset `json:"ruleset,omitempty"`
	// ID of a rule to bind to the connection. Default to the Workspace default ruleset
	RulesetID *string `json:"ruleset_id,omitempty"`
	// Source input object
	Source *CreateConnectionRequestBodySource `json:"source,omitempty"`
	// ID of a source to bind to the connection
	SourceID *string `json:"source_id,omitempty"`
}

type CreateConnectionResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	// A single connection
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

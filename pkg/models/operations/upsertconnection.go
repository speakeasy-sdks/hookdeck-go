// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

// UpsertConnectionRequestBodyDestinationRateLimitPeriod - Period to rate limit attempts
type UpsertConnectionRequestBodyDestinationRateLimitPeriod string

const (
	UpsertConnectionRequestBodyDestinationRateLimitPeriodSecond UpsertConnectionRequestBodyDestinationRateLimitPeriod = "second"
	UpsertConnectionRequestBodyDestinationRateLimitPeriodMinute UpsertConnectionRequestBodyDestinationRateLimitPeriod = "minute"
	UpsertConnectionRequestBodyDestinationRateLimitPeriodHour   UpsertConnectionRequestBodyDestinationRateLimitPeriod = "hour"
)

func (e UpsertConnectionRequestBodyDestinationRateLimitPeriod) ToPointer() *UpsertConnectionRequestBodyDestinationRateLimitPeriod {
	return &e
}

func (e *UpsertConnectionRequestBodyDestinationRateLimitPeriod) UnmarshalJSON(data []byte) error {
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
		*e = UpsertConnectionRequestBodyDestinationRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpsertConnectionRequestBodyDestinationRateLimitPeriod: %v", v)
	}
}

// UpsertConnectionRequestBodyDestination - Destination input object
type UpsertConnectionRequestBodyDestination struct {
	// Config for the destination's auth method
	AuthMethod *shared.DestinationAuthMethodConfig `json:"auth_method,omitempty"`
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
	RateLimitPeriod *UpsertConnectionRequestBodyDestinationRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Endpoint of the destination
	URL *string `json:"url,omitempty"`
}

func (o *UpsertConnectionRequestBodyDestination) GetAuthMethod() *shared.DestinationAuthMethodConfig {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *UpsertConnectionRequestBodyDestination) GetCliPath() *string {
	if o == nil {
		return nil
	}
	return o.CliPath
}

func (o *UpsertConnectionRequestBodyDestination) GetHTTPMethod() *shared.DestinationHTTPMethod {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *UpsertConnectionRequestBodyDestination) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpsertConnectionRequestBodyDestination) GetPathForwardingDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.PathForwardingDisabled
}

func (o *UpsertConnectionRequestBodyDestination) GetRateLimit() interface{} {
	if o == nil {
		return nil
	}
	return o.RateLimit
}

func (o *UpsertConnectionRequestBodyDestination) GetRateLimitPeriod() *UpsertConnectionRequestBodyDestinationRateLimitPeriod {
	if o == nil {
		return nil
	}
	return o.RateLimitPeriod
}

func (o *UpsertConnectionRequestBodyDestination) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// UpsertConnectionRequestBodyRuleset - Ruleset input object
type UpsertConnectionRequestBodyRuleset struct {
	IsTeamDefault *bool `json:"is_team_default,omitempty"`
	// Name for the ruleset
	Name string `json:"name"`
	// Array of rules to apply
	Rules []shared.Rule `json:"rules,omitempty"`
}

func (o *UpsertConnectionRequestBodyRuleset) GetIsTeamDefault() *bool {
	if o == nil {
		return nil
	}
	return o.IsTeamDefault
}

func (o *UpsertConnectionRequestBodyRuleset) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpsertConnectionRequestBodyRuleset) GetRules() []shared.Rule {
	if o == nil {
		return nil
	}
	return o.Rules
}

// UpsertConnectionRequestBodySource - Source input object
type UpsertConnectionRequestBodySource struct {
	// List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.
	AllowedHTTPMethods []shared.SourceAllowedHTTPMethod `json:"allowed_http_methods,omitempty"`
	// Custom response object
	CustomResponse *shared.SourceCustomResponse `json:"custom_response,omitempty"`
	// A unique name for the source
	Name string `json:"name"`
}

func (o *UpsertConnectionRequestBodySource) GetAllowedHTTPMethods() []shared.SourceAllowedHTTPMethod {
	if o == nil {
		return nil
	}
	return o.AllowedHTTPMethods
}

func (o *UpsertConnectionRequestBodySource) GetCustomResponse() *shared.SourceCustomResponse {
	if o == nil {
		return nil
	}
	return o.CustomResponse
}

func (o *UpsertConnectionRequestBodySource) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type UpsertConnectionRequestBody struct {
	// Destination input object
	Destination *UpsertConnectionRequestBodyDestination `json:"destination,omitempty"`
	// ID of a destination to bind to the connection
	DestinationID *string `json:"destination_id,omitempty"`
	// A unique name of the connection for the source
	Name string `json:"name"`
	// Array of rules to apply
	Rules []shared.Rule `json:"rules,omitempty"`
	// Ruleset input object
	Ruleset *UpsertConnectionRequestBodyRuleset `json:"ruleset,omitempty"`
	// ID of a rule to bind to the connection. Default to the Workspace default ruleset
	RulesetID *string `json:"ruleset_id,omitempty"`
	// Source input object
	Source *UpsertConnectionRequestBodySource `json:"source,omitempty"`
	// ID of a source to bind to the connection
	SourceID *string `json:"source_id,omitempty"`
}

func (o *UpsertConnectionRequestBody) GetDestination() *UpsertConnectionRequestBodyDestination {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *UpsertConnectionRequestBody) GetDestinationID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *UpsertConnectionRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpsertConnectionRequestBody) GetRules() []shared.Rule {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *UpsertConnectionRequestBody) GetRuleset() *UpsertConnectionRequestBodyRuleset {
	if o == nil {
		return nil
	}
	return o.Ruleset
}

func (o *UpsertConnectionRequestBody) GetRulesetID() *string {
	if o == nil {
		return nil
	}
	return o.RulesetID
}

func (o *UpsertConnectionRequestBody) GetSource() *UpsertConnectionRequestBodySource {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *UpsertConnectionRequestBody) GetSourceID() *string {
	if o == nil {
		return nil
	}
	return o.SourceID
}

type UpsertConnectionResponse struct {
	// A single connection
	Connection *shared.Connection
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpsertConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *UpsertConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

// CreateDestinationRateLimitPeriod - Period to rate limit attempts
type CreateDestinationRateLimitPeriod string

const (
	CreateDestinationRateLimitPeriodSecond CreateDestinationRateLimitPeriod = "second"
	CreateDestinationRateLimitPeriodMinute CreateDestinationRateLimitPeriod = "minute"
	CreateDestinationRateLimitPeriodHour   CreateDestinationRateLimitPeriod = "hour"
)

func (e CreateDestinationRateLimitPeriod) ToPointer() *CreateDestinationRateLimitPeriod {
	return &e
}

func (e *CreateDestinationRateLimitPeriod) UnmarshalJSON(data []byte) error {
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
		*e = CreateDestinationRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDestinationRateLimitPeriod: %v", v)
	}
}

type CreateDestinationRequestBody struct {
	// Config for the destination's auth method
	AuthMethod *components.DestinationAuthMethodConfig `json:"auth_method,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// HTTP method used on requests sent to the destination, overrides the method used on requests sent to the source.
	HTTPMethod *components.DestinationHTTPMethod `json:"http_method,omitempty"`
	// Name for the destination
	Name                   string `json:"name"`
	PathForwardingDisabled *bool  `json:"path_forwarding_disabled,omitempty"`
	// Limit event attempts to receive per period
	RateLimit *int64 `json:"rate_limit,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *CreateDestinationRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Endpoint of the destination
	URL *string `json:"url,omitempty"`
}

func (o *CreateDestinationRequestBody) GetAuthMethod() *components.DestinationAuthMethodConfig {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *CreateDestinationRequestBody) GetCliPath() *string {
	if o == nil {
		return nil
	}
	return o.CliPath
}

func (o *CreateDestinationRequestBody) GetHTTPMethod() *components.DestinationHTTPMethod {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *CreateDestinationRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateDestinationRequestBody) GetPathForwardingDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.PathForwardingDisabled
}

func (o *CreateDestinationRequestBody) GetRateLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.RateLimit
}

func (o *CreateDestinationRequestBody) GetRateLimitPeriod() *CreateDestinationRateLimitPeriod {
	if o == nil {
		return nil
	}
	return o.RateLimitPeriod
}

func (o *CreateDestinationRequestBody) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type CreateDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single destination
	Destination *components.Destination
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDestinationResponse) GetDestination() *components.Destination {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *CreateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

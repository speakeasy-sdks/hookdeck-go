// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

// UpsertDestinationRequestBodyRateLimitPeriod - Period to rate limit attempts
type UpsertDestinationRequestBodyRateLimitPeriod string

const (
	UpsertDestinationRequestBodyRateLimitPeriodSecond UpsertDestinationRequestBodyRateLimitPeriod = "second"
	UpsertDestinationRequestBodyRateLimitPeriodMinute UpsertDestinationRequestBodyRateLimitPeriod = "minute"
	UpsertDestinationRequestBodyRateLimitPeriodHour   UpsertDestinationRequestBodyRateLimitPeriod = "hour"
)

func (e UpsertDestinationRequestBodyRateLimitPeriod) ToPointer() *UpsertDestinationRequestBodyRateLimitPeriod {
	return &e
}

func (e *UpsertDestinationRequestBodyRateLimitPeriod) UnmarshalJSON(data []byte) error {
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
		*e = UpsertDestinationRequestBodyRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpsertDestinationRequestBodyRateLimitPeriod: %v", v)
	}
}

type UpsertDestinationRequestBody struct {
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
	RateLimitPeriod *UpsertDestinationRequestBodyRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Endpoint of the destination
	URL *string `json:"url,omitempty"`
}

type UpsertDestinationResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single destination
	Destination *shared.Destination
	StatusCode  int
	RawResponse *http.Response
}

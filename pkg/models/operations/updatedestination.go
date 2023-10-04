// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

// UpdateDestinationRequestBodyRateLimitPeriod - Period to rate limit attempts
type UpdateDestinationRequestBodyRateLimitPeriod string

const (
	UpdateDestinationRequestBodyRateLimitPeriodSecond UpdateDestinationRequestBodyRateLimitPeriod = "second"
	UpdateDestinationRequestBodyRateLimitPeriodMinute UpdateDestinationRequestBodyRateLimitPeriod = "minute"
	UpdateDestinationRequestBodyRateLimitPeriodHour   UpdateDestinationRequestBodyRateLimitPeriod = "hour"
)

func (e UpdateDestinationRequestBodyRateLimitPeriod) ToPointer() *UpdateDestinationRequestBodyRateLimitPeriod {
	return &e
}

func (e *UpdateDestinationRequestBodyRateLimitPeriod) UnmarshalJSON(data []byte) error {
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
		*e = UpdateDestinationRequestBodyRateLimitPeriod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateDestinationRequestBodyRateLimitPeriod: %v", v)
	}
}

type UpdateDestinationRequestBody struct {
	// Config for the destination's auth method
	AuthMethod *shared.DestinationAuthMethodConfig `json:"auth_method,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// HTTP method used on requests sent to the destination, overrides the method used on requests sent to the source.
	HTTPMethod *shared.DestinationHTTPMethod `json:"http_method,omitempty"`
	// Name for the destination
	Name                   *string `json:"name,omitempty"`
	PathForwardingDisabled *bool   `json:"path_forwarding_disabled,omitempty"`
	// Limit event attempts to receive per period
	RateLimit *int64 `json:"rate_limit,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *UpdateDestinationRequestBodyRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Endpoint of the destination
	URL *string `json:"url,omitempty"`
}

func (o *UpdateDestinationRequestBody) GetAuthMethod() *shared.DestinationAuthMethodConfig {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *UpdateDestinationRequestBody) GetCliPath() *string {
	if o == nil {
		return nil
	}
	return o.CliPath
}

func (o *UpdateDestinationRequestBody) GetHTTPMethod() *shared.DestinationHTTPMethod {
	if o == nil {
		return nil
	}
	return o.HTTPMethod
}

func (o *UpdateDestinationRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateDestinationRequestBody) GetPathForwardingDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.PathForwardingDisabled
}

func (o *UpdateDestinationRequestBody) GetRateLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.RateLimit
}

func (o *UpdateDestinationRequestBody) GetRateLimitPeriod() *UpdateDestinationRequestBodyRateLimitPeriod {
	if o == nil {
		return nil
	}
	return o.RateLimitPeriod
}

func (o *UpdateDestinationRequestBody) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

type UpdateDestinationRequest struct {
	RequestBody UpdateDestinationRequestBody `request:"mediaType=application/json"`
	ID          string                       `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateDestinationRequest) GetRequestBody() UpdateDestinationRequestBody {
	if o == nil {
		return UpdateDestinationRequestBody{}
	}
	return o.RequestBody
}

func (o *UpdateDestinationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single destination
	Destination *shared.Destination
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDestinationResponse) GetDestination() *shared.Destination {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *UpdateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

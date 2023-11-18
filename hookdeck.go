// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package hookdeckgo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production API
	"https://api.hookdeck.com/{version}",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

type Hookdeck struct {
	// An attempt is any request that Hookdeck makes on behalf of an event.
	Attempts *Attempts
	Attempt  *Attempt
	// A bookmark lets you conveniently store and replay a specific request.
	Bookmarks              *Bookmarks
	Bookmark               *Bookmark
	BulkRetryEvent         *BulkRetryEvent
	BulkRetryIgnoredEvents *BulkRetryIgnoredEvents
	BulkRetryIgnoredEvent  *BulkRetryIgnoredEvent
	BulkRetryRequests      *BulkRetryRequests
	RequestBulkRetry       *RequestBulkRetry
	// A destination is any endpoint to which your webhooks can be routed.
	Destinations *Destinations
	Destination  *Destination
	// An event is any request that Hookdeck receives from a source.
	Events       *Events
	Event        *Event
	EventRawBody *EventRawBody
	// An integration configures platform-specific behaviors, such as signature verification.
	Integrations  *Integrations
	Integration   *Integration
	IssueTriggers *IssueTriggers
	IssueTrigger  *IssueTrigger
	// Issues lets you track problems in your workspace and communicate resolution steps with your team.
	Issues               *Issues
	IssuesCount          *IssuesCount
	Issue                *Issue
	WebhookNotifications *WebhookNotifications
	// A request represent a webhook received by Hookdeck.
	Requests             *Requests
	Request              *Request
	RequestEvents        *RequestEvents
	RequestIgnoredEvents *RequestIgnoredEvents
	RequestRawBody       *RequestRawBody
	// A ruleset defines a group of rules that can be used across many connections.
	Rulesets *Rulesets
	Ruleset  *Ruleset
	// A source represents any third party that sends webhooks to Hookdeck.
	Sources       *Sources
	Source        *Source
	CustomDomain  *CustomDomain
	CustomDomains *CustomDomains
	// A transformation represents JavaScript code that will be executed on a connection's requests. Transformations are applied to connections using Rules.
	Transformations          *Transformations
	Transformation           *Transformation
	TransformationExecutions *TransformationExecutions
	TransformationExecution  *TransformationExecution
	// A connection lets you route webhooks from a source to a destination, using a ruleset.
	Connections            *Connections
	Connection             *Connection
	ConnectionNumberUpdate *ConnectionNumberUpdate

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Hookdeck)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Hookdeck) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Hookdeck) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Hookdeck) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

type ServerVersion string

const (
	ServerVersionTwoThousandAndTwentyThree0101 ServerVersion = "2023-01-01"
	ServerVersionTwoThousandAndTwentyTwo1101   ServerVersion = "2022-11-01"
	ServerVersionTwoThousandAndTwentyTwo1001   ServerVersion = "2022-10-01"
	ServerVersionTwoThousandAndTwentyTwo0701   ServerVersion = "2022-07-01"
	ServerVersionTwoThousandAndTwentyTwo0301   ServerVersion = "2022-03-01"
	ServerVersionTwoThousandAndTwentyOne0801   ServerVersion = "2021-08-01"
	ServerVersionTwoThousandAndTwenty0101      ServerVersion = "2020-01-01"
)

func (e ServerVersion) ToPointer() *ServerVersion {
	return &e
}

func (e *ServerVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "2023-01-01":
		fallthrough
	case "2022-11-01":
		fallthrough
	case "2022-10-01":
		fallthrough
	case "2022-07-01":
		fallthrough
	case "2022-03-01":
		fallthrough
	case "2021-08-01":
		fallthrough
	case "2020-01-01":
		*e = ServerVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerVersion: %v", v)
	}
}

// WithVersion allows setting the version variable for url substitution
func WithVersion(version ServerVersion) SDKOption {
	return func(sdk *Hookdeck) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["version"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["version"] = fmt.Sprintf("%v", version)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Hookdeck) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(security components.Security) SDKOption {
	return func(sdk *Hookdeck) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *Hookdeck) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Hookdeck) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Hookdeck {
	sdk := &Hookdeck{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "1.6.3",
			GenVersion:        "2.194.1",
			UserAgent:         "speakeasy-sdk/go 1.6.3 2.194.1 1.0.0 github.com/speakeasy-sdks/hookdeck-go",
			ServerDefaults: []map[string]string{
				{
					"version": "2023-01-01",
				},
			},
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Attempts = newAttempts(sdk.sdkConfiguration)

	sdk.Attempt = newAttempt(sdk.sdkConfiguration)

	sdk.Bookmarks = newBookmarks(sdk.sdkConfiguration)

	sdk.Bookmark = newBookmark(sdk.sdkConfiguration)

	sdk.BulkRetryEvent = newBulkRetryEvent(sdk.sdkConfiguration)

	sdk.BulkRetryIgnoredEvents = newBulkRetryIgnoredEvents(sdk.sdkConfiguration)

	sdk.BulkRetryIgnoredEvent = newBulkRetryIgnoredEvent(sdk.sdkConfiguration)

	sdk.BulkRetryRequests = newBulkRetryRequests(sdk.sdkConfiguration)

	sdk.RequestBulkRetry = newRequestBulkRetry(sdk.sdkConfiguration)

	sdk.Destinations = newDestinations(sdk.sdkConfiguration)

	sdk.Destination = newDestination(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Event = newEvent(sdk.sdkConfiguration)

	sdk.EventRawBody = newEventRawBody(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.Integration = newIntegration(sdk.sdkConfiguration)

	sdk.IssueTriggers = newIssueTriggers(sdk.sdkConfiguration)

	sdk.IssueTrigger = newIssueTrigger(sdk.sdkConfiguration)

	sdk.Issues = newIssues(sdk.sdkConfiguration)

	sdk.IssuesCount = newIssuesCount(sdk.sdkConfiguration)

	sdk.Issue = newIssue(sdk.sdkConfiguration)

	sdk.WebhookNotifications = newWebhookNotifications(sdk.sdkConfiguration)

	sdk.Requests = newRequests(sdk.sdkConfiguration)

	sdk.Request = newRequest(sdk.sdkConfiguration)

	sdk.RequestEvents = newRequestEvents(sdk.sdkConfiguration)

	sdk.RequestIgnoredEvents = newRequestIgnoredEvents(sdk.sdkConfiguration)

	sdk.RequestRawBody = newRequestRawBody(sdk.sdkConfiguration)

	sdk.Rulesets = newRulesets(sdk.sdkConfiguration)

	sdk.Ruleset = newRuleset(sdk.sdkConfiguration)

	sdk.Sources = newSources(sdk.sdkConfiguration)

	sdk.Source = newSource(sdk.sdkConfiguration)

	sdk.CustomDomain = newCustomDomain(sdk.sdkConfiguration)

	sdk.CustomDomains = newCustomDomains(sdk.sdkConfiguration)

	sdk.Transformations = newTransformations(sdk.sdkConfiguration)

	sdk.Transformation = newTransformation(sdk.sdkConfiguration)

	sdk.TransformationExecutions = newTransformationExecutions(sdk.sdkConfiguration)

	sdk.TransformationExecution = newTransformationExecution(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.Connection = newConnection(sdk.sdkConfiguration)

	sdk.ConnectionNumberUpdate = newConnectionNumberUpdate(sdk.sdkConfiguration)

	return sdk
}

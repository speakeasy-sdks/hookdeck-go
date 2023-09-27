// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package hookdeckgo

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
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
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

type Hookdeck struct {
	Attempt *attempt
	// An attempt is any request that Hookdeck makes on behalf of an event.
	Attempts *attempts
	Bookmark *bookmark
	// A bookmark lets you conveniently store and replay a specific request.
	Bookmarks              *bookmarks
	BulkRetryEvent         *bulkRetryEvent
	BulkRetryIgnoredEvent  *bulkRetryIgnoredEvent
	BulkRetryIgnoredEvents *bulkRetryIgnoredEvents
	BulkRetryRequests      *bulkRetryRequests
	Connection             *connection
	ConnectionNumberUpdate *connectionNumberUpdate
	// A connection lets you route webhooks from a source to a destination, using a ruleset.
	Connections   *connections
	CustomDomain  *customDomain
	CustomDomains *customDomains
	Destination   *destination
	// A destination is any endpoint to which your webhooks can be routed.
	Destinations *destinations
	Event        *event
	EventRawBody *eventRawBody
	// An event is any request that Hookdeck receives from a source.
	Events      *events
	Integration *integration
	// An integration configures platform-specific behaviors, such as signature verification.
	Integrations  *integrations
	Issue         *issue
	IssueTrigger  *issueTrigger
	IssueTriggers *issueTriggers
	// Issues lets you track problems in your workspace and communicate resolution steps with your team.
	Issues               *issues
	IssuesCount          *issuesCount
	Request              *request
	RequestBulkRetry     *requestBulkRetry
	RequestEvents        *requestEvents
	RequestIgnoredEvents *requestIgnoredEvents
	RequestRawBody       *requestRawBody
	// A request represent a webhook received by Hookdeck.
	Requests *requests
	Ruleset  *ruleset
	// A ruleset defines a group of rules that can be used across many connections.
	Rulesets *rulesets
	Source   *source
	// A source represents any third party that sends webhooks to Hookdeck.
	Sources                  *sources
	Transformation           *transformation
	TransformationExecution  *transformationExecution
	TransformationExecutions *transformationExecutions
	// A transformation represents JavaScript code that will be executed on a connection's requests. Transformations are applied to connections using Rules.
	Transformations      *transformations
	WebhookNotifications *webhookNotifications

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

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Hookdeck) {
		sdk.sdkConfiguration.Security = &security
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
			SDKVersion:        "1.6.0",
			GenVersion:        "2.131.1",
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

	sdk.Attempt = newAttempt(sdk.sdkConfiguration)

	sdk.Attempts = newAttempts(sdk.sdkConfiguration)

	sdk.Bookmark = newBookmark(sdk.sdkConfiguration)

	sdk.Bookmarks = newBookmarks(sdk.sdkConfiguration)

	sdk.BulkRetryEvent = newBulkRetryEvent(sdk.sdkConfiguration)

	sdk.BulkRetryIgnoredEvent = newBulkRetryIgnoredEvent(sdk.sdkConfiguration)

	sdk.BulkRetryIgnoredEvents = newBulkRetryIgnoredEvents(sdk.sdkConfiguration)

	sdk.BulkRetryRequests = newBulkRetryRequests(sdk.sdkConfiguration)

	sdk.Connection = newConnection(sdk.sdkConfiguration)

	sdk.ConnectionNumberUpdate = newConnectionNumberUpdate(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.CustomDomain = newCustomDomain(sdk.sdkConfiguration)

	sdk.CustomDomains = newCustomDomains(sdk.sdkConfiguration)

	sdk.Destination = newDestination(sdk.sdkConfiguration)

	sdk.Destinations = newDestinations(sdk.sdkConfiguration)

	sdk.Event = newEvent(sdk.sdkConfiguration)

	sdk.EventRawBody = newEventRawBody(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Integration = newIntegration(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.Issue = newIssue(sdk.sdkConfiguration)

	sdk.IssueTrigger = newIssueTrigger(sdk.sdkConfiguration)

	sdk.IssueTriggers = newIssueTriggers(sdk.sdkConfiguration)

	sdk.Issues = newIssues(sdk.sdkConfiguration)

	sdk.IssuesCount = newIssuesCount(sdk.sdkConfiguration)

	sdk.Request = newRequest(sdk.sdkConfiguration)

	sdk.RequestBulkRetry = newRequestBulkRetry(sdk.sdkConfiguration)

	sdk.RequestEvents = newRequestEvents(sdk.sdkConfiguration)

	sdk.RequestIgnoredEvents = newRequestIgnoredEvents(sdk.sdkConfiguration)

	sdk.RequestRawBody = newRequestRawBody(sdk.sdkConfiguration)

	sdk.Requests = newRequests(sdk.sdkConfiguration)

	sdk.Ruleset = newRuleset(sdk.sdkConfiguration)

	sdk.Rulesets = newRulesets(sdk.sdkConfiguration)

	sdk.Source = newSource(sdk.sdkConfiguration)

	sdk.Sources = newSources(sdk.sdkConfiguration)

	sdk.Transformation = newTransformation(sdk.sdkConfiguration)

	sdk.TransformationExecution = newTransformationExecution(sdk.sdkConfiguration)

	sdk.TransformationExecutions = newTransformationExecutions(sdk.sdkConfiguration)

	sdk.Transformations = newTransformations(sdk.sdkConfiguration)

	sdk.WebhookNotifications = newWebhookNotifications(sdk.sdkConfiguration)

	return sdk
}

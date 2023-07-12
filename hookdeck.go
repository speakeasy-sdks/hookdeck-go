// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package hookdeck

import (
	"encoding/json"
	"fmt"
	"hookdeck-go/pkg/models/shared"
	"hookdeck-go/pkg/utils"
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
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// Hookdeck
type Hookdeck struct {
	// Attempts - An attempt is any request that Hookdeck makes on behalf of an event.
	Attempts *attempts
	// Bookmarks - A bookmark lets you conveniently store and replay a specific request.
	Bookmarks              *bookmarks
	BulkRetryEvents        *bulkRetryEvents
	BulkRetryIgnoredEvents *bulkRetryIgnoredEvents
	BulkRetryRequests      *bulkRetryRequests
	// Connections - A connection lets you route webhooks from a source to a destination, using a ruleset.
	Connections *connections
	// Destinations - A destination is any endpoint to which your webhooks can be routed.
	Destinations *destinations
	// Events - An event is any request that Hookdeck receives from a source.
	Events *events
	// Integrations - An integration configures platform-specific behaviors, such as signature verification.
	Integrations  *integrations
	IssueTriggers *issueTriggers
	// Issues - Issues lets you track problems in your workspace and communicate resolution steps with your team.
	Issues *issues
	// Notifications - Notifications let your team receive alerts anytime an issue changes.
	Notifications *notifications
	// Requests - A request represent a webhook received by Hookdeck.
	Requests *requests
	// Rulesets - A ruleset defines a group of rules that can be used across many connections.
	Rulesets *rulesets
	// Sources - A source represents any third party that sends webhooks to Hookdeck.
	Sources *sources
	// Transformations - A transformation represents JavaScript code that will be executed on a connection's requests. Transformations are applied to connections using Rules.
	Transformations *transformations

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

// WithVersion allows setting the $name variable for url substitution
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

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Hookdeck {
	sdk := &Hookdeck{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "1.0.2",
			GenVersion:        "2.61.5",
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

	sdk.Bookmarks = newBookmarks(sdk.sdkConfiguration)

	sdk.BulkRetryEvents = newBulkRetryEvents(sdk.sdkConfiguration)

	sdk.BulkRetryIgnoredEvents = newBulkRetryIgnoredEvents(sdk.sdkConfiguration)

	sdk.BulkRetryRequests = newBulkRetryRequests(sdk.sdkConfiguration)

	sdk.Connections = newConnections(sdk.sdkConfiguration)

	sdk.Destinations = newDestinations(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.IssueTriggers = newIssueTriggers(sdk.sdkConfiguration)

	sdk.Issues = newIssues(sdk.sdkConfiguration)

	sdk.Notifications = newNotifications(sdk.sdkConfiguration)

	sdk.Requests = newRequests(sdk.sdkConfiguration)

	sdk.Rulesets = newRulesets(sdk.sdkConfiguration)

	sdk.Sources = newSources(sdk.sdkConfiguration)

	sdk.Transformations = newTransformations(sdk.sdkConfiguration)

	return sdk
}

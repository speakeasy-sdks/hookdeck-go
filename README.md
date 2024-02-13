<div align="center">
    <picture>
        <img src="https://github.com/speakeasy-sdks/hookdeck-go/assets/68016351/16cab688-a0e2-4444-82aa-ea3a8dd87956" width="500">
    </picture>
    <h1>Go SDK</h1>
   <p>Never miss a webhook</p>
   <a href="https://hookdeck.com/api-ref#getting-started"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/hookdeck-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/hookdeck-go?sort=semver&style=for-the-badge" /></a>
</div>

## Authentication

### Get an API key

You'll need to include an API key in every API request. If you do not yet have a Hookdeck account, please [sign up](https://dashboard.hookdeck.com/signup) for a free account. Your API key is located in your workspace settings.

To include your API key in requests, use either Bearer Token Authentication or Basic Authentication.

**Bearer Token Authentication**: Include the header Authorization: Bearer `$API_KEY`, replacing `$API_KEY` with your personal key
**Basic Authentication** : Set the username as your API key, and supply an empty password

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/hookdeck-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.EventAttemptPaginatedResult != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Attempts](docs/sdks/attempts/README.md)

* [Get](docs/sdks/attempts/README.md#get) - Get Attempts

### [Attempt](docs/sdks/attempt/README.md)

* [Get](docs/sdks/attempt/README.md#get) - Get a Single Attempt

### [Bookmarks](docs/sdks/bookmarks/README.md)

* [Get](docs/sdks/bookmarks/README.md#get) - Get Bookmarks

### [Bookmark](docs/sdks/bookmark/README.md)

* [Create](docs/sdks/bookmark/README.md#create) - Create a Bookmark
* [Delete](docs/sdks/bookmark/README.md#delete) - Delete a Bookmark
* [Get](docs/sdks/bookmark/README.md#get) - Get a Single Bookmark
* [Trigger](docs/sdks/bookmark/README.md#trigger) - Trigger a Bookmark
* [Update](docs/sdks/bookmark/README.md#update) - Update a Bookmark

### [BulkRetryEvent](docs/sdks/bulkretryevent/README.md)

* [Cancel](docs/sdks/bulkretryevent/README.md#cancel) - Cancel an events bulk retry
* [Create](docs/sdks/bulkretryevent/README.md#create) - Create an events bulk retry
* [Get](docs/sdks/bulkretryevent/README.md#get) - Get an events bulk retry

### [BulkRetryIgnoredEvents](docs/sdks/bulkretryignoredevents/README.md)

* [Get](docs/sdks/bulkretryignoredevents/README.md#get) - Get ignored events bulk retries

### [BulkRetryIgnoredEvent](docs/sdks/bulkretryignoredevent/README.md)

* [Cancel](docs/sdks/bulkretryignoredevent/README.md#cancel) - Cancel an ignored events bulk retry
* [Create](docs/sdks/bulkretryignoredevent/README.md#create) - Create an ignored events bulk retry
* [Generate](docs/sdks/bulkretryignoredevent/README.md#generate) - Generate an ignored events bulk retry plan
* [Get](docs/sdks/bulkretryignoredevent/README.md#get) - Get an ignored events bulk retry

### [BulkRetryRequests](docs/sdks/bulkretryrequests/README.md)

* [Create](docs/sdks/bulkretryrequests/README.md#create) - Create a requests bulk retry

### [RequestBulkRetry](docs/sdks/requestbulkretry/README.md)

* [Cancel](docs/sdks/requestbulkretry/README.md#cancel) - Cancel a requests bulk retry
* [Get](docs/sdks/requestbulkretry/README.md#get) - Get a requests bulk retry

### [Destinations](docs/sdks/destinations/README.md)

* [Get](docs/sdks/destinations/README.md#get) - Get Destinations

### [Destination](docs/sdks/destination/README.md)

* [Archive](docs/sdks/destination/README.md#archive) - Archive a Destination
* [Create](docs/sdks/destination/README.md#create) - Create a Destination
* [Delete](docs/sdks/destination/README.md#delete) - Delete a Destination
* [Get](docs/sdks/destination/README.md#get) - Get a Destination
* [Unarchive](docs/sdks/destination/README.md#unarchive) - Unarchive a Destination
* [Update](docs/sdks/destination/README.md#update) - Update a Destination
* [Upsert](docs/sdks/destination/README.md#upsert) - Update or Create a Destination

### [Events](docs/sdks/events/README.md)

* [Get](docs/sdks/events/README.md#get) - Get Events

### [Event](docs/sdks/event/README.md)

* [Get](docs/sdks/event/README.md#get) - Get an Event
* [MuteEvent](docs/sdks/event/README.md#muteevent) - Mute an Event
* [Retry](docs/sdks/event/README.md#retry) - Retry an Event

### [EventRawBody](docs/sdks/eventrawbody/README.md)

* [Get](docs/sdks/eventrawbody/README.md#get) - Get an Event Raw Body Data

### [Integrations](docs/sdks/integrations/README.md)

* [Get](docs/sdks/integrations/README.md#get) - Get integrations

### [Integration](docs/sdks/integration/README.md)

* [AttachIntegrationToSource](docs/sdks/integration/README.md#attachintegrationtosource) - Attach an integration to a source
* [Create](docs/sdks/integration/README.md#create) - Create an integration
* [Delete](docs/sdks/integration/README.md#delete) - Delete an integration
* [DetachIntegrationToSource](docs/sdks/integration/README.md#detachintegrationtosource) - Detach an integration from a source
* [Get](docs/sdks/integration/README.md#get) - Get an integration
* [Update](docs/sdks/integration/README.md#update) - Update an integration

### [IssueTriggers](docs/sdks/issuetriggers/README.md)

* [Get](docs/sdks/issuetriggers/README.md#get) - Get Issue Triggers

### [IssueTrigger](docs/sdks/issuetrigger/README.md)

* [Create](docs/sdks/issuetrigger/README.md#create) - Create an Issue Trigger
* [Delete](docs/sdks/issuetrigger/README.md#delete) - Delete an Issue Trigger
* [Disable](docs/sdks/issuetrigger/README.md#disable) - Disable an Issue Trigger
* [Enable](docs/sdks/issuetrigger/README.md#enable) - Enable an Issue Trigger
* [Get](docs/sdks/issuetrigger/README.md#get) - Get an Issue Trigger
* [Update](docs/sdks/issuetrigger/README.md#update) - Update an Issue Trigger
* [Upsert](docs/sdks/issuetrigger/README.md#upsert) - Create or Update an Issue Trigger

### [Issues](docs/sdks/issues/README.md)

* [Get](docs/sdks/issues/README.md#get) - Get issues

### [IssuesCount](docs/sdks/issuescount/README.md)

* [Get](docs/sdks/issuescount/README.md#get) - Get the number of issues

### [Issue](docs/sdks/issue/README.md)

* [Dismiss](docs/sdks/issue/README.md#dismiss) - Dismiss an issue
* [Get](docs/sdks/issue/README.md#get) - Get a single issue
* [Update](docs/sdks/issue/README.md#update) - Update issue

### [WebhookNotifications](docs/sdks/webhooknotifications/README.md)

* [Toggle](docs/sdks/webhooknotifications/README.md#toggle) - Toggle webhook notifications for the workspace

### [Requests](docs/sdks/requests/README.md)

* [Get](docs/sdks/requests/README.md#get) - Get requests

### [Request](docs/sdks/request/README.md)

* [Get](docs/sdks/request/README.md#get) - Get a request
* [Retry](docs/sdks/request/README.md#retry) - Retry a request

### [RequestEvents](docs/sdks/requestevents/README.md)

* [Get](docs/sdks/requestevents/README.md#get) - Get request events

### [RequestIgnoredEvents](docs/sdks/requestignoredevents/README.md)

* [Get](docs/sdks/requestignoredevents/README.md#get) - Get request ignored events

### [RequestRawBody](docs/sdks/requestrawbody/README.md)

* [Get](docs/sdks/requestrawbody/README.md#get) - Get a request raw body data

### [Rulesets](docs/sdks/rulesets/README.md)

* [Get](docs/sdks/rulesets/README.md#get) - Get rulesets

### [Ruleset](docs/sdks/ruleset/README.md)

* [Archive](docs/sdks/ruleset/README.md#archive) - Archive a ruleset
* [Create](docs/sdks/ruleset/README.md#create) - Create a ruleset
* [Get](docs/sdks/ruleset/README.md#get) - Get a ruleset
* [Unarchive](docs/sdks/ruleset/README.md#unarchive) - Unarchive a ruleset
* [Update](docs/sdks/ruleset/README.md#update) - Update a ruleset
* [Upsert](docs/sdks/ruleset/README.md#upsert) - Update or create a ruleset

### [Sources](docs/sdks/sources/README.md)

* [Get](docs/sdks/sources/README.md#get) - Get sources

### [Source](docs/sdks/source/README.md)

* [Archive](docs/sdks/source/README.md#archive) - Archive a source
* [Create](docs/sdks/source/README.md#create) - Create a source
* [Delete](docs/sdks/source/README.md#delete) - Delete a source
* [Get](docs/sdks/source/README.md#get) - Get a source
* [Unarchive](docs/sdks/source/README.md#unarchive) - Unarchive a source
* [Update](docs/sdks/source/README.md#update) - Update a source
* [Upsert](docs/sdks/source/README.md#upsert) - Update or create a source

### [CustomDomains](docs/sdks/customdomains/README.md)

* [List](docs/sdks/customdomains/README.md#list) - List all custom domains and their verification statuses for the workspace

### [CustomDomain](docs/sdks/customdomain/README.md)

* [Add](docs/sdks/customdomain/README.md#add) - Add a custom domain to the workspace
* [Delete](docs/sdks/customdomain/README.md#delete) - Removes a custom domain from the workspace

### [Transformations](docs/sdks/transformations/README.md)

* [Get](docs/sdks/transformations/README.md#get) - Get transformations

### [Transformation](docs/sdks/transformation/README.md)

* [Create](docs/sdks/transformation/README.md#create) - Create a transformation
* [Get](docs/sdks/transformation/README.md#get) - Get a transformation
* [Test](docs/sdks/transformation/README.md#test) - Test a transformation code
* [Update](docs/sdks/transformation/README.md#update) - Update a transformation
* [Upsert](docs/sdks/transformation/README.md#upsert) - Update or create a transformation

### [TransformationExecutions](docs/sdks/transformationexecutions/README.md)

* [Get](docs/sdks/transformationexecutions/README.md#get) - Get transformation executions

### [TransformationExecution](docs/sdks/transformationexecution/README.md)

* [Get](docs/sdks/transformationexecution/README.md#get) - Get a transformation execution

### [Connections](docs/sdks/connections/README.md)

* [Get](docs/sdks/connections/README.md#get) - Get connections

### [Connection](docs/sdks/connection/README.md)

* [Archive](docs/sdks/connection/README.md#archive) - Archive a connection
* [Create](docs/sdks/connection/README.md#create) - Create a connection
* [Delete](docs/sdks/connection/README.md#delete) - Delete a connection
* [Get](docs/sdks/connection/README.md#get) - Get a single connection
* [Pause](docs/sdks/connection/README.md#pause) - Pause a connection
* [Unarchive](docs/sdks/connection/README.md#unarchive) - Unarchive a connection
* [Unpause](docs/sdks/connection/README.md#unpause) - Unpause a connection
* [Update](docs/sdks/connection/README.md#update) - Update a connection

### [ConnectionNumberUpdate](docs/sdks/connectionnumberupdate/README.md)

* [Upsert](docs/sdks/connectionnumberupdate/README.md#upsert) - Update or create a connection
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

### Example

```go
package main

import (
	"context"
	"errors"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/sdkerrors"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{})
	if err != nil {

		var e *sdkerrors.APIErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.hookdeck.com/{version}` | `version` (default is `2023-01-01`) |

#### Example

```go
package main

import (
	"context"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithServerIndex(0),
		hookdeckgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.EventAttemptPaginatedResult != nil {
		// handle response
	}
}

```

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithVersion hookdeckgo.ServerVersion`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithServerURL("https://api.hookdeck.com/{version}"),
		hookdeckgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.EventAttemptPaginatedResult != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BasicAuth`  | http         | HTTP Basic   |
| `BearerAuth` | http         | HTTP Bearer  |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Password: "<YOUR_PASSWORD_HERE>",
				Username: "<YOUR_USERNAME_HERE>",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.EventAttemptPaginatedResult != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

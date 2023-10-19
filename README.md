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

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/hookdeck-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"log"
)

func main() {
	s := hookdeckgo.New(
		hookdeckgo.WithSecurity(shared.Security{
			BasicAuth: &shared.SchemeBasicAuth{
				Password: "",
				Username: "",
			},
		}),
	)

	var id string = "female"

	ctx := context.Background()
	res, err := s.Attempt.Get(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	if res.EventAttempt != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Attempt](docs/sdks/attempt/README.md)

* [Get](docs/sdks/attempt/README.md#get) - Get a Single Attempt

### [Attempts](docs/sdks/attempts/README.md)

* [Get](docs/sdks/attempts/README.md#get) - Get Attempts

### [Bookmark](docs/sdks/bookmark/README.md)

* [Create](docs/sdks/bookmark/README.md#create) - Create a Bookmark
* [Delete](docs/sdks/bookmark/README.md#delete) - Delete a Bookmark
* [Get](docs/sdks/bookmark/README.md#get) - Get a Single Bookmark
* [Trigger](docs/sdks/bookmark/README.md#trigger) - Trigger a Bookmark
* [Update](docs/sdks/bookmark/README.md#update) - Update a Bookmark

### [Bookmarks](docs/sdks/bookmarks/README.md)

* [Get](docs/sdks/bookmarks/README.md#get) - Get Bookmarks

### [BulkRetryEvent](docs/sdks/bulkretryevent/README.md)

* [Cancel](docs/sdks/bulkretryevent/README.md#cancel) - Cancel an events bulk retry
* [Create](docs/sdks/bulkretryevent/README.md#create) - Create an events bulk retry
* [Get](docs/sdks/bulkretryevent/README.md#get) - Get an events bulk retry

### [BulkRetryIgnoredEvent](docs/sdks/bulkretryignoredevent/README.md)

* [Cancel](docs/sdks/bulkretryignoredevent/README.md#cancel) - Cancel an ignored events bulk retry
* [Create](docs/sdks/bulkretryignoredevent/README.md#create) - Create an ignored events bulk retry
* [Generate](docs/sdks/bulkretryignoredevent/README.md#generate) - Generate an ignored events bulk retry plan
* [Get](docs/sdks/bulkretryignoredevent/README.md#get) - Get an ignored events bulk retry

### [BulkRetryIgnoredEvents](docs/sdks/bulkretryignoredevents/README.md)

* [Get](docs/sdks/bulkretryignoredevents/README.md#get) - Get ignored events bulk retries

### [BulkRetryRequests](docs/sdks/bulkretryrequests/README.md)

* [Create](docs/sdks/bulkretryrequests/README.md#create) - Create a requests bulk retry

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

### [Connections](docs/sdks/connections/README.md)

* [Get](docs/sdks/connections/README.md#get) - Get connections

### [CustomDomain](docs/sdks/customdomain/README.md)

* [Add](docs/sdks/customdomain/README.md#add) - Add a custom domain to the workspace
* [Delete](docs/sdks/customdomain/README.md#delete) - Removes a custom domain from the workspace

### [CustomDomains](docs/sdks/customdomains/README.md)

* [List](docs/sdks/customdomains/README.md#list) - List all custom domains and their verification statuses for the workspace

### [Destination](docs/sdks/destination/README.md)

* [Archive](docs/sdks/destination/README.md#archive) - Archive a Destination
* [Create](docs/sdks/destination/README.md#create) - Create a Destination
* [Delete](docs/sdks/destination/README.md#delete) - Delete a Destination
* [Get](docs/sdks/destination/README.md#get) - Get a Destination
* [Unarchive](docs/sdks/destination/README.md#unarchive) - Unarchive a Destination
* [Update](docs/sdks/destination/README.md#update) - Update a Destination
* [Upsert](docs/sdks/destination/README.md#upsert) - Update or Create a Destination

### [Destinations](docs/sdks/destinations/README.md)

* [Get](docs/sdks/destinations/README.md#get) - Get Destinations

### [Event](docs/sdks/event/README.md)

* [Get](docs/sdks/event/README.md#get) - Get an Event
* [MuteEvent](docs/sdks/event/README.md#muteevent) - Mute an Event
* [Retry](docs/sdks/event/README.md#retry) - Retry an Event

### [EventRawBody](docs/sdks/eventrawbody/README.md)

* [Get](docs/sdks/eventrawbody/README.md#get) - Get an Event Raw Body Data

### [Events](docs/sdks/events/README.md)

* [Get](docs/sdks/events/README.md#get) - Get Events

### [Integration](docs/sdks/integration/README.md)

* [AttachIntegrationToSource](docs/sdks/integration/README.md#attachintegrationtosource) - Attach an integration to a source
* [Create](docs/sdks/integration/README.md#create) - Create an integration
* [Delete](docs/sdks/integration/README.md#delete) - Delete an integration
* [DetachIntegrationToSource](docs/sdks/integration/README.md#detachintegrationtosource) - Detach an integration from a source
* [Get](docs/sdks/integration/README.md#get) - Get an integration
* [Update](docs/sdks/integration/README.md#update) - Update an integration

### [Integrations](docs/sdks/integrations/README.md)

* [Get](docs/sdks/integrations/README.md#get) - Get integrations

### [Issue](docs/sdks/issue/README.md)

* [Dismiss](docs/sdks/issue/README.md#dismiss) - Dismiss an issue
* [Get](docs/sdks/issue/README.md#get) - Get a single issue
* [Update](docs/sdks/issue/README.md#update) - Update issue

### [IssueTrigger](docs/sdks/issuetrigger/README.md)

* [Create](docs/sdks/issuetrigger/README.md#create) - Create an Issue Trigger
* [Delete](docs/sdks/issuetrigger/README.md#delete) - Delete an Issue Trigger
* [Disable](docs/sdks/issuetrigger/README.md#disable) - Disable an Issue Trigger
* [Enable](docs/sdks/issuetrigger/README.md#enable) - Enable an Issue Trigger
* [Get](docs/sdks/issuetrigger/README.md#get) - Get an Issue Trigger
* [Update](docs/sdks/issuetrigger/README.md#update) - Update an Issue Trigger
* [Upsert](docs/sdks/issuetrigger/README.md#upsert) - Create or Update an Issue Trigger

### [IssueTriggers](docs/sdks/issuetriggers/README.md)

* [Get](docs/sdks/issuetriggers/README.md#get) - Get Issue Triggers

### [Issues](docs/sdks/issues/README.md)

* [Get](docs/sdks/issues/README.md#get) - Get issues

### [IssuesCount](docs/sdks/issuescount/README.md)

* [Get](docs/sdks/issuescount/README.md#get) - Get the number of issues

### [Request](docs/sdks/request/README.md)

* [Get](docs/sdks/request/README.md#get) - Get a request
* [Retry](docs/sdks/request/README.md#retry) - Retry a request

### [RequestBulkRetry](docs/sdks/requestbulkretry/README.md)

* [Cancel](docs/sdks/requestbulkretry/README.md#cancel) - Cancel a requests bulk retry
* [Get](docs/sdks/requestbulkretry/README.md#get) - Get a requests bulk retry

### [RequestEvents](docs/sdks/requestevents/README.md)

* [Get](docs/sdks/requestevents/README.md#get) - Get request events

### [RequestIgnoredEvents](docs/sdks/requestignoredevents/README.md)

* [Get](docs/sdks/requestignoredevents/README.md#get) - Get request ignored events

### [RequestRawBody](docs/sdks/requestrawbody/README.md)

* [Get](docs/sdks/requestrawbody/README.md#get) - Get a request raw body data

### [Requests](docs/sdks/requests/README.md)

* [Get](docs/sdks/requests/README.md#get) - Get requests

### [Ruleset](docs/sdks/ruleset/README.md)

* [Archive](docs/sdks/ruleset/README.md#archive) - Archive a ruleset
* [Create](docs/sdks/ruleset/README.md#create) - Create a ruleset
* [Get](docs/sdks/ruleset/README.md#get) - Get a ruleset
* [Unarchive](docs/sdks/ruleset/README.md#unarchive) - Unarchive a ruleset
* [Update](docs/sdks/ruleset/README.md#update) - Update a ruleset
* [Upsert](docs/sdks/ruleset/README.md#upsert) - Update or create a ruleset

### [Rulesets](docs/sdks/rulesets/README.md)

* [Get](docs/sdks/rulesets/README.md#get) - Get rulesets

### [Source](docs/sdks/source/README.md)

* [Archive](docs/sdks/source/README.md#archive) - Archive a source
* [Create](docs/sdks/source/README.md#create) - Create a source
* [Delete](docs/sdks/source/README.md#delete) - Delete a source
* [Get](docs/sdks/source/README.md#get) - Get a source
* [Unarchive](docs/sdks/source/README.md#unarchive) - Unarchive a source
* [Update](docs/sdks/source/README.md#update) - Update a source
* [Upsert](docs/sdks/source/README.md#upsert) - Update or create a source

### [Sources](docs/sdks/sources/README.md)

* [Get](docs/sdks/sources/README.md#get) - Get sources

### [Transformation](docs/sdks/transformation/README.md)

* [Create](docs/sdks/transformation/README.md#create) - Create a transformation
* [Get](docs/sdks/transformation/README.md#get) - Get a transformation
* [Test](docs/sdks/transformation/README.md#test) - Test a transformation code
* [Update](docs/sdks/transformation/README.md#update) - Update a transformation
* [Upsert](docs/sdks/transformation/README.md#upsert) - Update or create a transformation

### [TransformationExecution](docs/sdks/transformationexecution/README.md)

* [Get](docs/sdks/transformationexecution/README.md#get) - Get a transformation execution

### [TransformationExecutions](docs/sdks/transformationexecutions/README.md)

* [Get](docs/sdks/transformationexecutions/README.md#get) - Get transformation executions

### [Transformations](docs/sdks/transformations/README.md)

* [Get](docs/sdks/transformations/README.md#get) - Get transformations

### [WebhookNotifications](docs/sdks/webhooknotifications/README.md)

* [Toggle](docs/sdks/webhooknotifications/README.md#toggle) - Toggle webhook notifications for the workspace
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

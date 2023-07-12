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

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
)

func main() {
    s := hookdeck.New(
        hookdeck.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Attempts.GetAttempt(ctx, operations.GetAttemptRequest{
        ID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
    })
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


### [Attempts](docs/sdks/attempts/README.md)

* [GetAttempt](docs/sdks/attempts/README.md#getattempt) - Get a Single Attempt
* [GetAttempts](docs/sdks/attempts/README.md#getattempts) - Get Attempts

### [Bookmarks](docs/sdks/bookmarks/README.md)

* [CreateBookmark](docs/sdks/bookmarks/README.md#createbookmark) - Create a Bookmark
* [DeleteBookmark](docs/sdks/bookmarks/README.md#deletebookmark) - Delete a Bookmark
* [GetBookmark](docs/sdks/bookmarks/README.md#getbookmark) - Get a Single Bookmark
* [GetBookmarks](docs/sdks/bookmarks/README.md#getbookmarks) - Get Bookmarks
* [TriggerBookmark](docs/sdks/bookmarks/README.md#triggerbookmark) - Trigger a Bookmark
* [UpdateBookmark](docs/sdks/bookmarks/README.md#updatebookmark) - Update a Bookmark

### [BulkRetryEvents](docs/sdks/bulkretryevents/README.md)

* [CancelEventBulkRetry](docs/sdks/bulkretryevents/README.md#canceleventbulkretry) - Cancel an events bulk retry
* [CreateEventBulkRetry](docs/sdks/bulkretryevents/README.md#createeventbulkretry) - Create an events bulk retry
* [GetEventBulkRetry](docs/sdks/bulkretryevents/README.md#geteventbulkretry) - Get an events bulk retry

### [BulkRetryIgnoredEvents](docs/sdks/bulkretryignoredevents/README.md)

* [CancelIgnoredEventBulkRetry](docs/sdks/bulkretryignoredevents/README.md#cancelignoredeventbulkretry) - Cancel an ignored events bulk retry
* [CreateIgnoredEventBulkRetry](docs/sdks/bulkretryignoredevents/README.md#createignoredeventbulkretry) - Create an ignored events bulk retry
* [GenerateIgnoredEventBulkRetryPlan](docs/sdks/bulkretryignoredevents/README.md#generateignoredeventbulkretryplan) - Generate an ignored events bulk retry plan
* [GetIgnoredEventBulkRetries](docs/sdks/bulkretryignoredevents/README.md#getignoredeventbulkretries) - Get ignored events bulk retries
* [GetIgnoredEventBulkRetry](docs/sdks/bulkretryignoredevents/README.md#getignoredeventbulkretry) - Get an ignored events bulk retry

### [BulkRetryRequests](docs/sdks/bulkretryrequests/README.md)

* [CancelRequestBulkRetry](docs/sdks/bulkretryrequests/README.md#cancelrequestbulkretry) - Cancel a requests bulk retry
* [CreateRequestBulkRetry](docs/sdks/bulkretryrequests/README.md#createrequestbulkretry) - Create a requests bulk retry
* [GetRequestBulkRetry](docs/sdks/bulkretryrequests/README.md#getrequestbulkretry) - Get a requests bulk retry

### [Connections](docs/sdks/connections/README.md)

* [ArchiveConnection](docs/sdks/connections/README.md#archiveconnection) - Archive a connection
* [CreateConnection](docs/sdks/connections/README.md#createconnection) - Create a connection
* [DeleteConnection](docs/sdks/connections/README.md#deleteconnection) - Delete a connection
* [GetConnection](docs/sdks/connections/README.md#getconnection) - Get a single connection
* [GetConnections](docs/sdks/connections/README.md#getconnections) - Get connections
* [PauseConnection](docs/sdks/connections/README.md#pauseconnection) - Pause a connection
* [UnarchiveConnection](docs/sdks/connections/README.md#unarchiveconnection) - Unarchive a connection
* [UnpauseConnection](docs/sdks/connections/README.md#unpauseconnection) - Unpause a connection
* [UpdateConnection](docs/sdks/connections/README.md#updateconnection) - Update a connection
* [UpsertConnection](docs/sdks/connections/README.md#upsertconnection) - Update or create a connection

### [Destinations](docs/sdks/destinations/README.md)

* [ArchiveDestination](docs/sdks/destinations/README.md#archivedestination) - Archive a Destination
* [CreateDestination](docs/sdks/destinations/README.md#createdestination) - Create a Destination
* [DeleteDestination](docs/sdks/destinations/README.md#deletedestination) - Delete a Destination
* [GetDestination](docs/sdks/destinations/README.md#getdestination) - Get a Destination
* [GetDestinations](docs/sdks/destinations/README.md#getdestinations) - Get Destinations
* [UnarchiveDestination](docs/sdks/destinations/README.md#unarchivedestination) - Unarchive a Destination
* [UpdateDestination](docs/sdks/destinations/README.md#updatedestination) - Update a Destination
* [UpsertDestination](docs/sdks/destinations/README.md#upsertdestination) - Update or Create a Destination

### [Events](docs/sdks/events/README.md)

* [GetEvent](docs/sdks/events/README.md#getevent) - Get an Event
* [GetEventRawBody](docs/sdks/events/README.md#geteventrawbody) - Get an Event Raw Body Data
* [GetEvents](docs/sdks/events/README.md#getevents) - Get Events
* [MuteEvent](docs/sdks/events/README.md#muteevent) - Mute an Event
* [RetryEvent](docs/sdks/events/README.md#retryevent) - Retry an Event

### [Integrations](docs/sdks/integrations/README.md)

* [AttachIntegrationToSource](docs/sdks/integrations/README.md#attachintegrationtosource) - Attach an integration to a source
* [CreateIntegration](docs/sdks/integrations/README.md#createintegration) - Create an integration
* [DeleteIntegration](docs/sdks/integrations/README.md#deleteintegration) - Delete an integration
* [DetachIntegrationToSource](docs/sdks/integrations/README.md#detachintegrationtosource) - Detach an integration from a source
* [GetIntegration](docs/sdks/integrations/README.md#getintegration) - Get an integration
* [GetIntegrations](docs/sdks/integrations/README.md#getintegrations) - Get integrations
* [UpdateIntegration](docs/sdks/integrations/README.md#updateintegration) - Update an integration

### [IssueTriggers](docs/sdks/issuetriggers/README.md)

* [CreateIssueTrigger](docs/sdks/issuetriggers/README.md#createissuetrigger) - Create an Issue Trigger
* [DeleteIssueTrigger](docs/sdks/issuetriggers/README.md#deleteissuetrigger) - Delete an Issue Trigger
* [DisableIssueTrigger](docs/sdks/issuetriggers/README.md#disableissuetrigger) - Disable an Issue Trigger
* [EnableIssueTrigger](docs/sdks/issuetriggers/README.md#enableissuetrigger) - Enable an Issue Trigger
* [GetIssueTrigger](docs/sdks/issuetriggers/README.md#getissuetrigger) - Get an Issue Trigger
* [GetIssueTriggers](docs/sdks/issuetriggers/README.md#getissuetriggers) - Get Issue Triggers
* [UpdateIssueTrigger](docs/sdks/issuetriggers/README.md#updateissuetrigger) - Update an Issue Trigger
* [UpsertIssueTrigger](docs/sdks/issuetriggers/README.md#upsertissuetrigger) - Create or Update an Issue Trigger

### [Issues](docs/sdks/issues/README.md)

* [DismissIssue](docs/sdks/issues/README.md#dismississue) - Dismiss an issue
* [GetIssue](docs/sdks/issues/README.md#getissue) - Get a single issue
* [GetIssueCount](docs/sdks/issues/README.md#getissuecount) - Get the number of issues
* [GetIssues](docs/sdks/issues/README.md#getissues) - Get issues
* [UpdateIssue](docs/sdks/issues/README.md#updateissue) - Update issue

### [Notifications](docs/sdks/notifications/README.md)

* [AddCustomDomain](docs/sdks/notifications/README.md#addcustomdomain) - Add a custom domain to the workspace
* [DeleteCustomDomain](docs/sdks/notifications/README.md#deletecustomdomain) - Removes a custom domain from the workspace
* [ListCustomDomains](docs/sdks/notifications/README.md#listcustomdomains) - List all custom domains and their verification statuses for the workspace
* [ToggleWebhookNotifications](docs/sdks/notifications/README.md#togglewebhooknotifications) - Toggle webhook notifications for the workspace

### [Requests](docs/sdks/requests/README.md)

* [GetRequest](docs/sdks/requests/README.md#getrequest) - Get a request
* [GetRequestEvents](docs/sdks/requests/README.md#getrequestevents) - Get request events
* [GetRequestIgnoredEvents](docs/sdks/requests/README.md#getrequestignoredevents) - Get request ignored events
* [GetRequestRawBody](docs/sdks/requests/README.md#getrequestrawbody) - Get a request raw body data
* [GetRequests](docs/sdks/requests/README.md#getrequests) - Get requests
* [RetryRequest](docs/sdks/requests/README.md#retryrequest) - Retry a request

### [Rulesets](docs/sdks/rulesets/README.md)

* [ArchiveRuleset](docs/sdks/rulesets/README.md#archiveruleset) - Archive a ruleset
* [CreateRuleset](docs/sdks/rulesets/README.md#createruleset) - Create a ruleset
* [GetRuleset](docs/sdks/rulesets/README.md#getruleset) - Get a ruleset
* [GetRulesets](docs/sdks/rulesets/README.md#getrulesets) - Get rulesets
* [UnarchiveRuleset](docs/sdks/rulesets/README.md#unarchiveruleset) - Unarchive a ruleset
* [UpdateRuleset](docs/sdks/rulesets/README.md#updateruleset) - Update a ruleset
* [UpsertRuleset](docs/sdks/rulesets/README.md#upsertruleset) - Update or create a ruleset

### [Sources](docs/sdks/sources/README.md)

* [ArchiveSource](docs/sdks/sources/README.md#archivesource) - Archive a source
* [CreateSource](docs/sdks/sources/README.md#createsource) - Create a source
* [DeleteSource](docs/sdks/sources/README.md#deletesource) - Delete a source
* [GetSource](docs/sdks/sources/README.md#getsource) - Get a source
* [GetSources](docs/sdks/sources/README.md#getsources) - Get sources
* [UnarchiveSource](docs/sdks/sources/README.md#unarchivesource) - Unarchive a source
* [UpdateSource](docs/sdks/sources/README.md#updatesource) - Update a source
* [UpsertSource](docs/sdks/sources/README.md#upsertsource) - Update or create a source

### [Transformations](docs/sdks/transformations/README.md)

* [CreateTransformation](docs/sdks/transformations/README.md#createtransformation) - Create a transformation
* [GetTransformation](docs/sdks/transformations/README.md#gettransformation) - Get a transformation
* [GetTransformationExecution](docs/sdks/transformations/README.md#gettransformationexecution) - Get a transformation execution
* [GetTransformationExecutions](docs/sdks/transformations/README.md#gettransformationexecutions) - Get transformation executions
* [GetTransformations](docs/sdks/transformations/README.md#gettransformations) - Get transformations
* [TestTransformation](docs/sdks/transformations/README.md#testtransformation) - Test a transformation code
* [UpdateTransformation](docs/sdks/transformations/README.md#updatetransformation) - Update a transformation
* [UpsertTransformation](docs/sdks/transformations/README.md#upserttransformation) - Update or create a transformation
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

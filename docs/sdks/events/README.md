# Events

## Overview

An event is any request that Hookdeck receives from a source.

### Available Operations

* [Get](#get) - Get Events

## Get

Retrieve a list of requests that Hookdeck receives from a source.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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
    res, err := s.Events.Get(ctx, operations.GetEventsRequest{
        Attempts: &operations.GetEventsAttempts{},
        Body: &operations.GetEventsBody{},
        BulkRetryID: &operations.GetEventsBulkRetryID{},
        CliID: &operations.GetEventsCliID{},
        CliUserID: &operations.GetEventsCliUserID{},
        CreatedAt: &operations.GetEventsCreatedAt{},
        DestinationID: &operations.GetEventsDestinationID{},
        Dir: &operations.GetEventsDir{},
        ErrorCode: &operations.GetEventsErrorCode{},
        EventDataID: &operations.GetEventsEventDataID{},
        Headers: &operations.GetEventsHeaders{},
        ID: &operations.GetEventsID{},
        Include: operations.GetEventsIncludeData.ToPointer(),
        IssueID: &operations.GetEventsIssueID{},
        LastAttemptAt: &operations.GetEventsLastAttemptAt{},
        Limit: hookdeck.Int64(978571),
        Next: hookdeck.String("rerum"),
        OrderBy: &operations.GetEventsOrderBy{},
        ParsedQuery: &operations.GetEventsParsedQuery{},
        Path: hookdeck.String("dicta"),
        Prev: hookdeck.String("magnam"),
        ResponseStatus: &operations.GetEventsResponseStatus{},
        SearchTerm: hookdeck.String("cumque"),
        SourceID: &operations.GetEventsSourceID{},
        Status: &operations.GetEventsStatus{},
        SuccessfulAt: &operations.GetEventsSuccessfulAt{},
        WebhookID: &operations.GetEventsWebhookID{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetEventsRequest](../../models/operations/geteventsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetEventsResponse](../../models/operations/geteventsresponse.md), error**


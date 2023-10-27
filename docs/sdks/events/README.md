# Events
(*Events*)

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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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

    ctx := context.Background()
    res, err := s.Events.Get(ctx, operations.GetEventsRequest{
        Attempts: operations.CreateGetEventsAttemptsGetEventsAttempts2(
                operations.GetEventsAttempts2{},
        ),
        Body: operations.CreateGetEventsBodyStr(
        "string",
        ),
        BulkRetryID: operations.CreateGetEventsBulkRetryIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliID: operations.CreateGetEventsCliIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliUserID: operations.CreateGetEventsCliUserIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CreatedAt: operations.CreateGetEventsCreatedAtGetEventsCreatedAt2(
                operations.GetEventsCreatedAt2{},
        ),
        DestinationID: operations.CreateGetEventsDestinationIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Dir: operations.CreateGetEventsDirGetEventsDir1(
        operations.GetEventsDir1Desc,
        ),
        ErrorCode: operations.CreateGetEventsErrorCodeArrayOfstr(
                []string{
                    "string",
                },
        ),
        EventDataID: operations.CreateGetEventsEventDataIDStr(
        "string",
        ),
        Headers: operations.CreateGetEventsHeadersGetEventsHeaders2(
                operations.GetEventsHeaders2{},
        ),
        ID: operations.CreateGetEventsIDStr(
        "string",
        ),
        IssueID: operations.CreateGetEventsIssueIDStr(
        "string",
        ),
        LastAttemptAt: operations.CreateGetEventsLastAttemptAtDateTime(
        types.MustTimeFromString("2023-09-12T22:44:18.030Z"),
        ),
        OrderBy: operations.CreateGetEventsOrderByGetEventsOrderBy1(
        operations.GetEventsOrderBy1LastAttemptAt,
        ),
        ParsedQuery: operations.CreateGetEventsParsedQueryGetEventsParsedQuery2(
                operations.GetEventsParsedQuery2{},
        ),
        ResponseStatus: operations.CreateGetEventsResponseStatusInteger(
        438142,
        ),
        SourceID: operations.CreateGetEventsSourceIDStr(
        "string",
        ),
        Status: operations.CreateGetEventsStatusArrayOfEventStatus(
                []shared.EventStatus{
                    shared.EventStatusSuccessful,
                },
        ),
        SuccessfulAt: operations.CreateGetEventsSuccessfulAtDateTime(
        types.MustTimeFromString("2023-12-06T14:41:34.659Z"),
        ),
        WebhookID: operations.CreateGetEventsWebhookIDStr(
        "string",
        ),
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


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
        "program",
        ),
        BulkRetryID: operations.CreateGetEventsBulkRetryIDArrayOfstr(
                []string{
                    "Kia",
                },
        ),
        CliID: operations.CreateGetEventsCliIDStr(
        "towards",
        ),
        CliUserID: operations.CreateGetEventsCliUserIDStr(
        "Xenon",
        ),
        CreatedAt: operations.CreateGetEventsCreatedAtDateTime(
        types.MustTimeFromString("2021-10-15T18:16:01.443Z"),
        ),
        DestinationID: operations.CreateGetEventsDestinationIDArrayOfstr(
                []string{
                    "Car",
                },
        ),
        Dir: operations.CreateGetEventsDirArrayOfgetEventsDir2(
                []operations.GetEventsDir2{
                    operations.GetEventsDir2Asc,
                },
        ),
        ErrorCode: operations.CreateGetEventsErrorCodeStr(
        "Neon",
        ),
        EventDataID: operations.CreateGetEventsEventDataIDArrayOfstr(
                []string{
                    "Table",
                },
        ),
        Headers: operations.CreateGetEventsHeadersStr(
        "quasi",
        ),
        ID: operations.CreateGetEventsIDStr(
        "input",
        ),
        IssueID: operations.CreateGetEventsIssueIDArrayOfstr(
                []string{
                    "grey",
                },
        ),
        LastAttemptAt: operations.CreateGetEventsLastAttemptAtDateTime(
        types.MustTimeFromString("2023-03-29T12:14:34.236Z"),
        ),
        OrderBy: operations.CreateGetEventsOrderByArrayOfgetEventsOrderBy2(
                []operations.GetEventsOrderBy2{
                    operations.GetEventsOrderBy2NextAttemptAt,
                },
        ),
        ParsedQuery: operations.CreateGetEventsParsedQueryStr(
        "bypass",
        ),
        ResponseStatus: operations.CreateGetEventsResponseStatusInteger(
        297154,
        ),
        SourceID: operations.CreateGetEventsSourceIDStr(
        "Devolved",
        ),
        Status: operations.CreateGetEventsStatusArrayOfEventStatus(
                []shared.EventStatus{
                    shared.EventStatusQueued,
                },
        ),
        SuccessfulAt: operations.CreateGetEventsSuccessfulAtDateTime(
        types.MustTimeFromString("2023-05-05T09:16:31.781Z"),
        ),
        WebhookID: operations.CreateGetEventsWebhookIDStr(
        "Holmium",
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


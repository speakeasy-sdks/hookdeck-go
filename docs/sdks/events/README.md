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
                operations.GetEventsAttempts2{
                    AdditionalProperties: map[string]interface{}{
                        "Northeast": "Hatchback",
                    },
                },
        ),
        Body: operations.CreateGetEventsBodyMapOfany(
                map[string]interface{}{
                    "protocol": "towards",
                },
        ),
        BulkRetryID: operations.CreateGetEventsBulkRetryIDStr(
        "Xenon",
        ),
        CliID: operations.CreateGetEventsCliIDGetEventsCliID2(
                operations.GetEventsCliID2{
                    AdditionalProperties: map[string]interface{}{
                        "Car": "magnetic",
                    },
                },
        ),
        CliUserID: operations.CreateGetEventsCliUserIDStr(
        "Gasoline",
        ),
        CreatedAt: operations.CreateGetEventsCreatedAtDateTime(
        types.MustTimeFromString("2023-12-06T14:41:34.659Z"),
        ),
        DestinationID: operations.CreateGetEventsDestinationIDStr(
        "Dollar",
        ),
        Dir: operations.CreateGetEventsDirGetEventsDir1(
        operations.GetEventsDir1Asc,
        ),
        ErrorCode: operations.CreateGetEventsErrorCodeArrayOfstr(
                []string{
                    "Bicycle",
                },
        ),
        EventDataID: operations.CreateGetEventsEventDataIDStr(
        "Reduced",
        ),
        Headers: operations.CreateGetEventsHeadersMapOfany(
                map[string]interface{}{
                    "woman": "Bedfordshire",
                },
        ),
        ID: operations.CreateGetEventsIDStr(
        "cyan",
        ),
        IssueID: operations.CreateGetEventsIssueIDStr(
        "West",
        ),
        LastAttemptAt: operations.CreateGetEventsLastAttemptAtDateTime(
        types.MustTimeFromString("2021-11-08T00:36:38.602Z"),
        ),
        OrderBy: operations.CreateGetEventsOrderByArrayOfgetEventsOrderBy2(
                []operations.GetEventsOrderBy2{
                    operations.GetEventsOrderBy2CreatedAt,
                },
        ),
        ParsedQuery: operations.CreateGetEventsParsedQueryMapOfany(
                map[string]interface{}{
                    "Bicycle": "functionalities",
                },
        ),
        ResponseStatus: operations.CreateGetEventsResponseStatusArrayOfinteger(
                []int64{
                    324692,
                },
        ),
        SourceID: operations.CreateGetEventsSourceIDArrayOfstr(
                []string{
                    "Maria",
                },
        ),
        Status: operations.CreateGetEventsStatusEventStatus(
        shared.EventStatusHold,
        ),
        SuccessfulAt: operations.CreateGetEventsSuccessfulAtGetEventsSuccessfulAt2(
                operations.GetEventsSuccessfulAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Iranian": "Aaliyah",
                    },
                },
        ),
        WebhookID: operations.CreateGetEventsWebhookIDStr(
        "Stamford",
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


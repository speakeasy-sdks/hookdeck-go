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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
	"log"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Events.Get(ctx, operations.GetEventsRequest{
        Attempts: operations.CreateQueryParamAttemptsGetEventsQueryParam2(
                operations.GetEventsQueryParam2{},
        ),
        Body: operations.CreateQueryParamBodyStr(
        "string",
        ),
        BulkRetryID: operations.CreateQueryParamBulkRetryIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliID: operations.CreateQueryParamCliIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliUserID: operations.CreateQueryParamCliUserIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CreatedAt: operations.CreateGetEventsQueryParamCreatedAtGetEventsQueryParamEventsCreatedAt2(
                operations.GetEventsQueryParamEventsCreatedAt2{},
        ),
        DestinationID: operations.CreateGetEventsQueryParamDestinationIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Dir: operations.CreateGetEventsQueryParamDirGetEventsQueryParam1(
        operations.GetEventsQueryParam1Desc,
        ),
        ErrorCode: operations.CreateQueryParamErrorCodeArrayOfstr(
                []string{
                    "string",
                },
        ),
        EventDataID: operations.CreateGetEventsQueryParamEventDataIDStr(
        "string",
        ),
        Headers: operations.CreateQueryParamHeadersGetEventsQueryParamEventsHeaders2(
                operations.GetEventsQueryParamEventsHeaders2{},
        ),
        ID: operations.CreateGetEventsQueryParamIDStr(
        "string",
        ),
        IssueID: operations.CreateQueryParamIssueIDStr(
        "string",
        ),
        LastAttemptAt: operations.CreateQueryParamLastAttemptAtDateTime(
        types.MustTimeFromString("2023-09-12T22:44:18.030Z"),
        ),
        OrderBy: operations.CreateGetEventsQueryParamOrderByGetEventsQueryParamEvents1(
        operations.GetEventsQueryParamEvents1LastAttemptAt,
        ),
        ParsedQuery: operations.CreateQueryParamParsedQueryGetEventsQueryParamEventsParsedQuery2(
                operations.GetEventsQueryParamEventsParsedQuery2{},
        ),
        ResponseStatus: operations.CreateQueryParamResponseStatusInteger(
        438142,
        ),
        SourceID: operations.CreateGetEventsQueryParamSourceIDStr(
        "string",
        ),
        Status: operations.CreateQueryParamStatusArrayOfEventStatus(
                []components.EventStatus{
                    components.EventStatusSuccessful,
                },
        ),
        SuccessfulAt: operations.CreateQueryParamSuccessfulAtDateTime(
        types.MustTimeFromString("2023-12-06T14:41:34.659Z"),
        ),
        WebhookID: operations.CreateGetEventsQueryParamWebhookIDStr(
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

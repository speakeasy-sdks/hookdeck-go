# RequestEvents
(*RequestEvents*)

### Available Operations

* [Get](#get) - Get request events

## Get

Get request events

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
    res, err := s.RequestEvents.Get(ctx, operations.GetRequestEventsRequest{
        Attempts: operations.CreateGetRequestEventsAttemptsGetRequestEventsAttempts2(
                operations.GetRequestEventsAttempts2{},
        ),
        Body: operations.CreateGetRequestEventsBodyStr(
        "string",
        ),
        BulkRetryID: operations.CreateGetRequestEventsBulkRetryIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliID: operations.CreateGetRequestEventsCliIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliUserID: operations.CreateGetRequestEventsCliUserIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CreatedAt: operations.CreateGetRequestEventsCreatedAtGetRequestEventsCreatedAt2(
                operations.GetRequestEventsCreatedAt2{},
        ),
        DestinationID: operations.CreateGetRequestEventsDestinationIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Dir: operations.CreateGetRequestEventsDirGetRequestEventsDir1(
        operations.GetRequestEventsDir1Desc,
        ),
        ErrorCode: operations.CreateGetRequestEventsErrorCodeArrayOfstr(
                []string{
                    "string",
                },
        ),
        EventDataID: operations.CreateGetRequestEventsEventDataIDStr(
        "string",
        ),
        Headers: operations.CreateGetRequestEventsHeadersGetRequestEventsHeaders2(
                operations.GetRequestEventsHeaders2{},
        ),
        IDPathParameter: "string",
        IDQueryParameter: operations.CreateGetRequestEventsIDStr(
        "string",
        ),
        IssueID: operations.CreateGetRequestEventsIssueIDStr(
        "string",
        ),
        LastAttemptAt: operations.CreateGetRequestEventsLastAttemptAtDateTime(
        types.MustTimeFromString("2023-09-12T22:44:18.030Z"),
        ),
        OrderBy: operations.CreateGetRequestEventsOrderByGetRequestEventsOrderBy1(
        operations.GetRequestEventsOrderBy1LastAttemptAt,
        ),
        ParsedQuery: operations.CreateGetRequestEventsParsedQueryGetRequestEventsParsedQuery2(
                operations.GetRequestEventsParsedQuery2{},
        ),
        ResponseStatus: operations.CreateGetRequestEventsResponseStatusInteger(
        438142,
        ),
        SourceID: operations.CreateGetRequestEventsSourceIDStr(
        "string",
        ),
        Status: operations.CreateGetRequestEventsStatusArrayOfEventStatus(
                []shared.EventStatus{
                    shared.EventStatusSuccessful,
                },
        ),
        SuccessfulAt: operations.CreateGetRequestEventsSuccessfulAtDateTime(
        types.MustTimeFromString("2023-12-06T14:41:34.659Z"),
        ),
        WebhookID: operations.CreateGetRequestEventsWebhookIDStr(
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetRequestEventsRequest](../../models/operations/getrequesteventsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetRequestEventsResponse](../../models/operations/getrequesteventsresponse.md), error**


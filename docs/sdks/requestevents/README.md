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
        "program",
        ),
        BulkRetryID: operations.CreateGetRequestEventsBulkRetryIDArrayOfstr(
                []string{
                    "Kia",
                },
        ),
        CliID: operations.CreateGetRequestEventsCliIDStr(
        "towards",
        ),
        CliUserID: operations.CreateGetRequestEventsCliUserIDStr(
        "Xenon",
        ),
        CreatedAt: operations.CreateGetRequestEventsCreatedAtDateTime(
        types.MustTimeFromString("2021-10-15T18:16:01.443Z"),
        ),
        DestinationID: operations.CreateGetRequestEventsDestinationIDArrayOfstr(
                []string{
                    "Car",
                },
        ),
        Dir: operations.CreateGetRequestEventsDirArrayOfgetRequestEventsDir2(
                []operations.GetRequestEventsDir2{
                    operations.GetRequestEventsDir2Asc,
                },
        ),
        ErrorCode: operations.CreateGetRequestEventsErrorCodeStr(
        "Neon",
        ),
        EventDataID: operations.CreateGetRequestEventsEventDataIDArrayOfstr(
                []string{
                    "Table",
                },
        ),
        Headers: operations.CreateGetRequestEventsHeadersStr(
        "quasi",
        ),
        IDPathParameter: "input Hybrid",
        IDQueryParameter: operations.CreateGetRequestEventsIDStr(
        "Reduced",
        ),
        IssueID: operations.CreateGetRequestEventsIssueIDArrayOfstr(
                []string{
                    "woman",
                },
        ),
        LastAttemptAt: operations.CreateGetRequestEventsLastAttemptAtDateTime(
        types.MustTimeFromString("2021-12-20T06:15:23.515Z"),
        ),
        OrderBy: operations.CreateGetRequestEventsOrderByGetRequestEventsOrderBy1(
        operations.GetRequestEventsOrderBy1LastAttemptAt,
        ),
        ParsedQuery: operations.CreateGetRequestEventsParsedQueryStr(
        "cyan",
        ),
        ResponseStatus: operations.CreateGetRequestEventsResponseStatusInteger(
        2108,
        ),
        SourceID: operations.CreateGetRequestEventsSourceIDArrayOfstr(
                []string{
                    "City",
                },
        ),
        Status: operations.CreateGetRequestEventsStatusArrayOfEventStatus(
                []shared.EventStatus{
                    shared.EventStatusHold,
                },
        ),
        SuccessfulAt: operations.CreateGetRequestEventsSuccessfulAtGetRequestEventsSuccessfulAt2(
                operations.GetRequestEventsSuccessfulAt2{},
        ),
        WebhookID: operations.CreateGetRequestEventsWebhookIDStr(
        "female",
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


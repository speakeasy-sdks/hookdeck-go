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
                operations.GetRequestEventsAttempts2{
                    AdditionalProperties: map[string]interface{}{
                        "Northeast": "Hatchback",
                    },
                },
        ),
        Body: operations.CreateGetRequestEventsBodyMapOfany(
                map[string]interface{}{
                    "protocol": "towards",
                },
        ),
        BulkRetryID: operations.CreateGetRequestEventsBulkRetryIDStr(
        "Xenon",
        ),
        CliID: operations.CreateGetRequestEventsCliIDGetRequestEventsCliID2(
                operations.GetRequestEventsCliID2{
                    AdditionalProperties: map[string]interface{}{
                        "Car": "magnetic",
                    },
                },
        ),
        CliUserID: operations.CreateGetRequestEventsCliUserIDStr(
        "Gasoline",
        ),
        CreatedAt: operations.CreateGetRequestEventsCreatedAtDateTime(
        types.MustTimeFromString("2023-12-06T14:41:34.659Z"),
        ),
        DestinationID: operations.CreateGetRequestEventsDestinationIDStr(
        "Dollar",
        ),
        Dir: operations.CreateGetRequestEventsDirGetRequestEventsDir1(
        operations.GetRequestEventsDir1Asc,
        ),
        ErrorCode: operations.CreateGetRequestEventsErrorCodeArrayOfstr(
                []string{
                    "Bicycle",
                },
        ),
        EventDataID: operations.CreateGetRequestEventsEventDataIDStr(
        "Reduced",
        ),
        Headers: operations.CreateGetRequestEventsHeadersMapOfany(
                map[string]interface{}{
                    "woman": "Bedfordshire",
                },
        ),
        IDPathParameter: "cyan",
        IDQueryParameter: operations.CreateGetRequestEventsIDStr(
        "West",
        ),
        IssueID: operations.CreateGetRequestEventsIssueIDStr(
        "Holmium",
        ),
        LastAttemptAt: operations.CreateGetRequestEventsLastAttemptAtGetRequestEventsLastAttemptAt2(
                operations.GetRequestEventsLastAttemptAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Oklahoma": "functionalities",
                    },
                },
        ),
        OrderBy: operations.CreateGetRequestEventsOrderByArrayOfgetRequestEventsOrderBy2(
                []operations.GetRequestEventsOrderBy2{
                    operations.GetRequestEventsOrderBy2LastAttemptAt,
                },
        ),
        ParsedQuery: operations.CreateGetRequestEventsParsedQueryMapOfany(
                map[string]interface{}{
                    "Maria": "Kendall",
                },
        ),
        ResponseStatus: operations.CreateGetRequestEventsResponseStatusArrayOfinteger(
                []int64{
                    424256,
                },
        ),
        SourceID: operations.CreateGetRequestEventsSourceIDStr(
        "Delaware",
        ),
        Status: operations.CreateGetRequestEventsStatusEventStatus(
        shared.EventStatusScheduled,
        ),
        SuccessfulAt: operations.CreateGetRequestEventsSuccessfulAtGetRequestEventsSuccessfulAt2(
                operations.GetRequestEventsSuccessfulAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Alexa": "Direct",
                    },
                },
        ),
        WebhookID: operations.CreateGetRequestEventsWebhookIDArrayOfstr(
                []string{
                    "Northwest",
                },
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


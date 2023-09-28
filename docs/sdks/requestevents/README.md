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
        Attempts: &operations.GetRequestEventsAttempts{},
        Body: &operations.GetRequestEventsBody{},
        BulkRetryID: &operations.GetRequestEventsBulkRetryID{},
        CliID: &operations.GetRequestEventsCliID{},
        CliUserID: &operations.GetRequestEventsCliUserID{},
        CreatedAt: &operations.GetRequestEventsCreatedAt{},
        DestinationID: &operations.GetRequestEventsDestinationID{},
        Dir: &operations.GetRequestEventsDir{},
        ErrorCode: &operations.GetRequestEventsErrorCode{},
        EventDataID: &operations.GetRequestEventsEventDataID{},
        Headers: &operations.GetRequestEventsHeaders{},
        IDPathParameter: "id",
        IDQueryParameter: &operations.GetRequestEventsID{},
        Include: hookdeckgo.String("labore"),
        IssueID: &operations.GetRequestEventsIssueID{},
        LastAttemptAt: &operations.GetRequestEventsLastAttemptAt{},
        Limit: hookdeckgo.Int64(290077),
        Next: hookdeckgo.String("suscipit"),
        OrderBy: &operations.GetRequestEventsOrderBy{},
        ParsedQuery: &operations.GetRequestEventsParsedQuery{},
        Path: hookdeckgo.String("natus"),
        Prev: hookdeckgo.String("nobis"),
        ResponseStatus: &operations.GetRequestEventsResponseStatus{},
        SearchTerm: hookdeckgo.String("eum"),
        SourceID: &operations.GetRequestEventsSourceID{},
        Status: &operations.GetRequestEventsStatus{},
        SuccessfulAt: &operations.GetRequestEventsSuccessfulAt{},
        WebhookID: &operations.GetRequestEventsWebhookID{},
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


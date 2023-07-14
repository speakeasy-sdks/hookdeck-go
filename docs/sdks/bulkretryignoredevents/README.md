# BulkRetryIgnoredEvents

### Available Operations

* [Get](#get) - Get ignored events bulk retries

## Get

Get ignored events bulk retries

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
    res, err := s.BulkRetryIgnoredEvents.Get(ctx, operations.GetIgnoredEventBulkRetriesRequest{
        CancelledAt: &operations.GetIgnoredEventBulkRetriesCancelledAt{},
        CompletedAt: &operations.GetIgnoredEventBulkRetriesCompletedAt{},
        CreatedAt: &operations.GetIgnoredEventBulkRetriesCreatedAt{},
        Dir: &operations.GetIgnoredEventBulkRetriesDir{},
        ID: &operations.GetIgnoredEventBulkRetriesID{},
        InProgress: hookdeck.Bool(false),
        Limit: hookdeck.Int64(140350),
        Next: hookdeck.String("at"),
        OrderBy: &operations.GetIgnoredEventBulkRetriesOrderBy{},
        Prev: hookdeck.String("at"),
        Query: &operations.GetIgnoredEventBulkRetriesQuery{
            Cause: &operations.GetIgnoredEventBulkRetriesQueryCause{},
            TransformationID: hookdeck.String("maiores"),
            WebhookID: &operations.GetIgnoredEventBulkRetriesQueryWebhookID{},
        },
        QueryPartialMatch: hookdeck.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperationPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetIgnoredEventBulkRetriesRequest](../../models/operations/getignoredeventbulkretriesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetIgnoredEventBulkRetriesResponse](../../models/operations/getignoredeventbulkretriesresponse.md), error**


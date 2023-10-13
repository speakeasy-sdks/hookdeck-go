# BulkRetryIgnoredEvents
(*BulkRetryIgnoredEvents*)

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
    res, err := s.BulkRetryIgnoredEvents.Get(ctx, operations.GetIgnoredEventBulkRetriesRequest{
        CancelledAt: operations.CreateGetIgnoredEventBulkRetriesCancelledAtGetIgnoredEventBulkRetriesCancelledAt2(
                operations.GetIgnoredEventBulkRetriesCancelledAt2{},
        ),
        CompletedAt: operations.CreateGetIgnoredEventBulkRetriesCompletedAtDateTime(
        types.MustTimeFromString("2022-09-04T22:09:08.769Z"),
        ),
        CreatedAt: operations.CreateGetIgnoredEventBulkRetriesCreatedAtGetIgnoredEventBulkRetriesCreatedAt2(
                operations.GetIgnoredEventBulkRetriesCreatedAt2{},
        ),
        Dir: operations.CreateGetIgnoredEventBulkRetriesDirArrayOfgetIgnoredEventBulkRetriesDir2(
                []operations.GetIgnoredEventBulkRetriesDir2{
                    operations.GetIgnoredEventBulkRetriesDir2Desc,
                },
        ),
        ID: operations.CreateGetIgnoredEventBulkRetriesIDArrayOfstr(
                []string{
                    "Cambridgeshire",
                },
        ),
        OrderBy: operations.CreateGetIgnoredEventBulkRetriesOrderByArrayOfgetIgnoredEventBulkRetriesOrderBy2(
                []operations.GetIgnoredEventBulkRetriesOrderBy2{
                    operations.GetIgnoredEventBulkRetriesOrderBy2CreatedAt,
                },
        ),
        Query: &operations.GetIgnoredEventBulkRetriesQuery{
            Cause: operations.CreateGetIgnoredEventBulkRetriesQueryCauseStr(
            "Xenon",
            ),
            WebhookID: operations.CreateGetIgnoredEventBulkRetriesQueryWebhookIDStr(
            "Car",
            ),
        },
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


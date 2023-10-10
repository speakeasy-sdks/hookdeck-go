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
                operations.GetIgnoredEventBulkRetriesCancelledAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Northeast": "Hatchback",
                    },
                },
        ),
        CompletedAt: operations.CreateGetIgnoredEventBulkRetriesCompletedAtGetIgnoredEventBulkRetriesCompletedAt2(
                operations.GetIgnoredEventBulkRetriesCompletedAt2{
                    AdditionalProperties: map[string]interface{}{
                        "protocol": "towards",
                    },
                },
        ),
        CreatedAt: operations.CreateGetIgnoredEventBulkRetriesCreatedAtDateTime(
        types.MustTimeFromString("2023-03-17T03:19:31.784Z"),
        ),
        Dir: operations.CreateGetIgnoredEventBulkRetriesDirGetIgnoredEventBulkRetriesDir1(
        operations.GetIgnoredEventBulkRetriesDir1Asc,
        ),
        ID: operations.CreateGetIgnoredEventBulkRetriesIDStr(
        "eventually",
        ),
        OrderBy: operations.CreateGetIgnoredEventBulkRetriesOrderByGetIgnoredEventBulkRetriesOrderBy1(
        operations.GetIgnoredEventBulkRetriesOrderBy1CreatedAt,
        ),
        Query: &operations.GetIgnoredEventBulkRetriesQuery{
            AdditionalProperties: map[string]interface{}{
                "ah": "Rupiah",
            },
            Cause: operations.CreateGetIgnoredEventBulkRetriesQueryCauseArrayOfstr(
                    []string{
                        "Neon",
                    },
            ),
            WebhookID: operations.CreateGetIgnoredEventBulkRetriesQueryWebhookIDArrayOfstr(
                    []string{
                        "Table",
                    },
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


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
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/v2/types"
	"log"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.BulkRetryIgnoredEvents.Get(ctx, operations.GetIgnoredEventBulkRetriesRequest{
        CancelledAt: operations.CreateQueryParamCancelledAtGetIgnoredEventBulkRetriesQueryParam2(
                operations.GetIgnoredEventBulkRetriesQueryParam2{},
        ),
        CompletedAt: operations.CreateQueryParamCompletedAtDateTime(
        types.MustTimeFromString("2023-09-05T11:33:52.010Z"),
        ),
        CreatedAt: operations.CreateGetIgnoredEventBulkRetriesQueryParamCreatedAtGetIgnoredEventBulkRetriesQueryParamBulkRetryIgnoredEventsCreatedAt2(
                operations.GetIgnoredEventBulkRetriesQueryParamBulkRetryIgnoredEventsCreatedAt2{},
        ),
        Dir: operations.CreateGetIgnoredEventBulkRetriesQueryParamDirArrayOfgetIgnoredEventBulkRetriesQueryParamBulkRetryIgnoredEventsDir2(
                []operations.GetIgnoredEventBulkRetriesQueryParamBulkRetryIgnoredEventsDir2{
                    operations.GetIgnoredEventBulkRetriesQueryParamBulkRetryIgnoredEventsDir2Desc,
                },
        ),
        ID: operations.CreateGetIgnoredEventBulkRetriesQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        OrderBy: operations.CreateGetIgnoredEventBulkRetriesQueryParamOrderByGetIgnoredEventBulkRetriesQueryParamBulkRetryIgnoredEvents1(
        operations.GetIgnoredEventBulkRetriesQueryParamBulkRetryIgnoredEvents1CreatedAt,
        ),
        Query: &operations.GetIgnoredEventBulkRetriesQueryParamQuery{
            Cause: operations.CreateGetIgnoredEventBulkRetriesQueryParamCauseArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            WebhookID: operations.CreateGetIgnoredEventBulkRetriesQueryParamWebhookIDArrayOfstr(
                    []string{
                        "string",
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

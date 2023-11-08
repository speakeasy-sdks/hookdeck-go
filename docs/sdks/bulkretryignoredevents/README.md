# BulkRetryIgnoredEvents
(*.BulkRetryIgnoredEvents*)

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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
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
    res, err := s.BulkRetryIgnoredEvents.Get(ctx, operations.GetIgnoredEventBulkRetriesRequest{
        CancelledAt: operations.CreateQueryParamCancelledAtGetIgnoredEventBulkRetriesQueryParam2(
                operations.GetIgnoredEventBulkRetriesQueryParam2{},
        ),
        CompletedAt: operations.CreateQueryParamCompletedAtDateTime(
        types.MustTimeFromString("2022-09-04T22:09:08.769Z"),
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

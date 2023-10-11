# BulkRetryRequests
(*BulkRetryRequests*)

### Available Operations

* [Create](#create) - Create a requests bulk retry

## Create

Create a requests bulk retry

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
    res, err := s.BulkRetryRequests.Create(ctx, operations.CreateRequestBulkRetryRequestBody{
        Query: &operations.CreateRequestBulkRetryRequestBodyQuery{
            Body: operations.CreateCreateRequestBulkRetryRequestBodyQueryBodyStr(
            "bluetooth",
            ),
            BulkRetryID: operations.CreateCreateRequestBulkRetryRequestBodyQueryBulkRetryIDArrayOfstr(
                    []string{
                        "Money",
                    },
            ),
            CreatedAt: operations.CreateCreateRequestBulkRetryRequestBodyQueryCreatedAtCreateRequestBulkRetryRequestBodyQueryCreatedAt2(
                    operations.CreateRequestBulkRetryRequestBodyQueryCreatedAt2{},
            ),
            EventsCount: operations.CreateCreateRequestBulkRetryRequestBodyQueryEventsCountInteger(
            996706,
            ),
            Headers: operations.CreateCreateRequestBulkRetryRequestBodyQueryHeadersCreateRequestBulkRetryRequestBodyQueryHeaders2(
                    operations.CreateRequestBulkRetryRequestBodyQueryHeaders2{},
            ),
            ID: operations.CreateCreateRequestBulkRetryRequestBodyQueryIDStr(
            "technology",
            ),
            IgnoredCount: operations.CreateCreateRequestBulkRetryRequestBodyQueryIgnoredCountInteger(
            455222,
            ),
            IngestedAt: operations.CreateCreateRequestBulkRetryRequestBodyQueryIngestedAtDateTime(
            types.MustTimeFromString("2021-11-27T03:41:30.033Z"),
            ),
            ParsedQuery: operations.CreateCreateRequestBulkRetryRequestBodyQueryParsedQueryStr(
            "male",
            ),
            RejectionCause: operations.CreateCreateRequestBulkRetryRequestBodyQueryRejectionCauseArrayOfRequestRejectionCause(
                    []shared.RequestRejectionCause{
                        shared.RequestRejectionCauseIngestionFatal,
                    },
            ),
            SourceID: operations.CreateCreateRequestBulkRetryRequestBodyQuerySourceIDArrayOfstr(
                    []string{
                        "Screen",
                    },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateRequestBulkRetryRequestBody](../../models/operations/createrequestbulkretryrequestbody.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.CreateRequestBulkRetryResponse](../../models/operations/createrequestbulkretryresponse.md), error**


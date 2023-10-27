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
            "string",
            ),
            BulkRetryID: operations.CreateCreateRequestBulkRetryRequestBodyQueryBulkRetryIDStr(
            "string",
            ),
            CreatedAt: operations.CreateCreateRequestBulkRetryRequestBodyQueryCreatedAtCreateRequestBulkRetryRequestBodyQueryCreatedAt2(
                    operations.CreateRequestBulkRetryRequestBodyQueryCreatedAt2{},
            ),
            EventsCount: operations.CreateCreateRequestBulkRetryRequestBodyQueryEventsCountArrayOfinteger(
                    []int64{
                        417458,
                    },
            ),
            Headers: operations.CreateCreateRequestBulkRetryRequestBodyQueryHeadersStr(
            "string",
            ),
            ID: operations.CreateCreateRequestBulkRetryRequestBodyQueryIDStr(
            "string",
            ),
            IgnoredCount: operations.CreateCreateRequestBulkRetryRequestBodyQueryIgnoredCountArrayOfinteger(
                    []int64{
                        69025,
                    },
            ),
            IngestedAt: operations.CreateCreateRequestBulkRetryRequestBodyQueryIngestedAtCreateRequestBulkRetryRequestBodyQueryIngestedAt2(
                    operations.CreateRequestBulkRetryRequestBodyQueryIngestedAt2{},
            ),
            ParsedQuery: operations.CreateCreateRequestBulkRetryRequestBodyQueryParsedQueryCreateRequestBulkRetryRequestBodyQueryParsedQuery2(
                    operations.CreateRequestBulkRetryRequestBodyQueryParsedQuery2{},
            ),
            RejectionCause: operations.CreateCreateRequestBulkRetryRequestBodyQueryRejectionCauseCreateRequestBulkRetryRequestBodyQueryRejectionCause2(
                    operations.CreateRequestBulkRetryRequestBodyQueryRejectionCause2{},
            ),
            SourceID: operations.CreateCreateRequestBulkRetryRequestBodyQuerySourceIDArrayOfstr(
                    []string{
                        "string",
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


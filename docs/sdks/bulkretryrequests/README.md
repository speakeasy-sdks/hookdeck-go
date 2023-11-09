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
    res, err := s.BulkRetryRequests.Create(ctx, operations.CreateRequestBulkRetryRequestBody{
        Query: &operations.CreateRequestBulkRetryQuery{
            Body: operations.CreateCreateRequestBulkRetryBodyStr(
            "string",
            ),
            BulkRetryID: operations.CreateCreateRequestBulkRetryBulkRetryIDStr(
            "string",
            ),
            CreatedAt: operations.CreateCreateRequestBulkRetryCreatedAtCreateRequestBulkRetryBulkRetryRequests2(
                    operations.CreateRequestBulkRetryBulkRetryRequests2{},
            ),
            EventsCount: operations.CreateEventsCountArrayOfinteger(
                    []int64{
                        417458,
                    },
            ),
            Headers: operations.CreateCreateRequestBulkRetryHeadersStr(
            "string",
            ),
            ID: operations.CreateCreateRequestBulkRetryIDStr(
            "string",
            ),
            IgnoredCount: operations.CreateIgnoredCountArrayOfinteger(
                    []int64{
                        69025,
                    },
            ),
            IngestedAt: operations.CreateIngestedAtCreateRequestBulkRetryBulkRetryRequestsRequestRequestBodyQueryIngestedAt2(
                    operations.CreateRequestBulkRetryBulkRetryRequestsRequestRequestBodyQueryIngestedAt2{},
            ),
            ParsedQuery: operations.CreateCreateRequestBulkRetryParsedQueryCreateRequestBulkRetryBulkRetryRequestsRequestRequestBodyQueryParsedQuery2(
                    operations.CreateRequestBulkRetryBulkRetryRequestsRequestRequestBodyQueryParsedQuery2{},
            ),
            RejectionCause: operations.CreateRejectionCauseCreateRequestBulkRetryBulkRetryRequestsRequestRequestBodyQueryRejectionCause2(
                    operations.CreateRequestBulkRetryBulkRetryRequestsRequestRequestBodyQueryRejectionCause2{},
            ),
            SourceID: operations.CreateCreateRequestBulkRetrySourceIDArrayOfstr(
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

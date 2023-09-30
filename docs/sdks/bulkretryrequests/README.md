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
            Body: &operations.CreateRequestBulkRetryRequestBodyQueryBody{},
            BulkRetryID: &operations.CreateRequestBulkRetryRequestBodyQueryBulkRetryID{},
            CreatedAt: &operations.CreateRequestBulkRetryRequestBodyQueryCreatedAt{},
            EventsCount: &operations.CreateRequestBulkRetryRequestBodyQueryEventsCount{},
            Headers: &operations.CreateRequestBulkRetryRequestBodyQueryHeaders{},
            ID: &operations.CreateRequestBulkRetryRequestBodyQueryID{},
            IgnoredCount: &operations.CreateRequestBulkRetryRequestBodyQueryIgnoredCount{},
            IngestedAt: &operations.CreateRequestBulkRetryRequestBodyQueryIngestedAt{},
            ParsedQuery: &operations.CreateRequestBulkRetryRequestBodyQueryParsedQuery{},
            Path: hookdeckgo.String("/private/var"),
            RejectionCause: &operations.CreateRequestBulkRetryRequestBodyQueryRejectionCause{},
            SearchTerm: hookdeckgo.String("Configuration Money"),
            SourceID: &operations.CreateRequestBulkRetryRequestBodyQuerySourceID{},
            Status: operations.CreateRequestBulkRetryRequestBodyQueryStatusRejected.ToPointer(),
            Verified: hookdeckgo.Bool(false),
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


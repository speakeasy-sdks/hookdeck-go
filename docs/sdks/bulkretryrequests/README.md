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
        AdditionalProperties: map[string]interface{}{
            "online": "Configuration",
        },
        Query: &operations.CreateRequestBulkRetryRequestBodyQuery{
            AdditionalProperties: map[string]interface{}{
                "Money": "blue",
            },
            Body: operations.CreateCreateRequestBulkRetryRequestBodyQueryBodyMapOfany(
                    map[string]interface{}{
                        "grey": "technology",
                    },
            ),
            BulkRetryID: operations.CreateCreateRequestBulkRetryRequestBodyQueryBulkRetryIDStr(
            "deposit",
            ),
            CreatedAt: operations.CreateCreateRequestBulkRetryRequestBodyQueryCreatedAtDateTime(
            types.MustTimeFromString("2021-04-09T12:15:56.056Z"),
            ),
            EventsCount: operations.CreateCreateRequestBulkRetryRequestBodyQueryEventsCountArrayOfinteger(
                    []int64{
                        792620,
                    },
            ),
            Headers: operations.CreateCreateRequestBulkRetryRequestBodyQueryHeadersMapOfany(
                    map[string]interface{}{
                        "Gasoline": "Screen",
                    },
            ),
            ID: operations.CreateCreateRequestBulkRetryRequestBodyQueryIDStr(
            "physical",
            ),
            IgnoredCount: operations.CreateCreateRequestBulkRetryRequestBodyQueryIgnoredCountCreateRequestBulkRetryRequestBodyQueryIgnoredCount2(
                    operations.CreateRequestBulkRetryRequestBodyQueryIgnoredCount2{
                        AdditionalProperties: map[string]interface{}{
                            "Durham": "after",
                        },
                    },
            ),
            IngestedAt: operations.CreateCreateRequestBulkRetryRequestBodyQueryIngestedAtCreateRequestBulkRetryRequestBodyQueryIngestedAt2(
                    operations.CreateRequestBulkRetryRequestBodyQueryIngestedAt2{
                        AdditionalProperties: map[string]interface{}{
                            "Intelligent": "Fish",
                        },
                    },
            ),
            ParsedQuery: operations.CreateCreateRequestBulkRetryRequestBodyQueryParsedQueryStr(
            "Fiat",
            ),
            RejectionCause: operations.CreateCreateRequestBulkRetryRequestBodyQueryRejectionCauseArrayOfRequestRejectionCause(
                    []shared.RequestRejectionCause{
                        shared.RequestRejectionCauseNoWebhook,
                    },
            ),
            SourceID: operations.CreateCreateRequestBulkRetryRequestBodyQuerySourceIDStr(
            "Borders",
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


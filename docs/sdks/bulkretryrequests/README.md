# BulkRetryRequests

### Available Operations

* [CancelRequestBulkRetry](#cancelrequestbulkretry) - Cancel a requests bulk retry
* [CreateRequestBulkRetry](#createrequestbulkretry) - Create a requests bulk retry
* [GetRequestBulkRetry](#getrequestbulkretry) - Get a requests bulk retry

## CancelRequestBulkRetry

Cancel a requests bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.BulkRetryRequests.CancelRequestBulkRetry(ctx, operations.CancelRequestBulkRetryRequest{
        ID: "7053202c-73d5-4fe9-b90c-28909b3fe49a",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CancelRequestBulkRetryRequest](../../models/operations/cancelrequestbulkretryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CancelRequestBulkRetryResponse](../../models/operations/cancelrequestbulkretryresponse.md), error**


## CreateRequestBulkRetry

Create a requests bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.BulkRetryRequests.CreateRequestBulkRetry(ctx, operations.CreateRequestBulkRetryRequestBody{
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
            Path: hookdeck.String("deleniti"),
            RejectionCause: &operations.CreateRequestBulkRetryRequestBodyQueryRejectionCause{},
            SearchTerm: hookdeck.String("pariatur"),
            SourceID: &operations.CreateRequestBulkRetryRequestBodyQuerySourceID{},
            Status: operations.CreateRequestBulkRetryRequestBodyQueryStatusRejected.ToPointer(),
            Verified: hookdeck.Bool(false),
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


## GetRequestBulkRetry

Get a requests bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.BulkRetryRequests.GetRequestBulkRetry(ctx, operations.GetRequestBulkRetryRequest{
        ID: "cbf48633-323f-49b7-bf3a-4100674ebf69",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetRequestBulkRetryRequest](../../models/operations/getrequestbulkretryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetRequestBulkRetryResponse](../../models/operations/getrequestbulkretryresponse.md), error**


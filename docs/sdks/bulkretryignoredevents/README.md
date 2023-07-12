# BulkRetryIgnoredEvents

### Available Operations

* [CancelIgnoredEventBulkRetry](#cancelignoredeventbulkretry) - Cancel an ignored events bulk retry
* [CreateIgnoredEventBulkRetry](#createignoredeventbulkretry) - Create an ignored events bulk retry
* [GenerateIgnoredEventBulkRetryPlan](#generateignoredeventbulkretryplan) - Generate an ignored events bulk retry plan
* [GetIgnoredEventBulkRetries](#getignoredeventbulkretries) - Get ignored events bulk retries
* [GetIgnoredEventBulkRetry](#getignoredeventbulkretry) - Get an ignored events bulk retry

## CancelIgnoredEventBulkRetry

Cancel an ignored events bulk retry

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
    res, err := s.BulkRetryIgnoredEvents.CancelIgnoredEventBulkRetry(ctx, operations.CancelIgnoredEventBulkRetryRequest{
        ID: "0afa563e-2516-4fe4-88b7-11e5b7fd2ed0",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CancelIgnoredEventBulkRetryRequest](../../models/operations/cancelignoredeventbulkretryrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.CancelIgnoredEventBulkRetryResponse](../../models/operations/cancelignoredeventbulkretryresponse.md), error**


## CreateIgnoredEventBulkRetry

Create an ignored events bulk retry

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
    res, err := s.BulkRetryIgnoredEvents.CreateIgnoredEventBulkRetry(ctx, operations.CreateIgnoredEventBulkRetryRequestBody{
        Query: &operations.CreateIgnoredEventBulkRetryRequestBodyQuery{
            Cause: &operations.CreateIgnoredEventBulkRetryRequestBodyQueryCause{},
            TransformationID: hookdeck.String("consequuntur"),
            WebhookID: &operations.CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID{},
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CreateIgnoredEventBulkRetryRequestBody](../../models/operations/createignoredeventbulkretryrequestbody.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.CreateIgnoredEventBulkRetryResponse](../../models/operations/createignoredeventbulkretryresponse.md), error**


## GenerateIgnoredEventBulkRetryPlan

Generate an ignored events bulk retry plan

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
    res, err := s.BulkRetryIgnoredEvents.GenerateIgnoredEventBulkRetryPlan(ctx, operations.GenerateIgnoredEventBulkRetryPlanRequest{
        Query: &operations.GenerateIgnoredEventBulkRetryPlanQuery{
            Cause: &operations.GenerateIgnoredEventBulkRetryPlanQueryCause{},
            TransformationID: hookdeck.String("praesentium"),
            WebhookID: &operations.GenerateIgnoredEventBulkRetryPlanQueryWebhookID{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperationPlan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GenerateIgnoredEventBulkRetryPlanRequest](../../models/operations/generateignoredeventbulkretryplanrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.GenerateIgnoredEventBulkRetryPlanResponse](../../models/operations/generateignoredeventbulkretryplanresponse.md), error**


## GetIgnoredEventBulkRetries

Get ignored events bulk retries

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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
    res, err := s.BulkRetryIgnoredEvents.GetIgnoredEventBulkRetries(ctx, operations.GetIgnoredEventBulkRetriesRequest{
        CancelledAt: &operations.GetIgnoredEventBulkRetriesCancelledAt{},
        CompletedAt: &operations.GetIgnoredEventBulkRetriesCompletedAt{},
        CreatedAt: &operations.GetIgnoredEventBulkRetriesCreatedAt{},
        Dir: &operations.GetIgnoredEventBulkRetriesDir{},
        ID: &operations.GetIgnoredEventBulkRetriesID{},
        InProgress: hookdeck.Bool(false),
        Limit: hookdeck.Int64(615560),
        Next: hookdeck.String("magni"),
        OrderBy: &operations.GetIgnoredEventBulkRetriesOrderBy{},
        Prev: hookdeck.String("sunt"),
        Query: &operations.GetIgnoredEventBulkRetriesQuery{
            Cause: &operations.GetIgnoredEventBulkRetriesQueryCause{},
            TransformationID: hookdeck.String("quo"),
            WebhookID: &operations.GetIgnoredEventBulkRetriesQueryWebhookID{},
        },
        QueryPartialMatch: hookdeck.Bool(false),
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


## GetIgnoredEventBulkRetry

Get an ignored events bulk retry

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
    res, err := s.BulkRetryIgnoredEvents.GetIgnoredEventBulkRetry(ctx, operations.GetIgnoredEventBulkRetryRequest{
        ID: "ddc69260-1fb5-476b-8d5f-0d30c5fbb258",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetIgnoredEventBulkRetryRequest](../../models/operations/getignoredeventbulkretryrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetIgnoredEventBulkRetryResponse](../../models/operations/getignoredeventbulkretryresponse.md), error**


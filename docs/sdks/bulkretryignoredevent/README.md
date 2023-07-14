# BulkRetryIgnoredEvent

### Available Operations

* [Cancel](#cancel) - Cancel an ignored events bulk retry
* [Create](#create) - Create an ignored events bulk retry
* [Generate](#generate) - Generate an ignored events bulk retry plan
* [Get](#get) - Get an ignored events bulk retry

## Cancel

Cancel an ignored events bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.BulkRetryIgnoredEvent.Cancel(ctx, "ipsam")
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.CancelIgnoredEventBulkRetryResponse](../../models/operations/cancelignoredeventbulkretryresponse.md), error**


## Create

Create an ignored events bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.BulkRetryIgnoredEvent.Create(ctx, operations.CreateIgnoredEventBulkRetryRequestBody{
        Query: &operations.CreateIgnoredEventBulkRetryRequestBodyQuery{
            Cause: &operations.CreateIgnoredEventBulkRetryRequestBodyQueryCause{},
            TransformationID: hookdeck.String("repellendus"),
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


## Generate

Generate an ignored events bulk retry plan

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.BulkRetryIgnoredEvent.Generate(ctx, &operations.GenerateIgnoredEventBulkRetryPlanQuery{
        Cause: &operations.GenerateIgnoredEventBulkRetryPlanQueryCause{},
        TransformationID: hookdeck.String("sapiente"),
        WebhookID: &operations.GenerateIgnoredEventBulkRetryPlanQueryWebhookID{},
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

| Parameter                                                                                                               | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                   | :heavy_check_mark:                                                                                                      | The context to use for the request.                                                                                     |
| `query`                                                                                                                 | [*operations.GenerateIgnoredEventBulkRetryPlanQuery](../../models/operations/generateignoredeventbulkretryplanquery.md) | :heavy_minus_sign:                                                                                                      | Filter by the bulk retry ignored event query object                                                                     |


### Response

**[*operations.GenerateIgnoredEventBulkRetryPlanResponse](../../models/operations/generateignoredeventbulkretryplanresponse.md), error**


## Get

Get an ignored events bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.BulkRetryIgnoredEvent.Get(ctx, "quo")
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetIgnoredEventBulkRetryResponse](../../models/operations/getignoredeventbulkretryresponse.md), error**


# BulkRetryIgnoredEvent
(*.BulkRetryIgnoredEvent*)

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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
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


    var id string = "string"

    ctx := context.Background()
    res, err := s.BulkRetryIgnoredEvent.Cancel(ctx, id)
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
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
    res, err := s.BulkRetryIgnoredEvent.Create(ctx, operations.CreateIgnoredEventBulkRetryRequestBody{
        Query: &operations.CreateIgnoredEventBulkRetryQuery{
            Cause: operations.CreateCauseStr(
            "string",
            ),
            WebhookID: operations.CreateCreateIgnoredEventBulkRetryWebhookIDStr(
            "string",
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
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


    query := &operations.QueryParamQuery{
        Cause: operations.CreateQueryParamCauseArrayOfstr(
                []string{
                    "string",
                },
        ),
        WebhookID: operations.CreateGenerateIgnoredEventBulkRetryPlanQueryParamWebhookIDArrayOfstr(
                []string{
                    "string",
                },
        ),
    }

    ctx := context.Background()
    res, err := s.BulkRetryIgnoredEvent.Generate(ctx, query)
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperationPlan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `query`                                                                   | [*operations.QueryParamQuery](../../models/operations/queryparamquery.md) | :heavy_minus_sign:                                                        | Filter by the bulk retry ignored event query object                       |


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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
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


    var id string = "string"

    ctx := context.Background()
    res, err := s.BulkRetryIgnoredEvent.Get(ctx, id)
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


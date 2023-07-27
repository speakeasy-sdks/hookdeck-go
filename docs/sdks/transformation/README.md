# Transformation

### Available Operations

* [Create](#create) - Create a transformation
* [Get](#get) - Get a transformation
* [Test](#test) - Test a transformation code
* [Update](#update) - Update a transformation
* [Upsert](#upsert) - Update or create a transformation

## Create

Create a transformation

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
    res, err := s.Transformation.Create(ctx, operations.CreateTransformationRequestBody{
        Code: "quod",
        Env: map[string]interface{}{
            "qui": "dolorum",
            "a": "esse",
            "harum": "iusto",
            "ipsum": "quisquam",
        },
        Name: "Marvin Renner",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transformation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateTransformationRequestBody](../../models/operations/createtransformationrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.CreateTransformationResponse](../../models/operations/createtransformationresponse.md), error**


## Get

Get a transformation

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
    res, err := s.Transformation.Get(ctx, "enim")
    if err != nil {
        log.Fatal(err)
    }

    if res.Transformation != nil {
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

**[*operations.GetTransformationResponse](../../models/operations/gettransformationresponse.md), error**


## Test

Test a transformation code

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
    res, err := s.Transformation.Test(ctx, operations.TestTransformationRequestBody{
        Code: hookdeck.String("dolorem"),
        Env: &operations.TestTransformationRequestBodyEnv{},
        EventID: hookdeck.String("sapiente"),
        Request: &operations.TestTransformationRequestBodyRequest{
            Body: &operations.TestTransformationRequestBodyRequestBody{},
            Headers: map[string]string{
                "nihil": "sit",
                "expedita": "neque",
                "sed": "vel",
            },
            ParsedQuery: &operations.TestTransformationRequestBodyRequestParsedQuery{},
            Path: hookdeck.String("libero"),
            Query: hookdeck.String("voluptas"),
        },
        TransformationID: hookdeck.String("deserunt"),
        WebhookID: hookdeck.String("quam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransformationExecutorOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.TestTransformationRequestBody](../../models/operations/testtransformationrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.TestTransformationResponse](../../models/operations/testtransformationresponse.md), error**


## Update

Update a transformation

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
    res, err := s.Transformation.Update(ctx, operations.UpdateTransformationRequestBody{
        Code: hookdeck.String("ipsum"),
        Env: map[string]interface{}{
            "qui": "cupiditate",
            "maxime": "pariatur",
        },
        Name: hookdeck.String("Keith Padberg"),
    }, "aspernatur")
    if err != nil {
        log.Fatal(err)
    }

    if res.Transformation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `requestBody`                                                                                            | [operations.UpdateTransformationRequestBody](../../models/operations/updatetransformationrequestbody.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `id`                                                                                                     | *string*                                                                                                 | :heavy_check_mark:                                                                                       | N/A                                                                                                      |


### Response

**[*operations.UpdateTransformationResponse](../../models/operations/updatetransformationresponse.md), error**


## Upsert

Update or create a transformation

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
    res, err := s.Transformation.Upsert(ctx, operations.UpsertTransformationRequestBody{
        Code: "dolores",
        Env: map[string]interface{}{
            "facilis": "aliquid",
            "quam": "molestias",
            "temporibus": "qui",
        },
        Name: "Beverly Cummerata II",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transformation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpsertTransformationRequestBody](../../models/operations/upserttransformationrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.UpsertTransformationResponse](../../models/operations/upserttransformationresponse.md), error**


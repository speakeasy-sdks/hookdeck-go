# Transformation
(*Transformation*)

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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformation.Create(ctx, operations.CreateTransformationRequestBody{
        Code: "bluetooth Extended",
        Env: map[string]interface{}{
            "aspernatur": "blue",
        },
        Name: "grey technology East",
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    id := "female"

    ctx := context.Background()
    res, err := s.Transformation.Get(ctx, id)
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformation.Test(ctx, operations.TestTransformationRequestBody{
        Code: hookdeckgo.String("Home CSS Liaison"),
        Env: &operations.TestTransformationRequestBodyEnv{},
        EventID: hookdeckgo.String("transmit"),
        Request: &operations.TestTransformationRequestBodyRequest{
            Body: &operations.TestTransformationRequestBodyRequestBody{},
            Headers: map[string]string{
                "maiores": "Card",
            },
            ParsedQuery: &operations.TestTransformationRequestBodyRequestParsedQuery{},
            Path: hookdeckgo.String("/net"),
            Query: hookdeckgo.String("edible"),
        },
        TransformationID: hookdeckgo.String("withdrawal Gasoline Group"),
        WebhookID: hookdeckgo.String("Officer Cotton Intelligent"),
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    requestBody := operations.UpdateTransformationRequestBody{
        Code: hookdeckgo.String("New Reactive dock"),
        Env: map[string]interface{}{
            "similique": "Analyst",
        },
        Name: hookdeckgo.String("cheater Islands"),
    }
    id := "online"

    ctx := context.Background()
    res, err := s.Transformation.Update(ctx, requestBody, id)
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformation.Upsert(ctx, operations.UpsertTransformationRequestBody{
        Code: "Pound extranet",
        Env: map[string]interface{}{
            "iusto": "kilogram",
        },
        Name: "Triathlon/Time",
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


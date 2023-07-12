# Transformations

## Overview

A transformation represents JavaScript code that will be executed on a connection's requests. Transformations are applied to connections using Rules.

### Available Operations

* [CreateTransformation](#createtransformation) - Create a transformation
* [GetTransformation](#gettransformation) - Get a transformation
* [GetTransformationExecution](#gettransformationexecution) - Get a transformation execution
* [GetTransformationExecutions](#gettransformationexecutions) - Get transformation executions
* [GetTransformations](#gettransformations) - Get transformations
* [TestTransformation](#testtransformation) - Test a transformation code
* [UpdateTransformation](#updatetransformation) - Update a transformation
* [UpsertTransformation](#upserttransformation) - Update or create a transformation

## CreateTransformation

Create a transformation

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.CreateTransformation(ctx, operations.CreateTransformationRequestBody{
        Code: "velit",
        Env: map[string]interface{}{
            "provident": "consectetur",
            "eligendi": "dignissimos",
            "consectetur": "soluta",
            "natus": "temporibus",
        },
        Name: "Marvin White",
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


## GetTransformation

Get a transformation

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.GetTransformation(ctx, operations.GetTransformationRequest{
        ID: "eda7e23f-2257-4411-baf4-b7544e472e80",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetTransformationRequest](../../models/operations/gettransformationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetTransformationResponse](../../models/operations/gettransformationresponse.md), error**


## GetTransformationExecution

Get a transformation execution

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.GetTransformationExecution(ctx, operations.GetTransformationExecutionRequest{
        ExecutionID: "odit",
        ID: "857a5b40-463a-47d5-b5f1-400e764ad733",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransformationExecution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetTransformationExecutionRequest](../../models/operations/gettransformationexecutionrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetTransformationExecutionResponse](../../models/operations/gettransformationexecutionresponse.md), error**


## GetTransformationExecutions

Get transformation executions

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/types"
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
    res, err := s.Transformations.GetTransformationExecutions(ctx, operations.GetTransformationExecutionsRequest{
        CreatedAt: &operations.GetTransformationExecutionsCreatedAt{},
        Dir: &operations.GetTransformationExecutionsDir{},
        ID: "4ec1b781-b36a-4080-88d1-00efada200ef",
        IssueID: &operations.GetTransformationExecutionsIssueID{},
        Limit: hookdeck.Int64(27982),
        LogLevel: &operations.GetTransformationExecutionsLogLevel{},
        Next: hookdeck.String("incidunt"),
        OrderBy: &operations.GetTransformationExecutionsOrderBy{},
        Prev: hookdeck.String("qui"),
        WebhookID: &operations.GetTransformationExecutionsWebhookID{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransformationExecutionPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetTransformationExecutionsRequest](../../models/operations/gettransformationexecutionsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetTransformationExecutionsResponse](../../models/operations/gettransformationexecutionsresponse.md), error**


## GetTransformations

Get transformations

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.GetTransformations(ctx, operations.GetTransformationsRequest{
        Dir: &operations.GetTransformationsDir{},
        ID: &operations.GetTransformationsID{},
        Limit: hookdeck.Int64(185897),
        Name: &operations.GetTransformationsName{},
        Next: hookdeck.String("necessitatibus"),
        OrderBy: &operations.GetTransformationsOrderBy{},
        Prev: hookdeck.String("harum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransformationPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetTransformationsRequest](../../models/operations/gettransformationsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetTransformationsResponse](../../models/operations/gettransformationsresponse.md), error**


## TestTransformation

Test a transformation code

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.TestTransformation(ctx, operations.TestTransformationRequestBody{
        Code: hookdeck.String("explicabo"),
        Env: &operations.TestTransformationRequestBodyEnv{},
        EventID: hookdeck.String("beatae"),
        Request: &operations.TestTransformationRequestBodyRequest{
            Body: &operations.TestTransformationRequestBodyRequestBody{},
            Headers: map[string]string{
                "modi": "optio",
                "voluptatibus": "molestias",
            },
            ParsedQuery: &operations.TestTransformationRequestBodyRequestParsedQuery{},
            Path: hookdeck.String("officia"),
            Query: hookdeck.String("libero"),
        },
        TransformationID: hookdeck.String("totam"),
        WebhookID: hookdeck.String("sequi"),
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


## UpdateTransformation

Update a transformation

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.UpdateTransformation(ctx, operations.UpdateTransformationRequest{
        RequestBody: operations.UpdateTransformationRequestBody{
            Code: hookdeck.String("aliquid"),
            Env: map[string]interface{}{
                "impedit": "ducimus",
                "odit": "velit",
            },
            Name: hookdeck.String("Jerald Stoltenberg"),
        },
        ID: "e06bee48-25c1-4fc0-a115-c80bff918544",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateTransformationRequest](../../models/operations/updatetransformationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpdateTransformationResponse](../../models/operations/updatetransformationresponse.md), error**


## UpsertTransformation

Update or create a transformation

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.UpsertTransformation(ctx, operations.UpsertTransformationRequestBody{
        Code: "itaque",
        Env: map[string]interface{}{
            "modi": "consequuntur",
            "assumenda": "vero",
            "doloribus": "impedit",
            "porro": "accusamus",
        },
        Name: "Abel Bergstrom",
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


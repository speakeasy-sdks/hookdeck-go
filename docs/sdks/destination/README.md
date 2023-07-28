# Destination

### Available Operations

* [Archive](#archive) - Archive a Destination
* [Create](#create) - Create a Destination
* [Delete](#delete) - Delete a Destination
* [Get](#get) - Get a Destination
* [Unarchive](#unarchive) - Unarchive a Destination
* [Update](#update) - Update a Destination
* [Upsert](#upsert) - Update or Create a Destination

## Archive

Archive an unused endpoint.

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
    id := "quos"

    ctx := context.Background()
    res, err := s.Destination.Archive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
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

**[*operations.ArchiveDestinationResponse](../../models/operations/archivedestinationresponse.md), error**


## Create

Create a new endpoint to which your webhooks can be routed.

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
    res, err := s.Destination.Create(ctx, operations.CreateDestinationRequestBody{
        AuthMethod: &shared.HookdeckSignature{
            Config: &shared.DestinationAuthMethodSignatureConfig{},
            Type: shared.HookdeckSignatureTypeHookdeckSignature,
        },
        CliPath: hookdeck.String("magni"),
        HTTPMethod: shared.DestinationHTTPMethodPatch.ToPointer(),
        Name: "Linda Corkery",
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(703737),
        RateLimitPeriod: operations.CreateDestinationRequestBodyRateLimitPeriodHour.ToPointer(),
        URL: hookdeck.String("labore"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateDestinationRequestBody](../../models/operations/createdestinationrequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.CreateDestinationResponse](../../models/operations/createdestinationresponse.md), error**


## Delete

Delete an endpoint to which your webhooks can be routed.

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
    id := "delectus"

    ctx := context.Background()
    res, err := s.Destination.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDestination200ApplicationJSONObject != nil {
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

**[*operations.DeleteDestinationResponse](../../models/operations/deletedestinationresponse.md), error**


## Get

Retrieve an endpoint to which your webhooks can be routed.

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
    id := "eum"

    ctx := context.Background()
    res, err := s.Destination.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
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

**[*operations.GetDestinationResponse](../../models/operations/getdestinationresponse.md), error**


## Unarchive

Unarchive an endpoint.

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
    id := "non"

    ctx := context.Background()
    res, err := s.Destination.Unarchive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
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

**[*operations.UnarchiveDestinationResponse](../../models/operations/unarchivedestinationresponse.md), error**


## Update

Update an existing endpoint to which your webhooks can be routed.

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
    requestBody := operations.UpdateDestinationRequestBody{
        AuthMethod: &shared.BearerToken{
            Config: &shared.DestinationAuthMethodBearerTokenConfig{
                Token: "sint",
            },
            Type: shared.BearerTokenTypeBearerToken,
        },
        CliPath: hookdeck.String("aliquid"),
        HTTPMethod: shared.DestinationHTTPMethodPut.ToPointer(),
        Name: hookdeck.String("Perry Nikolaus"),
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(680056),
        RateLimitPeriod: operations.UpdateDestinationRequestBodyRateLimitPeriodMinute.ToPointer(),
        URL: hookdeck.String("in"),
    }
    id := "illum"

    ctx := context.Background()
    res, err := s.Destination.Update(ctx, requestBody, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `requestBody`                                                                                      | [operations.UpdateDestinationRequestBody](../../models/operations/updatedestinationrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |


### Response

**[*operations.UpdateDestinationResponse](../../models/operations/updatedestinationresponse.md), error**


## Upsert

Update or create a new endpoint to which your webhooks can be routed.

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
    res, err := s.Destination.Upsert(ctx, operations.UpsertDestinationRequestBody{
        AuthMethod: &shared.CustomSignature{
            Config: shared.DestinationAuthMethodCustomSignatureConfig{
                Key: "rerum",
                SigningSecret: hookdeck.String("dicta"),
            },
            Type: shared.CustomSignatureTypeCustomSignature,
        },
        CliPath: hookdeck.String("magnam"),
        HTTPMethod: shared.DestinationHTTPMethodPatch.ToPointer(),
        Name: "Nathaniel Hyatt",
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(581273),
        RateLimitPeriod: operations.UpsertDestinationRequestBodyRateLimitPeriodSecond.ToPointer(),
        URL: hookdeck.String("accusamus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpsertDestinationRequestBody](../../models/operations/upsertdestinationrequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.UpsertDestinationResponse](../../models/operations/upsertdestinationresponse.md), error**


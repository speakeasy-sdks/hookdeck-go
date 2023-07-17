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

    ctx := context.Background()
    res, err := s.Destination.Archive(ctx, "magni")
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
        AuthMethod: &shared.CustomSignature{
            Config: shared.DestinationAuthMethodCustomSignatureConfig{
                Key: "ipsam",
                SigningSecret: hookdeck.String("alias"),
            },
            Type: shared.CustomSignatureTypeCustomSignature,
        },
        CliPath: hookdeck.String("fugit"),
        HTTPMethod: shared.DestinationHTTPMethodPatch.ToPointer(),
        Name: "Eddie Prosacco",
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(433288),
        RateLimitPeriod: operations.CreateDestinationRequestBodyRateLimitPeriodSecond.ToPointer(),
        URL: hookdeck.String("eligendi"),
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

    ctx := context.Background()
    res, err := s.Destination.Delete(ctx, "sint")
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

    ctx := context.Background()
    res, err := s.Destination.Get(ctx, "aliquid")
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

    ctx := context.Background()
    res, err := s.Destination.Unarchive(ctx, "provident")
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

    ctx := context.Background()
    res, err := s.Destination.Update(ctx, operations.UpdateDestinationRequestBody{
        AuthMethod: &shared.CustomSignature{
            Config: shared.DestinationAuthMethodCustomSignatureConfig{
                Key: "sint",
                SigningSecret: hookdeck.String("officia"),
            },
            Type: shared.CustomSignatureTypeCustomSignature,
        },
        CliPath: hookdeck.String("dolor"),
        HTTPMethod: shared.DestinationHTTPMethodDelete.ToPointer(),
        Name: hookdeck.String("Wilbur King"),
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(699479),
        RateLimitPeriod: operations.UpdateDestinationRequestBodyRateLimitPeriodSecond.ToPointer(),
        URL: hookdeck.String("magnam"),
    }, "cumque")
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
                Key: "ea",
                SigningSecret: hookdeck.String("aliquid"),
            },
            Type: shared.CustomSignatureTypeCustomSignature,
        },
        CliPath: hookdeck.String("laborum"),
        HTTPMethod: shared.DestinationHTTPMethodDelete.ToPointer(),
        Name: "Toni Haley",
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(588465),
        RateLimitPeriod: operations.UpsertDestinationRequestBodyRateLimitPeriodHour.ToPointer(),
        URL: hookdeck.String("id"),
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


# Destinations

## Overview

A destination is any endpoint to which your webhooks can be routed.

### Available Operations

* [ArchiveDestination](#archivedestination) - Archive a Destination
* [CreateDestination](#createdestination) - Create a Destination
* [DeleteDestination](#deletedestination) - Delete a Destination
* [GetDestination](#getdestination) - Get a Destination
* [GetDestinations](#getdestinations) - Get Destinations
* [UnarchiveDestination](#unarchivedestination) - Unarchive a Destination
* [UpdateDestination](#updatedestination) - Update a Destination
* [UpsertDestination](#upsertdestination) - Update or Create a Destination

## ArchiveDestination

Archive an unused endpoint.

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
    res, err := s.Destinations.ArchiveDestination(ctx, operations.ArchiveDestinationRequest{
        ID: "c0ab3c20-c4f3-4789-bd87-1f99dd2efd12",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ArchiveDestinationRequest](../../models/operations/archivedestinationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ArchiveDestinationResponse](../../models/operations/archivedestinationresponse.md), error**


## CreateDestination

Create a new endpoint to which your webhooks can be routed.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/models/shared"
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
    res, err := s.Destinations.CreateDestination(ctx, operations.CreateDestinationRequestBody{
        AuthMethod: &shared.HookdeckSignature{
            Config: &shared.DestinationAuthMethodSignatureConfig{},
            Type: shared.HookdeckSignatureTypeHookdeckSignature,
        },
        CliPath: hookdeck.String("similique"),
        HTTPMethod: shared.DestinationHTTPMethodPut.ToPointer(),
        Name: "Mandy Berge",
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(258684),
        RateLimitPeriod: operations.CreateDestinationRequestBodyRateLimitPeriodHour.ToPointer(),
        URL: hookdeck.String("illum"),
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


## DeleteDestination

Delete an endpoint to which your webhooks can be routed.

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
    res, err := s.Destinations.DeleteDestination(ctx, operations.DeleteDestinationRequest{
        ID: "b04f1575-6082-4d68-aa19-f1d17051339d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDestination200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteDestinationRequest](../../models/operations/deletedestinationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteDestinationResponse](../../models/operations/deletedestinationresponse.md), error**


## GetDestination

Retrieve an endpoint to which your webhooks can be routed.

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
    res, err := s.Destinations.GetDestination(ctx, operations.GetDestinationRequest{
        ID: "08086a18-4039-44c2-a071-f93f5f0642da",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetDestinationRequest](../../models/operations/getdestinationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetDestinationResponse](../../models/operations/getdestinationresponse.md), error**


## GetDestinations

Retrieve a list of endpoints to which your webhooks can be routed.

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
    res, err := s.Destinations.GetDestinations(ctx, operations.GetDestinationsRequest{
        Archived: hookdeck.Bool(false),
        ArchivedAt: &operations.GetDestinationsArchivedAt{},
        CliPath: &operations.GetDestinationsCliPath{},
        Dir: &operations.GetDestinationsDir{},
        ID: &operations.GetDestinationsID{},
        Limit: hookdeck.Int64(807023),
        Name: &operations.GetDestinationsName{},
        Next: hookdeck.String("dignissimos"),
        OrderBy: &operations.GetDestinationsOrderBy{},
        Prev: hookdeck.String("officia"),
        URL: &operations.GetDestinationsURL{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DestinationPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDestinationsRequest](../../models/operations/getdestinationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetDestinationsResponse](../../models/operations/getdestinationsresponse.md), error**


## UnarchiveDestination

Unarchive an endpoint.

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
    res, err := s.Destinations.UnarchiveDestination(ctx, operations.UnarchiveDestinationRequest{
        ID: "f515cc41-3aa6-43aa-a8d6-7864dbb675fd",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UnarchiveDestinationRequest](../../models/operations/unarchivedestinationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UnarchiveDestinationResponse](../../models/operations/unarchivedestinationresponse.md), error**


## UpdateDestination

Update an existing endpoint to which your webhooks can be routed.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/models/shared"
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
    res, err := s.Destinations.UpdateDestination(ctx, operations.UpdateDestinationRequest{
        RequestBody: operations.UpdateDestinationRequestBody{
            AuthMethod: &shared.BasicAuth{
                Config: &shared.DestinationAuthMethodBasicAuthConfig{
                    Password: "recusandae",
                    Username: "Graham.Bayer44",
                },
                Type: shared.BasicAuthTypeBasicAuth,
            },
            CliPath: hookdeck.String("exercitationem"),
            HTTPMethod: shared.DestinationHTTPMethodDelete.ToPointer(),
            Name: hookdeck.String("Jesus Yost"),
            PathForwardingDisabled: hookdeck.Bool(false),
            RateLimit: hookdeck.Int64(904949),
            RateLimitPeriod: operations.UpdateDestinationRequestBodyRateLimitPeriodHour.ToPointer(),
            URL: hookdeck.String("dolore"),
        },
        ID: "1f33317f-e35b-460e-b1ea-426555ba3c28",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateDestinationRequest](../../models/operations/updatedestinationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateDestinationResponse](../../models/operations/updatedestinationresponse.md), error**


## UpsertDestination

Update or create a new endpoint to which your webhooks can be routed.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/models/shared"
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
    res, err := s.Destinations.UpsertDestination(ctx, operations.UpsertDestinationRequestBody{
        AuthMethod: &shared.APIKey{
            Config: &shared.DestinationAuthMethodAPIKeyConfig{
                APIKey: "dolore",
                Key: "aliquam",
                To: shared.DestinationAuthMethodAPIKeyConfigToQuery.ToPointer(),
            },
            Type: shared.APIKeyTypeAPIKey,
        },
        CliPath: hookdeck.String("temporibus"),
        HTTPMethod: shared.DestinationHTTPMethodPost.ToPointer(),
        Name: "Karla Kuvalis",
        PathForwardingDisabled: hookdeck.Bool(false),
        RateLimit: hookdeck.Int64(633998),
        RateLimitPeriod: operations.UpsertDestinationRequestBodyRateLimitPeriodMinute.ToPointer(),
        URL: hookdeck.String("pariatur"),
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


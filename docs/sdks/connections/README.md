# Connections

## Overview

A connection lets you route webhooks from a source to a destination, using a ruleset.

### Available Operations

* [ArchiveConnection](#archiveconnection) - Archive a connection
* [CreateConnection](#createconnection) - Create a connection
* [DeleteConnection](#deleteconnection) - Delete a connection
* [GetConnection](#getconnection) - Get a single connection
* [GetConnections](#getconnections) - Get connections
* [PauseConnection](#pauseconnection) - Pause a connection
* [UnarchiveConnection](#unarchiveconnection) - Unarchive a connection
* [UnpauseConnection](#unpauseconnection) - Unpause a connection
* [UpdateConnection](#updateconnection) - Update a connection
* [UpsertConnection](#upsertconnection) - Update or create a connection

## ArchiveConnection

Archive a connection

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
    res, err := s.Connections.ArchiveConnection(ctx, operations.ArchiveConnectionRequest{
        ID: "280d1ba7-7a89-4ebf-b37a-e4203ce5e6a9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ArchiveConnectionRequest](../../models/operations/archiveconnectionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ArchiveConnectionResponse](../../models/operations/archiveconnectionresponse.md), error**


## CreateConnection

Create a connection

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
    res, err := s.Connections.CreateConnection(ctx, operations.CreateConnectionRequestBody{
        Destination: &operations.CreateConnectionRequestBodyDestination{
            AuthMethod: &shared.BasicAuth{
                Config: &shared.DestinationAuthMethodBasicAuthConfig{
                    Password: "repellendus",
                    Username: "Josianne87",
                },
                Type: shared.BasicAuthTypeBasicAuth,
            },
            CliPath: hookdeck.String("quaerat"),
            HTTPMethod: shared.DestinationHTTPMethodGet.ToPointer(),
            Name: "Bernadette Torp",
            PathForwardingDisabled: hookdeck.Bool(false),
            RateLimit: hookdeck.Int64(456130),
            RateLimitPeriod: operations.CreateConnectionRequestBodyDestinationRateLimitPeriodHour.ToPointer(),
            URL: hookdeck.String("iusto"),
        },
        DestinationID: hookdeck.String("ipsum"),
        Name: "Saul Fay",
        Rules: []interface{}{
            shared.AlertRule{
                Strategy: shared.AlertStrategyEachAttempt,
                Type: shared.AlertRuleTypeAlert,
            },
            shared.DelayRule{
                Delay: 518201,
                Type: shared.DelayRuleTypeDelay,
            },
        },
        Ruleset: &operations.CreateConnectionRequestBodyRuleset{
            IsTeamDefault: hookdeck.Bool(false),
            Name: "Karen Rath",
            Rules: []interface{}{
                shared.TransformReference{
                    TransformationID: "deserunt",
                    Type: shared.TransformReferenceTypeTransform,
                },
                shared.FilterRule{
                    Body: &shared.ConnectionFilterProperty{},
                    Headers: &shared.ConnectionFilterProperty{},
                    Path: &shared.ConnectionFilterProperty{},
                    Query: &shared.ConnectionFilterProperty{},
                    Type: shared.FilterRuleTypeFilter,
                },
            },
        },
        RulesetID: hookdeck.String("ipsum"),
        Source: &operations.CreateConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodGet,
                shared.SourceAllowedHTTPMethodPut,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "maxime",
                ContentType: shared.SourceCustomResponseContentTypeXML,
            },
            Name: "Keith Padberg",
        },
        SourceID: hookdeck.String("aspernatur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateConnectionRequestBody](../../models/operations/createconnectionrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateConnectionResponse](../../models/operations/createconnectionresponse.md), error**


## DeleteConnection

Delete a connection

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
    res, err := s.Connections.DeleteConnection(ctx, operations.DeleteConnectionRequest{
        ID: "2bb679d2-3227-415b-b0cb-b1e31b8b90f3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteConnection200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteConnectionRequest](../../models/operations/deleteconnectionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteConnectionResponse](../../models/operations/deleteconnectionresponse.md), error**


## GetConnection

Get a single connection

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
    res, err := s.Connections.GetConnection(ctx, operations.GetConnectionRequest{
        ID: "443a1108-e0ad-4cf4-b921-879fce953f73",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetConnectionRequest](../../models/operations/getconnectionrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetConnectionResponse](../../models/operations/getconnectionresponse.md), error**


## GetConnections

Get connections

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
    res, err := s.Connections.GetConnections(ctx, operations.GetConnectionsRequest{
        Archived: hookdeck.Bool(false),
        ArchivedAt: &operations.GetConnectionsArchivedAt{},
        DestinationID: &operations.GetConnectionsDestinationID{},
        Dir: &operations.GetConnectionsDir{},
        FullName: hookdeck.String("vero"),
        ID: &operations.GetConnectionsID{},
        Limit: hookdeck.Int64(949319),
        Name: &operations.GetConnectionsName{},
        Next: hookdeck.String("dignissimos"),
        OrderBy: &operations.GetConnectionsOrderBy{},
        PausedAt: &operations.GetConnectionsPausedAt{},
        Prev: hookdeck.String("hic"),
        SourceID: &operations.GetConnectionsSourceID{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetConnectionsRequest](../../models/operations/getconnectionsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetConnectionsResponse](../../models/operations/getconnectionsresponse.md), error**


## PauseConnection

Pause a connection

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
    res, err := s.Connections.PauseConnection(ctx, operations.PauseConnectionRequest{
        ID: "bc7abd74-dd39-4c0f-9d2c-ff7c70a45626",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PauseConnectionRequest](../../models/operations/pauseconnectionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PauseConnectionResponse](../../models/operations/pauseconnectionresponse.md), error**


## UnarchiveConnection

Unarchive a connection

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
    res, err := s.Connections.UnarchiveConnection(ctx, operations.UnarchiveConnectionRequest{
        ID: "d436813f-16d9-4f5f-8e6c-556146c3e250",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UnarchiveConnectionRequest](../../models/operations/unarchiveconnectionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UnarchiveConnectionResponse](../../models/operations/unarchiveconnectionresponse.md), error**


## UnpauseConnection

Unpause a connection

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
    res, err := s.Connections.UnpauseConnection(ctx, operations.UnpauseConnectionRequest{
        ID: "fb008c42-e141-4aac-b66c-8dd6b1442907",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UnpauseConnectionRequest](../../models/operations/unpauseconnectionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UnpauseConnectionResponse](../../models/operations/unpauseconnectionresponse.md), error**


## UpdateConnection

Update a connection

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
    res, err := s.Connections.UpdateConnection(ctx, operations.UpdateConnectionRequest{
        RequestBody: operations.UpdateConnectionRequestBody{
            Name: hookdeck.String("Viola Gibson"),
            Rules: []interface{}{
                shared.TransformReference{
                    TransformationID: "quidem",
                    Type: shared.TransformReferenceTypeTransform,
                },
                shared.DelayRule{
                    Delay: 283519,
                    Type: shared.DelayRuleTypeDelay,
                },
                shared.FilterRule{
                    Body: &shared.ConnectionFilterProperty{},
                    Headers: &shared.ConnectionFilterProperty{},
                    Path: &shared.ConnectionFilterProperty{},
                    Query: &shared.ConnectionFilterProperty{},
                    Type: shared.FilterRuleTypeFilter,
                },
            },
            Ruleset: &operations.UpdateConnectionRequestBodyRuleset{
                IsTeamDefault: hookdeck.Bool(false),
                Name: "Angelina Davis",
                Rules: []interface{}{
                    shared.RetryRule{
                        Count: hookdeck.Int64(660040),
                        Interval: hookdeck.Int64(696997),
                        Strategy: shared.RetryStrategyLinear,
                        Type: shared.RetryRuleTypeRetry,
                    },
                },
            },
            RulesetID: hookdeck.String("quo"),
        },
        ID: "dca42519-04e5-423c-be0b-c7178e4796f2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateConnectionRequest](../../models/operations/updateconnectionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateConnectionResponse](../../models/operations/updateconnectionresponse.md), error**


## UpsertConnection

Update or create a connection

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
    res, err := s.Connections.UpsertConnection(ctx, operations.UpsertConnectionRequestBody{
        Destination: &operations.UpsertConnectionRequestBodyDestination{
            AuthMethod: &shared.BearerToken{
                Config: &shared.DestinationAuthMethodBearerTokenConfig{
                    Token: "molestiae",
                },
                Type: shared.BearerTokenTypeBearerToken,
            },
            CliPath: hookdeck.String("accusantium"),
            HTTPMethod: shared.DestinationHTTPMethodPatch.ToPointer(),
            Name: "Dianne Langosh",
            PathForwardingDisabled: hookdeck.Bool(false),
            RateLimit: hookdeck.Int64(681393),
            RateLimitPeriod: operations.UpsertConnectionRequestBodyDestinationRateLimitPeriodMinute.ToPointer(),
            URL: hookdeck.String("incidunt"),
        },
        DestinationID: hookdeck.String("atque"),
        Name: "Cathy Huel",
        Rules: []interface{}{
            shared.RetryRule{
                Count: hookdeck.Int64(129412),
                Interval: hookdeck.Int64(903984),
                Strategy: shared.RetryStrategyExponential,
                Type: shared.RetryRuleTypeRetry,
            },
        },
        Ruleset: &operations.UpsertConnectionRequestBodyRuleset{
            IsTeamDefault: hookdeck.Bool(false),
            Name: "Carl Koch",
            Rules: []interface{}{
                shared.FilterRule{
                    Body: &shared.ConnectionFilterProperty{},
                    Headers: &shared.ConnectionFilterProperty{},
                    Path: &shared.ConnectionFilterProperty{},
                    Query: &shared.ConnectionFilterProperty{},
                    Type: shared.FilterRuleTypeFilter,
                },
            },
        },
        RulesetID: hookdeck.String("quod"),
        Source: &operations.UpsertConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodDelete,
                shared.SourceAllowedHTTPMethodPost,
                shared.SourceAllowedHTTPMethodGet,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "saepe",
                ContentType: shared.SourceCustomResponseContentTypeText,
            },
            Name: "Javier Price",
        },
        SourceID: hookdeck.String("distinctio"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpsertConnectionRequestBody](../../models/operations/upsertconnectionrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpsertConnectionResponse](../../models/operations/upsertconnectionresponse.md), error**


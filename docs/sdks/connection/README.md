# Connection

### Available Operations

* [Archive](#archive) - Archive a connection
* [Create](#create) - Create a connection
* [Delete](#delete) - Delete a connection
* [Get](#get) - Get a single connection
* [Pause](#pause) - Pause a connection
* [Unarchive](#unarchive) - Unarchive a connection
* [Unpause](#unpause) - Unpause a connection
* [Update](#update) - Update a connection

## Archive

Archive a connection

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
    id := "esse"

    ctx := context.Background()
    res, err := s.Connection.Archive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.ArchiveConnectionResponse](../../models/operations/archiveconnectionresponse.md), error**


## Create

Create a connection

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
    res, err := s.Connection.Create(ctx, operations.CreateConnectionRequestBody{
        Destination: &operations.CreateConnectionRequestBodyDestination{
            AuthMethod: &shared.APIKey{
                Config: &shared.DestinationAuthMethodAPIKeyConfig{
                    APIKey: "porro",
                    Key: "dolorum",
                    To: shared.DestinationAuthMethodAPIKeyConfigToHeader.ToPointer(),
                },
                Type: shared.APIKeyTypeAPIKey,
            },
            CliPath: hookdeck.String("nam"),
            HTTPMethod: shared.DestinationHTTPMethodPut.ToPointer(),
            Name: "Wayne Lind",
            PathForwardingDisabled: hookdeck.Bool(false),
            RateLimit: hookdeck.Int64(105907),
            RateLimitPeriod: operations.CreateConnectionRequestBodyDestinationRateLimitPeriodMinute.ToPointer(),
            URL: hookdeck.String("molestiae"),
        },
        DestinationID: hookdeck.String("modi"),
        Name: "Krista Rippin",
        Rules: []interface{}{
            shared.FilterRule{
                Body: &shared.ConnectionFilterProperty{},
                Headers: &shared.ConnectionFilterProperty{},
                Path: &shared.ConnectionFilterProperty{},
                Query: &shared.ConnectionFilterProperty{},
                Type: shared.FilterRuleTypeFilter,
            },
        },
        Ruleset: &operations.CreateConnectionRequestBodyRuleset{
            IsTeamDefault: hookdeck.Bool(false),
            Name: "Dorothy Hane",
            Rules: []interface{}{
                shared.TransformReference{
                    TransformationID: "natus",
                    Type: shared.TransformReferenceTypeTransform,
                },
            },
        },
        RulesetID: hookdeck.String("laboriosam"),
        Source: &operations.CreateConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodDelete,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "saepe",
                ContentType: shared.SourceCustomResponseContentTypeXML,
            },
            Name: "Stacy Moore",
        },
        SourceID: hookdeck.String("quidem"),
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


## Delete

Delete a connection

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
    id := "architecto"

    ctx := context.Background()
    res, err := s.Connection.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteConnection200ApplicationJSONObject != nil {
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

**[*operations.DeleteConnectionResponse](../../models/operations/deleteconnectionresponse.md), error**


## Get

Get a single connection

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
    id := "ipsa"

    ctx := context.Background()
    res, err := s.Connection.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.GetConnectionResponse](../../models/operations/getconnectionresponse.md), error**


## Pause

Pause a connection

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
    id := "reiciendis"

    ctx := context.Background()
    res, err := s.Connection.Pause(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.PauseConnectionResponse](../../models/operations/pauseconnectionresponse.md), error**


## Unarchive

Unarchive a connection

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
    id := "est"

    ctx := context.Background()
    res, err := s.Connection.Unarchive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.UnarchiveConnectionResponse](../../models/operations/unarchiveconnectionresponse.md), error**


## Unpause

Unpause a connection

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
    id := "mollitia"

    ctx := context.Background()
    res, err := s.Connection.Unpause(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.UnpauseConnectionResponse](../../models/operations/unpauseconnectionresponse.md), error**


## Update

Update a connection

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
    requestBody := operations.UpdateConnectionRequestBody{
        Name: hookdeck.String("Ernest Ebert"),
        Rules: []interface{}{
            shared.TransformReference{
                TransformationID: "omnis",
                Type: shared.TransformReferenceTypeTransform,
            },
        },
        Ruleset: &operations.UpdateConnectionRequestBodyRuleset{
            IsTeamDefault: hookdeck.Bool(false),
            Name: "Ms. Cathy Marks",
            Rules: []interface{}{
                shared.DelayRule{
                    Delay: 958950,
                    Type: shared.DelayRuleTypeDelay,
                },
            },
        },
        RulesetID: hookdeck.String("architecto"),
    }
    id := "mollitia"

    ctx := context.Background()
    res, err := s.Connection.Update(ctx, requestBody, id)
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
| `requestBody`                                                                                    | [operations.UpdateConnectionRequestBody](../../models/operations/updateconnectionrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |


### Response

**[*operations.UpdateConnectionResponse](../../models/operations/updateconnectionresponse.md), error**


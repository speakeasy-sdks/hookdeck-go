# Ruleset

### Available Operations

* [Archive](#archive) - Archive a ruleset
* [Create](#create) - Create a ruleset
* [Get](#get) - Get a ruleset
* [Unarchive](#unarchive) - Unarchive a ruleset
* [Update](#update) - Update a ruleset
* [Upsert](#upsert) - Update or create a ruleset

## Archive

Archive a ruleset

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
    res, err := s.Ruleset.Archive(ctx, "asperiores")
    if err != nil {
        log.Fatal(err)
    }

    if res.Ruleset != nil {
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

**[*operations.ArchiveRulesetResponse](../../models/operations/archiverulesetresponse.md), error**


## Create

Create a ruleset

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
    res, err := s.Ruleset.Create(ctx, operations.CreateRulesetRequestBody{
        IsTeamDefault: hookdeck.Bool(false),
        Name: "Edwin Morar",
        Rules: []interface{}{
            shared.FilterRule{
                Body: &shared.ConnectionFilterProperty{},
                Headers: &shared.ConnectionFilterProperty{},
                Path: &shared.ConnectionFilterProperty{},
                Query: &shared.ConnectionFilterProperty{},
                Type: shared.FilterRuleTypeFilter,
            },
            shared.TransformFull{
                Transformation: &shared.TransformFullTransformation{
                    Code: "delectus",
                    Env: map[string]string{
                        "quos": "aliquid",
                        "dolorem": "dolorem",
                    },
                    Name: "Norma Erdman",
                },
                TransformationID: hookdeck.String("cum"),
                Type: shared.TransformFullTypeTransform,
            },
            shared.FilterRule{
                Body: &shared.ConnectionFilterProperty{},
                Headers: &shared.ConnectionFilterProperty{},
                Path: &shared.ConnectionFilterProperty{},
                Query: &shared.ConnectionFilterProperty{},
                Type: shared.FilterRuleTypeFilter,
            },
            shared.FilterRule{
                Body: &shared.ConnectionFilterProperty{},
                Headers: &shared.ConnectionFilterProperty{},
                Path: &shared.ConnectionFilterProperty{},
                Query: &shared.ConnectionFilterProperty{},
                Type: shared.FilterRuleTypeFilter,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Ruleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateRulesetRequestBody](../../models/operations/createrulesetrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateRulesetResponse](../../models/operations/createrulesetresponse.md), error**


## Get

Get a ruleset

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
    res, err := s.Ruleset.Get(ctx, "reiciendis")
    if err != nil {
        log.Fatal(err)
    }

    if res.Ruleset != nil {
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

**[*operations.GetRulesetResponse](../../models/operations/getrulesetresponse.md), error**


## Unarchive

Unarchive a ruleset

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
    res, err := s.Ruleset.Unarchive(ctx, "amet")
    if err != nil {
        log.Fatal(err)
    }

    if res.Ruleset != nil {
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

**[*operations.UnarchiveRulesetResponse](../../models/operations/unarchiverulesetresponse.md), error**


## Update

Update a ruleset

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
    res, err := s.Ruleset.Update(ctx, operations.UpdateRulesetRequestBody{
        IsTeamDefault: hookdeck.Bool(false),
        Name: hookdeck.String("Mr. Bradley Bogan"),
        Rules: []interface{}{
            shared.AlertRule{
                Strategy: shared.AlertStrategyLastAttempt,
                Type: shared.AlertRuleTypeAlert,
            },
            shared.TransformFull{
                Transformation: &shared.TransformFullTransformation{
                    Code: "voluptas",
                    Env: map[string]string{
                        "eos": "atque",
                        "sit": "fugiat",
                        "ab": "soluta",
                    },
                    Name: "Ted Kling",
                },
                TransformationID: hookdeck.String("omnis"),
                Type: shared.TransformFullTypeTransform,
            },
        },
    }, "necessitatibus")
    if err != nil {
        log.Fatal(err)
    }

    if res.Ruleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `requestBody`                                                                              | [operations.UpdateRulesetRequestBody](../../models/operations/updaterulesetrequestbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |


### Response

**[*operations.UpdateRulesetResponse](../../models/operations/updaterulesetresponse.md), error**


## Upsert

Update or create a ruleset

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
    res, err := s.Ruleset.Upsert(ctx, operations.UpsertRulesetRequestBody{
        IsTeamDefault: hookdeck.Bool(false),
        Name: "Emmett Kovacek",
        Rules: []interface{}{
            shared.DelayRule{
                Delay: 263322,
                Type: shared.DelayRuleTypeDelay,
            },
            shared.RetryRule{
                Count: hookdeck.Int64(20651),
                Interval: hookdeck.Int64(229219),
                Strategy: shared.RetryStrategyExponential,
                Type: shared.RetryRuleTypeRetry,
            },
            shared.DelayRule{
                Delay: 320017,
                Type: shared.DelayRuleTypeDelay,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Ruleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpsertRulesetRequestBody](../../models/operations/upsertrulesetrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpsertRulesetResponse](../../models/operations/upsertrulesetresponse.md), error**

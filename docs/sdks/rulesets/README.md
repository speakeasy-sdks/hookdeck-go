# Rulesets

## Overview

A ruleset defines a group of rules that can be used across many connections.

### Available Operations

* [ArchiveRuleset](#archiveruleset) - Archive a ruleset
* [CreateRuleset](#createruleset) - Create a ruleset
* [GetRuleset](#getruleset) - Get a ruleset
* [GetRulesets](#getrulesets) - Get rulesets
* [UnarchiveRuleset](#unarchiveruleset) - Unarchive a ruleset
* [UpdateRuleset](#updateruleset) - Update a ruleset
* [UpsertRuleset](#upsertruleset) - Update or create a ruleset

## ArchiveRuleset

Archive a ruleset

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
    res, err := s.Rulesets.ArchiveRuleset(ctx, operations.ArchiveRulesetRequest{
        ID: "084cb067-2d1a-4d87-9eeb-9665b85efbd0",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ArchiveRulesetRequest](../../models/operations/archiverulesetrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ArchiveRulesetResponse](../../models/operations/archiverulesetresponse.md), error**


## CreateRuleset

Create a ruleset

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
    res, err := s.Rulesets.CreateRuleset(ctx, operations.CreateRulesetRequestBody{
        IsTeamDefault: hookdeck.Bool(false),
        Name: "Yvette Paucek MD",
        Rules: []interface{}{
            shared.RetryRule{
                Count: hookdeck.Int64(844854),
                Interval: hookdeck.Int64(483518),
                Strategy: shared.RetryStrategyExponential,
                Type: shared.RetryRuleTypeRetry,
            },
            shared.RetryRule{
                Count: hookdeck.Int64(127688),
                Interval: hookdeck.Int64(358995),
                Strategy: shared.RetryStrategyExponential,
                Type: shared.RetryRuleTypeRetry,
            },
            shared.DelayRule{
                Delay: 239337,
                Type: shared.DelayRuleTypeDelay,
            },
            shared.DelayRule{
                Delay: 630871,
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
| `request`                                                                                  | [operations.CreateRulesetRequestBody](../../models/operations/createrulesetrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateRulesetResponse](../../models/operations/createrulesetresponse.md), error**


## GetRuleset

Get a ruleset

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
    res, err := s.Rulesets.GetRuleset(ctx, operations.GetRulesetRequest{
        ID: "4b5197f9-2443-4da7-8e52-b895c537c645",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetRulesetRequest](../../models/operations/getrulesetrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetRulesetResponse](../../models/operations/getrulesetresponse.md), error**


## GetRulesets

Get rulesets

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
    res, err := s.Rulesets.GetRulesets(ctx, operations.GetRulesetsRequest{
        Archived: hookdeck.Bool(false),
        ArchivedAt: &operations.GetRulesetsArchivedAt{},
        Dir: &operations.GetRulesetsDir{},
        ID: &operations.GetRulesetsID{},
        Limit: hookdeck.Int64(298264),
        Name: &operations.GetRulesetsName2{
            Any: hookdeck.Bool(false),
            Contains: &operations.GetRulesetsName2Contains{},
            Gt: &operations.GetRulesetsName2Gt{},
            Gte: &operations.GetRulesetsName2Gte{},
            Le: &operations.GetRulesetsName2Le{},
            Lte: &operations.GetRulesetsName2Lte{},
        },
        Next: hookdeck.String("maiores"),
        OrderBy: &operations.GetRulesetsOrderBy{},
        Prev: hookdeck.String("tempore"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RulesetPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetRulesetsRequest](../../models/operations/getrulesetsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetRulesetsResponse](../../models/operations/getrulesetsresponse.md), error**


## UnarchiveRuleset

Unarchive a ruleset

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
    res, err := s.Rulesets.UnarchiveRuleset(ctx, operations.UnarchiveRulesetRequest{
        ID: "0b34896c-3ca5-4acf-be2f-d57075779291",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UnarchiveRulesetRequest](../../models/operations/unarchiverulesetrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UnarchiveRulesetResponse](../../models/operations/unarchiverulesetresponse.md), error**


## UpdateRuleset

Update a ruleset

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
    res, err := s.Rulesets.UpdateRuleset(ctx, operations.UpdateRulesetRequest{
        RequestBody: operations.UpdateRulesetRequestBody{
            IsTeamDefault: hookdeck.Bool(false),
            Name: hookdeck.String("Colleen Streich"),
            Rules: []interface{}{
                shared.FilterRule{
                    Body: &shared.ConnectionFilterProperty{},
                    Headers: &shared.ConnectionFilterProperty{},
                    Path: &shared.ConnectionFilterProperty{},
                    Query: &shared.ConnectionFilterProperty{},
                    Type: shared.FilterRuleTypeFilter,
                },
                shared.AlertRule{
                    Strategy: shared.AlertStrategyEachAttempt,
                    Type: shared.AlertRuleTypeAlert,
                },
                shared.DelayRule{
                    Delay: 810839,
                    Type: shared.DelayRuleTypeDelay,
                },
                shared.TransformReference{
                    TransformationID: "quam",
                    Type: shared.TransformReferenceTypeTransform,
                },
            },
        },
        ID: "3409e3eb-1e5a-42b1-aeb0-7f116db99545",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateRulesetRequest](../../models/operations/updaterulesetrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateRulesetResponse](../../models/operations/updaterulesetresponse.md), error**


## UpsertRuleset

Update or create a ruleset

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
    res, err := s.Rulesets.UpsertRuleset(ctx, operations.UpsertRulesetRequestBody{
        IsTeamDefault: hookdeck.Bool(false),
        Name: "Guillermo Marvin",
        Rules: []interface{}{
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
| `request`                                                                                  | [operations.UpsertRulesetRequestBody](../../models/operations/upsertrulesetrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpsertRulesetResponse](../../models/operations/upsertrulesetresponse.md), error**


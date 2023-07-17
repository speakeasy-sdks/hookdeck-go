# Rulesets

## Overview

A ruleset defines a group of rules that can be used across many connections.

### Available Operations

* [Get](#get) - Get rulesets

## Get

Get rulesets

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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
    res, err := s.Rulesets.Get(ctx, operations.GetRulesetsRequest{
        Archived: hookdeck.Bool(false),
        ArchivedAt: &operations.GetRulesetsArchivedAt{},
        Dir: &operations.GetRulesetsDir{},
        ID: &operations.GetRulesetsID{},
        Limit: hookdeck.Int64(896672),
        Name: &operations.GetRulesetsName2{
            Any: hookdeck.Bool(false),
            Contains: &operations.GetRulesetsName2Contains{},
            Gt: &operations.GetRulesetsName2Gt{},
            Gte: &operations.GetRulesetsName2Gte{},
            Le: &operations.GetRulesetsName2Le{},
            Lte: &operations.GetRulesetsName2Lte{},
        },
        Next: hookdeck.String("asperiores"),
        OrderBy: &operations.GetRulesetsOrderBy{},
        Prev: hookdeck.String("nihil"),
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


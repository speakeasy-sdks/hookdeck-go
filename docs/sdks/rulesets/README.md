# Rulesets
(*Rulesets*)

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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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
    res, err := s.Rulesets.Get(ctx, operations.GetRulesetsRequest{
        Archived: hookdeckgo.Bool(false),
        ArchivedAt: &operations.GetRulesetsArchivedAt{},
        Dir: &operations.GetRulesetsDir{},
        ID: &operations.GetRulesetsID{},
        Limit: hookdeckgo.Int64(700347),
        Name: &operations.GetRulesetsName{
            Any: hookdeckgo.Bool(false),
            Contains: &operations.GetRulesetsNameContains{},
            Gt: &operations.GetRulesetsNameGt{},
            Gte: &operations.GetRulesetsNameGte{},
            Le: &operations.GetRulesetsNameLe{},
            Lte: &operations.GetRulesetsNameLte{},
        },
        Next: hookdeckgo.String("program"),
        OrderBy: &operations.GetRulesetsOrderBy{},
        Prev: hookdeckgo.String("Kia Cambridgeshire"),
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


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
        ArchivedAt: operations.CreateGetRulesetsArchivedAtGetRulesetsArchivedAt2(
                operations.GetRulesetsArchivedAt2{},
        ),
        Dir: operations.CreateGetRulesetsDirGetRulesetsDir1(
        operations.GetRulesetsDir1Desc,
        ),
        ID: operations.CreateGetRulesetsIDArrayOfstr(
                []string{
                    "transmit",
                },
        ),
        Name: operations.CreateGetRulesetsNameGetRulesetsName2(
                operations.GetRulesetsName2{
                    Contains: operations.CreateGetRulesetsName2ContainsStr(
                    "towards",
                    ),
                    Gt: operations.CreateGetRulesetsName2GtStr(
                    "Xenon",
                    ),
                    Gte: operations.CreateGetRulesetsName2GteStr(
                    "Car",
                    ),
                    Le: operations.CreateGetRulesetsName2LeStr(
                    "Rupiah",
                    ),
                    Lte: operations.CreateGetRulesetsName2LteArrayOfstr(
                            []string{
                                "Neon",
                            },
                    ),
                },
        ),
        OrderBy: operations.CreateGetRulesetsOrderByArrayOfgetRulesetsOrderBy2(
                []operations.GetRulesetsOrderBy2{
                    operations.GetRulesetsOrderBy2CreatedAt,
                },
        ),
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


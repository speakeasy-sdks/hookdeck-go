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
                    "string",
                },
        ),
        Name: operations.CreateGetRulesetsNameGetRulesetsName2(
                operations.GetRulesetsName2{
                    Contains: operations.CreateGetRulesetsName2ContainsArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                    Gt: operations.CreateGetRulesetsName2GtArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                    Gte: operations.CreateGetRulesetsName2GteStr(
                    "string",
                    ),
                    Le: operations.CreateGetRulesetsName2LeArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                    Lte: operations.CreateGetRulesetsName2LteArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                },
        ),
        OrderBy: operations.CreateGetRulesetsOrderByGetRulesetsOrderBy1(
        operations.GetRulesetsOrderBy1CreatedAt,
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


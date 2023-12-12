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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"log"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Rulesets.Get(ctx, operations.GetRulesetsRequest{
        ArchivedAt: operations.CreateGetRulesetsQueryParamArchivedAtGetRulesetsQueryParam2(
                operations.GetRulesetsQueryParam2{},
        ),
        Dir: operations.CreateGetRulesetsQueryParamDirGetRulesetsQueryParam1(
        operations.GetRulesetsQueryParam1Desc,
        ),
        ID: operations.CreateGetRulesetsQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Name: operations.CreateGetRulesetsQueryParamNameGetRulesetsQueryParamRulesetsName2(
                operations.GetRulesetsQueryParamRulesetsName2{
                    Contains: operations.CreateContainsArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                    Gt: operations.CreateGtArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                    Gte: operations.CreateGteStr(
                    "string",
                    ),
                    Le: operations.CreateLeArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                    Lte: operations.CreateLteArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                },
        ),
        OrderBy: operations.CreateGetRulesetsQueryParamOrderByGetRulesetsQueryParamRulesetsOrderBy1(
        operations.GetRulesetsQueryParamRulesetsOrderBy1CreatedAt,
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

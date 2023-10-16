# Sources
(*Sources*)

## Overview

A source represents any third party that sends webhooks to Hookdeck.

### Available Operations

* [Get](#get) - Get sources

## Get

Get sources

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
    res, err := s.Sources.Get(ctx, operations.GetSourcesRequest{
        ArchivedAt: operations.CreateGetSourcesArchivedAtGetSourcesArchivedAt2(
                operations.GetSourcesArchivedAt2{},
        ),
        Dir: operations.CreateGetSourcesDirGetSourcesDir1(
        operations.GetSourcesDir1Desc,
        ),
        ID: operations.CreateGetSourcesIDArrayOfstr(
                []string{
                    "transmit",
                },
        ),
        IntegrationID: operations.CreateGetSourcesIntegrationIDGetSourcesIntegrationID2(
                operations.GetSourcesIntegrationID2{},
        ),
        Name: operations.CreateGetSourcesNameStr(
        "towards",
        ),
        OrderBy: operations.CreateGetSourcesOrderByGetSourcesOrderBy1(
        operations.GetSourcesOrderBy1CreatedAt,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SourcePaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetSourcesRequest](../../models/operations/getsourcesrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetSourcesResponse](../../models/operations/getsourcesresponse.md), error**


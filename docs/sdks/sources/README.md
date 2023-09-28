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
        Archived: hookdeckgo.Bool(false),
        ArchivedAt: &operations.GetSourcesArchivedAt{},
        Dir: &operations.GetSourcesDir{},
        ID: &operations.GetSourcesID{},
        IntegrationID: &operations.GetSourcesIntegrationID{},
        Limit: hookdeckgo.Int64(975522),
        Name: &operations.GetSourcesName{},
        Next: hookdeckgo.String("perferendis"),
        OrderBy: &operations.GetSourcesOrderBy{},
        Prev: hookdeckgo.String("fugiat"),
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


# Sources

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
    res, err := s.Sources.Get(ctx, operations.GetSourcesRequest{
        Archived: hookdeck.Bool(false),
        ArchivedAt: &operations.GetSourcesArchivedAt{},
        Dir: &operations.GetSourcesDir{},
        ID: &operations.GetSourcesID{},
        IntegrationID: &operations.GetSourcesIntegrationID{},
        Limit: hookdeck.Int64(452109),
        Name: &operations.GetSourcesName{},
        Next: hookdeck.String("dignissimos"),
        OrderBy: &operations.GetSourcesOrderBy{},
        Prev: hookdeck.String("reiciendis"),
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


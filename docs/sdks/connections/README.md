# Connections

## Overview

A connection lets you route webhooks from a source to a destination, using a ruleset.

### Available Operations

* [Get](#get) - Get connections

## Get

Get connections

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
    res, err := s.Connections.Get(ctx, operations.GetConnectionsRequest{
        Archived: hookdeck.Bool(false),
        ArchivedAt: &operations.GetConnectionsArchivedAt{},
        DestinationID: &operations.GetConnectionsDestinationID{},
        Dir: &operations.GetConnectionsDir{},
        FullName: hookdeck.String("deserunt"),
        ID: &operations.GetConnectionsID{},
        Name: &operations.GetConnectionsName{},
        OrderBy: &operations.GetConnectionsOrderBy{},
        PausedAt: &operations.GetConnectionsPausedAt{},
        SourceID: &operations.GetConnectionsSourceID{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConnectionPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetConnectionsRequest](../../models/operations/getconnectionsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetConnectionsResponse](../../models/operations/getconnectionsresponse.md), error**


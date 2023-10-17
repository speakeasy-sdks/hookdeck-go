# Connections
(*Connections*)

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
    res, err := s.Connections.Get(ctx, operations.GetConnectionsRequest{
        ArchivedAt: operations.CreateGetConnectionsArchivedAtGetConnectionsArchivedAt2(
                operations.GetConnectionsArchivedAt2{},
        ),
        DestinationID: operations.CreateGetConnectionsDestinationIDStr(
        "program",
        ),
        Dir: operations.CreateGetConnectionsDirArrayOfgetConnectionsDir2(
                []operations.GetConnectionsDir2{
                    operations.GetConnectionsDir2Desc,
                },
        ),
        ID: operations.CreateGetConnectionsIDArrayOfstr(
                []string{
                    "Cambridgeshire",
                },
        ),
        Name: operations.CreateGetConnectionsNameGetConnectionsName2(
                operations.GetConnectionsName2{},
        ),
        OrderBy: operations.CreateGetConnectionsOrderByGetConnectionsOrderBy1(
        operations.GetConnectionsOrderBy1CreatedAt,
        ),
        PausedAt: operations.CreateGetConnectionsPausedAtGetConnectionsPausedAt2(
                operations.GetConnectionsPausedAt2{},
        ),
        SourceID: operations.CreateGetConnectionsSourceIDStr(
        "Account",
        ),
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


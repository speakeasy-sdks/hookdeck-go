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
        "string",
        ),
        Dir: operations.CreateGetConnectionsDirArrayOfgetConnectionsDir2(
                []operations.GetConnectionsDir2{
                    operations.GetConnectionsDir2Desc,
                },
        ),
        ID: operations.CreateGetConnectionsIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Name: operations.CreateGetConnectionsNameGetConnectionsName2(
                operations.GetConnectionsName2{},
        ),
        OrderBy: operations.CreateGetConnectionsOrderByArrayOfgetConnectionsOrderBy2(
                []operations.GetConnectionsOrderBy2{
                    operations.GetConnectionsOrderBy2CreatedAt,
                },
        ),
        PausedAt: operations.CreateGetConnectionsPausedAtDateTime(
        types.MustTimeFromString("2023-11-29T02:34:03.781Z"),
        ),
        SourceID: operations.CreateGetConnectionsSourceIDArrayOfstr(
                []string{
                    "string",
                },
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


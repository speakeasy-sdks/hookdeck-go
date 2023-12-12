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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
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
    res, err := s.Connections.Get(ctx, operations.GetConnectionsRequest{
        ArchivedAt: operations.CreateArchivedAtGetConnectionsQueryParam2(
                operations.GetConnectionsQueryParam2{},
        ),
        DestinationID: operations.CreateGetConnectionsQueryParamDestinationIDStr(
        "string",
        ),
        Dir: operations.CreateGetConnectionsQueryParamDirArrayOfgetConnectionsQueryParamConnections2(
                []operations.GetConnectionsQueryParamConnections2{
                    operations.GetConnectionsQueryParamConnections2Desc,
                },
        ),
        ID: operations.CreateGetConnectionsQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Name: operations.CreateQueryParamNameGetConnectionsQueryParamConnectionsName2(
                operations.GetConnectionsQueryParamConnectionsName2{},
        ),
        OrderBy: operations.CreateGetConnectionsQueryParamOrderByArrayOfgetConnectionsQueryParamConnectionsOrderBy2(
                []operations.GetConnectionsQueryParamConnectionsOrderBy2{
                    operations.GetConnectionsQueryParamConnectionsOrderBy2CreatedAt,
                },
        ),
        PausedAt: operations.CreatePausedAtDateTime(
        types.MustTimeFromString("2023-11-29T02:34:03.781Z"),
        ),
        SourceID: operations.CreateGetConnectionsQueryParamSourceIDArrayOfstr(
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

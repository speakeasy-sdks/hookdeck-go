# Destinations
(*Destinations*)

## Overview

A destination is any endpoint to which your webhooks can be routed.

### Available Operations

* [Get](#get) - Get Destinations

## Get

Retrieve a list of endpoints to which your webhooks can be routed.

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
    res, err := s.Destinations.Get(ctx, operations.GetDestinationsRequest{
        Archived: hookdeckgo.Bool(false),
        ArchivedAt: &operations.GetDestinationsArchivedAt{},
        CliPath: &operations.GetDestinationsCliPath{},
        Dir: &operations.GetDestinationsDir{},
        ID: &operations.GetDestinationsID{},
        Limit: hookdeckgo.Int64(700347),
        Name: &operations.GetDestinationsName{},
        Next: hookdeckgo.String("program"),
        OrderBy: &operations.GetDestinationsOrderBy{},
        Prev: hookdeckgo.String("Kia Cambridgeshire"),
        URL: &operations.GetDestinationsURL{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DestinationPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDestinationsRequest](../../models/operations/getdestinationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetDestinationsResponse](../../models/operations/getdestinationsresponse.md), error**


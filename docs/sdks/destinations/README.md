# Destinations
(*.Destinations*)

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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Destinations.Get(ctx, operations.GetDestinationsRequest{
        ArchivedAt: operations.CreateQueryParamArchivedAtGetDestinationsQueryParam2(
                operations.GetDestinationsQueryParam2{},
        ),
        CliPath: operations.CreateCliPathStr(
        "string",
        ),
        Dir: operations.CreateGetDestinationsQueryParamDirArrayOfgetDestinationsQueryParamDestinationsDir2(
                []operations.GetDestinationsQueryParamDestinationsDir2{
                    operations.GetDestinationsQueryParamDestinationsDir2Desc,
                },
        ),
        ID: operations.CreateGetDestinationsQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Name: operations.CreateGetDestinationsQueryParamNameGetDestinationsQueryParamDestinationsName2(
                operations.GetDestinationsQueryParamDestinationsName2{},
        ),
        OrderBy: operations.CreateGetDestinationsQueryParamOrderByArrayOfgetDestinationsQueryParamDestinationsOrderBy2(
                []operations.GetDestinationsQueryParamDestinationsOrderBy2{
                    operations.GetDestinationsQueryParamDestinationsOrderBy2CreatedAt,
                },
        ),
        URL: operations.CreateURLStr(
        "string",
        ),
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


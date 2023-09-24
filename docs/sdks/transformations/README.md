# Transformations

## Overview

A transformation represents JavaScript code that will be executed on a connection's requests. Transformations are applied to connections using Rules.

### Available Operations

* [Get](#get) - Get transformations

## Get

Get transformations

### Example Usage

```go
package main

import(
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.Get(ctx, operations.GetTransformationsRequest{
        Dir: &operations.GetTransformationsDir{},
        ID: &operations.GetTransformationsID{},
        Limit: hookdeckgo.Int64(696344),
        Name: &operations.GetTransformationsName{},
        Next: hookdeckgo.String("voluptatibus"),
        OrderBy: &operations.GetTransformationsOrderBy{},
        Prev: hookdeckgo.String("voluptas"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransformationPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetTransformationsRequest](../../models/operations/gettransformationsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetTransformationsResponse](../../models/operations/gettransformationsresponse.md), error**


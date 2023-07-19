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
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Transformations.Get(ctx, operations.GetTransformationsRequest{
        Dir: &operations.GetTransformationsDir{},
        ID: &operations.GetTransformationsID{},
        Limit: hookdeck.Int64(512393),
        Name: &operations.GetTransformationsName{},
        Next: hookdeck.String("odio"),
        OrderBy: &operations.GetTransformationsOrderBy{},
        Prev: hookdeck.String("occaecati"),
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


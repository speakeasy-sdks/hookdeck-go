# Transformations
(*Transformations*)

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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
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
    res, err := s.Transformations.Get(ctx, operations.GetTransformationsRequest{
        Dir: operations.CreateGetTransformationsQueryParamDirArrayOfgetTransformationsQueryParam2(
                []operations.GetTransformationsQueryParam2{
                    operations.GetTransformationsQueryParam2Asc,
                },
        ),
        ID: operations.CreateGetTransformationsQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Name: operations.CreateGetTransformationsQueryParamNameArrayOfstr(
                []string{
                    "string",
                },
        ),
        OrderBy: operations.CreateGetTransformationsQueryParamOrderByArrayOfgetTransformationsQueryParamTransformations2(
                []operations.GetTransformationsQueryParamTransformations2{
                    operations.GetTransformationsQueryParamTransformations2CreatedAt,
                },
        ),
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

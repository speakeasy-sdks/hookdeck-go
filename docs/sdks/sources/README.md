# Sources
(*Sources*)

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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
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
    res, err := s.Sources.Get(ctx, operations.GetSourcesRequest{
        ArchivedAt: operations.CreateGetSourcesQueryParamArchivedAtGetSourcesQueryParam2(
                operations.GetSourcesQueryParam2{},
        ),
        Dir: operations.CreateGetSourcesQueryParamDirGetSourcesQueryParam1(
        operations.GetSourcesQueryParam1Desc,
        ),
        ID: operations.CreateGetSourcesQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        IntegrationID: operations.CreateIntegrationIDGetSourcesQueryParamSourcesIntegrationID2(
                operations.GetSourcesQueryParamSourcesIntegrationID2{},
        ),
        Name: operations.CreateGetSourcesQueryParamNameGetSourcesQueryParamSourcesName2(
                operations.GetSourcesQueryParamSourcesName2{},
        ),
        OrderBy: operations.CreateGetSourcesQueryParamOrderByArrayOfgetSourcesQueryParamSourcesOrderBy2(
                []operations.GetSourcesQueryParamSourcesOrderBy2{
                    operations.GetSourcesQueryParamSourcesOrderBy2CreatedAt,
                },
        ),
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

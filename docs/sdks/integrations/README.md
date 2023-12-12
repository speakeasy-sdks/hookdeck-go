# Integrations
(*Integrations*)

## Overview

An integration configures platform-specific behaviors, such as signature verification.

### Available Operations

* [Get](#get) - Get integrations

## Get

Get integrations

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
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


    var label *string = "string"

    var provider *components.IntegrationProvider = components.IntegrationProviderOura

    ctx := context.Background()
    res, err := s.Integrations.Get(ctx, label, provider)
    if err != nil {
        log.Fatal(err)
    }

    if res.IntegrationPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `label`                                                                           | **string*                                                                         | :heavy_minus_sign:                                                                | N/A                                                                               |
| `provider`                                                                        | [*components.IntegrationProvider](../../models/components/integrationprovider.md) | :heavy_minus_sign:                                                                | The provider name                                                                 |


### Response

**[*operations.GetIntegrationsResponse](../../models/operations/getintegrationsresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

# Integrations

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
    label := "ad"
    provider := shared.IntegrationProviderSvix

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

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `label`                                                                   | **string*                                                                 | :heavy_minus_sign:                                                        | N/A                                                                       |
| `provider`                                                                | [*shared.IntegrationProvider](../../models/shared/integrationprovider.md) | :heavy_minus_sign:                                                        | The provider name                                                         |


### Response

**[*operations.GetIntegrationsResponse](../../models/operations/getintegrationsresponse.md), error**


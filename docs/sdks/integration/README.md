# Integration

### Available Operations

* [AttachIntegrationToSource](#attachintegrationtosource) - Attach an integration to a source
* [Create](#create) - Create an integration
* [Delete](#delete) - Delete an integration
* [DetachIntegrationToSource](#detachintegrationtosource) - Detach an integration from a source
* [Get](#get) - Get an integration
* [Update](#update) - Update an integration

## AttachIntegrationToSource

Attach an integration to a source

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
    res, err := s.Integration.AttachIntegrationToSource(ctx, "vero", "aspernatur")
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachedIntegrationToSource != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `sourceID`                                            | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.AttachIntegrationToSourceResponse](../../models/operations/attachintegrationtosourceresponse.md), error**


## Create

Create an integration

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
    res, err := s.Integration.Create(ctx, operations.CreateIntegrationRequestBody{
        Configs: &operations.CreateIntegrationRequestBodyConfigs1{},
        Features: []shared.IntegrationFeature{
            shared.IntegrationFeatureVerification,
            shared.IntegrationFeatureHandshake,
        },
        Label: hookdeck.String("ullam"),
        Provider: shared.IntegrationProviderGitlab.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Integration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateIntegrationRequestBody](../../models/operations/createintegrationrequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.CreateIntegrationResponse](../../models/operations/createintegrationresponse.md), error**


## Delete

Delete an integration

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
    res, err := s.Integration.Delete(ctx, "quos")
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletedIntegration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.DeleteIntegrationResponse](../../models/operations/deleteintegrationresponse.md), error**


## DetachIntegrationToSource

Detach an integration from a source

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
    res, err := s.Integration.DetachIntegrationToSource(ctx, "sint", "accusantium")
    if err != nil {
        log.Fatal(err)
    }

    if res.DetachedIntegrationFromSource != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `sourceID`                                            | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.DetachIntegrationToSourceResponse](../../models/operations/detachintegrationtosourceresponse.md), error**


## Get

Get an integration

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
    res, err := s.Integration.Get(ctx, "mollitia")
    if err != nil {
        log.Fatal(err)
    }

    if res.Integration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetIntegrationResponse](../../models/operations/getintegrationresponse.md), error**


## Update

Update an integration

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
    res, err := s.Integration.Update(ctx, operations.UpdateIntegrationRequestBody{
        Configs: &shared.ShopifyIntegrationConfigs{
            APIKey: hookdeck.String("mollitia"),
            APISecret: hookdeck.String("ad"),
            RateLimit: hookdeck.Float32(4314.18),
            RateLimitPeriod: shared.ShopifyIntegrationConfigsRateLimitPeriodLessThanNilGreaterThan.ToPointer(),
            Shop: hookdeck.String("necessitatibus"),
            WebhookSecretKey: "odit",
        },
        Features: []shared.IntegrationFeature{
            shared.IntegrationFeatureVerification,
            shared.IntegrationFeatureHandshake,
        },
        Label: hookdeck.String("doloribus"),
        Provider: shared.IntegrationProviderWorkos.ToPointer(),
    }, "eius")
    if err != nil {
        log.Fatal(err)
    }

    if res.Integration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `requestBody`                                                                                      | [operations.UpdateIntegrationRequestBody](../../models/operations/updateintegrationrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |


### Response

**[*operations.UpdateIntegrationResponse](../../models/operations/updateintegrationresponse.md), error**


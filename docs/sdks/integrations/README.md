# Integrations

## Overview

An integration configures platform-specific behaviors, such as signature verification.

### Available Operations

* [AttachIntegrationToSource](#attachintegrationtosource) - Attach an integration to a source
* [CreateIntegration](#createintegration) - Create an integration
* [DeleteIntegration](#deleteintegration) - Delete an integration
* [DetachIntegrationToSource](#detachintegrationtosource) - Detach an integration from a source
* [GetIntegration](#getintegration) - Get an integration
* [GetIntegrations](#getintegrations) - Get integrations
* [UpdateIntegration](#updateintegration) - Update an integration

## AttachIntegrationToSource

Attach an integration to a source

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Integrations.AttachIntegrationToSource(ctx, operations.AttachIntegrationToSourceRequest{
        ID: "abf603a7-9f9d-4fe0-ab7d-a8a50ce187f8",
        SourceID: "suscipit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachedIntegrationToSource != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AttachIntegrationToSourceRequest](../../models/operations/attachintegrationtosourcerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.AttachIntegrationToSourceResponse](../../models/operations/attachintegrationtosourceresponse.md), error**


## CreateIntegration

Create an integration

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/models/shared"
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
    res, err := s.Integrations.CreateIntegration(ctx, operations.CreateIntegrationRequestBody{
        Configs: &shared.BasicAuthIntegrationConfigs{
            Name: "Henry Koepp",
            Password: "ea",
        },
        Features: []shared.IntegrationFeature{
            shared.IntegrationFeatureHandshake,
            shared.IntegrationFeaturePolling,
            shared.IntegrationFeaturePolling,
        },
        Label: hookdeck.String("accusamus"),
        Provider: shared.IntegrationProviderPropertyFinder.ToPointer(),
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


## DeleteIntegration

Delete an integration

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Integrations.DeleteIntegration(ctx, operations.DeleteIntegrationRequest{
        ID: "526f8d98-6e88-41ea-94f0-e1012563f94e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletedIntegration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteIntegrationRequest](../../models/operations/deleteintegrationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Integrations.DetachIntegrationToSource(ctx, operations.DetachIntegrationToSourceRequest{
        ID: "29e973e9-22a5-47a1-9be3-e060807e2b6e",
        SourceID: "ratione",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DetachedIntegrationFromSource != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DetachIntegrationToSourceRequest](../../models/operations/detachintegrationtosourcerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.DetachIntegrationToSourceResponse](../../models/operations/detachintegrationtosourceresponse.md), error**


## GetIntegration

Get an integration

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Integrations.GetIntegration(ctx, operations.GetIntegrationRequest{
        ID: "ab8845f0-597a-460f-b2a5-4a31e94764a3",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetIntegrationRequest](../../models/operations/getintegrationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetIntegrationResponse](../../models/operations/getintegrationresponse.md), error**


## GetIntegrations

Get integrations

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/models/shared"
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
    res, err := s.Integrations.GetIntegrations(ctx, operations.GetIntegrationsRequest{
        Label: hookdeck.String("debitis"),
        Provider: shared.IntegrationProviderAkeneo.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IntegrationPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetIntegrationsRequest](../../models/operations/getintegrationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetIntegrationsResponse](../../models/operations/getintegrationsresponse.md), error**


## UpdateIntegration

Update an integration

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/models/shared"
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
    res, err := s.Integrations.UpdateIntegration(ctx, operations.UpdateIntegrationRequest{
        RequestBody: operations.UpdateIntegrationRequestBody{
            Configs: &shared.APIKeyIntegrationConfigs{
                APIKey: "nemo",
                HeaderKey: "recusandae",
            },
            Features: []shared.IntegrationFeature{
                shared.IntegrationFeatureHandshake,
                shared.IntegrationFeatureHandshake,
            },
            Label: hookdeck.String("eum"),
            Provider: shared.IntegrationProviderAwsSns.ToPointer(),
        },
        ID: "9251a5a9-da66-40ff-97bf-aad4f9efc1b4",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateIntegrationRequest](../../models/operations/updateintegrationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateIntegrationResponse](../../models/operations/updateintegrationresponse.md), error**


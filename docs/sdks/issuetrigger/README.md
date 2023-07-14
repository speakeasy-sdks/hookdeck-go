# IssueTrigger

### Available Operations

* [Create](#create) - Create an Issue Trigger
* [Delete](#delete) - Delete an Issue Trigger
* [Disable](#disable) - Disable an Issue Trigger
* [Enable](#enable) - Enable an Issue Trigger
* [Get](#get) - Get an Issue Trigger
* [Update](#update) - Update an Issue Trigger
* [Upsert](#upsert) - Create or Update an Issue Trigger

## Create

Create a new issue trigger.

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
    res, err := s.IssueTrigger.Create(ctx, operations.CreateIssueTriggerRequestBody{
        Channels: shared.IssueTriggerChannels{
            Email: &shared.IssueTriggerEmailChannel{},
            Opsgenie: &shared.IssueTriggerIntegrationChannel{},
            Slack: &shared.IssueTriggerSlackChannel{
                ChannelName: "repudiandae",
            },
        },
        Configs: &shared.IssueTriggerTransformationConfigs{
            LogLevel: shared.TransformationExecutionLogLevelError,
            Transformations: shared.IssueTriggerTransformationConfigsTransformations{},
        },
        Name: hookdeck.String("Kristie Spencer"),
        Type: shared.IssueTypeBackpressure,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTrigger != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateIssueTriggerRequestBody](../../models/operations/createissuetriggerrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateIssueTriggerResponse](../../models/operations/createissuetriggerresponse.md), error**


## Delete

Delete an existing issue trigger.

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
    res, err := s.IssueTrigger.Delete(ctx, "accusantium")
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletedIssueTriggerResponse != nil {
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

**[*operations.DeleteIssueTriggerResponse](../../models/operations/deleteissuetriggerresponse.md), error**


## Disable

Disable an existing issue trigger.

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
    res, err := s.IssueTrigger.Disable(ctx, "consequuntur")
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTrigger != nil {
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

**[*operations.DisableIssueTriggerResponse](../../models/operations/disableissuetriggerresponse.md), error**


## Enable

Enable an existing issue trigger.

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
    res, err := s.IssueTrigger.Enable(ctx, "praesentium")
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTrigger != nil {
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

**[*operations.EnableIssueTriggerResponse](../../models/operations/enableissuetriggerresponse.md), error**


## Get

Get a single issue trigger details.

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
    res, err := s.IssueTrigger.Get(ctx, "natus")
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTrigger != nil {
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

**[*operations.GetIssueTriggerResponse](../../models/operations/getissuetriggerresponse.md), error**


## Update

Update the details of an issue trigger.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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
    res, err := s.IssueTrigger.Update(ctx, operations.UpdateIssueTriggerRequestBody{
        Channels: &shared.IssueTriggerChannels{
            Email: &shared.IssueTriggerEmailChannel{},
            Opsgenie: &shared.IssueTriggerIntegrationChannel{},
            Slack: &shared.IssueTriggerSlackChannel{
                ChannelName: "magni",
            },
        },
        Configs: &shared.IssueTriggerDeliveryConfigs{
            Connections: shared.IssueTriggerDeliveryConfigsConnections{},
            Strategy: shared.IssueTriggerStrategyFinalAttempt,
        },
        DisabledAt: types.MustTimeFromString("2020-05-28T21:33:10.895Z"),
        Name: hookdeck.String("Nathaniel Marks"),
    }, "accusantium")
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTrigger != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `requestBody`                                                                                        | [operations.UpdateIssueTriggerRequestBody](../../models/operations/updateissuetriggerrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `id`                                                                                                 | *string*                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |


### Response

**[*operations.UpdateIssueTriggerResponse](../../models/operations/updateissuetriggerresponse.md), error**


## Upsert

Create or update an existing issue trigger.

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
    res, err := s.IssueTrigger.Upsert(ctx, operations.UpsertIssueTriggerRequestBody{
        Channels: shared.IssueTriggerChannels{
            Email: &shared.IssueTriggerEmailChannel{},
            Opsgenie: &shared.IssueTriggerIntegrationChannel{},
            Slack: &shared.IssueTriggerSlackChannel{
                ChannelName: "ab",
            },
        },
        Configs: &shared.IssueTriggerBackpressureConfigs{
            Delay: 697429,
            Destinations: shared.IssueTriggerBackpressureConfigsDestinations{},
        },
        Name: "Colleen Johnston PhD",
        Type: shared.IssueTypeTransformation,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTrigger != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpsertIssueTriggerRequestBody](../../models/operations/upsertissuetriggerrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.UpsertIssueTriggerResponse](../../models/operations/upsertissuetriggerresponse.md), error**


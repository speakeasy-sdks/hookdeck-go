# IssueTriggers

### Available Operations

* [CreateIssueTrigger](#createissuetrigger) - Create an Issue Trigger
* [DeleteIssueTrigger](#deleteissuetrigger) - Delete an Issue Trigger
* [DisableIssueTrigger](#disableissuetrigger) - Disable an Issue Trigger
* [EnableIssueTrigger](#enableissuetrigger) - Enable an Issue Trigger
* [GetIssueTrigger](#getissuetrigger) - Get an Issue Trigger
* [GetIssueTriggers](#getissuetriggers) - Get Issue Triggers
* [UpdateIssueTrigger](#updateissuetrigger) - Update an Issue Trigger
* [UpsertIssueTrigger](#upsertissuetrigger) - Create or Update an Issue Trigger

## CreateIssueTrigger

Create a new issue trigger.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.IssueTriggers.CreateIssueTrigger(ctx, operations.CreateIssueTriggerRequestBody{
        Channels: shared.IssueTriggerChannels{
            Email: &shared.IssueTriggerEmailChannel{},
            Opsgenie: &shared.IssueTriggerIntegrationChannel{},
            Slack: &shared.IssueTriggerSlackChannel{
                ChannelName: "quis",
            },
        },
        Configs: &shared.IssueTriggerDeliveryConfigs{
            Connections: shared.IssueTriggerDeliveryConfigsConnections{},
            Strategy: shared.IssueTriggerStrategyFirstAttempt,
        },
        Name: hookdeck.String("Scott Bahringer"),
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
| `request`                                                                                            | [operations.CreateIssueTriggerRequestBody](../../models/operations/createissuetriggerrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateIssueTriggerResponse](../../models/operations/createissuetriggerresponse.md), error**


## DeleteIssueTrigger

Delete an existing issue trigger.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.IssueTriggers.DeleteIssueTrigger(ctx, operations.DeleteIssueTriggerRequest{
        ID: "48dc2f61-5199-4ebf-90e9-fe6c632ca3ae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletedIssueTriggerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteIssueTriggerRequest](../../models/operations/deleteissuetriggerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteIssueTriggerResponse](../../models/operations/deleteissuetriggerresponse.md), error**


## DisableIssueTrigger

Disable an existing issue trigger.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.IssueTriggers.DisableIssueTrigger(ctx, operations.DisableIssueTriggerRequest{
        ID: "d0117996-312f-4de0-8771-778ff61d0174",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DisableIssueTriggerRequest](../../models/operations/disableissuetriggerrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DisableIssueTriggerResponse](../../models/operations/disableissuetriggerresponse.md), error**


## EnableIssueTrigger

Enable an existing issue trigger.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.IssueTriggers.EnableIssueTrigger(ctx, operations.EnableIssueTriggerRequest{
        ID: "76360a15-db6a-4660-a59a-1adeaab5851d",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.EnableIssueTriggerRequest](../../models/operations/enableissuetriggerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.EnableIssueTriggerResponse](../../models/operations/enableissuetriggerresponse.md), error**


## GetIssueTrigger

Get a single issue trigger details.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.IssueTriggers.GetIssueTrigger(ctx, operations.GetIssueTriggerRequest{
        ID: "6c645b08-b618-491b-aa0f-e1ade008e6f8",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetIssueTriggerRequest](../../models/operations/getissuetriggerrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetIssueTriggerResponse](../../models/operations/getissuetriggerresponse.md), error**


## GetIssueTriggers

Retrieve a list of issue triggers.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.IssueTriggers.GetIssueTriggers(ctx, operations.GetIssueTriggersRequest{
        Dir: &operations.GetIssueTriggersDir{},
        DisabledAt: &operations.GetIssueTriggersDisabledAt{},
        Limit: hookdeck.Int64(796320),
        Name: hookdeck.String("Ollie Dicki PhD"),
        Next: hookdeck.String("totam"),
        OrderBy: &operations.GetIssueTriggersOrderBy{},
        Prev: hookdeck.String("impedit"),
        Type: shared.IssueTypeBackpressure.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTriggerPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetIssueTriggersRequest](../../models/operations/getissuetriggersrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetIssueTriggersResponse](../../models/operations/getissuetriggersresponse.md), error**


## UpdateIssueTrigger

Update the details of an issue trigger.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.IssueTriggers.UpdateIssueTrigger(ctx, operations.UpdateIssueTriggerRequest{
        RequestBody: operations.UpdateIssueTriggerRequestBody{
            Channels: &shared.IssueTriggerChannels{
                Email: &shared.IssueTriggerEmailChannel{},
                Opsgenie: &shared.IssueTriggerIntegrationChannel{},
                Slack: &shared.IssueTriggerSlackChannel{
                    ChannelName: "nam",
                },
            },
            Configs: &shared.IssueTriggerTransformationConfigs{
                LogLevel: shared.TransformationExecutionLogLevelError,
                Transformations: shared.IssueTriggerTransformationConfigsTransformations{},
            },
            DisabledAt: types.MustTimeFromString("2022-09-10T19:23:06.016Z"),
            Name: hookdeck.String("Cassandra Bogan"),
        },
        ID: "01042181-3d52-408e-8e7e-253b668451c6",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateIssueTriggerRequest](../../models/operations/updateissuetriggerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateIssueTriggerResponse](../../models/operations/updateissuetriggerresponse.md), error**


## UpsertIssueTrigger

Create or update an existing issue trigger.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.IssueTriggers.UpsertIssueTrigger(ctx, operations.UpsertIssueTriggerRequestBody{
        Channels: shared.IssueTriggerChannels{
            Email: &shared.IssueTriggerEmailChannel{},
            Opsgenie: &shared.IssueTriggerIntegrationChannel{},
            Slack: &shared.IssueTriggerSlackChannel{
                ChannelName: "nobis",
            },
        },
        Configs: &shared.IssueTriggerTransformationConfigs{
            LogLevel: shared.TransformationExecutionLogLevelFatal,
            Transformations: shared.IssueTriggerTransformationConfigsTransformations{},
        },
        Name: "Helen Heller III",
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
| `request`                                                                                            | [operations.UpsertIssueTriggerRequestBody](../../models/operations/upsertissuetriggerrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.UpsertIssueTriggerResponse](../../models/operations/upsertissuetriggerresponse.md), error**


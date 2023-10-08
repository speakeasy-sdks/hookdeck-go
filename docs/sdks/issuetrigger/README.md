# IssueTrigger
(*IssueTrigger*)

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
    res, err := s.IssueTrigger.Create(ctx, operations.CreateIssueTriggerRequestBody{
        AdditionalProperties: map[string]interface{}{
            "online": "Configuration",
        },
        Channels: &shared.IssueTriggerChannels{
            AdditionalProperties: map[string]interface{}{
                "Money": "blue",
            },
            Email: map[string]interface{}{
                "shred": "abnormally",
            },
            Opsgenie: map[string]interface{}{
                "deposit": "evolve",
            },
            Slack: &shared.IssueTriggerSlackChannel{
                AdditionalProperties: map[string]interface{}{
                    "male": "SUV",
                },
                ChannelName: "Screen mobile",
            },
        },
        Configs: operations.CreateCreateIssueTriggerRequestBodyConfigsIssueTriggerTransformationConfigs(
                shared.IssueTriggerTransformationConfigs{
                    AdditionalProperties: map[string]interface{}{
                        "Ameliorated": "Fresh",
                    },
                    LogLevel: shared.TransformationExecutionLogLevelDebug,
                    Transformations: shared.CreateIssueTriggerTransformationConfigsTransformationsArrayOfstr(
                            []string{
                                "Intelligent",
                            },
                    ),
                },
        ),
        Type: shared.IssueTypeDelivery,
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    var id string = "program"

    ctx := context.Background()
    res, err := s.IssueTrigger.Delete(ctx, id)
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    var id string = "Bicycle"

    ctx := context.Background()
    res, err := s.IssueTrigger.Disable(ctx, id)
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    var id string = "male"

    ctx := context.Background()
    res, err := s.IssueTrigger.Enable(ctx, id)
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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    var id string = "female"

    ctx := context.Background()
    res, err := s.IssueTrigger.Get(ctx, id)
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
    requestBody := operations.UpdateIssueTriggerRequestBody{
        AdditionalProperties: map[string]interface{}{
            "Van": "East",
        },
        Channels: &shared.IssueTriggerChannels{
            AdditionalProperties: map[string]interface{}{
                "male": "Metal",
            },
            Email: map[string]interface{}{
                "cheater": "Islands",
            },
            Opsgenie: map[string]interface{}{
                "online": "dynamic",
            },
            Slack: &shared.IssueTriggerSlackChannel{
                AdditionalProperties: map[string]interface{}{
                    "white": "bifurcated",
                },
                ChannelName: "silver immediately",
            },
        },
        Configs: operations.CreateUpdateIssueTriggerRequestBodyConfigsIssueTriggerDeliveryConfigs(
                shared.IssueTriggerDeliveryConfigs{
                    AdditionalProperties: map[string]interface{}{
                        "East": "Baht",
                    },
                    Connections: shared.CreateIssueTriggerDeliveryConfigsConnectionsArrayOfstr(
                            []string{
                                "Representative",
                            },
                    ),
                    Strategy: shared.IssueTriggerStrategyFirstAttempt,
                },
        ),
    }
    var id string = "driver"

    ctx := context.Background()
    res, err := s.IssueTrigger.Update(ctx, requestBody, id)
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
    res, err := s.IssueTrigger.Upsert(ctx, operations.UpsertIssueTriggerRequestBody{
        AdditionalProperties: map[string]interface{}{
            "index": "PNG",
        },
        Channels: &shared.IssueTriggerChannels{
            AdditionalProperties: map[string]interface{}{
                "extranet": "SSD",
            },
            Email: map[string]interface{}{
                "Small": "Triathlon/Time",
            },
            Opsgenie: map[string]interface{}{
                "Cisgender": "Northwest",
            },
            Slack: &shared.IssueTriggerSlackChannel{
                AdditionalProperties: map[string]interface{}{
                    "Facilitator": "transmitting",
                },
                ChannelName: "calculating",
            },
        },
        Configs: operations.CreateUpsertIssueTriggerRequestBodyConfigsIssueTriggerDeliveryConfigs(
                shared.IssueTriggerDeliveryConfigs{
                    AdditionalProperties: map[string]interface{}{
                        "Market": "South",
                    },
                    Connections: shared.CreateIssueTriggerDeliveryConfigsConnectionsArrayOfstr(
                            []string{
                                "female",
                            },
                    ),
                    Strategy: shared.IssueTriggerStrategyFinalAttempt,
                },
        ),
        Name: "altruistic",
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


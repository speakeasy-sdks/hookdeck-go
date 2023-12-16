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
    res, err := s.IssueTrigger.Create(ctx, operations.CreateIssueTriggerRequestBody{
        Channels: &components.IssueTriggerChannels{
            Email: &components.IssueTriggerEmailChannel{},
            Opsgenie: &components.IssueTriggerIntegrationChannel{},
            Slack: &components.IssueTriggerSlackChannel{
                ChannelName: "string",
            },
        },
        Configs: operations.CreateCreateIssueTriggerConfigsIssueTriggerTransformationConfigs(
                components.IssueTriggerTransformationConfigs{
                    LogLevel: components.TransformationExecutionLogLevelWarn,
                    Transformations: components.CreateTransformationsArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                },
        ),
        Type: components.IssueTypeBackpressure,
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Delete

Delete an existing issue trigger.

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


    var id string = "string"

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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Disable

Disable an existing issue trigger.

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


    var id string = "string"

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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Enable

Enable an existing issue trigger.

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


    var id string = "string"

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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Get

Get a single issue trigger details.

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


    var id string = "string"

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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Update

Update the details of an issue trigger.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
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


    requestBody := operations.UpdateIssueTriggerRequestBody{
        Channels: &components.IssueTriggerChannels{
            Email: &components.IssueTriggerEmailChannel{},
            Opsgenie: &components.IssueTriggerIntegrationChannel{},
            Slack: &components.IssueTriggerSlackChannel{
                ChannelName: "string",
            },
        },
        Configs: operations.CreateUpdateIssueTriggerConfigsIssueTriggerBackpressureConfigs(
                components.IssueTriggerBackpressureConfigs{
                    Delay: 24555,
                    Destinations: components.CreateDestinationsArrayOfstr(
                            []string{
                                "string",
                            },
                    ),
                },
        ),
    }

    var id string = "string"

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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Upsert

Create or update an existing issue trigger.

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
    res, err := s.IssueTrigger.Upsert(ctx, operations.UpsertIssueTriggerRequestBody{
        Channels: &components.IssueTriggerChannels{
            Email: &components.IssueTriggerEmailChannel{},
            Opsgenie: &components.IssueTriggerIntegrationChannel{},
            Slack: &components.IssueTriggerSlackChannel{
                ChannelName: "string",
            },
        },
        Configs: operations.CreateUpsertIssueTriggerConfigsIssueTriggerTransformationConfigs(
                components.IssueTriggerTransformationConfigs{
                    LogLevel: components.TransformationExecutionLogLevelWarn,
                    Transformations: components.CreateTransformationsStr(
                    "string",
                    ),
                },
        ),
        Name: "string",
        Type: components.IssueTypeTransformation,
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

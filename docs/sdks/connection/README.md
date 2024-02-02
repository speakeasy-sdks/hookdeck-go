# Connection
(*Connection*)

### Available Operations

* [Archive](#archive) - Archive a connection
* [Create](#create) - Create a connection
* [Delete](#delete) - Delete a connection
* [Get](#get) - Get a single connection
* [Pause](#pause) - Pause a connection
* [Unarchive](#unarchive) - Unarchive a connection
* [Unpause](#unpause) - Unpause a connection
* [Update](#update) - Update a connection

## Archive

Archive a connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
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
    res, err := s.Connection.Archive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.ArchiveConnectionResponse](../../models/operations/archiveconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Create

Create a connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
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
    res, err := s.Connection.Create(ctx, operations.CreateConnectionRequestBody{
        Destination: &operations.CreateConnectionDestination{
            AuthMethod: components.CreateDestinationAuthMethodConfigAPIKey(
                    components.APIKey{
                        Config: &components.DestinationAuthMethodAPIKeyConfig{
                            APIKey: "string",
                            Key: "<key>",
                        },
                        Type: components.TypeAPIKey,
                    },
            ),
            Name: "string",
        },
        Name: "string",
        Rules: []components.Rule{
            components.CreateRuleFilterRule(
                components.FilterRule{
                    Body: components.CreateConnectionFilterPropertyBoolean(
                    false,
                    ),
                    Headers: components.CreateConnectionFilterPropertyFour(
                            components.Four{},
                    ),
                    Path: components.CreateConnectionFilterPropertyFloat32(
                    2884.08,
                    ),
                    Query: components.CreateConnectionFilterPropertyStr(
                    "string",
                    ),
                    Type: components.FilterRuleTypeFilter,
                },
            ),
        },
        Ruleset: &operations.CreateConnectionRuleset{
            Name: "string",
            Rules: []components.Rule{
                components.CreateRuleTransformRule(
                    components.CreateTransformRuleTransformReference(
                            components.TransformReference{
                                TransformationID: "string",
                                Type: components.TransformReferenceTypeTransform,
                            },
                    ),
                ),
            },
        },
        Source: &operations.CreateConnectionSource{
            AllowedHTTPMethods: []components.SourceAllowedHTTPMethod{
                components.SourceAllowedHTTPMethodDelete,
            },
            CustomResponse: &components.SourceCustomResponse{
                Body: "string",
                ContentType: components.SourceCustomResponseContentTypeXML,
            },
            Name: "string",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateConnectionRequestBody](../../models/operations/createconnectionrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CreateConnectionResponse](../../models/operations/createconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Delete

Delete a connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
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
    res, err := s.Connection.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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

**[*operations.DeleteConnectionResponse](../../models/operations/deleteconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Get

Get a single connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
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
    res, err := s.Connection.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.GetConnectionResponse](../../models/operations/getconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404,410                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Pause

Pause a connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
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
    res, err := s.Connection.Pause(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.PauseConnectionResponse](../../models/operations/pauseconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Unarchive

Unarchive a connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
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
    res, err := s.Connection.Unarchive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.UnarchiveConnectionResponse](../../models/operations/unarchiveconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Unpause

Unpause a connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
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
    res, err := s.Connection.Unpause(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
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

**[*operations.UnpauseConnectionResponse](../../models/operations/unpauseconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Update

Update a connection

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
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


    requestBody := operations.UpdateConnectionRequestBody{
        Rules: []components.Rule{
            components.CreateRuleDelayRule(
                components.DelayRule{
                    Delay: 24555,
                    Type: components.DelayRuleTypeDelay,
                },
            ),
        },
        Ruleset: &operations.UpdateConnectionRuleset{
            Name: "string",
            Rules: []components.Rule{
                components.CreateRuleFilterRule(
                    components.FilterRule{
                        Body: components.CreateConnectionFilterPropertyStr(
                        "string",
                        ),
                        Headers: components.CreateConnectionFilterPropertyFloat32(
                        7084.55,
                        ),
                        Path: components.CreateConnectionFilterPropertyFour(
                                components.Four{},
                        ),
                        Query: components.CreateConnectionFilterPropertyFloat32(
                        6276.9,
                        ),
                        Type: components.FilterRuleTypeFilter,
                    },
                ),
            },
        },
    }

    var id string = "string"

    ctx := context.Background()
    res, err := s.Connection.Update(ctx, requestBody, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Connection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `requestBody`                                                                                    | [operations.UpdateConnectionRequestBody](../../models/operations/updateconnectionrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |


### Response

**[*operations.UpdateConnectionResponse](../../models/operations/updateconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,404,422                | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"log"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Create

Create a connection

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
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Delete

Delete a connection

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
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Get

Get a single connection

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
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Pause

Pause a connection

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
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Unarchive

Unarchive a connection

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
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Unpause

Unpause a connection

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
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

## Update

Update a connection

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
                Password: "",
                Username: "",
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
| sdkerrors.SDKError         | 400-600                    | */*                        |

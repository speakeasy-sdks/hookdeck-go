# Ruleset
(*Ruleset*)

### Available Operations

* [Archive](#archive) - Archive a ruleset
* [Create](#create) - Create a ruleset
* [Get](#get) - Get a ruleset
* [Unarchive](#unarchive) - Unarchive a ruleset
* [Update](#update) - Update a ruleset
* [Upsert](#upsert) - Update or create a ruleset

## Archive

Archive a ruleset

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


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Ruleset.Archive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Ruleset != nil {
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

**[*operations.ArchiveRulesetResponse](../../models/operations/archiverulesetresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Create

Create a ruleset

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
    res, err := s.Ruleset.Create(ctx, operations.CreateRulesetRequestBody{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Ruleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateRulesetRequestBody](../../models/operations/createrulesetrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateRulesetResponse](../../models/operations/createrulesetresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Get

Get a ruleset

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


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Ruleset.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Ruleset != nil {
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

**[*operations.GetRulesetResponse](../../models/operations/getrulesetresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Unarchive

Unarchive a ruleset

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


    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Ruleset.Unarchive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Ruleset != nil {
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

**[*operations.UnarchiveRulesetResponse](../../models/operations/unarchiverulesetresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Update

Update a ruleset

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


    requestBody := operations.UpdateRulesetRequestBody{}

    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Ruleset.Update(ctx, requestBody, id)
    if err != nil {
        log.Fatal(err)
    }
    if res.Ruleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `requestBody`                                                                              | [operations.UpdateRulesetRequestBody](../../models/operations/updaterulesetrequestbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |


### Response

**[*operations.UpdateRulesetResponse](../../models/operations/updaterulesetresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,404,422                | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Upsert

Update or create a ruleset

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
    res, err := s.Ruleset.Upsert(ctx, operations.UpsertRulesetRequestBody{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Ruleset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpsertRulesetRequestBody](../../models/operations/upsertrulesetrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpsertRulesetResponse](../../models/operations/upsertrulesetresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

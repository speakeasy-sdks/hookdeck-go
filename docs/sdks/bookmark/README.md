# Bookmark
(*Bookmark*)

### Available Operations

* [Create](#create) - Create a Bookmark
* [Delete](#delete) - Delete a Bookmark
* [Get](#get) - Get a Single Bookmark
* [Trigger](#trigger) - Trigger a Bookmark
* [Update](#update) - Update a Bookmark

## Create

Create a new bookmark.

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
    res, err := s.Bookmark.Create(ctx, operations.CreateBookmarkRequestBody{
        EventDataID: "corrupti",
        Label: "illum",
        Name: hookdeckgo.String("Sabrina Oberbrunner"),
        WebhookID: "magnam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Bookmark != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateBookmarkRequestBody](../../models/operations/createbookmarkrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateBookmarkResponse](../../models/operations/createbookmarkresponse.md), error**


## Delete

Delete an existing bookmark

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
    id := "debitis"

    ctx := context.Background()
    res, err := s.Bookmark.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletedBookmarkResponse != nil {
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

**[*operations.DeleteBookmarkResponse](../../models/operations/deletebookmarkresponse.md), error**


## Get

Retrieve an existing bookmark details.

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
    id := "ipsa"

    ctx := context.Background()
    res, err := s.Bookmark.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Bookmark != nil {
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

**[*operations.GetBookmarkResponse](../../models/operations/getbookmarkresponse.md), error**


## Trigger

Trigger a bookmark operation to store and replay a specific request.

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
    requestBody := operations.TriggerBookmarkRequestBody{
        Target: operations.TriggerBookmarkRequestBodyTargetCli.ToPointer(),
    }
    id := "tempora"

    ctx := context.Background()
    res, err := s.Bookmark.Trigger(ctx, requestBody, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventArray != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `requestBody`                                                                                  | [operations.TriggerBookmarkRequestBody](../../models/operations/triggerbookmarkrequestbody.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |


### Response

**[*operations.TriggerBookmarkResponse](../../models/operations/triggerbookmarkresponse.md), error**


## Update

Update an existing bookmark information.

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
    requestBody := operations.UpdateBookmarkRequestBody{
        EventDataID: hookdeckgo.String("suscipit"),
        Label: hookdeckgo.String("molestiae"),
        Name: hookdeckgo.String("Irving Lehner"),
        WebhookID: hookdeckgo.String("nisi"),
    }
    id := "recusandae"

    ctx := context.Background()
    res, err := s.Bookmark.Update(ctx, requestBody, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Bookmark != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `requestBody`                                                                                | [operations.UpdateBookmarkRequestBody](../../models/operations/updatebookmarkrequestbody.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |


### Response

**[*operations.UpdateBookmarkResponse](../../models/operations/updatebookmarkresponse.md), error**


# Source
(*Source*)

### Available Operations

* [Archive](#archive) - Archive a source
* [Create](#create) - Create a source
* [Delete](#delete) - Delete a source
* [Get](#get) - Get a source
* [Unarchive](#unarchive) - Unarchive a source
* [Update](#update) - Update a source
* [Upsert](#upsert) - Update or create a source

## Archive

Archive a source

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
    var id string = "meek"

    ctx := context.Background()
    res, err := s.Source.Archive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
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

**[*operations.ArchiveSourceResponse](../../models/operations/archivesourceresponse.md), error**


## Create

Create a source

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
    res, err := s.Source.Create(ctx, operations.CreateSourceRequestBody{
        AdditionalProperties: map[string]interface{}{
            "online": "Configuration",
        },
        AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
            shared.SourceAllowedHTTPMethodPut,
        },
        CustomResponse: &shared.SourceCustomResponse{
            AdditionalProperties: map[string]interface{}{
                "innovative": "blue",
            },
            Body: "grey technology East",
            ContentType: shared.SourceCustomResponseContentTypeJSON,
        },
        Name: "Northwest",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateSourceRequestBody](../../models/operations/createsourcerequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateSourceResponse](../../models/operations/createsourceresponse.md), error**


## Delete

Delete a source

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
    res, err := s.Source.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteSource200ApplicationJSONObject != nil {
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

**[*operations.DeleteSourceResponse](../../models/operations/deletesourceresponse.md), error**


## Get

Get a source

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
    res, err := s.Source.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
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

**[*operations.GetSourceResponse](../../models/operations/getsourceresponse.md), error**


## Unarchive

Unarchive a source

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
    var id string = "Bedfordshire"

    ctx := context.Background()
    res, err := s.Source.Unarchive(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
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

**[*operations.UnarchiveSourceResponse](../../models/operations/unarchivesourceresponse.md), error**


## Update

Update a source

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
    requestBody := operations.UpdateSourceRequestBody{
        AdditionalProperties: map[string]interface{}{
            "Van": "East",
        },
        AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
            shared.SourceAllowedHTTPMethodPatch,
        },
        CustomResponse: &shared.SourceCustomResponse{
            AdditionalProperties: map[string]interface{}{
                "dock": "Quality",
            },
            Body: "invoice Arizona",
            ContentType: shared.SourceCustomResponseContentTypeJSON,
        },
    }
    var id string = "mostly"

    ctx := context.Background()
    res, err := s.Source.Update(ctx, requestBody, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `requestBody`                                                                            | [operations.UpdateSourceRequestBody](../../models/operations/updatesourcerequestbody.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |


### Response

**[*operations.UpdateSourceResponse](../../models/operations/updatesourceresponse.md), error**


## Upsert

Update or create a source

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
    res, err := s.Source.Upsert(ctx, operations.UpsertSourceRequestBody{
        AdditionalProperties: map[string]interface{}{
            "index": "PNG",
        },
        AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
            shared.SourceAllowedHTTPMethodPost,
        },
        CustomResponse: &shared.SourceCustomResponse{
            AdditionalProperties: map[string]interface{}{
                "grow": "kilogram",
            },
            Body: "Triathlon/Time",
            ContentType: shared.SourceCustomResponseContentTypeText,
        },
        Name: "North",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpsertSourceRequestBody](../../models/operations/upsertsourcerequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpsertSourceResponse](../../models/operations/upsertsourceresponse.md), error**


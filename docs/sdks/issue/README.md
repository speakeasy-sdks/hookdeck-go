# Issue
(*Issue*)

### Available Operations

* [Dismiss](#dismiss) - Dismiss an issue
* [Get](#get) - Get a single issue
* [Update](#update) - Update issue

## Dismiss

Dismiss an issue

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
    res, err := s.Issue.Dismiss(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Issue != nil {
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

**[*operations.DismissIssueResponse](../../models/operations/dismississueresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Get

Get a single issue

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
    res, err := s.Issue.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueWithData != nil {
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

**[*operations.GetIssueResponse](../../models/operations/getissueresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Update

Update issue

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


    requestBody := operations.UpdateIssueRequestBody{
        Status: operations.UpdateIssueStatusResolved,
    }

    var id string = "<value>"

    ctx := context.Background()
    res, err := s.Issue.Update(ctx, requestBody, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Issue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `requestBody`                                                                          | [operations.UpdateIssueRequestBody](../../models/operations/updateissuerequestbody.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |


### Response

**[*operations.UpdateIssueResponse](../../models/operations/updateissueresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

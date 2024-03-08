# CustomDomains
(*CustomDomains*)

### Available Operations

* [List](#list) - List all custom domains and their verification statuses for the workspace

## List

List all custom domains and their verification statuses for the workspace

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


    var teamID string = "<value>"

    ctx := context.Background()
    res, err := s.CustomDomains.List(ctx, teamID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCustomDomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `teamID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.ListCustomDomainsResponse](../../models/operations/listcustomdomainsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

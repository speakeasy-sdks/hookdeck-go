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
    teamID := "temporibus"

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


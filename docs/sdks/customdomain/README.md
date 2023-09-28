# CustomDomain
(*CustomDomain*)

### Available Operations

* [Add](#add) - Add a custom domain to the workspace
* [Delete](#delete) - Removes a custom domain from the workspace

## Add

Add a custom domain to the workspace

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
    addCustomHostname := shared.AddCustomHostname{
        Hostname: "happy-pink.net",
    }
    teamID := "aut"

    ctx := context.Background()
    res, err := s.CustomDomain.Add(ctx, addCustomHostname, teamID)
    if err != nil {
        log.Fatal(err)
    }

    if res.AddCustomHostname != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `addCustomHostname`                                                  | [shared.AddCustomHostname](../../models/shared/addcustomhostname.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `teamID`                                                             | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |


### Response

**[*operations.AddCustomDomainResponse](../../models/operations/addcustomdomainresponse.md), error**


## Delete

Removes a custom domain from the workspace

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
    domainID := "quasi"
    teamID := "error"

    ctx := context.Background()
    res, err := s.CustomDomain.Delete(ctx, domainID, teamID)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCustomDomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `domainID`                                            | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `teamID`                                              | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.DeleteCustomDomainResponse](../../models/operations/deletecustomdomainresponse.md), error**


# CustomDomain
(*.CustomDomain*)

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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
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


    addCustomHostname := components.AddCustomHostname{
        Hostname: "murky-sole.name",
    }

    var teamID string = "string"

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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `addCustomHostname`                                                      | [components.AddCustomHostname](../../models/shared/addcustomhostname.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `teamID`                                                                 | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |


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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
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


    var domainID string = "string"

    var teamID string = "string"

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


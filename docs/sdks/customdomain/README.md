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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `addCustomHostname`                                                          | [components.AddCustomHostname](../../models/components/addcustomhostname.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `teamID`                                                                     | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |


### Response

**[*operations.AddCustomDomainResponse](../../models/operations/addcustomdomainresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

Removes a custom domain from the workspace

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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

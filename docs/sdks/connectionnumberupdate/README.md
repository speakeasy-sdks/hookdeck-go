# ConnectionNumberUpdate
(*ConnectionNumberUpdate*)

### Available Operations

* [Upsert](#upsert) - Update or create a connection

## Upsert

Update or create a connection

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
    res, err := s.ConnectionNumberUpdate.Upsert(ctx, operations.UpsertConnectionRequestBody{
        Destination: &operations.Destination{
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
                    Body: components.CreateConnectionFilterPropertyFloat32(
                    5344.19,
                    ),
                    Headers: components.CreateConnectionFilterPropertyFloat32(
                    3135.5,
                    ),
                    Path: components.CreateConnectionFilterPropertyFloat32(
                    7400.64,
                    ),
                    Query: components.CreateConnectionFilterPropertyStr(
                    "string",
                    ),
                    Type: components.FilterRuleTypeFilter,
                },
            ),
        },
        Ruleset: &operations.Ruleset{
            Name: "string",
            Rules: []components.Rule{
                components.CreateRuleRetryRule(
                    components.RetryRule{
                        Strategy: components.RetryStrategyExponential,
                        Type: components.RetryRuleTypeRetry,
                    },
                ),
            },
        },
        Source: &operations.Source{
            AllowedHTTPMethods: []components.SourceAllowedHTTPMethod{
                components.SourceAllowedHTTPMethodDelete,
            },
            CustomResponse: &components.SourceCustomResponse{
                Body: "string",
                ContentType: components.SourceCustomResponseContentTypeJSON,
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
| `request`                                                                                        | [operations.UpsertConnectionRequestBody](../../models/operations/upsertconnectionrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpsertConnectionResponse](../../models/operations/upsertconnectionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

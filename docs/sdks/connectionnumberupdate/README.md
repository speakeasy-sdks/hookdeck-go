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
    res, err := s.ConnectionNumberUpdate.Upsert(ctx, operations.UpsertConnectionRequestBody{
        Destination: &operations.UpsertConnectionRequestBodyDestination{
            AuthMethod: shared.CreateDestinationAuthMethodConfigAPIKey(
                    shared.APIKey{
                        Config: &shared.DestinationAuthMethodAPIKeyConfig{
                            APIKey: "string",
                            Key: "<key>",
                        },
                        Type: shared.APIKeyTypeAPIKey,
                    },
            ),
            Name: "string",
        },
        Name: "string",
        Rules: []shared.Rule{
            shared.CreateRuleFilterRule(
                shared.FilterRule{
                    Body: shared.CreateConnectionFilterPropertyFloat32(
                    5344.19,
                    ),
                    Headers: shared.CreateConnectionFilterPropertyFloat32(
                    3135.5,
                    ),
                    Path: shared.CreateConnectionFilterPropertyFloat32(
                    7400.64,
                    ),
                    Query: shared.CreateConnectionFilterPropertyStr(
                    "string",
                    ),
                    Type: shared.FilterRuleTypeFilter,
                },
            ),
        },
        Ruleset: &operations.UpsertConnectionRequestBodyRuleset{
            Name: "string",
            Rules: []shared.Rule{
                shared.CreateRuleRetryRule(
                    shared.RetryRule{
                        Strategy: shared.RetryStrategyExponential,
                        Type: shared.RetryRuleTypeRetry,
                    },
                ),
            },
        },
        Source: &operations.UpsertConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodDelete,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "string",
                ContentType: shared.SourceCustomResponseContentTypeJSON,
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


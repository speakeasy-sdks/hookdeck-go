# ConnectionNumberUpdate

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
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
)

func main() {
    s := hookdeck.New(
        hookdeck.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.ConnectionNumberUpdate.Upsert(ctx, operations.UpsertConnectionRequestBody{
        Destination: &operations.UpsertConnectionRequestBodyDestination{
            AuthMethod: &shared.HookdeckSignature{
                Config: &shared.DestinationAuthMethodSignatureConfig{},
                Type: shared.HookdeckSignatureTypeHookdeckSignature,
            },
            CliPath: hookdeck.String("reiciendis"),
            HTTPMethod: shared.DestinationHTTPMethodDelete.ToPointer(),
            Name: "Jessie Langosh V",
            PathForwardingDisabled: hookdeck.Bool(false),
            RateLimit: hookdeck.Int64(739264),
            RateLimitPeriod: operations.UpsertConnectionRequestBodyDestinationRateLimitPeriodSecond.ToPointer(),
            URL: hookdeck.String("doloremque"),
        },
        DestinationID: hookdeck.String("reprehenderit"),
        Name: "Shawna Carter",
        Rules: []interface{}{
            shared.RetryRule{
                Count: hookdeck.Int64(688661),
                Interval: hookdeck.Int64(317983),
                Strategy: shared.RetryStrategyExponential,
                Type: shared.RetryRuleTypeRetry,
            },
            shared.FilterRule{
                Body: &shared.ConnectionFilterProperty{},
                Headers: &shared.ConnectionFilterProperty{},
                Path: &shared.ConnectionFilterProperty{},
                Query: &shared.ConnectionFilterProperty{},
                Type: shared.FilterRuleTypeFilter,
            },
        },
        Ruleset: &operations.UpsertConnectionRequestBodyRuleset{
            IsTeamDefault: hookdeck.Bool(false),
            Name: "Eric Emmerich",
            Rules: []interface{}{
                shared.DelayRule{
                    Delay: 265389,
                    Type: shared.DelayRuleTypeDelay,
                },
                shared.FilterRule{
                    Body: &shared.ConnectionFilterProperty{},
                    Headers: &shared.ConnectionFilterProperty{},
                    Path: &shared.ConnectionFilterProperty{},
                    Query: &shared.ConnectionFilterProperty{},
                    Type: shared.FilterRuleTypeFilter,
                },
                shared.FilterRule{
                    Body: &shared.ConnectionFilterProperty{},
                    Headers: &shared.ConnectionFilterProperty{},
                    Path: &shared.ConnectionFilterProperty{},
                    Query: &shared.ConnectionFilterProperty{},
                    Type: shared.FilterRuleTypeFilter,
                },
            },
        },
        RulesetID: hookdeck.String("voluptates"),
        Source: &operations.UpsertConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodDelete,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "sint",
                ContentType: shared.SourceCustomResponseContentTypeJSON,
            },
            Name: "Miss Randall Hamill",
        },
        SourceID: hookdeck.String("explicabo"),
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


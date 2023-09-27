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
            AuthMethod: &shared.DestinationAuthMethodConfig{},
            CliPath: hookdeckgo.String("enim"),
            HTTPMethod: shared.DestinationHTTPMethodPut.ToPointer(),
            Name: "Ms. Cathy Marks",
            PathForwardingDisabled: hookdeckgo.Bool(false),
            RateLimit: hookdeckgo.Int64(958950),
            RateLimitPeriod: operations.UpsertConnectionRequestBodyDestinationRateLimitPeriodSecond.ToPointer(),
            URL: hookdeckgo.String("mollitia"),
        },
        DestinationID: hookdeckgo.String("dolorem"),
        Name: "Carlos Ziemann",
        Rules: []shared.Rule{
            shared.Rule{},
        },
        Ruleset: &operations.UpsertConnectionRequestBodyRuleset{
            IsTeamDefault: hookdeckgo.Bool(false),
            Name: "Lucy Konopelski",
            Rules: []shared.Rule{
                shared.Rule{},
            },
        },
        RulesetID: hookdeckgo.String("error"),
        Source: &operations.UpsertConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodGet,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "quis",
                ContentType: shared.SourceCustomResponseContentTypeJSON,
            },
            Name: "Matt Hamill",
        },
        SourceID: hookdeckgo.String("sequi"),
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


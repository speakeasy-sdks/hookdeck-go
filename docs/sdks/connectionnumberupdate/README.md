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
                            APIKey: "Pound",
                            Key: "<key>",
                        },
                        Type: shared.APIKeyTypeAPIKey,
                    },
            ),
            Name: "extranet",
        },
        Name: "SSD",
        Rules: []shared.Rule{
            shared.CreateRuleAlertRule(
                shared.AlertRule{
                    Strategy: shared.AlertStrategyEachAttempt,
                    Type: shared.AlertRuleTypeAlert,
                },
            ),
        },
        Ruleset: &operations.UpsertConnectionRequestBodyRuleset{
            Name: "Triathlon/Time",
            Rules: []shared.Rule{
                shared.CreateRuleTransformRule(
                    shared.CreateTransformRuleTransformReference(
                            shared.TransformReference{
                                TransformationID: "North",
                                Type: shared.TransformReferenceTypeTransform,
                            },
                    ),
                ),
            },
        },
        Source: &operations.UpsertConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodPost,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "Facilitator",
                ContentType: shared.SourceCustomResponseContentTypeText,
            },
            Name: "Van",
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


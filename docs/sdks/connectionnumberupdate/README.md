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
        AdditionalProperties: map[string]interface{}{
            "index": "PNG",
        },
        Destination: &operations.UpsertConnectionRequestBodyDestination{
            AdditionalProperties: map[string]interface{}{
                "extranet": "SSD",
            },
            AuthMethod: shared.CreateDestinationAuthMethodConfigBasicAuth(
                    shared.BasicAuth{
                        AdditionalProperties: map[string]interface{}{
                            "West": "consequently",
                        },
                        Config: &shared.DestinationAuthMethodBasicAuthConfig{
                            AdditionalProperties: map[string]interface{}{
                                "Cisgender": "Northwest",
                            },
                            Password: "crQrPcHBWicBftW",
                            Username: "Rodger_Crona3",
                        },
                        Type: shared.BasicAuthTypeBasicAuth,
                    },
            ),
            Name: "lobby snowman Sarasota",
        },
        Name: "Electric",
        Rules: []shared.Rule{
            shared.CreateRuleDelayRule(
                shared.DelayRule{
                    AdditionalProperties: map[string]interface{}{
                        "up": "Mews",
                    },
                    Delay: 930505,
                    Type: shared.DelayRuleTypeDelay,
                },
            ),
        },
        Ruleset: &operations.UpsertConnectionRequestBodyRuleset{
            AdditionalProperties: map[string]interface{}{
                "District": "back",
            },
            Name: "Northwest",
            Rules: []shared.Rule{
                shared.CreateRuleFilterRule(
                    shared.FilterRule{
                        AdditionalProperties: map[string]interface{}{
                            "program": "Northeast",
                        },
                        Body: shared.CreateConnectionFilterPropertyBoolean(
                        false,
                        ),
                        Headers: shared.CreateConnectionFilterPropertyBoolean(
                        false,
                        ),
                        Path: shared.CreateConnectionFilterPropertyMapOfany(
                                map[string]interface{}{
                                    "Palos": "worthless",
                                },
                        ),
                        Query: shared.CreateConnectionFilterPropertyFloat32(
                        2178.06,
                        ),
                        Type: shared.FilterRuleTypeFilter,
                    },
                ),
            },
        },
        Source: &operations.UpsertConnectionRequestBodySource{
            AdditionalProperties: map[string]interface{}{
                "Romania": "red",
            },
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodPost,
            },
            CustomResponse: &shared.SourceCustomResponse{
                AdditionalProperties: map[string]interface{}{
                    "South": "Tuna",
                },
                Body: "navigate frightfully",
                ContentType: shared.SourceCustomResponseContentTypeJSON,
            },
            Name: "pixel farad Rubber",
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


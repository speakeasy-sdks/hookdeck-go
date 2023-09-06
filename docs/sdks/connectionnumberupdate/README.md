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
            AuthMethod: &shared.BasicAuth{
                Config: &shared.DestinationAuthMethodBasicAuthConfig{
                    Password: "culpa",
                    Username: "Caroline_Ziemann",
                },
                Type: shared.BasicAuthTypeBasicAuth,
            },
            CliPath: hookdeck.String("numquam"),
            HTTPMethod: shared.DestinationHTTPMethodPost.ToPointer(),
            Name: "Nellie Frami",
            PathForwardingDisabled: hookdeck.Bool(false),
            RateLimit: hookdeck.Int64(110375),
            RateLimitPeriod: operations.UpsertConnectionRequestBodyDestinationRateLimitPeriodHour.ToPointer(),
            URL: hookdeck.String("animi"),
        },
        DestinationID: hookdeck.String("enim"),
        Name: "Angelica Dietrich",
        Rules: []interface{}{
            shared.TransformFull{
                Transformation: &shared.TransformFullTransformation{
                    Code: "aut",
                    Env: map[string]string{
                        "quasi": "error",
                    },
                    Name: "Neal Boyer",
                },
                TransformationID: hookdeck.String("vero"),
                Type: shared.TransformFullTypeTransform,
            },
        },
        Ruleset: &operations.UpsertConnectionRequestBodyRuleset{
            IsTeamDefault: hookdeck.Bool(false),
            Name: "Miss Irma Wolff",
            Rules: []interface{}{
                shared.TransformReference{
                    TransformationID: "doloremque",
                    Type: shared.TransformReferenceTypeTransform,
                },
            },
        },
        RulesetID: hookdeck.String("reprehenderit"),
        Source: &operations.UpsertConnectionRequestBodySource{
            AllowedHTTPMethods: []shared.SourceAllowedHTTPMethod{
                shared.SourceAllowedHTTPMethodPost,
            },
            CustomResponse: &shared.SourceCustomResponse{
                Body: "maiores",
                ContentType: shared.SourceCustomResponseContentTypeJSON,
            },
            Name: "Miss Valerie Kshlerin",
        },
        SourceID: hookdeck.String("accusamus"),
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


# TransformationExecutions

### Available Operations

* [Get](#get) - Get transformation executions

## Get

Get transformation executions

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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
    res, err := s.TransformationExecutions.Get(ctx, operations.GetTransformationExecutionsRequest{
        CreatedAt: &operations.GetTransformationExecutionsCreatedAt{},
        Dir: &operations.GetTransformationExecutionsDir{},
        ID: "adcf4b92-1879-4fce-953f-73ef7fbc7abd",
        IssueID: &operations.GetTransformationExecutionsIssueID{},
        Limit: hookdeck.Int64(498140),
        LogLevel: &operations.GetTransformationExecutionsLogLevel{},
        Next: hookdeck.String("dolore"),
        OrderBy: &operations.GetTransformationExecutionsOrderBy{},
        Prev: hookdeck.String("quibusdam"),
        WebhookID: &operations.GetTransformationExecutionsWebhookID{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransformationExecutionPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetTransformationExecutionsRequest](../../models/operations/gettransformationexecutionsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetTransformationExecutionsResponse](../../models/operations/gettransformationexecutionsresponse.md), error**


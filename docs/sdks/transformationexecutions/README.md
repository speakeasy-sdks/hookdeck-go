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
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
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
    res, err := s.TransformationExecutions.Get(ctx, operations.GetTransformationExecutionsRequest{
        CreatedAt: &operations.GetTransformationExecutionsCreatedAt{},
        Dir: &operations.GetTransformationExecutionsDir{},
        ID: "e49a8d9c-bf48-4633-b23f-9b77f3a41006",
        IssueID: &operations.GetTransformationExecutionsIssueID{},
        Limit: hookdeckgo.Int64(487838),
        LogLevel: &operations.GetTransformationExecutionsLogLevel{},
        Next: hookdeckgo.String("quaerat"),
        OrderBy: &operations.GetTransformationExecutionsOrderBy{},
        Prev: hookdeckgo.String("accusamus"),
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


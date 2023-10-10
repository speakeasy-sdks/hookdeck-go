# TransformationExecutions
(*TransformationExecutions*)

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
        CreatedAt: operations.CreateGetTransformationExecutionsCreatedAtGetTransformationExecutionsCreatedAt2(
                operations.GetTransformationExecutionsCreatedAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Northeast": "Hatchback",
                    },
                },
        ),
        Dir: operations.CreateGetTransformationExecutionsDirArrayOfgetTransformationExecutionsDir2(
                []operations.GetTransformationExecutionsDir2{
                    operations.GetTransformationExecutionsDir2Desc,
                },
        ),
        ID: "<ID>",
        IssueID: operations.CreateGetTransformationExecutionsIssueIDStr(
        "towards",
        ),
        LogLevel: operations.CreateGetTransformationExecutionsLogLevelGetTransformationExecutionsLogLevel1(
        operations.GetTransformationExecutionsLogLevel1Error,
        ),
        OrderBy: operations.CreateGetTransformationExecutionsOrderByGetTransformationExecutionsOrderBy1(
        operations.GetTransformationExecutionsOrderBy1CreatedAt,
        ),
        WebhookID: operations.CreateGetTransformationExecutionsWebhookIDStr(
        "Car",
        ),
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


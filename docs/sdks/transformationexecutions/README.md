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
                operations.GetTransformationExecutionsCreatedAt2{},
        ),
        Dir: operations.CreateGetTransformationExecutionsDirGetTransformationExecutionsDir1(
        operations.GetTransformationExecutionsDir1Desc,
        ),
        ID: "<ID>",
        IssueID: operations.CreateGetTransformationExecutionsIssueIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        LogLevel: operations.CreateGetTransformationExecutionsLogLevelArrayOfgetTransformationExecutionsLogLevel2(
                []operations.GetTransformationExecutionsLogLevel2{
                    operations.GetTransformationExecutionsLogLevel2Error,
                },
        ),
        OrderBy: operations.CreateGetTransformationExecutionsOrderByArrayOfgetTransformationExecutionsOrderBy2(
                []operations.GetTransformationExecutionsOrderBy2{
                    operations.GetTransformationExecutionsOrderBy2CreatedAt,
                },
        ),
        WebhookID: operations.CreateGetTransformationExecutionsWebhookIDStr(
        "string",
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


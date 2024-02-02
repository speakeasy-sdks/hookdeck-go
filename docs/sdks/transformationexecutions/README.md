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
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
	"log"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.TransformationExecutions.Get(ctx, operations.GetTransformationExecutionsRequest{
        CreatedAt: operations.CreateGetTransformationExecutionsQueryParamCreatedAtGetTransformationExecutionsQueryParam2(
                operations.GetTransformationExecutionsQueryParam2{},
        ),
        Dir: operations.CreateGetTransformationExecutionsQueryParamDirGetTransformationExecutionsQueryParam1(
        operations.GetTransformationExecutionsQueryParam1Desc,
        ),
        ID: "<ID>",
        IssueID: operations.CreateGetTransformationExecutionsQueryParamIssueIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        LogLevel: operations.CreateLogLevelArrayOfgetTransformationExecutionsQueryParamTransformationExecutionsLogLevel2(
                []operations.GetTransformationExecutionsQueryParamTransformationExecutionsLogLevel2{
                    operations.GetTransformationExecutionsQueryParamTransformationExecutionsLogLevel2Error,
                },
        ),
        OrderBy: operations.CreateGetTransformationExecutionsQueryParamOrderByArrayOfgetTransformationExecutionsQueryParamTransformationExecutionsOrderBy2(
                []operations.GetTransformationExecutionsQueryParamTransformationExecutionsOrderBy2{
                    operations.GetTransformationExecutionsQueryParamTransformationExecutionsOrderBy2CreatedAt,
                },
        ),
        WebhookID: operations.CreateGetTransformationExecutionsQueryParamWebhookIDStr(
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

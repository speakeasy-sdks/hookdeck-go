# TransformationExecution
(*TransformationExecution*)

### Available Operations

* [Get](#get) - Get a transformation execution

## Get

Get a transformation execution

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
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


    var executionID string = "string"

    var id string = "string"

    ctx := context.Background()
    res, err := s.TransformationExecution.Get(ctx, executionID, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransformationExecution != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `executionID`                                         | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetTransformationExecutionResponse](../../models/operations/gettransformationexecutionresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

# TransformationExecution
(*.TransformationExecution*)

### Available Operations

* [Get](#get) - Get a transformation execution

## Get

Get a transformation execution

### Example Usage

```go
package main

import(
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "",
                Username: "",
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


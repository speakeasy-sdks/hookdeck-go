# RequestRawBody
(*RequestRawBody*)

### Available Operations

* [Get](#get) - Get a request raw body data

## Get

Get a request raw body data

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


    var id string = "string"

    ctx := context.Background()
    res, err := s.RequestRawBody.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.RawBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetRequestRawBodyResponse](../../models/operations/getrequestrawbodyresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

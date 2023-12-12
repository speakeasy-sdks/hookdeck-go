# EventRawBody
(*EventRawBody*)

### Available Operations

* [Get](#get) - Get an Event Raw Body Data

## Get

Retrieve a raw body data of a request that Hookdeck receives from a source.

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


    var id string = "string"

    ctx := context.Background()
    res, err := s.EventRawBody.Get(ctx, id)
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

**[*operations.GetEventRawBodyResponse](../../models/operations/geteventrawbodyresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

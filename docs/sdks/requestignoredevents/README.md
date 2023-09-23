# RequestIgnoredEvents

### Available Operations

* [Get](#get) - Get request ignored events

## Get

Get request ignored events

### Example Usage

```go
package main

import(
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.RequestIgnoredEvents.Get(ctx, operations.GetRequestIgnoredEventsRequest{
        Dir: &operations.GetRequestIgnoredEventsDir{},
        IDPathParameter: "vero",
        IDQueryParameter: &operations.GetRequestIgnoredEventsID{},
        Limit: hookdeckgo.Int64(135474),
        Next: hookdeckgo.String("architecto"),
        OrderBy: &operations.GetRequestIgnoredEventsOrderBy{},
        Prev: hookdeckgo.String("magnam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IgnoredEventPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetRequestIgnoredEventsRequest](../../models/operations/getrequestignoredeventsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetRequestIgnoredEventsResponse](../../models/operations/getrequestignoredeventsresponse.md), error**


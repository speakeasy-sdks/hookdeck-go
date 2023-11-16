# RequestIgnoredEvents
(*RequestIgnoredEvents*)

### Available Operations

* [Get](#get) - Get request ignored events

## Get

Get request ignored events

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"log"
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

    ctx := context.Background()
    res, err := s.RequestIgnoredEvents.Get(ctx, operations.GetRequestIgnoredEventsRequest{
        Dir: operations.CreateGetRequestIgnoredEventsQueryParamDirArrayOfgetRequestIgnoredEventsQueryParam2(
                []operations.GetRequestIgnoredEventsQueryParam2{
                    operations.GetRequestIgnoredEventsQueryParam2Asc,
                },
        ),
        IDPathParameter: "string",
        IDQueryParameter: operations.CreateGetRequestIgnoredEventsQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        OrderBy: operations.CreateGetRequestIgnoredEventsQueryParamOrderByArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2(
                []operations.GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2{
                    operations.GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2CreatedAt,
                },
        ),
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,404,422                | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

# IssueTriggers
(*IssueTriggers*)

### Available Operations

* [Get](#get) - Get Issue Triggers

## Get

Retrieve a list of issue triggers.

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
                Password: "<YOUR_PASSWORD_HERE>",
                Username: "<YOUR_USERNAME_HERE>",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.IssueTriggers.Get(ctx, operations.GetIssueTriggersRequest{
        Dir: operations.CreateGetIssueTriggersQueryParamDirArrayOfgetIssueTriggersQueryParam2(
                []operations.GetIssueTriggersQueryParam2{
                    operations.GetIssueTriggersQueryParam2Asc,
                },
        ),
        DisabledAt: operations.CreateDisabledAtGetIssueTriggersQueryParamIssueTriggers2(
                operations.GetIssueTriggersQueryParamIssueTriggers2{},
        ),
        OrderBy: operations.CreateGetIssueTriggersQueryParamOrderByArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2(
                []operations.GetIssueTriggersQueryParamIssueTriggersOrderBy2{
                    operations.GetIssueTriggersQueryParamIssueTriggersOrderBy2Type,
                },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueTriggerPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetIssueTriggersRequest](../../models/operations/getissuetriggersrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetIssueTriggersResponse](../../models/operations/getissuetriggersresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

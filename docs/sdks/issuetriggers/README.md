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
    res, err := s.IssueTriggers.Get(ctx, operations.GetIssueTriggersRequest{
        Dir: operations.CreateGetIssueTriggersDirArrayOfgetIssueTriggersDir2(
                []operations.GetIssueTriggersDir2{
                    operations.GetIssueTriggersDir2Asc,
                },
        ),
        DisabledAt: operations.CreateGetIssueTriggersDisabledAtGetIssueTriggersDisabledAt2(
                operations.GetIssueTriggersDisabledAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Hatchback": "Kia",
                    },
                },
        ),
        OrderBy: operations.CreateGetIssueTriggersOrderByGetIssueTriggersOrderBy1(
        operations.GetIssueTriggersOrderBy1Type,
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


# IssuesCount
(*IssuesCount*)

### Available Operations

* [Get](#get) - Get the number of issues

## Get

Get the number of issues

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
    res, err := s.IssuesCount.Get(ctx, operations.GetIssueCountRequest{
        AggregationKeys: &operations.GetIssueCountAggregationKeys{
            ErrorCode: &operations.GetIssueCountAggregationKeysErrorCode{},
            ResponseStatus: &operations.GetIssueCountAggregationKeysResponseStatus{},
            WebhookID: &operations.GetIssueCountAggregationKeysWebhookID{},
        },
        CreatedAt: &operations.GetIssueCountCreatedAt{},
        Dir: &operations.GetIssueCountDir{},
        DismissedAt: &operations.GetIssueCountDismissedAt{},
        FirstSeenAt: &operations.GetIssueCountFirstSeenAt{},
        ID: &operations.GetIssueCountID{},
        IssueTriggerID: &operations.GetIssueCountIssueTriggerID{},
        LastSeenAt: &operations.GetIssueCountLastSeenAt{},
        Limit: hookdeckgo.Int64(700347),
        MergedWith: &operations.GetIssueCountMergedWith{},
        Next: hookdeckgo.String("program"),
        OrderBy: &operations.GetIssueCountOrderBy{},
        Prev: hookdeckgo.String("Kia Cambridgeshire"),
        Status: &operations.GetIssueCountStatus{},
        Type: &operations.GetIssueCountType{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueCount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetIssueCountRequest](../../models/operations/getissuecountrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetIssueCountResponse](../../models/operations/getissuecountresponse.md), error**


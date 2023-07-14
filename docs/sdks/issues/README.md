# Issues

## Overview

Issues lets you track problems in your workspace and communicate resolution steps with your team.

### Available Operations

* [Get](#get) - Get issues

## Get

Get issues

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
)

func main() {
    s := hookdeck.New(
        hookdeck.WithSecurity(shared.Security{
            BasicAuth: &shared.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.Issues.Get(ctx, operations.GetIssuesRequest{
        AggregationKeys: &operations.GetIssuesAggregationKeys{
            ErrorCode: &operations.GetIssuesAggregationKeysErrorCode{},
            ResponseStatus: &operations.GetIssuesAggregationKeysResponseStatus{},
            WebhookID: &operations.GetIssuesAggregationKeysWebhookID{},
        },
        CreatedAt: &operations.GetIssuesCreatedAt{},
        Dir: &operations.GetIssuesDir{},
        DismissedAt: &operations.GetIssuesDismissedAt{},
        FirstSeenAt: &operations.GetIssuesFirstSeenAt{},
        ID: &operations.GetIssuesID{},
        IssueTriggerID: &operations.GetIssuesIssueTriggerID{},
        LastSeenAt: &operations.GetIssuesLastSeenAt{},
        Limit: hookdeck.Int64(171629),
        MergedWith: &operations.GetIssuesMergedWith{},
        Next: hookdeck.String("quis"),
        OrderBy: &operations.GetIssuesOrderBy{},
        Prev: hookdeck.String("totam"),
        Status: &operations.GetIssuesStatus{},
        Type: &operations.GetIssuesType{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueWithDataPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetIssuesRequest](../../models/operations/getissuesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetIssuesResponse](../../models/operations/getissuesresponse.md), error**


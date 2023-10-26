# Issues
(*Issues*)

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
    res, err := s.Issues.Get(ctx, operations.GetIssuesRequest{
        AggregationKeys: &operations.GetIssuesAggregationKeys{
            ErrorCode: operations.CreateGetIssuesAggregationKeysErrorCodeArrayOfAttemptErrorCode(
                    []shared.AttemptErrorCode{
                        shared.AttemptErrorCodeTimeout,
                    },
            ),
            ResponseStatus: operations.CreateGetIssuesAggregationKeysResponseStatusArrayOffloat32(
                    []float32{
                        8441.99,
                    },
            ),
            WebhookID: operations.CreateGetIssuesAggregationKeysWebhookIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
        },
        CreatedAt: operations.CreateGetIssuesCreatedAtGetIssuesCreatedAt2(
                operations.GetIssuesCreatedAt2{},
        ),
        Dir: operations.CreateGetIssuesDirArrayOfgetIssuesDir2(
                []operations.GetIssuesDir2{
                    operations.GetIssuesDir2Asc,
                },
        ),
        DismissedAt: operations.CreateGetIssuesDismissedAtGetIssuesDismissedAt2(
                operations.GetIssuesDismissedAt2{},
        ),
        FirstSeenAt: operations.CreateGetIssuesFirstSeenAtGetIssuesFirstSeenAt2(
                operations.GetIssuesFirstSeenAt2{},
        ),
        ID: operations.CreateGetIssuesIDStr(
        "iss_YXKv5OdJXCiVwkPhGy",
        ),
        IssueTriggerID: operations.CreateGetIssuesIssueTriggerIDArrayOfstr(
                []string{
                    "i",
                    "t",
                    "_",
                    "B",
                    "X",
                    "K",
                    "v",
                    "5",
                    "O",
                    "d",
                    "J",
                    "X",
                    "C",
                    "i",
                    "V",
                    "w",
                    "k",
                    "P",
                    "h",
                    "G",
                    "y",
                },
        ),
        LastSeenAt: operations.CreateGetIssuesLastSeenAtDateTime(
        types.MustTimeFromString("2022-03-29T23:04:28.455Z"),
        ),
        MergedWith: operations.CreateGetIssuesMergedWithStr(
        "iss_AXKv3OdJXCiKlkPhDz",
        ),
        OrderBy: operations.CreateGetIssuesOrderByArrayOfgetIssuesOrderBy2(
                []operations.GetIssuesOrderBy2{
                    operations.GetIssuesOrderBy2FirstSeenAt,
                },
        ),
        Status: operations.CreateGetIssuesStatusGetIssuesStatus1(
        operations.GetIssuesStatus1Opened,
        ),
        Type: operations.CreateGetIssuesTypeArrayOfgetIssuesType2(
                []operations.GetIssuesType2{
                    operations.GetIssuesType2Delivery,
                    operations.GetIssuesType2Transformation,
                    operations.GetIssuesType2Transformation,
                    operations.GetIssuesType2Backpressure,
                    operations.GetIssuesType2Backpressure,
                    operations.GetIssuesType2Delivery,
                    operations.GetIssuesType2Backpressure,
                    operations.GetIssuesType2Delivery,
                },
        ),
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


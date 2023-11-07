# Issues
(*.Issues*)

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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
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
    res, err := s.Issues.Get(ctx, operations.GetIssuesRequest{
        AggregationKeys: &operations.AggregationKeys{
            ErrorCode: operations.CreateGetIssuesQueryParamErrorCodeArrayOfAttemptErrorCode(
                    []components.AttemptErrorCode{
                        components.AttemptErrorCodeTimeout,
                    },
            ),
            ResponseStatus: operations.CreateGetIssuesQueryParamResponseStatusArrayOffloat32(
                    []float32{
                        8441.99,
                    },
            ),
            WebhookID: operations.CreateGetIssuesQueryParamWebhookIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
        },
        CreatedAt: operations.CreateGetIssuesQueryParamCreatedAtGetIssuesQueryParam2(
                operations.GetIssuesQueryParam2{},
        ),
        Dir: operations.CreateGetIssuesQueryParamDirArrayOfgetIssuesQueryParamIssues2(
                []operations.GetIssuesQueryParamIssues2{
                    operations.GetIssuesQueryParamIssues2Asc,
                },
        ),
        DismissedAt: operations.CreateDismissedAtGetIssuesQueryParamIssuesDismissedAt2(
                operations.GetIssuesQueryParamIssuesDismissedAt2{},
        ),
        FirstSeenAt: operations.CreateFirstSeenAtGetIssuesQueryParamIssuesFirstSeenAt2(
                operations.GetIssuesQueryParamIssuesFirstSeenAt2{},
        ),
        ID: operations.CreateGetIssuesQueryParamIDStr(
        "iss_YXKv5OdJXCiVwkPhGy",
        ),
        IssueTriggerID: operations.CreateIssueTriggerIDArrayOfstr(
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
        LastSeenAt: operations.CreateLastSeenAtDateTime(
        types.MustTimeFromString("2022-03-29T23:04:28.455Z"),
        ),
        MergedWith: operations.CreateMergedWithStr(
        "iss_AXKv3OdJXCiKlkPhDz",
        ),
        OrderBy: operations.CreateGetIssuesQueryParamOrderByArrayOfgetIssuesQueryParamIssuesOrderBy2(
                []operations.GetIssuesQueryParamIssuesOrderBy2{
                    operations.GetIssuesQueryParamIssuesOrderBy2FirstSeenAt,
                },
        ),
        Status: operations.CreateGetIssuesQueryParamStatusGetIssuesQueryParamIssuesStatus1(
        operations.GetIssuesQueryParamIssuesStatus1Opened,
        ),
        Type: operations.CreateTypeArrayOfgetIssuesQueryParamIssuesType2(
                []operations.GetIssuesQueryParamIssuesType2{
                    operations.GetIssuesQueryParamIssuesType2Delivery,
                    operations.GetIssuesQueryParamIssuesType2Transformation,
                    operations.GetIssuesQueryParamIssuesType2Transformation,
                    operations.GetIssuesQueryParamIssuesType2Backpressure,
                    operations.GetIssuesQueryParamIssuesType2Backpressure,
                    operations.GetIssuesQueryParamIssuesType2Delivery,
                    operations.GetIssuesQueryParamIssuesType2Backpressure,
                    operations.GetIssuesQueryParamIssuesType2Delivery,
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


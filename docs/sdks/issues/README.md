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
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
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
                    "it_BXKv5OdJXCiVwkPhGy",
                },
        ),
        LastSeenAt: operations.CreateLastSeenAtDateTime(
        types.MustTimeFromString("2023-03-30T09:00:08.974Z"),
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

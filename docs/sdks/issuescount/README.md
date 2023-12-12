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
    res, err := s.IssuesCount.Get(ctx, operations.GetIssueCountRequest{
        AggregationKeys: &operations.QueryParamAggregationKeys{
            ErrorCode: operations.CreateGetIssueCountQueryParamErrorCodeArrayOfAttemptErrorCode(
                    []components.AttemptErrorCode{
                        components.AttemptErrorCodeTimeout,
                    },
            ),
            ResponseStatus: operations.CreateGetIssueCountQueryParamResponseStatusArrayOffloat32(
                    []float32{
                        8441.99,
                    },
            ),
            WebhookID: operations.CreateGetIssueCountQueryParamWebhookIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
        },
        CreatedAt: operations.CreateGetIssueCountQueryParamCreatedAtGetIssueCountQueryParam2(
                operations.GetIssueCountQueryParam2{},
        ),
        Dir: operations.CreateGetIssueCountQueryParamDirArrayOfgetIssueCountQueryParamIssuesCount2(
                []operations.GetIssueCountQueryParamIssuesCount2{
                    operations.GetIssueCountQueryParamIssuesCount2Asc,
                },
        ),
        DismissedAt: operations.CreateQueryParamDismissedAtGetIssueCountQueryParamIssuesCountDismissedAt2(
                operations.GetIssueCountQueryParamIssuesCountDismissedAt2{},
        ),
        FirstSeenAt: operations.CreateQueryParamFirstSeenAtGetIssueCountQueryParamIssuesCountFirstSeenAt2(
                operations.GetIssueCountQueryParamIssuesCountFirstSeenAt2{},
        ),
        ID: operations.CreateGetIssueCountQueryParamIDStr(
        "iss_YXKv5OdJXCiVwkPhGy",
        ),
        IssueTriggerID: operations.CreateQueryParamIssueTriggerIDArrayOfstr(
                []string{
                    "it_BXKv5OdJXCiVwkPhGy",
                },
        ),
        LastSeenAt: operations.CreateQueryParamLastSeenAtDateTime(
        types.MustTimeFromString("2022-03-29T23:04:28.455Z"),
        ),
        MergedWith: operations.CreateQueryParamMergedWithStr(
        "iss_AXKv3OdJXCiKlkPhDz",
        ),
        OrderBy: operations.CreateGetIssueCountQueryParamOrderByArrayOfgetIssueCountQueryParamIssuesCountOrderBy2(
                []operations.GetIssueCountQueryParamIssuesCountOrderBy2{
                    operations.GetIssueCountQueryParamIssuesCountOrderBy2FirstSeenAt,
                },
        ),
        Status: operations.CreateGetIssueCountQueryParamStatusGetIssueCountQueryParamIssuesCountStatus1(
        operations.GetIssueCountQueryParamIssuesCountStatus1Opened,
        ),
        Type: operations.CreateQueryParamTypeArrayOfgetIssueCountQueryParamIssuesCountType2(
                []operations.GetIssueCountQueryParamIssuesCountType2{
                    operations.GetIssueCountQueryParamIssuesCountType2Delivery,
                },
        ),
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 422                        | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

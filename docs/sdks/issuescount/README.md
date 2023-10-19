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
            ErrorCode: operations.CreateGetIssueCountAggregationKeysErrorCodeArrayOfAttemptErrorCode(
                    []shared.AttemptErrorCode{
                        shared.AttemptErrorCodeTimeout,
                    },
            ),
            ResponseStatus: operations.CreateGetIssueCountAggregationKeysResponseStatusArrayOffloat32(
                    []float32{
                        8441.99,
                    },
            ),
            WebhookID: operations.CreateGetIssueCountAggregationKeysWebhookIDArrayOfstr(
                    []string{
                        "Kia",
                    },
            ),
        },
        CreatedAt: operations.CreateGetIssueCountCreatedAtDateTime(
        types.MustTimeFromString("2023-11-29T02:34:03.781Z"),
        ),
        Dir: operations.CreateGetIssueCountDirArrayOfgetIssueCountDir2(
                []operations.GetIssueCountDir2{
                    operations.GetIssueCountDir2Asc,
                },
        ),
        DismissedAt: operations.CreateGetIssueCountDismissedAtGetIssueCountDismissedAt2(
                operations.GetIssueCountDismissedAt2{},
        ),
        FirstSeenAt: operations.CreateGetIssueCountFirstSeenAtDateTime(
        types.MustTimeFromString("2022-03-29T23:04:28.455Z"),
        ),
        ID: operations.CreateGetIssueCountIDStr(
        "iss_YXKv5OdJXCiVwkPhGy",
        ),
        IssueTriggerID: operations.CreateGetIssueCountIssueTriggerIDArrayOfstr(
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
        LastSeenAt: operations.CreateGetIssueCountLastSeenAtDateTime(
        types.MustTimeFromString("2021-03-16T20:28:26.941Z"),
        ),
        MergedWith: operations.CreateGetIssueCountMergedWithArrayOfstr(
                []string{
                    "i",
                    "s",
                    "s",
                    "_",
                    "A",
                    "X",
                    "K",
                    "v",
                    "3",
                    "O",
                    "d",
                    "J",
                    "X",
                    "C",
                    "i",
                    "K",
                    "l",
                    "k",
                    "P",
                    "h",
                    "D",
                    "z",
                },
        ),
        OrderBy: operations.CreateGetIssueCountOrderByGetIssueCountOrderBy1(
        operations.GetIssueCountOrderBy1LastSeenAt,
        ),
        Status: operations.CreateGetIssueCountStatusGetIssueCountStatus1(
        operations.GetIssueCountStatus1Opened,
        ),
        Type: operations.CreateGetIssueCountTypeArrayOfgetIssueCountType2(
                []operations.GetIssueCountType2{
                    operations.GetIssueCountType2Backpressure,
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Backpressure,
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Transformation,
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Transformation,
                    operations.GetIssueCountType2Delivery,
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


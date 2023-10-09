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
            AdditionalProperties: map[string]interface{}{
                "female": "program",
            },
            ErrorCode: operations.CreateGetIssueCountAggregationKeysErrorCodeArrayOfAttemptErrorCode(
                    []shared.AttemptErrorCode{
                        shared.AttemptErrorCodeHostUnreachable,
                    },
            ),
            ResponseStatus: operations.CreateGetIssueCountAggregationKeysResponseStatusArrayOffloat32(
                    []float32{
                        785.92,
                    },
            ),
            WebhookID: operations.CreateGetIssueCountAggregationKeysWebhookIDArrayOfstr(
                    []string{
                        "joyous",
                    },
            ),
        },
        CreatedAt: operations.CreateGetIssueCountCreatedAtGetIssueCountCreatedAt2(
                operations.GetIssueCountCreatedAt2{
                    AdditionalProperties: map[string]interface{}{
                        "withdrawal": "Car",
                    },
                },
        ),
        Dir: operations.CreateGetIssueCountDirGetIssueCountDir1(
        operations.GetIssueCountDir1Asc,
        ),
        DismissedAt: operations.CreateGetIssueCountDismissedAtDateTime(
        types.MustTimeFromString("2023-05-28T12:03:58.127Z"),
        ),
        FirstSeenAt: operations.CreateGetIssueCountFirstSeenAtGetIssueCountFirstSeenAt2(
                operations.GetIssueCountFirstSeenAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Cambridgeshire": "Table",
                    },
                },
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
        types.MustTimeFromString("2022-01-14T00:57:47.709Z"),
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
        OrderBy: operations.CreateGetIssueCountOrderByArrayOfgetIssueCountOrderBy2(
                []operations.GetIssueCountOrderBy2{
                    operations.GetIssueCountOrderBy2Status,
                },
        ),
        Status: operations.CreateGetIssueCountStatusArrayOfgetIssueCountStatus2(
                []operations.GetIssueCountStatus2{
                    operations.GetIssueCountStatus2Ignored,
                    operations.GetIssueCountStatus2Ignored,
                    operations.GetIssueCountStatus2Acknowledged,
                    operations.GetIssueCountStatus2Resolved,
                    operations.GetIssueCountStatus2Acknowledged,
                    operations.GetIssueCountStatus2Opened,
                },
        ),
        Type: operations.CreateGetIssueCountTypeArrayOfgetIssueCountType2(
                []operations.GetIssueCountType2{
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Transformation,
                    operations.GetIssueCountType2Delivery,
                    operations.GetIssueCountType2Backpressure,
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


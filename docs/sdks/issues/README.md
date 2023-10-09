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
            AdditionalProperties: map[string]interface{}{
                "female": "program",
            },
            ErrorCode: operations.CreateGetIssuesAggregationKeysErrorCodeArrayOfAttemptErrorCode(
                    []shared.AttemptErrorCode{
                        shared.AttemptErrorCodeHostUnreachable,
                    },
            ),
            ResponseStatus: operations.CreateGetIssuesAggregationKeysResponseStatusArrayOffloat32(
                    []float32{
                        785.92,
                    },
            ),
            WebhookID: operations.CreateGetIssuesAggregationKeysWebhookIDArrayOfstr(
                    []string{
                        "joyous",
                    },
            ),
        },
        CreatedAt: operations.CreateGetIssuesCreatedAtGetIssuesCreatedAt2(
                operations.GetIssuesCreatedAt2{
                    AdditionalProperties: map[string]interface{}{
                        "withdrawal": "Car",
                    },
                },
        ),
        Dir: operations.CreateGetIssuesDirGetIssuesDir1(
        operations.GetIssuesDir1Asc,
        ),
        DismissedAt: operations.CreateGetIssuesDismissedAtDateTime(
        types.MustTimeFromString("2023-05-28T12:03:58.127Z"),
        ),
        FirstSeenAt: operations.CreateGetIssuesFirstSeenAtGetIssuesFirstSeenAt2(
                operations.GetIssuesFirstSeenAt2{
                    AdditionalProperties: map[string]interface{}{
                        "Cambridgeshire": "Table",
                    },
                },
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
        types.MustTimeFromString("2022-01-14T00:57:47.709Z"),
        ),
        MergedWith: operations.CreateGetIssuesMergedWithArrayOfstr(
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
        OrderBy: operations.CreateGetIssuesOrderByArrayOfgetIssuesOrderBy2(
                []operations.GetIssuesOrderBy2{
                    operations.GetIssuesOrderBy2Status,
                },
        ),
        Status: operations.CreateGetIssuesStatusArrayOfgetIssuesStatus2(
                []operations.GetIssuesStatus2{
                    operations.GetIssuesStatus2Ignored,
                    operations.GetIssuesStatus2Ignored,
                    operations.GetIssuesStatus2Acknowledged,
                    operations.GetIssuesStatus2Resolved,
                    operations.GetIssuesStatus2Acknowledged,
                    operations.GetIssuesStatus2Opened,
                },
        ),
        Type: operations.CreateGetIssuesTypeArrayOfgetIssuesType2(
                []operations.GetIssuesType2{
                    operations.GetIssuesType2Delivery,
                    operations.GetIssuesType2Delivery,
                    operations.GetIssuesType2Delivery,
                    operations.GetIssuesType2Delivery,
                    operations.GetIssuesType2Transformation,
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


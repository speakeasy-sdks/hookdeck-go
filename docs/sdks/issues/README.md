# Issues

## Overview

Issues lets you track problems in your workspace and communicate resolution steps with your team.

### Available Operations

* [DismissIssue](#dismississue) - Dismiss an issue
* [GetIssue](#getissue) - Get a single issue
* [GetIssueCount](#getissuecount) - Get the number of issues
* [GetIssues](#getissues) - Get issues
* [UpdateIssue](#updateissue) - Update issue

## DismissIssue

Dismiss an issue

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Issues.DismissIssue(ctx, operations.DismissIssueRequest{
        ID: "eab3fec9-578a-4645-8427-3a8418d16230",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Issue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DismissIssueRequest](../../models/operations/dismississuerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DismissIssueResponse](../../models/operations/dismississueresponse.md), error**


## GetIssue

Get a single issue

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Issues.GetIssue(ctx, operations.GetIssueRequest{
        ID: "9fb09299-21ae-4fb9-b58c-4d86e68e4be0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssueWithData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetIssueRequest](../../models/operations/getissuerequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetIssueResponse](../../models/operations/getissueresponse.md), error**


## GetIssueCount

Get the number of issues

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.Issues.GetIssueCount(ctx, operations.GetIssueCountRequest{
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
        Limit: hookdeck.Int64(371919),
        MergedWith: &operations.GetIssueCountMergedWith{},
        Next: hookdeck.String("vel"),
        OrderBy: &operations.GetIssueCountOrderBy{},
        Prev: hookdeck.String("alias"),
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


## GetIssues

Get issues

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.Issues.GetIssues(ctx, operations.GetIssuesRequest{
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
        Limit: hookdeck.Int64(93894),
        MergedWith: &operations.GetIssuesMergedWith{},
        Next: hookdeck.String("non"),
        OrderBy: &operations.GetIssuesOrderBy{},
        Prev: hookdeck.String("maiores"),
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


## UpdateIssue

Update issue

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Issues.UpdateIssue(ctx, operations.UpdateIssueRequest{
        RequestBody: operations.UpdateIssueRequestBody{
            Status: operations.UpdateIssueRequestBodyStatusIgnored,
        },
        ID: "9da757a5-9ecf-4ef6-aef1-caa3383c2beb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Issue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpdateIssueRequest](../../models/operations/updateissuerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.UpdateIssueResponse](../../models/operations/updateissueresponse.md), error**


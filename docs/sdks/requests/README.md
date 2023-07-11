# Requests

## Overview

A request represent a webhook received by Hookdeck.

### Available Operations

* [GetRequest](#getrequest) - Get a request
* [GetRequestEvents](#getrequestevents) - Get request events
* [GetRequestIgnoredEvents](#getrequestignoredevents) - Get request ignored events
* [GetRequestRawBody](#getrequestrawbody) - Get a request raw body data
* [GetRequests](#getrequests) - Get requests
* [RetryRequest](#retryrequest) - Retry a request

## GetRequest

Get a request

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.Requests.GetRequest(ctx, operations.GetRequestRequest{
        ID: "4d1db1f2-c431-4066-9e96-349e1cf9e06e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Request != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetRequestRequest](../../models/operations/getrequestrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetRequestResponse](../../models/operations/getrequestresponse.md), error**


## GetRequestEvents

Get request events

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
	"hookdeck/pkg/types"
	"hookdeck/pkg/models/shared"
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
    res, err := s.Requests.GetRequestEvents(ctx, operations.GetRequestEventsRequest{
        Attempts: &operations.GetRequestEventsAttempts{},
        Body: &operations.GetRequestEventsBody{},
        BulkRetryID: &operations.GetRequestEventsBulkRetryID{},
        CliID: &operations.GetRequestEventsCliID{},
        CliUserID: &operations.GetRequestEventsCliUserID{},
        CreatedAt: &operations.GetRequestEventsCreatedAt{},
        DestinationID: &operations.GetRequestEventsDestinationID{},
        Dir: &operations.GetRequestEventsDir{},
        ErrorCode: &operations.GetRequestEventsErrorCode{},
        EventDataID: &operations.GetRequestEventsEventDataID{},
        Headers: &operations.GetRequestEventsHeaders{},
        IDPathParameter: "velit",
        IDQueryParameter: &operations.GetRequestEventsID{},
        Include: operations.GetRequestEventsIncludeData.ToPointer(),
        IssueID: &operations.GetRequestEventsIssueID{},
        LastAttemptAt: &operations.GetRequestEventsLastAttemptAt{},
        Limit: hookdeck.Int64(673838),
        Next: hookdeck.String("non"),
        OrderBy: &operations.GetRequestEventsOrderBy{},
        ParsedQuery: &operations.GetRequestEventsParsedQuery{},
        Path: hookdeck.String("dolor"),
        Prev: hookdeck.String("iusto"),
        ResponseStatus: &operations.GetRequestEventsResponseStatus{},
        SearchTerm: hookdeck.String("sit"),
        SourceID: &operations.GetRequestEventsSourceID{},
        Status: &operations.GetRequestEventsStatus{},
        SuccessfulAt: &operations.GetRequestEventsSuccessfulAt{},
        WebhookID: &operations.GetRequestEventsWebhookID{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetRequestEventsRequest](../../models/operations/getrequesteventsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetRequestEventsResponse](../../models/operations/getrequesteventsresponse.md), error**


## GetRequestIgnoredEvents

Get request ignored events

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.Requests.GetRequestIgnoredEvents(ctx, operations.GetRequestIgnoredEventsRequest{
        Dir: &operations.GetRequestIgnoredEventsDir{},
        IDPathParameter: "doloremque",
        IDQueryParameter: &operations.GetRequestIgnoredEventsID{},
        Limit: hookdeck.Int64(7468),
        Next: hookdeck.String("officia"),
        OrderBy: &operations.GetRequestIgnoredEventsOrderBy{},
        Prev: hookdeck.String("recusandae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IgnoredEventPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetRequestIgnoredEventsRequest](../../models/operations/getrequestignoredeventsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetRequestIgnoredEventsResponse](../../models/operations/getrequestignoredeventsresponse.md), error**


## GetRequestRawBody

Get a request raw body data

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.Requests.GetRequestRawBody(ctx, operations.GetRequestRawBodyRequest{
        ID: "6b6bc9b8-f759-4eac-95a9-741d31135296",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RawBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetRequestRawBodyRequest](../../models/operations/getrequestrawbodyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetRequestRawBodyResponse](../../models/operations/getrequestrawbodyresponse.md), error**


## GetRequests

Get requests

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
	"hookdeck/pkg/types"
	"hookdeck/pkg/models/shared"
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
    res, err := s.Requests.GetRequests(ctx, operations.GetRequestsRequest{
        Body: &operations.GetRequestsBody{},
        BulkRetryID: &operations.GetRequestsBulkRetryID{},
        CreatedAt: &operations.GetRequestsCreatedAt{},
        Dir: &operations.GetRequestsDir{},
        EventsCount: &operations.GetRequestsEventsCount{},
        Headers: &operations.GetRequestsHeaders{},
        ID: &operations.GetRequestsID{},
        IgnoredCount: &operations.GetRequestsIgnoredCount{},
        Include: operations.GetRequestsIncludeData.ToPointer(),
        IngestedAt: &operations.GetRequestsIngestedAt{},
        Limit: hookdeck.Int64(367626),
        Next: hookdeck.String("soluta"),
        OrderBy: &operations.GetRequestsOrderBy{},
        ParsedQuery: &operations.GetRequestsParsedQuery{},
        Path: hookdeck.String("libero"),
        Prev: hookdeck.String("rem"),
        RejectionCause: &operations.GetRequestsRejectionCause{},
        SearchTerm: hookdeck.String("dolorum"),
        SourceID: &operations.GetRequestsSourceID{},
        Status: operations.GetRequestsStatusAccepted.ToPointer(),
        Verified: hookdeck.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetRequestsRequest](../../models/operations/getrequestsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetRequestsResponse](../../models/operations/getrequestsresponse.md), error**


## RetryRequest

Retry a request

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.Requests.RetryRequest(ctx, operations.RetryRequestRequest{
        RequestBody: operations.RetryRequestRequestBody{
            WebhookIds: []string{
                "alias",
            },
        },
        ID: "2611435e-139d-4bc2-a59b-1abda8c070e1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RetryRequest != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.RetryRequestRequest](../../models/operations/retryrequestrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.RetryRequestResponse](../../models/operations/retryrequestresponse.md), error**


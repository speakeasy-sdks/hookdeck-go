# Requests

## Overview

A request represent a webhook received by Hookdeck.

### Available Operations

* [Get](#get) - Get requests

## Get

Get requests

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
    res, err := s.Requests.Get(ctx, operations.GetRequestsRequest{
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
        Limit: hookdeck.Int64(793698),
        Next: hookdeck.String("quam"),
        OrderBy: &operations.GetRequestsOrderBy{},
        ParsedQuery: &operations.GetRequestsParsedQuery{},
        Path: hookdeck.String("dolor"),
        Prev: hookdeck.String("vero"),
        RejectionCause: &operations.GetRequestsRejectionCause{},
        SearchTerm: hookdeck.String("nostrum"),
        SourceID: &operations.GetRequestsSourceID{},
        Status: operations.GetRequestsStatusRejected.ToPointer(),
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


# Requests
(*Requests*)

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
    res, err := s.Requests.Get(ctx, operations.GetRequestsRequest{
        Body: operations.CreateGetRequestsBodyGetRequestsBody2(
                operations.GetRequestsBody2{},
        ),
        BulkRetryID: operations.CreateGetRequestsBulkRetryIDStr(
        "program",
        ),
        CreatedAt: operations.CreateGetRequestsCreatedAtGetRequestsCreatedAt2(
                operations.GetRequestsCreatedAt2{},
        ),
        Dir: operations.CreateGetRequestsDirArrayOfgetRequestsDir2(
                []operations.GetRequestsDir2{
                    operations.GetRequestsDir2Desc,
                },
        ),
        EventsCount: operations.CreateGetRequestsEventsCountInteger(
        969961,
        ),
        Headers: operations.CreateGetRequestsHeadersGetRequestsHeaders2(
                operations.GetRequestsHeaders2{},
        ),
        ID: operations.CreateGetRequestsIDStr(
        "Xenon",
        ),
        IgnoredCount: operations.CreateGetRequestsIgnoredCountGetRequestsIgnoredCount2(
                operations.GetRequestsIgnoredCount2{},
        ),
        IngestedAt: operations.CreateGetRequestsIngestedAtDateTime(
        types.MustTimeFromString("2023-09-12T22:44:18.030Z"),
        ),
        OrderBy: operations.CreateGetRequestsOrderByGetRequestsOrderBy1(
        operations.GetRequestsOrderBy1IngestedAt,
        ),
        ParsedQuery: operations.CreateGetRequestsParsedQueryGetRequestsParsedQuery2(
                operations.GetRequestsParsedQuery2{},
        ),
        RejectionCause: operations.CreateGetRequestsRejectionCauseGetRequestsRejectionCauseRequestRejectionCause(
        operations.GetRequestsRejectionCauseRequestRejectionCauseUnsupportedHTTPMethod,
        ),
        SourceID: operations.CreateGetRequestsSourceIDStr(
        "Gasoline",
        ),
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


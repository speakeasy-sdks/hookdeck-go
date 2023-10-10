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
        Body: operations.CreateGetRequestsBodyMapOfany(
                map[string]interface{}{
                    "Northeast": "Hatchback",
                },
        ),
        BulkRetryID: operations.CreateGetRequestsBulkRetryIDArrayOfstr(
                []string{
                    "protocol",
                },
        ),
        CreatedAt: operations.CreateGetRequestsCreatedAtGetRequestsCreatedAt2(
                operations.GetRequestsCreatedAt2{
                    AdditionalProperties: map[string]interface{}{
                        "joyous": "Xenon",
                    },
                },
        ),
        Dir: operations.CreateGetRequestsDirGetRequestsDir1(
        operations.GetRequestsDir1Asc,
        ),
        EventsCount: operations.CreateGetRequestsEventsCountArrayOfinteger(
                []int64{
                    213457,
                },
        ),
        Headers: operations.CreateGetRequestsHeadersStr(
        "ah",
        ),
        ID: operations.CreateGetRequestsIDStr(
        "Neon",
        ),
        IgnoredCount: operations.CreateGetRequestsIgnoredCountArrayOfinteger(
                []int64{
                    219956,
                },
        ),
        IngestedAt: operations.CreateGetRequestsIngestedAtDateTime(
        types.MustTimeFromString("2021-12-10T05:49:57.828Z"),
        ),
        OrderBy: operations.CreateGetRequestsOrderByArrayOfgetRequestsOrderBy2(
                []operations.GetRequestsOrderBy2{
                    operations.GetRequestsOrderBy2IngestedAt,
                },
        ),
        ParsedQuery: operations.CreateGetRequestsParsedQueryStr(
        "input",
        ),
        RejectionCause: operations.CreateGetRequestsRejectionCauseArrayOfRequestRejectionCause(
                []shared.RequestRejectionCause{
                    shared.RequestRejectionCauseIngestionFatal,
                },
        ),
        SourceID: operations.CreateGetRequestsSourceIDStr(
        "Reduced",
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


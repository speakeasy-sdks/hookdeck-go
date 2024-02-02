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
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
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
    res, err := s.Requests.Get(ctx, operations.GetRequestsRequest{
        Body: operations.CreateGetRequestsQueryParamBodyGetRequestsQueryParam2(
                operations.GetRequestsQueryParam2{},
        ),
        BulkRetryID: operations.CreateGetRequestsQueryParamBulkRetryIDStr(
        "string",
        ),
        CreatedAt: operations.CreateGetRequestsQueryParamCreatedAtGetRequestsQueryParamRequests2(
                operations.GetRequestsQueryParamRequests2{},
        ),
        Dir: operations.CreateGetRequestsQueryParamDirArrayOfgetRequestsQueryParamRequestsDir2(
                []operations.GetRequestsQueryParamRequestsDir2{
                    operations.GetRequestsQueryParamRequestsDir2Desc,
                },
        ),
        EventsCount: operations.CreateQueryParamEventsCountArrayOfinteger(
                []int64{
                    521235,
                },
        ),
        Headers: operations.CreateGetRequestsQueryParamHeadersStr(
        "string",
        ),
        ID: operations.CreateGetRequestsQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        IgnoredCount: operations.CreateQueryParamIgnoredCountArrayOfinteger(
                []int64{
                    458049,
                },
        ),
        IngestedAt: operations.CreateQueryParamIngestedAtGetRequestsQueryParamRequestsIngestedAt2(
                operations.GetRequestsQueryParamRequestsIngestedAt2{},
        ),
        OrderBy: operations.CreateGetRequestsQueryParamOrderByGetRequestsQueryParamRequests1(
        operations.GetRequestsQueryParamRequests1IngestedAt,
        ),
        ParsedQuery: operations.CreateGetRequestsQueryParamParsedQueryStr(
        "string",
        ),
        RejectionCause: operations.CreateQueryParamRejectionCauseArrayOfRequestRejectionCause(
                []components.RequestRejectionCause{
                    components.RequestRejectionCauseNoWebhook,
                },
        ),
        SourceID: operations.CreateGetRequestsQueryParamSourceIDStr(
        "string",
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

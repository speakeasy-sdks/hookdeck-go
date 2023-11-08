# RequestEvents
(*.RequestEvents*)

### Available Operations

* [Get](#get) - Get request events

## Get

Get request events

### Example Usage

```go
package main

import(
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
)

func main() {
    s := hookdeckgo.New(
        hookdeckgo.WithSecurity(components.Security{
            BasicAuth: &components.SchemeBasicAuth{
                Password: "",
                Username: "",
            },
        }),
    )

    ctx := context.Background()
    res, err := s.RequestEvents.Get(ctx, operations.GetRequestEventsRequest{
        Attempts: operations.CreateGetRequestEventsQueryParamAttemptsGetRequestEventsQueryParam2(
                operations.GetRequestEventsQueryParam2{},
        ),
        Body: operations.CreateGetRequestEventsQueryParamBodyStr(
        "string",
        ),
        BulkRetryID: operations.CreateGetRequestEventsQueryParamBulkRetryIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliID: operations.CreateGetRequestEventsQueryParamCliIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CliUserID: operations.CreateGetRequestEventsQueryParamCliUserIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        CreatedAt: operations.CreateGetRequestEventsQueryParamCreatedAtGetRequestEventsQueryParamRequestEventsCreatedAt2(
                operations.GetRequestEventsQueryParamRequestEventsCreatedAt2{},
        ),
        DestinationID: operations.CreateGetRequestEventsQueryParamDestinationIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Dir: operations.CreateGetRequestEventsQueryParamDirGetRequestEventsQueryParam1(
        operations.GetRequestEventsQueryParam1Desc,
        ),
        ErrorCode: operations.CreateGetRequestEventsQueryParamErrorCodeArrayOfstr(
                []string{
                    "string",
                },
        ),
        EventDataID: operations.CreateGetRequestEventsQueryParamEventDataIDStr(
        "string",
        ),
        Headers: operations.CreateGetRequestEventsQueryParamHeadersGetRequestEventsQueryParamRequestEventsHeaders2(
                operations.GetRequestEventsQueryParamRequestEventsHeaders2{},
        ),
        IDPathParameter: "string",
        IDQueryParameter: operations.CreateGetRequestEventsQueryParamIDStr(
        "string",
        ),
        IssueID: operations.CreateGetRequestEventsQueryParamIssueIDStr(
        "string",
        ),
        LastAttemptAt: operations.CreateGetRequestEventsQueryParamLastAttemptAtDateTime(
        types.MustTimeFromString("2023-09-12T22:44:18.030Z"),
        ),
        OrderBy: operations.CreateGetRequestEventsQueryParamOrderByGetRequestEventsQueryParamRequestEvents1(
        operations.GetRequestEventsQueryParamRequestEvents1LastAttemptAt,
        ),
        ParsedQuery: operations.CreateGetRequestEventsQueryParamParsedQueryGetRequestEventsQueryParamRequestEventsParsedQuery2(
                operations.GetRequestEventsQueryParamRequestEventsParsedQuery2{},
        ),
        ResponseStatus: operations.CreateGetRequestEventsQueryParamResponseStatusInteger(
        438142,
        ),
        SourceID: operations.CreateGetRequestEventsQueryParamSourceIDStr(
        "string",
        ),
        Status: operations.CreateGetRequestEventsQueryParamStatusArrayOfEventStatus(
                []components.EventStatus{
                    components.EventStatusSuccessful,
                },
        ),
        SuccessfulAt: operations.CreateGetRequestEventsQueryParamSuccessfulAtDateTime(
        types.MustTimeFromString("2023-12-06T14:41:34.659Z"),
        ),
        WebhookID: operations.CreateGetRequestEventsQueryParamWebhookIDStr(
        "string",
        ),
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,404,422                | application/json           |
| sdkerrors.SDKError         | 400-600                    | */*                        |

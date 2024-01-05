# BulkRetryEvent
(*BulkRetryEvent*)

### Available Operations

* [Cancel](#cancel) - Cancel an events bulk retry
* [Create](#create) - Create an events bulk retry
* [Get](#get) - Get an events bulk retry

## Cancel

Cancel an events bulk retry

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
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


    var id string = "string"

    ctx := context.Background()
    res, err := s.BulkRetryEvent.Cancel(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.CancelEventBulkRetryResponse](../../models/operations/canceleventbulkretryresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Create

Create an events bulk retry

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/types"
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
    res, err := s.BulkRetryEvent.Create(ctx, operations.CreateEventBulkRetryRequestBody{
        Query: &operations.Query{
            Attempts: operations.CreateAttemptsInteger(
            489382,
            ),
            Body: operations.CreateBodyCreateEventBulkRetry2(
                    operations.CreateEventBulkRetry2{},
            ),
            BulkRetryID: operations.CreateBulkRetryIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            CliID: operations.CreateCliIDCreateEventBulkRetryBulkRetryEvent2(
                    operations.CreateEventBulkRetryBulkRetryEvent2{},
            ),
            CliUserID: operations.CreateCliUserIDStr(
            "string",
            ),
            CreatedAt: operations.CreateCreatedAtDateTime(
            types.MustTimeFromString("2024-05-12T01:18:11.295Z"),
            ),
            DestinationID: operations.CreateDestinationIDStr(
            "string",
            ),
            ErrorCode: operations.CreateErrorCodeArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            EventDataID: operations.CreateEventDataIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            Headers: operations.CreateHeadersStr(
            "string",
            ),
            ID: operations.CreateIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            IssueID: operations.CreateIssueIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            LastAttemptAt: operations.CreateLastAttemptAtDateTime(
            types.MustTimeFromString("2023-05-14T22:08:51.375Z"),
            ),
            ParsedQuery: operations.CreateParsedQueryStr(
            "string",
            ),
            ResponseStatus: operations.CreateResponseStatusInteger(
            89964,
            ),
            SourceID: operations.CreateSourceIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            Status: operations.CreateStatusArrayOfEventStatus(
                    []components.EventStatus{
                        components.EventStatusFailed,
                    },
            ),
            SuccessfulAt: operations.CreateSuccessfulAtCreateEventBulkRetryBulkRetryEventRequestRequestBodyQuerySuccessfulAt2(
                    operations.CreateEventBulkRetryBulkRetryEventRequestRequestBodyQuerySuccessfulAt2{},
            ),
            WebhookID: operations.CreateWebhookIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateEventBulkRetryRequestBody](../../models/operations/createeventbulkretryrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.CreateEventBulkRetryResponse](../../models/operations/createeventbulkretryresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

## Get

Get an events bulk retry

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"context"
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


    var id string = "string"

    ctx := context.Background()
    res, err := s.BulkRetryEvent.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.BatchOperation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.GetEventBulkRetryResponse](../../models/operations/geteventbulkretryresponse.md), error**
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 404                        | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

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
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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


## Create

Create an events bulk retry

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
    res, err := s.BulkRetryEvent.Create(ctx, operations.CreateEventBulkRetryRequestBody{
        Query: &operations.CreateEventBulkRetryRequestBodyQuery{
            Attempts: operations.CreateCreateEventBulkRetryRequestBodyQueryAttemptsInteger(
            489382,
            ),
            Body: operations.CreateCreateEventBulkRetryRequestBodyQueryBodyCreateEventBulkRetryRequestBodyQueryBody2(
                    operations.CreateEventBulkRetryRequestBodyQueryBody2{},
            ),
            BulkRetryID: operations.CreateCreateEventBulkRetryRequestBodyQueryBulkRetryIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            CliID: operations.CreateCreateEventBulkRetryRequestBodyQueryCliIDCreateEventBulkRetryRequestBodyQueryCliID2(
                    operations.CreateEventBulkRetryRequestBodyQueryCliID2{},
            ),
            CliUserID: operations.CreateCreateEventBulkRetryRequestBodyQueryCliUserIDStr(
            "string",
            ),
            CreatedAt: operations.CreateCreateEventBulkRetryRequestBodyQueryCreatedAtDateTime(
            types.MustTimeFromString("2023-05-12T06:25:33.730Z"),
            ),
            DestinationID: operations.CreateCreateEventBulkRetryRequestBodyQueryDestinationIDStr(
            "string",
            ),
            ErrorCode: operations.CreateCreateEventBulkRetryRequestBodyQueryErrorCodeArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            EventDataID: operations.CreateCreateEventBulkRetryRequestBodyQueryEventDataIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            Headers: operations.CreateCreateEventBulkRetryRequestBodyQueryHeadersStr(
            "string",
            ),
            ID: operations.CreateCreateEventBulkRetryRequestBodyQueryIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            IssueID: operations.CreateCreateEventBulkRetryRequestBodyQueryIssueIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            LastAttemptAt: operations.CreateCreateEventBulkRetryRequestBodyQueryLastAttemptAtDateTime(
            types.MustTimeFromString("2022-05-14T11:13:20.233Z"),
            ),
            ParsedQuery: operations.CreateCreateEventBulkRetryRequestBodyQueryParsedQueryStr(
            "string",
            ),
            ResponseStatus: operations.CreateCreateEventBulkRetryRequestBodyQueryResponseStatusInteger(
            89964,
            ),
            SourceID: operations.CreateCreateEventBulkRetryRequestBodyQuerySourceIDArrayOfstr(
                    []string{
                        "string",
                    },
            ),
            Status: operations.CreateCreateEventBulkRetryRequestBodyQueryStatusArrayOfEventStatus(
                    []shared.EventStatus{
                        shared.EventStatusFailed,
                    },
            ),
            SuccessfulAt: operations.CreateCreateEventBulkRetryRequestBodyQuerySuccessfulAtCreateEventBulkRetryRequestBodyQuerySuccessfulAt2(
                    operations.CreateEventBulkRetryRequestBodyQuerySuccessfulAt2{},
            ),
            WebhookID: operations.CreateCreateEventBulkRetryRequestBodyQueryWebhookIDArrayOfstr(
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


## Get

Get an events bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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


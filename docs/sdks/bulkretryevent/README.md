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


    var id string = "Clifton"

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
                        "Money",
                    },
            ),
            CliID: operations.CreateCreateEventBulkRetryRequestBodyQueryCliIDArrayOfstr(
                    []string{
                        "Cambridgeshire",
                    },
            ),
            CliUserID: operations.CreateCreateEventBulkRetryRequestBodyQueryCliUserIDArrayOfstr(
                    []string{
                        "abnormally",
                    },
            ),
            CreatedAt: operations.CreateCreateEventBulkRetryRequestBodyQueryCreatedAtDateTime(
            types.MustTimeFromString("2021-07-05T20:25:56.241Z"),
            ),
            DestinationID: operations.CreateCreateEventBulkRetryRequestBodyQueryDestinationIDStr(
            "Northwest",
            ),
            ErrorCode: operations.CreateCreateEventBulkRetryRequestBodyQueryErrorCodeArrayOfstr(
                    []string{
                        "SUV",
                    },
            ),
            EventDataID: operations.CreateCreateEventBulkRetryRequestBodyQueryEventDataIDArrayOfstr(
                    []string{
                        "Screen",
                    },
            ),
            Headers: operations.CreateCreateEventBulkRetryRequestBodyQueryHeadersStr(
            "physical",
            ),
            ID: operations.CreateCreateEventBulkRetryRequestBodyQueryIDStr(
            "Durham",
            ),
            IssueID: operations.CreateCreateEventBulkRetryRequestBodyQueryIssueIDArrayOfstr(
                    []string{
                        "South",
                    },
            ),
            LastAttemptAt: operations.CreateCreateEventBulkRetryRequestBodyQueryLastAttemptAtDateTime(
            types.MustTimeFromString("2021-08-16T06:48:41.243Z"),
            ),
            ParsedQuery: operations.CreateCreateEventBulkRetryRequestBodyQueryParsedQueryStr(
            "female",
            ),
            ResponseStatus: operations.CreateCreateEventBulkRetryRequestBodyQueryResponseStatusArrayOfinteger(
                    []int64{
                        322997,
                    },
            ),
            SourceID: operations.CreateCreateEventBulkRetryRequestBodyQuerySourceIDArrayOfstr(
                    []string{
                        "Grocery",
                    },
            ),
            Status: operations.CreateCreateEventBulkRetryRequestBodyQueryStatusEventStatus(
            shared.EventStatusHold,
            ),
            SuccessfulAt: operations.CreateCreateEventBulkRetryRequestBodyQuerySuccessfulAtDateTime(
            types.MustTimeFromString("2022-01-27T17:59:46.224Z"),
            ),
            WebhookID: operations.CreateCreateEventBulkRetryRequestBodyQueryWebhookIDArrayOfstr(
                    []string{
                        "Kentucky",
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


    var id string = "female"

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


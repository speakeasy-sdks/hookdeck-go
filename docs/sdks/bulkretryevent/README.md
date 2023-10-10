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
        AdditionalProperties: map[string]interface{}{
            "online": "Configuration",
        },
        Query: &operations.CreateEventBulkRetryRequestBodyQuery{
            AdditionalProperties: map[string]interface{}{
                "Money": "blue",
            },
            Attempts: operations.CreateCreateEventBulkRetryRequestBodyQueryAttemptsCreateEventBulkRetryRequestBodyQueryAttempts2(
                    operations.CreateEventBulkRetryRequestBodyQueryAttempts2{
                        AdditionalProperties: map[string]interface{}{
                            "grey": "technology",
                        },
                    },
            ),
            Body: operations.CreateCreateEventBulkRetryRequestBodyQueryBodyStr(
            "deposit",
            ),
            BulkRetryID: operations.CreateCreateEventBulkRetryRequestBodyQueryBulkRetryIDStr(
            "Northwest",
            ),
            CliID: operations.CreateCreateEventBulkRetryRequestBodyQueryCliIDArrayOfstr(
                    []string{
                        "SUV",
                    },
            ),
            CliUserID: operations.CreateCreateEventBulkRetryRequestBodyQueryCliUserIDArrayOfstr(
                    []string{
                        "Screen",
                    },
            ),
            CreatedAt: operations.CreateCreateEventBulkRetryRequestBodyQueryCreatedAtDateTime(
            types.MustTimeFromString("2023-11-15T08:41:39.306Z"),
            ),
            DestinationID: operations.CreateCreateEventBulkRetryRequestBodyQueryDestinationIDArrayOfstr(
                    []string{
                        "Ameliorated",
                    },
            ),
            ErrorCode: operations.CreateCreateEventBulkRetryRequestBodyQueryErrorCodeStr(
            "after",
            ),
            EventDataID: operations.CreateCreateEventBulkRetryRequestBodyQueryEventDataIDArrayOfstr(
                    []string{
                        "Intelligent",
                    },
            ),
            Headers: operations.CreateCreateEventBulkRetryRequestBodyQueryHeadersStr(
            "female",
            ),
            ID: operations.CreateCreateEventBulkRetryRequestBodyQueryIDArrayOfstr(
                    []string{
                        "functionalities",
                    },
            ),
            IssueID: operations.CreateCreateEventBulkRetryRequestBodyQueryIssueIDStr(
            "Account",
            ),
            LastAttemptAt: operations.CreateCreateEventBulkRetryRequestBodyQueryLastAttemptAtDateTime(
            types.MustTimeFromString("2022-01-27T17:59:46.224Z"),
            ),
            ParsedQuery: operations.CreateCreateEventBulkRetryRequestBodyQueryParsedQueryMapOfany(
                    map[string]interface{}{
                        "Kentucky": "animated",
                    },
            ),
            ResponseStatus: operations.CreateCreateEventBulkRetryRequestBodyQueryResponseStatusArrayOfinteger(
                    []int64{
                        621636,
                    },
            ),
            SourceID: operations.CreateCreateEventBulkRetryRequestBodyQuerySourceIDArrayOfstr(
                    []string{
                        "Senior",
                    },
            ),
            Status: operations.CreateCreateEventBulkRetryRequestBodyQueryStatusEventStatus(
            shared.EventStatusScheduled,
            ),
            SuccessfulAt: operations.CreateCreateEventBulkRetryRequestBodyQuerySuccessfulAtDateTime(
            types.MustTimeFromString("2023-09-22T09:00:36.734Z"),
            ),
            WebhookID: operations.CreateCreateEventBulkRetryRequestBodyQueryWebhookIDArrayOfstr(
                    []string{
                        "Towels",
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


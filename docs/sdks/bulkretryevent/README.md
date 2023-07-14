# BulkRetryEvent

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
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.BulkRetryEvent.Cancel(ctx, "quis")
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
    res, err := s.BulkRetryEvent.Create(ctx, operations.CreateEventBulkRetryRequestBody{
        Query: &operations.CreateEventBulkRetryRequestBodyQuery{
            Attempts: &operations.CreateEventBulkRetryRequestBodyQueryAttempts{},
            Body: &operations.CreateEventBulkRetryRequestBodyQueryBody{},
            BulkRetryID: &operations.CreateEventBulkRetryRequestBodyQueryBulkRetryID{},
            CliID: &operations.CreateEventBulkRetryRequestBodyQueryCliID{},
            CliUserID: &operations.CreateEventBulkRetryRequestBodyQueryCliUserID{},
            CreatedAt: &operations.CreateEventBulkRetryRequestBodyQueryCreatedAt{},
            DestinationID: &operations.CreateEventBulkRetryRequestBodyQueryDestinationID{},
            ErrorCode: &operations.CreateEventBulkRetryRequestBodyQueryErrorCode{},
            EventDataID: &operations.CreateEventBulkRetryRequestBodyQueryEventDataID{},
            Headers: &operations.CreateEventBulkRetryRequestBodyQueryHeaders{},
            ID: &operations.CreateEventBulkRetryRequestBodyQueryID{},
            IssueID: &operations.CreateEventBulkRetryRequestBodyQueryIssueID{},
            LastAttemptAt: &operations.CreateEventBulkRetryRequestBodyQueryLastAttemptAt{},
            ParsedQuery: &operations.CreateEventBulkRetryRequestBodyQueryParsedQuery{},
            Path: hookdeck.String("veritatis"),
            ResponseStatus: &operations.CreateEventBulkRetryRequestBodyQueryResponseStatus{},
            SearchTerm: hookdeck.String("deserunt"),
            SourceID: &operations.CreateEventBulkRetryRequestBodyQuerySourceID{},
            Status: &operations.CreateEventBulkRetryRequestBodyQueryStatus{},
            SuccessfulAt: &operations.CreateEventBulkRetryRequestBodyQuerySuccessfulAt{},
            WebhookID: &operations.CreateEventBulkRetryRequestBodyQueryWebhookID{},
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
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.BulkRetryEvent.Get(ctx, "perferendis")
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


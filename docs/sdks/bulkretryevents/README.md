# BulkRetryEvents

### Available Operations

* [CancelEventBulkRetry](#canceleventbulkretry) - Cancel an events bulk retry
* [CreateEventBulkRetry](#createeventbulkretry) - Create an events bulk retry
* [GetEventBulkRetry](#geteventbulkretry) - Get an events bulk retry

## CancelEventBulkRetry

Cancel an events bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.BulkRetryEvents.CancelEventBulkRetry(ctx, operations.CancelEventBulkRetryRequest{
        ID: "c969e9a3-efa7-47df-b14c-d66ae395efb9",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CancelEventBulkRetryRequest](../../models/operations/canceleventbulkretryrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CancelEventBulkRetryResponse](../../models/operations/canceleventbulkretryresponse.md), error**


## CreateEventBulkRetry

Create an events bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
	"hookdeck/pkg/types"
	"hookdeck/pkg/models/shared"
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
    res, err := s.BulkRetryEvents.CreateEventBulkRetry(ctx, operations.CreateEventBulkRetryRequestBody{
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
            Path: hookdeck.String("nam"),
            ResponseStatus: &operations.CreateEventBulkRetryRequestBodyQueryResponseStatus{},
            SearchTerm: hookdeck.String("id"),
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


## GetEventBulkRetry

Get an events bulk retry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.BulkRetryEvents.GetEventBulkRetry(ctx, operations.GetEventBulkRetryRequest{
        ID: "88f3a669-9707-44ba-8469-b6e214195989",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetEventBulkRetryRequest](../../models/operations/geteventbulkretryrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetEventBulkRetryResponse](../../models/operations/geteventbulkretryresponse.md), error**


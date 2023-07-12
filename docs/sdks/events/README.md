# Events

## Overview

An event is any request that Hookdeck receives from a source.

### Available Operations

* [GetEvent](#getevent) - Get an Event
* [GetEventRawBody](#geteventrawbody) - Get an Event Raw Body Data
* [GetEvents](#getevents) - Get Events
* [MuteEvent](#muteevent) - Mute an Event
* [RetryEvent](#retryevent) - Retry an Event

## GetEvent

Retrieve a request that Hookdeck receives from a source.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.Events.GetEvent(ctx, operations.GetEventRequest{
        ID: "8f5c0b2f-2fb7-4b19-8a27-6b26916fe1f0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Event != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetEventRequest](../../models/operations/geteventrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetEventResponse](../../models/operations/geteventresponse.md), error**


## GetEventRawBody

Retrieve a raw body data of a request that Hookdeck receives from a source.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.Events.GetEventRawBody(ctx, operations.GetEventRawBodyRequest{
        ID: "8f4294e3-698f-4447-b603-e8b445e80ca5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RawBody != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetEventRawBodyRequest](../../models/operations/geteventrawbodyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetEventRawBodyResponse](../../models/operations/geteventrawbodyresponse.md), error**


## GetEvents

Retrieve a list of requests that Hookdeck receives from a source.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/types"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
    res, err := s.Events.GetEvents(ctx, operations.GetEventsRequest{
        Attempts: &operations.GetEventsAttempts{},
        Body: &operations.GetEventsBody{},
        BulkRetryID: &operations.GetEventsBulkRetryID{},
        CliID: &operations.GetEventsCliID{},
        CliUserID: &operations.GetEventsCliUserID{},
        CreatedAt: &operations.GetEventsCreatedAt{},
        DestinationID: &operations.GetEventsDestinationID{},
        Dir: &operations.GetEventsDir{},
        ErrorCode: &operations.GetEventsErrorCode{},
        EventDataID: &operations.GetEventsEventDataID{},
        Headers: &operations.GetEventsHeaders{},
        ID: &operations.GetEventsID{},
        Include: operations.GetEventsIncludeData.ToPointer(),
        IssueID: &operations.GetEventsIssueID{},
        LastAttemptAt: &operations.GetEventsLastAttemptAt{},
        Limit: hookdeck.Int64(329543),
        Next: hookdeck.String("recusandae"),
        OrderBy: &operations.GetEventsOrderBy{},
        ParsedQuery: &operations.GetEventsParsedQuery{},
        Path: hookdeck.String("reiciendis"),
        Prev: hookdeck.String("nulla"),
        ResponseStatus: &operations.GetEventsResponseStatus{},
        SearchTerm: hookdeck.String("magni"),
        SourceID: &operations.GetEventsSourceID{},
        Status: &operations.GetEventsStatus{},
        SuccessfulAt: &operations.GetEventsSuccessfulAt{},
        WebhookID: &operations.GetEventsWebhookID{},
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetEventsRequest](../../models/operations/geteventsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetEventsResponse](../../models/operations/geteventsresponse.md), error**


## MuteEvent

Mute a request that Hookdeck receives from a specific source.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.Events.MuteEvent(ctx, operations.MuteEventRequest{
        ID: "0e457e18-58b6-4a89-bbe3-a5aa8e4824d0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Event != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.MuteEventRequest](../../models/operations/muteeventrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.MuteEventResponse](../../models/operations/muteeventresponse.md), error**


## RetryEvent

Retry the operation to retrieve a request that Hookdeck receives from a specific source.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hookdeck-go"
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
    res, err := s.Events.RetryEvent(ctx, operations.RetryEventRequest{
        ID: "ab407508-8e51-4862-865e-904f3b1194b8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RetriedEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.RetryEventRequest](../../models/operations/retryeventrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.RetryEventResponse](../../models/operations/retryeventresponse.md), error**


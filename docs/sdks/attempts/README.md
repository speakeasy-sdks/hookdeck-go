# Attempts

## Overview

An attempt is any request that Hookdeck makes on behalf of an event.

### Available Operations

* [GetAttempt](#getattempt) - Get a Single Attempt
* [GetAttempts](#getattempts) - Get Attempts

## GetAttempt

Retrive a single attempt details.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Attempts.GetAttempt(ctx, operations.GetAttemptRequest{
        ID: "a05dfc2d-df7c-4c78-8a1b-a928fc816742",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventAttempt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetAttemptRequest](../../models/operations/getattemptrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetAttemptResponse](../../models/operations/getattemptresponse.md), error**


## GetAttempts

Retrieve a list of attempts.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
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
    res, err := s.Attempts.GetAttempts(ctx, operations.GetAttemptsRequest{
        Dir: &operations.GetAttemptsDir{},
        EventID: &operations.GetAttemptsEventID{},
        Limit: hookdeck.Int64(774234),
        Next: hookdeck.String("cum"),
        OrderBy: &operations.GetAttemptsOrderBy{},
        Prev: hookdeck.String("esse"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventAttemptPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetAttemptsRequest](../../models/operations/getattemptsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetAttemptsResponse](../../models/operations/getattemptsresponse.md), error**


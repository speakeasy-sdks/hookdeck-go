# Attempts

## Overview

An attempt is any request that Hookdeck makes on behalf of an event.

### Available Operations

* [Get](#get) - Get Attempts

## Get

Retrieve a list of attempts.

### Example Usage

```go
package main

import(
	"context"
	"log"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
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
    res, err := s.Attempts.Get(ctx, operations.GetAttemptsRequest{
        Dir: &operations.GetAttemptsDir{},
        EventID: &operations.GetAttemptsEventID{},
        Limit: hookdeckgo.Int64(844266),
        Next: hookdeckgo.String("unde"),
        OrderBy: &operations.GetAttemptsOrderBy{},
        Prev: hookdeckgo.String("nulla"),
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


# Bookmarks

## Overview

A bookmark lets you conveniently store and replay a specific request.

### Available Operations

* [Get](#get) - Get Bookmarks

## Get

Retrieve a list of bookmarks.

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
    res, err := s.Bookmarks.Get(ctx, operations.GetBookmarksRequest{
        Dir: &operations.GetBookmarksDir{},
        EventDataID: &operations.GetBookmarksEventDataID{},
        ID: &operations.GetBookmarksID{},
        Label: &operations.GetBookmarksLabel{},
        LastUsedAt: &operations.GetBookmarksLastUsedAt{},
        Limit: hookdeck.Int64(925597),
        Name: &operations.GetBookmarksName{},
        Next: hookdeck.String("temporibus"),
        OrderBy: &operations.GetBookmarksOrderBy{},
        Prev: hookdeck.String("ab"),
        WebhookID: &operations.GetBookmarksWebhookID{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BookmarkPaginatedResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetBookmarksRequest](../../models/operations/getbookmarksrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetBookmarksResponse](../../models/operations/getbookmarksresponse.md), error**


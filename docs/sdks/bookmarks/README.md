# Bookmarks
(*Bookmarks*)

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
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	hookdeckgo "github.com/speakeasy-sdks/hookdeck-go/v2"
	"context"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/operations"
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
    res, err := s.Bookmarks.Get(ctx, operations.GetBookmarksRequest{
        Dir: operations.CreateQueryParamDirArrayOfgetBookmarksQueryParam2(
                []operations.GetBookmarksQueryParam2{
                    operations.GetBookmarksQueryParam2Asc,
                },
        ),
        EventDataID: operations.CreateQueryParamEventDataIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        ID: operations.CreateQueryParamIDArrayOfstr(
                []string{
                    "string",
                },
        ),
        Label: operations.CreateLabelArrayOfstr(
                []string{
                    "string",
                },
        ),
        LastUsedAt: operations.CreateLastUsedAtGetBookmarksQueryParamBookmarks2(
                operations.GetBookmarksQueryParamBookmarks2{},
        ),
        Name: operations.CreateNameArrayOfstr(
                []string{
                    "string",
                },
        ),
        OrderBy: operations.CreateQueryParamOrderByGetBookmarksQueryParamBookmarks1(
        operations.GetBookmarksQueryParamBookmarks1CreatedAt,
        ),
        WebhookID: operations.CreateQueryParamWebhookIDArrayOfstr(
                []string{
                    "string",
                },
        ),
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
| Error Object               | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| sdkerrors.APIErrorResponse | 400,422                    | application/json           |
| sdkerrors.SDKError         | 4xx-5xx                    | */*                        |

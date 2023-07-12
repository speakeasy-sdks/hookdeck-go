# Bookmarks

## Overview

A bookmark lets you conveniently store and replay a specific request.

### Available Operations

* [CreateBookmark](#createbookmark) - Create a Bookmark
* [DeleteBookmark](#deletebookmark) - Delete a Bookmark
* [GetBookmark](#getbookmark) - Get a Single Bookmark
* [GetBookmarks](#getbookmarks) - Get Bookmarks
* [TriggerBookmark](#triggerbookmark) - Trigger a Bookmark
* [UpdateBookmark](#updatebookmark) - Update a Bookmark

## CreateBookmark

Create a new bookmark.

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
    res, err := s.Bookmarks.CreateBookmark(ctx, operations.CreateBookmarkRequestBody{
        EventDataID: "ipsum",
        Label: "excepturi",
        Name: hookdeck.String("Dorothy Hane"),
        WebhookID: "iste",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Bookmark != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateBookmarkRequestBody](../../models/operations/createbookmarkrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateBookmarkResponse](../../models/operations/createbookmarkresponse.md), error**


## DeleteBookmark

Delete an existing bookmark

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
    res, err := s.Bookmarks.DeleteBookmark(ctx, operations.DeleteBookmarkRequest{
        ID: "396fea75-96eb-410f-aaa2-352c5955907a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletedBookmarkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteBookmarkRequest](../../models/operations/deletebookmarkrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteBookmarkResponse](../../models/operations/deletebookmarkresponse.md), error**


## GetBookmark

Retrieve an existing bookmark details.

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
    res, err := s.Bookmarks.GetBookmark(ctx, operations.GetBookmarkRequest{
        ID: "ff1a3a2f-a946-4773-9251-aa52c3f5ad01",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Bookmark != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetBookmarkRequest](../../models/operations/getbookmarkrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetBookmarkResponse](../../models/operations/getbookmarkresponse.md), error**


## GetBookmarks

Retrieve a list of bookmarks.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck-go"
	"hookdeck-go/pkg/models/operations"
	"hookdeck-go/pkg/types"
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
    res, err := s.Bookmarks.GetBookmarks(ctx, operations.GetBookmarksRequest{
        Dir: &operations.GetBookmarksDir{},
        EventDataID: &operations.GetBookmarksEventDataID{},
        ID: &operations.GetBookmarksID{},
        Label: &operations.GetBookmarksLabel{},
        LastUsedAt: &operations.GetBookmarksLastUsedAt{},
        Limit: hookdeck.Int64(622846),
        Name: &operations.GetBookmarksName{},
        Next: hookdeck.String("temporibus"),
        OrderBy: &operations.GetBookmarksOrderBy{},
        Prev: hookdeck.String("laborum"),
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


## TriggerBookmark

Trigger a bookmark operation to store and replay a specific request.

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
    res, err := s.Bookmarks.TriggerBookmark(ctx, operations.TriggerBookmarkRequest{
        RequestBody: operations.TriggerBookmarkRequestBody{
            Target: operations.TriggerBookmarkRequestBodyTargetHTTP.ToPointer(),
        },
        ID: "ffe78f09-7b00-474f-9547-1b5e6e13b99d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventArray != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.TriggerBookmarkRequest](../../models/operations/triggerbookmarkrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.TriggerBookmarkResponse](../../models/operations/triggerbookmarkresponse.md), error**


## UpdateBookmark

Update an existing bookmark information.

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
    res, err := s.Bookmarks.UpdateBookmark(ctx, operations.UpdateBookmarkRequest{
        RequestBody: operations.UpdateBookmarkRequestBody{
            EventDataID: hookdeck.String("modi"),
            Label: hookdeck.String("praesentium"),
            Name: hookdeck.String("Grady Botsford"),
            WebhookID: hookdeck.String("veritatis"),
        },
        ID: "e450ad2a-bd44-4269-802d-502a94bb4f63",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Bookmark != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateBookmarkRequest](../../models/operations/updatebookmarkrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateBookmarkResponse](../../models/operations/updatebookmarkresponse.md), error**


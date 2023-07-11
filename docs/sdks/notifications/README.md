# Notifications

## Overview

Notifications let your team receive alerts anytime an issue changes.

### Available Operations

* [AddCustomDomain](#addcustomdomain) - Add a custom domain to the workspace
* [DeleteCustomDomain](#deletecustomdomain) - Removes a custom domain from the workspace
* [ListCustomDomains](#listcustomdomains) - List all custom domains and their verification statuses for the workspace
* [ToggleWebhookNotifications](#togglewebhooknotifications) - Toggle webhook notifications for the workspace

## AddCustomDomain

Add a custom domain to the workspace

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.Notifications.AddCustomDomain(ctx, operations.AddCustomDomainRequest{
        AddCustomHostname: shared.AddCustomHostname{
            Hostname: "flowery-jug.info",
        },
        TeamID: "sequi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddCustomHostname != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AddCustomDomainRequest](../../models/operations/addcustomdomainrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.AddCustomDomainResponse](../../models/operations/addcustomdomainresponse.md), error**


## DeleteCustomDomain

Removes a custom domain from the workspace

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
    res, err := s.Notifications.DeleteCustomDomain(ctx, operations.DeleteCustomDomainRequest{
        DomainID: "dignissimos",
        TeamID: "neque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCustomDomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteCustomDomainRequest](../../models/operations/deletecustomdomainrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteCustomDomainResponse](../../models/operations/deletecustomdomainresponse.md), error**


## ListCustomDomains

List all custom domains and their verification statuses for the workspace

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
    res, err := s.Notifications.ListCustomDomains(ctx, operations.ListCustomDomainsRequest{
        TeamID: "quo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCustomDomainSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListCustomDomainsRequest](../../models/operations/listcustomdomainsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListCustomDomainsResponse](../../models/operations/listcustomdomainsresponse.md), error**


## ToggleWebhookNotifications

Toggle webhook notifications for the workspace

### Example Usage

```go
package main

import(
	"context"
	"log"
	"hookdeck"
	"hookdeck/pkg/models/operations"
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
    res, err := s.Notifications.ToggleWebhookNotifications(ctx, operations.ToggleWebhookNotificationsRequestBody{
        Enabled: hookdeck.Bool(false),
        SourceID: hookdeck.String("deleniti"),
        Topics: []shared.TopicsValue{
            shared.TopicsValueIssueUpdated,
            shared.TopicsValueIssueOpened,
            shared.TopicsValueEventSuccessful,
            shared.TopicsValueIssueUpdated,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ToggleWebhookNotifications != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ToggleWebhookNotificationsRequestBody](../../models/operations/togglewebhooknotificationsrequestbody.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.ToggleWebhookNotificationsResponse](../../models/operations/togglewebhooknotificationsresponse.md), error**


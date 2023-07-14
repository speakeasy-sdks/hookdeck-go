# WebhookNotifications

### Available Operations

* [Toggle](#toggle) - Toggle webhook notifications for the workspace

## Toggle

Toggle webhook notifications for the workspace

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
    res, err := s.WebhookNotifications.Toggle(ctx, operations.ToggleWebhookNotificationsRequestBody{
        Enabled: hookdeck.Bool(false),
        SourceID: hookdeck.String("impedit"),
        Topics: []shared.TopicsValue{
            shared.TopicsValueEventSuccessful,
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


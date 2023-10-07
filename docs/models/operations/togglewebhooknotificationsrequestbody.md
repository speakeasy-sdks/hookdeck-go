# ToggleWebhookNotificationsRequestBody


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AdditionalProperties`                                     | map[string]*interface{}*                                   | :heavy_minus_sign:                                         | N/A                                                        |
| `Enabled`                                                  | **bool*                                                    | :heavy_minus_sign:                                         | Enable or disable webhook notifications on the workspace   |
| `SourceID`                                                 | **string*                                                  | :heavy_minus_sign:                                         | The Hookdeck Source to send the webhook to                 |
| `Topics`                                                   | [][shared.TopicsValue](../../models/shared/topicsvalue.md) | :heavy_minus_sign:                                         | List of topics to send notifications for                   |
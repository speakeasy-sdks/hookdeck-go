# IssueTriggerChannels

Notification channels object for the specific channel type


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AdditionalProperties`                                                       | map[string]*interface{}*                                                     | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Email`                                                                      | map[string]*interface{}*                                                     | :heavy_minus_sign:                                                           | Email channel for an issue trigger                                           |
| `Opsgenie`                                                                   | map[string]*interface{}*                                                     | :heavy_minus_sign:                                                           | Integration channel for an issue trigger                                     |
| `Slack`                                                                      | [*IssueTriggerSlackChannel](../../models/shared/issuetriggerslackchannel.md) | :heavy_minus_sign:                                                           | Slack channel for an issue trigger                                           |
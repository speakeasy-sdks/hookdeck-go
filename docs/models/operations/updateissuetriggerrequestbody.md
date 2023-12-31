# UpdateIssueTriggerRequestBody


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Channels`                                                                  | [*shared.IssueTriggerChannels](../../models/shared/issuetriggerchannels.md) | :heavy_minus_sign:                                                          | Notification channels object for the specific channel type                  |
| `Configs`                                                                   | *interface{}*                                                               | :heavy_minus_sign:                                                          | Configuration object for the specific issue type selected                   |
| `DisabledAt`                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                  | :heavy_minus_sign:                                                          | Date when the issue trigger was disabled                                    |
| `Name`                                                                      | **string*                                                                   | :heavy_minus_sign:                                                          | Optional unique name to use as reference when using the API                 |
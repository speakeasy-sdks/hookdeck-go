# UpsertIssueTriggerRequestBody


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Channels`                                                                 | [shared.IssueTriggerChannels](../../models/shared/issuetriggerchannels.md) | :heavy_check_mark:                                                         | Notification channels object for the specific channel type                 |
| `Configs`                                                                  | *interface{}*                                                              | :heavy_minus_sign:                                                         | Configuration object for the specific issue type selected                  |
| `Name`                                                                     | *string*                                                                   | :heavy_check_mark:                                                         | Required unique name to use as reference when using the API                |
| `Type`                                                                     | [shared.IssueType](../../models/shared/issuetype.md)                       | :heavy_check_mark:                                                         | Issue type                                                                 |
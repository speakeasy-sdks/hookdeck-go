# CreateIssueTriggerRequestBody


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Channels`                                                                     | [components.IssueTriggerChannels](../../models/shared/issuetriggerchannels.md) | :heavy_check_mark:                                                             | Notification channels object for the specific channel type                     |
| `Configs`                                                                      | [*operations.Configs](../../models/operations/configs.md)                      | :heavy_minus_sign:                                                             | Configuration object for the specific issue type selected                      |
| `Name`                                                                         | **string*                                                                      | :heavy_minus_sign:                                                             | Optional unique name to use as reference when using the API                    |
| `Type`                                                                         | [components.IssueType](../../models/shared/issuetype.md)                       | :heavy_check_mark:                                                             | Issue type                                                                     |
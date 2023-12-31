# IssueTrigger

A single issue trigger


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Channels`                                                           | [*IssueTriggerChannels](../../models/shared/issuetriggerchannels.md) | :heavy_minus_sign:                                                   | Notification channels object for the specific channel type           |
| `Configs`                                                            | *interface{}*                                                        | :heavy_check_mark:                                                   | Configuration object for the specific issue type selected            |
| `CreatedAt`                                                          | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | ISO timestamp for when the issue trigger was created                 |
| `DeletedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | ISO timestamp for when the issue trigger was deleted                 |
| `DisabledAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | ISO timestamp for when the issue trigger was disabled                |
| `ID`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | ID of the issue trigger                                              |
| `Name`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | Optional unique name to use as reference when using the API          |
| `TeamID`                                                             | **string*                                                            | :heavy_minus_sign:                                                   | ID of the workspace                                                  |
| `Type`                                                               | [IssueType](../../models/shared/issuetype.md)                        | :heavy_check_mark:                                                   | Issue type                                                           |
| `UpdatedAt`                                                          | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | ISO timestamp for when the issue trigger was last updated            |
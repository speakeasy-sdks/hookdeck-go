# DeliveryIssue

Delivery issue


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `AggregationKeys`                                                                   | [DeliveryIssueAggregationKeys](../../models/shared/deliveryissueaggregationkeys.md) | :heavy_check_mark:                                                                  | Keys used as the aggregation keys a 'delivery' type issue                           |                                                                                     |
| `AutoResolvedAt`                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                          | :heavy_minus_sign:                                                                  | N/A                                                                                 |                                                                                     |
| `CreatedAt`                                                                         | *string*                                                                            | :heavy_check_mark:                                                                  | ISO timestamp for when the issue was created                                        |                                                                                     |
| `DismissedAt`                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                          | :heavy_minus_sign:                                                                  | ISO timestamp for when the issue was dismissed                                      |                                                                                     |
| `FirstSeenAt`                                                                       | [time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_check_mark:                                                                  | ISO timestamp for when the issue was first opened                                   |                                                                                     |
| `ID`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Issue ID                                                                            | iss_YXKv5OdJXCiVwkPhGy                                                              |
| `LastSeenAt`                                                                        | [time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_check_mark:                                                                  | ISO timestamp for when the issue last occured                                       |                                                                                     |
| `LastUpdatedBy`                                                                     | **string*                                                                           | :heavy_minus_sign:                                                                  | ID of the team member who last updated the issue status                             |                                                                                     |
| `MergedWith`                                                                        | **string*                                                                           | :heavy_minus_sign:                                                                  | N/A                                                                                 |                                                                                     |
| `OpenedAt`                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_check_mark:                                                                  | ISO timestamp for when the issue was last opened                                    |                                                                                     |
| `Reference`                                                                         | [DeliveryIssueReference](../../models/shared/deliveryissuereference.md)             | :heavy_check_mark:                                                                  | Reference to the event and attempt an issue is being created for.                   |                                                                                     |
| `Status`                                                                            | [IssueStatus](../../models/shared/issuestatus.md)                                   | :heavy_check_mark:                                                                  | Issue status                                                                        |                                                                                     |
| `TeamID`                                                                            | *string*                                                                            | :heavy_check_mark:                                                                  | ID of the workspace                                                                 |                                                                                     |
| `Type`                                                                              | [DeliveryIssueType](../../models/shared/deliveryissuetype.md)                       | :heavy_check_mark:                                                                  | N/A                                                                                 |                                                                                     |
| `UpdatedAt`                                                                         | *string*                                                                            | :heavy_check_mark:                                                                  | ISO timestamp for when the issue was last updated                                   |                                                                                     |
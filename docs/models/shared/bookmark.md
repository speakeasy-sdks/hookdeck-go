# Bookmark

A single bookmark


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Alias`                                                  | **string*                                                | :heavy_minus_sign:                                       | Alternate alias for the bookmark                         |
| `CreatedAt`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | Date the bookmark was created                            |
| `Data`                                                   | [*ShortEventData](../../models/shared/shorteventdata.md) | :heavy_minus_sign:                                       | N/A                                                      |
| `EventDataID`                                            | *string*                                                 | :heavy_check_mark:                                       | ID of the bookmarked event data                          |
| `ID`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the bookmark                                       |
| `Label`                                                  | *string*                                                 | :heavy_check_mark:                                       | Descriptive name of the bookmark                         |
| `LastUsedAt`                                             | [*time.Time](https://pkg.go.dev/time#Time)               | :heavy_minus_sign:                                       | Date the bookmark was last manually triggered            |
| `TeamID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the workspace                                      |
| `UpdatedAt`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | Date the bookmark was last updated                       |
| `WebhookID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the associated connection                          |
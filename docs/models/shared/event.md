# Event

A single event


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Attempts`                                        | *int64*                                           | :heavy_check_mark:                                | Number of delivery attempts made                  |
| `CliID`                                           | **string*                                         | :heavy_minus_sign:                                | ID of the CLI the event is sent to                |
| `CreatedAt`                                       | [time.Time](https://pkg.go.dev/time#Time)         | :heavy_check_mark:                                | Date the event was created                        |
| `CreatedAtID`                                     | **string*                                         | :heavy_minus_sign:                                | N/A                                               |
| `Data`                                            | [*EventData](../../models/shared/eventdata.md)    | :heavy_minus_sign:                                | N/A                                               |
| `DestinationID`                                   | *string*                                          | :heavy_check_mark:                                | ID of the associated destination                  |
| `EventDataID`                                     | *string*                                          | :heavy_check_mark:                                | ID of the request data                            |
| `ID`                                              | *string*                                          | :heavy_check_mark:                                | ID of the event                                   |
| `LastAttemptAt`                                   | [*time.Time](https://pkg.go.dev/time#Time)        | :heavy_minus_sign:                                | Date of the most recently attempted retry         |
| `LastAttemptAtID`                                 | **string*                                         | :heavy_minus_sign:                                | N/A                                               |
| `NextAttemptAt`                                   | [*time.Time](https://pkg.go.dev/time#Time)        | :heavy_minus_sign:                                | Date of the next scheduled retry                  |
| `RequestID`                                       | *string*                                          | :heavy_check_mark:                                | ID of the request that created the event          |
| `ResponseStatus`                                  | **int64*                                          | :heavy_minus_sign:                                | Event status                                      |
| `SourceID`                                        | *string*                                          | :heavy_check_mark:                                | ID of the associated source                       |
| `Status`                                          | [EventStatus](../../models/shared/eventstatus.md) | :heavy_check_mark:                                | N/A                                               |
| `SuccessfulAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)        | :heavy_minus_sign:                                | Date of the latest successful attempt             |
| `TeamID`                                          | *string*                                          | :heavy_check_mark:                                | ID of the workspace                               |
| `UpdatedAt`                                       | [time.Time](https://pkg.go.dev/time#Time)         | :heavy_check_mark:                                | Date the event was last updated                   |
| `WebhookID`                                       | *string*                                          | :heavy_check_mark:                                | ID of the associated connection                   |
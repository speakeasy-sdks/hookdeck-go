# EventAttempt

A single attempt


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ArchivedAt`                                                                               | **string*                                                                                  | :heavy_minus_sign:                                                                         | Date the attempt was archived                                                              |
| `AttemptNumber`                                                                            | **int64*                                                                                   | :heavy_minus_sign:                                                                         | Sequential number of attempts (up to and including this one) made for the associated event |
| `Body`                                                                                     | [*EventAttemptBody](../../models/shared/eventattemptbody.md)                               | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `BulkRetryID`                                                                              | **string*                                                                                  | :heavy_minus_sign:                                                                         | ID of associated bulk retry                                                                |
| `CreatedAt`                                                                                | [time.Time](https://pkg.go.dev/time#Time)                                                  | :heavy_check_mark:                                                                         | Date the attempt was created                                                               |
| `DeliveredAt`                                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_minus_sign:                                                                         | Date the attempt was delivered                                                             |
| `DeliveryLatency`                                                                          | **int64*                                                                                   | :heavy_minus_sign:                                                                         | Time elapsed between attempt initiation and final delivery (in ms)                         |
| `DestinationID`                                                                            | **string*                                                                                  | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `ErrorCode`                                                                                | [*AttemptErrorCode](../../models/shared/attempterrorcode.md)                               | :heavy_minus_sign:                                                                         | Attempt could not complete because of an error                                             |
| `EventID`                                                                                  | *string*                                                                                   | :heavy_check_mark:                                                                         | Event ID                                                                                   |
| `HTTPMethod`                                                                               | [*EventAttemptHTTPMethod](../../models/shared/eventattempthttpmethod.md)                   | :heavy_minus_sign:                                                                         | HTTP method used to deliver the attempt                                                    |
| `ID`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | Attempt ID                                                                                 |
| `RequestedURL`                                                                             | **string*                                                                                  | :heavy_minus_sign:                                                                         | URL of the destination where delivery was attempted                                        |
| `RespondedAt`                                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_minus_sign:                                                                         | Date the destination responded to this attempt                                             |
| `ResponseLatency`                                                                          | **int64*                                                                                   | :heavy_minus_sign:                                                                         | Time elapsed between attempt initiation and a response from the destination (in ms)        |
| `ResponseStatus`                                                                           | **int64*                                                                                   | :heavy_minus_sign:                                                                         | Attempt's HTTP response code                                                               |
| `State`                                                                                    | [*AttemptState](../../models/shared/attemptstate.md)                                       | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Status`                                                                                   | [AttemptStatus](../../models/shared/attemptstatus.md)                                      | :heavy_check_mark:                                                                         | Attempt status                                                                             |
| `SuccessfulAt`                                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_minus_sign:                                                                         | Date the attempt was successful                                                            |
| `TeamID`                                                                                   | *string*                                                                                   | :heavy_check_mark:                                                                         | Team ID                                                                                    |
| `Trigger`                                                                                  | [*AttemptTrigger](../../models/shared/attempttrigger.md)                                   | :heavy_minus_sign:                                                                         | How the attempt was triggered                                                              |
| `UpdatedAt`                                                                                | [time.Time](https://pkg.go.dev/time#Time)                                                  | :heavy_check_mark:                                                                         | Date the attempt was last updated                                                          |
# UpsertDestinationRequestBody


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `AuthMethod`                                                                                                           | *interface{}*                                                                                                          | :heavy_minus_sign:                                                                                                     | Config for the destination's auth method                                                                               |
| `CliPath`                                                                                                              | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | Path for the CLI destination                                                                                           |
| `HTTPMethod`                                                                                                           | [*shared.DestinationHTTPMethod](../../models/shared/destinationhttpmethod.md)                                          | :heavy_minus_sign:                                                                                                     | HTTP method used on requests sent to the destination, overrides the method used on requests sent to the source.        |
| `Name`                                                                                                                 | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Name for the destination                                                                                               |
| `PathForwardingDisabled`                                                                                               | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | N/A                                                                                                                    |
| `RateLimit`                                                                                                            | *interface{}*                                                                                                          | :heavy_minus_sign:                                                                                                     | Limit event attempts to receive per period                                                                             |
| `RateLimitPeriod`                                                                                                      | [*UpsertDestinationRequestBodyRateLimitPeriod](../../models/operations/upsertdestinationrequestbodyratelimitperiod.md) | :heavy_minus_sign:                                                                                                     | Period to rate limit attempts                                                                                          |
| `URL`                                                                                                                  | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | Endpoint of the destination                                                                                            |
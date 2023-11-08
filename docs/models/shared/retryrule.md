# RetryRule


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Count`                                                          | **int64*                                                         | :heavy_minus_sign:                                               | Maximum number of retries to attempt                             |
| `Interval`                                                       | **int64*                                                         | :heavy_minus_sign:                                               | Time in MS between each retry                                    |
| `Strategy`                                                       | [components.RetryStrategy](../../models/shared/retrystrategy.md) | :heavy_check_mark:                                               | Algorithm to use when calculating delay between retries          |
| `Type`                                                           | [components.RetryRuleType](../../models/shared/retryruletype.md) | :heavy_check_mark:                                               | A retry rule must be of type `retry`                             |
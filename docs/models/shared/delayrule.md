# DelayRule


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AdditionalProperties`                                | map[string]*interface{}*                              | :heavy_minus_sign:                                    | N/A                                                   |
| `Delay`                                               | *int64*                                               | :heavy_check_mark:                                    | Delay to introduce in MS                              |
| `Type`                                                | [DelayRuleType](../../models/shared/delayruletype.md) | :heavy_check_mark:                                    | A delay rule must be of type `delay`                  |
# AlertRule


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `AdditionalProperties`                                | map[string]*interface{}*                              | :heavy_minus_sign:                                    | N/A                                                   |
| `Strategy`                                            | [AlertStrategy](../../models/shared/alertstrategy.md) | :heavy_check_mark:                                    | Alert strategy to use                                 |
| `Type`                                                | [AlertRuleType](../../models/shared/alertruletype.md) | :heavy_check_mark:                                    | An alert rule must be of type `alert`                 |
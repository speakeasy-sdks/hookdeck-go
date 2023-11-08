# CreateConnectionRuleset

Ruleset input object


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `IsTeamDefault`                                  | **bool*                                          | :heavy_minus_sign:                               | N/A                                              |
| `Name`                                           | *string*                                         | :heavy_check_mark:                               | Name for the ruleset                             |
| `Rules`                                          | [][components.Rule](../../models/shared/rule.md) | :heavy_minus_sign:                               | Array of rules to apply                          |
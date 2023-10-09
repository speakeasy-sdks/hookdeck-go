# UpdateRulesetRequestBody


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `AdditionalProperties`                       | map[string]*interface{}*                     | :heavy_minus_sign:                           | N/A                                          |
| `IsTeamDefault`                              | **bool*                                      | :heavy_minus_sign:                           | N/A                                          |
| `Name`                                       | **string*                                    | :heavy_minus_sign:                           | Name for the ruleset                         |
| `Rules`                                      | [][shared.Rule](../../models/shared/rule.md) | :heavy_minus_sign:                           | Array of rules to apply                      |
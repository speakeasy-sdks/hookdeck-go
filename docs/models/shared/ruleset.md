# Ruleset

Associated [Ruleset](#ruleset-object) object


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `ArchivedAt`                               | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | Date the ruleset was archived              |
| `CreatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | Date the ruleset was created               |
| `ID`                                       | *string*                                   | :heavy_check_mark:                         | ID of the ruleset                          |
| `IsTeamDefault`                            | *bool*                                     | :heavy_check_mark:                         | Default ruleset of Workspace               |
| `Name`                                     | *string*                                   | :heavy_check_mark:                         | A unique name for the ruleset              |
| `Rules`                                    | [][Rule](../../models/shared/rule.md)      | :heavy_check_mark:                         | Array of rules to apply                    |
| `TeamID`                                   | *string*                                   | :heavy_check_mark:                         | ID of the workspace                        |
| `UpdatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | Date the ruleset was last updated          |
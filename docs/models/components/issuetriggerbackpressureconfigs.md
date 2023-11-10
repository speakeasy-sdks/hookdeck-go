# IssueTriggerBackpressureConfigs

Configurations for a 'Backpressure' issue trigger


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Delay`                                                                                                    | *int64*                                                                                                    | :heavy_check_mark:                                                                                         | The minimum delay (backpressure) to open the issue for min of 1 minute (60000) and max of 1 day (86400000) |
| `Destinations`                                                                                             | [components.Destinations](../../models/components/destinations.md)                                         | :heavy_check_mark:                                                                                         | A pattern to match on the destination name or array of destination IDs. Use `*` as wildcard.               |
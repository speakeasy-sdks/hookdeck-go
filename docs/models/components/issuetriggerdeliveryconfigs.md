# IssueTriggerDeliveryConfigs

Configurations for a 'delivery' issue trigger


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Connections`                                                                              | [components.Connections](../../models/components/connections.md)                           | :heavy_check_mark:                                                                         | A pattern to match on the connection name or array of connection IDs. Use `*` as wildcard. |
| `Strategy`                                                                                 | [components.IssueTriggerStrategy](../../models/components/issuetriggerstrategy.md)         | :heavy_check_mark:                                                                         | The strategy uses to open the issue                                                        |
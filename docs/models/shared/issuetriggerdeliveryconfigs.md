# IssueTriggerDeliveryConfigs

Configurations for a 'delivery' issue trigger


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `Connections`                                                                                           | [IssueTriggerDeliveryConfigsConnections](../../models/shared/issuetriggerdeliveryconfigsconnections.md) | :heavy_check_mark:                                                                                      | A pattern to match on the connection name or array of connection IDs. Use `*` as wildcard.              |
| `Strategy`                                                                                              | [IssueTriggerStrategy](../../models/shared/issuetriggerstrategy.md)                                     | :heavy_check_mark:                                                                                      | The strategy uses to open the issue                                                                     |
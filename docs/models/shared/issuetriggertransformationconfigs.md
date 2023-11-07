# IssueTriggerTransformationConfigs

Configurations for a 'Transformation' issue trigger


## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `LogLevel`                                                                                           | [components.TransformationExecutionLogLevel](../../models/shared/transformationexecutionloglevel.md) | :heavy_check_mark:                                                                                   | The minimum log level to open the issue on                                                           |
| `Transformations`                                                                                    | [components.Transformations](../../models/shared/transformations.md)                                 | :heavy_check_mark:                                                                                   | A pattern to match on the transformation name or array of transformation IDs. Use `*` as wildcard.   |
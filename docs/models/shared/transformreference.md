# TransformReference


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `TransformationID`                                                                            | *string*                                                                                      | :heavy_check_mark:                                                                            | ID of the attached transformation object. Optional input, always set once the rule is defined |
| `Type`                                                                                        | [TransformReferenceType](../../models/shared/transformreferencetype.md)                       | :heavy_check_mark:                                                                            | A transformation rule must be of type `transformation`                                        |
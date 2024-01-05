# TransformFullTransformation

You can optionally define a new transformation while creating a transform rule


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `Code`                                                                                        | *string*                                                                                      | :heavy_check_mark:                                                                            | A string representation of your JavaScript (ES6) code to run                                  |
| `Env`                                                                                         | map[string]*string*                                                                           | :heavy_minus_sign:                                                                            | A key-value object of environment variables to encrypt and expose to your transformation code |
| `Name`                                                                                        | *string*                                                                                      | :heavy_check_mark:                                                                            | The unique name of the transformation                                                         |
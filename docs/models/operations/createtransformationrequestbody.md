# CreateTransformationRequestBody


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Code`                                                             | *string*                                                           | :heavy_check_mark:                                                 | JavaScript code to be executed as string                           |
| `Env`                                                              | map[string]*interface{}*                                           | :heavy_minus_sign:                                                 | Key-value environment variables to be passed to the transformation |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | A unique, human-friendly name for the transformation               |
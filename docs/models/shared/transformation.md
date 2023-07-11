# Transformation

A single transformation


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Code`                                                             | *string*                                                           | :heavy_check_mark:                                                 | JavaScript code to be executed                                     |
| `CreatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | Date the transformation was created                                |
| `EncryptedEnv`                                                     | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `Env`                                                              | map[string]*string*                                                | :heavy_minus_sign:                                                 | Key-value environment variables to be passed to the transformation |
| `ID`                                                               | *string*                                                           | :heavy_check_mark:                                                 | ID of the transformation                                           |
| `Iv`                                                               | **string*                                                          | :heavy_minus_sign:                                                 | N/A                                                                |
| `Name`                                                             | *string*                                                           | :heavy_check_mark:                                                 | A unique, human-friendly name for the transformation               |
| `TeamID`                                                           | *string*                                                           | :heavy_check_mark:                                                 | ID of the workspace                                                |
| `UpdatedAt`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | Date the transformation was last updated                           |
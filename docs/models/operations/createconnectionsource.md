# CreateConnectionSource

Source input object


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `AllowedHTTPMethods`                                                                       | [][components.SourceAllowedHTTPMethod](../../models/components/sourceallowedhttpmethod.md) | :heavy_minus_sign:                                                                         | List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.                        |
| `CustomResponse`                                                                           | [*components.SourceCustomResponse](../../models/components/sourcecustomresponse.md)        | :heavy_minus_sign:                                                                         | Custom response object                                                                     |
| `Name`                                                                                     | *string*                                                                                   | :heavy_check_mark:                                                                         | A unique name for the source                                                               |
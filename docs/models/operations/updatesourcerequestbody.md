# UpdateSourceRequestBody


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `AllowedHTTPMethods`                                                               | [][shared.SourceAllowedHTTPMethod](../../models/shared/sourceallowedhttpmethod.md) | :heavy_minus_sign:                                                                 | List of allowed HTTP methods. Defaults to PUT, POST, PATCH, DELETE.                |
| `CustomResponse`                                                                   | [*shared.SourceCustomResponse](../../models/shared/sourcecustomresponse.md)        | :heavy_minus_sign:                                                                 | Custom response object                                                             |
| `Name`                                                                             | **string*                                                                          | :heavy_minus_sign:                                                                 | A unique name for the source                                                       |
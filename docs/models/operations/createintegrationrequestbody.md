# CreateIntegrationRequestBody


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Configs`                                                                     | *interface{}*                                                                 | :heavy_minus_sign:                                                            | Decrypted Key/Value object of the associated configuration for that provider  |
| `Features`                                                                    | [][components.IntegrationFeature](../../models/shared/integrationfeature.md)  | :heavy_minus_sign:                                                            | List of features to enable (see features list above)                          |
| `Label`                                                                       | **string*                                                                     | :heavy_minus_sign:                                                            | Label of the integration                                                      |
| `Provider`                                                                    | [*components.IntegrationProvider](../../models/shared/integrationprovider.md) | :heavy_minus_sign:                                                            | The provider name                                                             |
# DestinationAuthMethodCustomSignatureConfig

Custom signature config for the destination's auth method


## Fields

| Field                                                                                           | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `AdditionalProperties`                                                                          | map[string]*interface{}*                                                                        | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `Key`                                                                                           | *string*                                                                                        | :heavy_check_mark:                                                                              | Key for the custom signature auth                                                               |
| `SigningSecret`                                                                                 | **string*                                                                                       | :heavy_minus_sign:                                                                              | Signing secret for the custom signature auth. If left empty a secret will be generated for you. |
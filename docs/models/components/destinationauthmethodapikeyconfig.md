# DestinationAuthMethodAPIKeyConfig

API key config for the destination's auth method


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `APIKey`                                                            | *string*                                                            | :heavy_check_mark:                                                  | API key for the API key auth                                        |
| `Key`                                                               | *string*                                                            | :heavy_check_mark:                                                  | Key for the API key auth                                            |
| `To`                                                                | [*components.To](../../models/components/to.md)                     | :heavy_minus_sign:                                                  | Whether the API key should be sent as a header or a query parameter |
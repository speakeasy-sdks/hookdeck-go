# DeleteIntegrationResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `DeletedIntegration`                                                    | [*shared.DeletedIntegration](../../models/shared/deletedintegration.md) | :heavy_minus_sign:                                                      | An object with deleted integration id                                   |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |
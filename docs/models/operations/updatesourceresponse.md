# UpdateSourceResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `APIErrorResponse`                                                  | [*shared.APIErrorResponse](../../models/shared/apierrorresponse.md) | :heavy_minus_sign:                                                  | Bad Request                                                         |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `Source`                                                            | [*shared.Source](../../models/shared/source.md)                     | :heavy_minus_sign:                                                  | A single source                                                     |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |
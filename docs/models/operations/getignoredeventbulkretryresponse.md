# GetIgnoredEventBulkRetryResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `APIErrorResponse`                                                  | [*shared.APIErrorResponse](../../models/shared/apierrorresponse.md) | :heavy_minus_sign:                                                  | Not Found                                                           |
| `BatchOperation`                                                    | [*shared.BatchOperation](../../models/shared/batchoperation.md)     | :heavy_minus_sign:                                                  | A single ignored events bulk retry                                  |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |
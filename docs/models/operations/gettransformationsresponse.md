# GetTransformationsResponse


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `APIErrorResponse`                                                                            | [*shared.APIErrorResponse](../../models/shared/apierrorresponse.md)                           | :heavy_minus_sign:                                                                            | Bad Request                                                                                   |
| `ContentType`                                                                                 | *string*                                                                                      | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `StatusCode`                                                                                  | *int*                                                                                         | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `RawResponse`                                                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)                                        | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `TransformationPaginatedResult`                                                               | [*shared.TransformationPaginatedResult](../../models/shared/transformationpaginatedresult.md) | :heavy_minus_sign:                                                                            | List of transformations                                                                       |
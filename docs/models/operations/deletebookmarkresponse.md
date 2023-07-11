# DeleteBookmarkResponse


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `APIErrorResponse`                                                                | [*shared.APIErrorResponse](../../models/shared/apierrorresponse.md)               | :heavy_minus_sign:                                                                | Not Found                                                                         |
| `ContentType`                                                                     | *string*                                                                          | :heavy_check_mark:                                                                | N/A                                                                               |
| `DeletedBookmarkResponse`                                                         | [*shared.DeletedBookmarkResponse](../../models/shared/deletedbookmarkresponse.md) | :heavy_minus_sign:                                                                | An object with deleted bookmark's id                                              |
| `StatusCode`                                                                      | *int*                                                                             | :heavy_check_mark:                                                                | N/A                                                                               |
| `RawResponse`                                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)                            | :heavy_minus_sign:                                                                | N/A                                                                               |
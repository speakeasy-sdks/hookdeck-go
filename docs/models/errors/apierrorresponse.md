# APIErrorResponse

Error response model


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |
| `Code`                                                 | *string*                                               | :heavy_check_mark:                                     | Error code                                             |
| `Data`                                                 | **APIErrorResponseData*                                | :heavy_minus_sign:                                     | N/A                                                    |
| `Message`                                              | *string*                                               | :heavy_check_mark:                                     | Error description                                      |
| `Status`                                               | *float32*                                              | :heavy_check_mark:                                     | Status code                                            |
# EnableIssueTriggerResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `APIErrorResponse`                                                  | [*shared.APIErrorResponse](../../models/shared/apierrorresponse.md) | :heavy_minus_sign:                                                  | Not Found                                                           |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `IssueTrigger`                                                      | [*shared.IssueTrigger](../../models/shared/issuetrigger.md)         | :heavy_minus_sign:                                                  | A single issue trigger                                              |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |
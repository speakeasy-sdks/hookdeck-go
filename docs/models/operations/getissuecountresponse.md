# GetIssueCountResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `APIErrorResponse`                                                  | [*shared.APIErrorResponse](../../models/shared/apierrorresponse.md) | :heavy_minus_sign:                                                  | Unprocessable Entity                                                |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `IssueCount`                                                        | [*shared.IssueCount](../../models/shared/issuecount.md)             | :heavy_minus_sign:                                                  | Issue count                                                         |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |
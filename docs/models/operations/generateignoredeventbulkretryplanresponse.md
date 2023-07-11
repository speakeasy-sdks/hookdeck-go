# GenerateIgnoredEventBulkRetryPlanResponse


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `APIErrorResponse`                                                      | [*shared.APIErrorResponse](../../models/shared/apierrorresponse.md)     | :heavy_minus_sign:                                                      | Bad Request                                                             |
| `BatchOperationPlan`                                                    | [*shared.BatchOperationPlan](../../models/shared/batchoperationplan.md) | :heavy_minus_sign:                                                      | Ignored events bulk retry plan                                          |
| `ContentType`                                                           | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `StatusCode`                                                            | *int*                                                                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `RawResponse`                                                           | [*http.Response](https://pkg.go.dev/net/http#Response)                  | :heavy_minus_sign:                                                      | N/A                                                                     |
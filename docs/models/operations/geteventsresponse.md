# GetEventsResponse


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ContentType`                                                               | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `EventPaginatedResult`                                                      | [*shared.EventPaginatedResult](../../models/shared/eventpaginatedresult.md) | :heavy_minus_sign:                                                          | List of events                                                              |
| `StatusCode`                                                                | *int*                                                                       | :heavy_check_mark:                                                          | N/A                                                                         |
| `RawResponse`                                                               | [*http.Response](https://pkg.go.dev/net/http#Response)                      | :heavy_minus_sign:                                                          | N/A                                                                         |
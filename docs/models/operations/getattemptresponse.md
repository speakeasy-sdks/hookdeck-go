# GetAttemptResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | HTTP response content type for this operation                       |
| `EventAttempt`                                                      | [*components.EventAttempt](../../models/components/eventattempt.md) | :heavy_minus_sign:                                                  | A single attempt                                                    |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | HTTP response status code for this operation                        |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_check_mark:                                                  | Raw HTTP response; suitable for custom response parsing             |
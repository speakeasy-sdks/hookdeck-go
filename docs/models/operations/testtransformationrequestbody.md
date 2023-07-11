# TestTransformationRequestBody


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `Code`                                                                                                   | **string*                                                                                                | :heavy_minus_sign:                                                                                       | JavaScript code to be executed                                                                           |
| `Env`                                                                                                    | [*TestTransformationRequestBodyEnv](../../models/operations/testtransformationrequestbodyenv.md)         | :heavy_minus_sign:                                                                                       | Key-value environment variables to be passed to the transformation                                       |
| `EventID`                                                                                                | **string*                                                                                                | :heavy_minus_sign:                                                                                       | N/A                                                                                                      |
| `Request`                                                                                                | [*TestTransformationRequestBodyRequest](../../models/operations/testtransformationrequestbodyrequest.md) | :heavy_minus_sign:                                                                                       | Request input to use for the transformation execution                                                    |
| `TransformationID`                                                                                       | **string*                                                                                                | :heavy_minus_sign:                                                                                       | Transformation ID                                                                                        |
| `WebhookID`                                                                                              | **string*                                                                                                | :heavy_minus_sign:                                                                                       | ID of the connection to use for the execution `context`                                                  |
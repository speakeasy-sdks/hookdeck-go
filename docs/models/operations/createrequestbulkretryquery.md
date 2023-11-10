# CreateRequestBulkRetryQuery

Filter properties for the events to be included in the bulk retry, use query parameters of [Requests](#requests)


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `Body`                                                                                                        | [*operations.CreateRequestBulkRetryBody](../../models/operations/createrequestbulkretrybody.md)               | :heavy_minus_sign:                                                                                            | URL Encoded string of the JSON to match to the data body                                                      |
| `BulkRetryID`                                                                                                 | [*operations.CreateRequestBulkRetryBulkRetryID](../../models/operations/createrequestbulkretrybulkretryid.md) | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `CreatedAt`                                                                                                   | [*operations.CreateRequestBulkRetryCreatedAt](../../models/operations/createrequestbulkretrycreatedat.md)     | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `EventsCount`                                                                                                 | [*operations.EventsCount](../../models/operations/eventscount.md)                                             | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `Headers`                                                                                                     | [*operations.CreateRequestBulkRetryHeaders](../../models/operations/createrequestbulkretryheaders.md)         | :heavy_minus_sign:                                                                                            | URL Encoded string of the JSON to match to the data headers                                                   |
| `ID`                                                                                                          | [*operations.CreateRequestBulkRetryID](../../models/operations/createrequestbulkretryid.md)                   | :heavy_minus_sign:                                                                                            | Filter by requests IDs                                                                                        |
| `IgnoredCount`                                                                                                | [*operations.IgnoredCount](../../models/operations/ignoredcount.md)                                           | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `IngestedAt`                                                                                                  | [*operations.IngestedAt](../../models/operations/ingestedat.md)                                               | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `ParsedQuery`                                                                                                 | [*operations.CreateRequestBulkRetryParsedQuery](../../models/operations/createrequestbulkretryparsedquery.md) | :heavy_minus_sign:                                                                                            | URL Encoded string of the JSON to match to the parsed query (JSON representation of the query)                |
| `Path`                                                                                                        | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | URL Encoded string of the string to match partially to the path                                               |
| `RejectionCause`                                                                                              | [*operations.RejectionCause](../../models/operations/rejectioncause.md)                                       | :heavy_minus_sign:                                                                                            | Filter by rejection cause                                                                                     |
| `SearchTerm`                                                                                                  | **string*                                                                                                     | :heavy_minus_sign:                                                                                            | URL Encoded string of the string to match partially to the body, headers, parsed_query or path                |
| `SourceID`                                                                                                    | [*operations.CreateRequestBulkRetrySourceID](../../models/operations/createrequestbulkretrysourceid.md)       | :heavy_minus_sign:                                                                                            | Filter by source IDs                                                                                          |
| `Status`                                                                                                      | [*operations.CreateRequestBulkRetryStatus](../../models/operations/createrequestbulkretrystatus.md)           | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `Verified`                                                                                                    | **bool*                                                                                                       | :heavy_minus_sign:                                                                                            | Filter by verification status                                                                                 |
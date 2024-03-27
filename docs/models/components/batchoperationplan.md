# BatchOperationPlan


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `EstimatedBatch`                                      | **int64*                                              | :heavy_minus_sign:                                    | Number of batches required to complete the bulk retry |
| `EstimatedCount`                                      | **int64*                                              | :heavy_minus_sign:                                    | Number of estimated events to be retried              |
| `Progress`                                            | **float32*                                            | :heavy_minus_sign:                                    | Progression of the batch operations, values 0 - 1     |
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TransformationExecutionPaginatedResult struct {
	Count      *int64                    `json:"count,omitempty"`
	Models     []TransformationExecution `json:"models,omitempty"`
	Pagination *SeekPagination           `json:"pagination,omitempty"`
}

func (o *TransformationExecutionPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *TransformationExecutionPaginatedResult) GetModels() []TransformationExecution {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *TransformationExecutionPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

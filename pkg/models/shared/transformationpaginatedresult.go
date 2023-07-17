// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TransformationPaginatedResult - List of transformations
type TransformationPaginatedResult struct {
	Count      *int64           `json:"count,omitempty"`
	Models     []Transformation `json:"models,omitempty"`
	Pagination *SeekPagination  `json:"pagination,omitempty"`
}

func (o *TransformationPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *TransformationPaginatedResult) GetModels() []Transformation {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *TransformationPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

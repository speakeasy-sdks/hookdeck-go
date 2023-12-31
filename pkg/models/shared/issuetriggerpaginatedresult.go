// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// IssueTriggerPaginatedResult - List of issue triggers
type IssueTriggerPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []IssueTrigger  `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

func (o *IssueTriggerPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *IssueTriggerPaginatedResult) GetModels() []IssueTrigger {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *IssueTriggerPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

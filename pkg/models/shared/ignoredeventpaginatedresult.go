// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// IgnoredEventPaginatedResult - List of ignored events
type IgnoredEventPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []IgnoredEvent  `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

func (o *IgnoredEventPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *IgnoredEventPaginatedResult) GetModels() []IgnoredEvent {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *IgnoredEventPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

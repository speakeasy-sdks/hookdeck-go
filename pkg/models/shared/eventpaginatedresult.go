// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EventPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []Event         `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

func (o *EventPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *EventPaginatedResult) GetModels() []Event {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *EventPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

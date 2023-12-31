// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EventAttemptPaginatedResult - List of attempts
type EventAttemptPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []EventAttempt  `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

func (o *EventAttemptPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *EventAttemptPaginatedResult) GetModels() []EventAttempt {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *EventAttemptPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

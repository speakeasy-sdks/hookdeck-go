// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// IgnoredEventPaginatedResult - List of ignored events
type IgnoredEventPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []IgnoredEvent  `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

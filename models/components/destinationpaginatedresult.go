// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type DestinationPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []Destination   `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

func (o *DestinationPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DestinationPaginatedResult) GetModels() []Destination {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *DestinationPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

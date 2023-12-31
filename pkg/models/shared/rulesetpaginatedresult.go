// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RulesetPaginatedResult - List of rulesets
type RulesetPaginatedResult struct {
	Count      *int64          `json:"count,omitempty"`
	Models     []Ruleset       `json:"models,omitempty"`
	Pagination *SeekPagination `json:"pagination,omitempty"`
}

func (o *RulesetPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *RulesetPaginatedResult) GetModels() []Ruleset {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *RulesetPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

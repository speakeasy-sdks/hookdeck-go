// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
)

type IssueWithDataPaginatedResult struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Count                *int64                 `json:"count,omitempty"`
	Models               []IssueWithData        `json:"models,omitempty"`
	Pagination           *SeekPagination        `json:"pagination,omitempty"`
}

func (i IssueWithDataPaginatedResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IssueWithDataPaginatedResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IssueWithDataPaginatedResult) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *IssueWithDataPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *IssueWithDataPaginatedResult) GetModels() []IssueWithData {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *IssueWithDataPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

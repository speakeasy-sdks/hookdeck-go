// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
)

type DestinationPaginatedResult struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Count                *int64                 `json:"count,omitempty"`
	Models               []Destination          `json:"models,omitempty"`
	Pagination           *SeekPagination        `json:"pagination,omitempty"`
}

func (d DestinationPaginatedResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPaginatedResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPaginatedResult) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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

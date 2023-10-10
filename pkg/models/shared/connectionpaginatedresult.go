// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
)

type ConnectionPaginatedResult struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Count                *int64                 `json:"count,omitempty"`
	Models               []Connection           `json:"models,omitempty"`
	Pagination           *SeekPagination        `json:"pagination,omitempty"`
}

func (c ConnectionPaginatedResult) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionPaginatedResult) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionPaginatedResult) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *ConnectionPaginatedResult) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ConnectionPaginatedResult) GetModels() []Connection {
	if o == nil {
		return nil
	}
	return o.Models
}

func (o *ConnectionPaginatedResult) GetPagination() *SeekPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

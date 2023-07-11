// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceAllowedHTTPMethod string

const (
	SourceAllowedHTTPMethodGet    SourceAllowedHTTPMethod = "GET"
	SourceAllowedHTTPMethodPost   SourceAllowedHTTPMethod = "POST"
	SourceAllowedHTTPMethodPut    SourceAllowedHTTPMethod = "PUT"
	SourceAllowedHTTPMethodPatch  SourceAllowedHTTPMethod = "PATCH"
	SourceAllowedHTTPMethodDelete SourceAllowedHTTPMethod = "DELETE"
)

func (e SourceAllowedHTTPMethod) ToPointer() *SourceAllowedHTTPMethod {
	return &e
}

func (e *SourceAllowedHTTPMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "PUT":
		fallthrough
	case "PATCH":
		fallthrough
	case "DELETE":
		*e = SourceAllowedHTTPMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAllowedHTTPMethod: %v", v)
	}
}
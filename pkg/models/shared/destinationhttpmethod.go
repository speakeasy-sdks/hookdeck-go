// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationHTTPMethod - HTTP method used on requests sent to the destination, overrides the method used on requests sent to the source.
type DestinationHTTPMethod string

const (
	DestinationHTTPMethodLessThanNilGreaterThan DestinationHTTPMethod = "<nil>"
	DestinationHTTPMethodGet                    DestinationHTTPMethod = "GET"
	DestinationHTTPMethodPost                   DestinationHTTPMethod = "POST"
	DestinationHTTPMethodPut                    DestinationHTTPMethod = "PUT"
	DestinationHTTPMethodPatch                  DestinationHTTPMethod = "PATCH"
	DestinationHTTPMethodDelete                 DestinationHTTPMethod = "DELETE"
)

func (e DestinationHTTPMethod) ToPointer() *DestinationHTTPMethod {
	return &e
}

func (e *DestinationHTTPMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "<nil>":
		fallthrough
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "PUT":
		fallthrough
	case "PATCH":
		fallthrough
	case "DELETE":
		*e = DestinationHTTPMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationHTTPMethod: %v", v)
	}
}
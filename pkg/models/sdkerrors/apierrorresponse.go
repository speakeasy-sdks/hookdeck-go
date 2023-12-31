// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type APIErrorResponseData struct {
	RawResponse *http.Response `json:"-"`
}

var _ error = &APIErrorResponseData{}

func (e *APIErrorResponseData) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// APIErrorResponse - Error response model
type APIErrorResponse struct {
	RawResponse *http.Response `json:"-"`
	// Error code
	Code string                `json:"code"`
	Data *APIErrorResponseData `json:"data,omitempty"`
	// Error description
	Message string `json:"message"`
	// Status code
	Status float32 `json:"status"`
}

var _ error = &APIErrorResponse{}

func (e *APIErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

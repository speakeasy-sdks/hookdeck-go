// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"net/http"
)

type GetEventBulkRetryRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetEventBulkRetryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetEventBulkRetryResponse struct {
	// A single events bulk retry
	BatchOperation *components.BatchOperation
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetEventBulkRetryResponse) GetBatchOperation() *components.BatchOperation {
	if o == nil {
		return nil
	}
	return o.BatchOperation
}

func (o *GetEventBulkRetryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEventBulkRetryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEventBulkRetryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

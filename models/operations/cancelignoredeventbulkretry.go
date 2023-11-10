// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type CancelIgnoredEventBulkRetryRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *CancelIgnoredEventBulkRetryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type CancelIgnoredEventBulkRetryResponse struct {
	// A single ignored events bulk retry
	BatchOperation *components.BatchOperation
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CancelIgnoredEventBulkRetryResponse) GetBatchOperation() *components.BatchOperation {
	if o == nil {
		return nil
	}
	return o.BatchOperation
}

func (o *CancelIgnoredEventBulkRetryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CancelIgnoredEventBulkRetryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CancelIgnoredEventBulkRetryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

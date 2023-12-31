// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
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
	BatchOperation *shared.BatchOperation
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetEventBulkRetryResponse) GetBatchOperation() *shared.BatchOperation {
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

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetIgnoredEventBulkRetryRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetIgnoredEventBulkRetryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetIgnoredEventBulkRetryResponse struct {
	// A single ignored events bulk retry
	BatchOperation *shared.BatchOperation
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetIgnoredEventBulkRetryResponse) GetBatchOperation() *shared.BatchOperation {
	if o == nil {
		return nil
	}
	return o.BatchOperation
}

func (o *GetIgnoredEventBulkRetryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIgnoredEventBulkRetryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIgnoredEventBulkRetryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

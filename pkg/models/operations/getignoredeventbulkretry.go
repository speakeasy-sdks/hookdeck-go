// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetIgnoredEventBulkRetryRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetIgnoredEventBulkRetryResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	// A single ignored events bulk retry
	BatchOperation *shared.BatchOperation
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

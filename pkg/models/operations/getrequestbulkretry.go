// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck/pkg/models/shared"
	"net/http"
)

type GetRequestBulkRetryRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetRequestBulkRetryResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	// A single requests bulk retry
	BatchOperation *shared.BatchOperation
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

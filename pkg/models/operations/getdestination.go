// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetDestinationRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetDestinationResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single destination
	Destination *shared.Destination
	StatusCode  int
	RawResponse *http.Response
}

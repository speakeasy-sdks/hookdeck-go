// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetAttemptRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAttemptResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single attempt
	EventAttempt *shared.EventAttempt
	StatusCode   int
	RawResponse  *http.Response
}

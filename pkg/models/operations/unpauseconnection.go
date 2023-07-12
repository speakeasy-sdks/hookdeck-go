// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type UnpauseConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UnpauseConnectionResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	// A single connection
	Connection  *shared.Connection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

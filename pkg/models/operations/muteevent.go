// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type MuteEventRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type MuteEventResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single event
	Event       *shared.Event
	StatusCode  int
	RawResponse *http.Response
}

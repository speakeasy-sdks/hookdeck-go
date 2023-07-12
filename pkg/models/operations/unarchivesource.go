// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type UnarchiveSourceRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UnarchiveSourceResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// A single source
	Source      *shared.Source
	StatusCode  int
	RawResponse *http.Response
}

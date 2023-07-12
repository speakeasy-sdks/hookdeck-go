// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type GetTransformationExecutionRequest struct {
	ExecutionID string `pathParam:"style=simple,explode=false,name=execution_id"`
	ID          string `pathParam:"style=simple,explode=false,name=id"`
}

type GetTransformationExecutionResponse struct {
	// Not Found
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	StatusCode       int
	RawResponse      *http.Response
	// A single transformation execution
	TransformationExecution *shared.TransformationExecution
}

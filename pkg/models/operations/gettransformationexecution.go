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

func (o *GetTransformationExecutionRequest) GetExecutionID() string {
	if o == nil {
		return ""
	}
	return o.ExecutionID
}

func (o *GetTransformationExecutionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTransformationExecutionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A single transformation execution
	TransformationExecution *shared.TransformationExecution
}

func (o *GetTransformationExecutionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransformationExecutionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransformationExecutionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransformationExecutionResponse) GetTransformationExecution() *shared.TransformationExecution {
	if o == nil {
		return nil
	}
	return o.TransformationExecution
}

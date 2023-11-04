// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
	"net/http"
)

// TestTransformationRequestBodyEnv - Key-value environment variables to be passed to the transformation
type TestTransformationRequestBodyEnv struct {
}

// TestTransformationRequestBodyRequestBody1 - Body of the request
type TestTransformationRequestBodyRequestBody1 struct {
}

type TestTransformationRequestBodyRequestBodyType string

const (
	TestTransformationRequestBodyRequestBodyTypeTestTransformationRequestBodyRequestBody1 TestTransformationRequestBodyRequestBodyType = "testTransformation_requestBody_request_body_1"
	TestTransformationRequestBodyRequestBodyTypeStr                                       TestTransformationRequestBodyRequestBodyType = "str"
)

type TestTransformationRequestBodyRequestBody struct {
	TestTransformationRequestBodyRequestBody1 *TestTransformationRequestBodyRequestBody1
	Str                                       *string

	Type TestTransformationRequestBodyRequestBodyType
}

func CreateTestTransformationRequestBodyRequestBodyTestTransformationRequestBodyRequestBody1(testTransformationRequestBodyRequestBody1 TestTransformationRequestBodyRequestBody1) TestTransformationRequestBodyRequestBody {
	typ := TestTransformationRequestBodyRequestBodyTypeTestTransformationRequestBodyRequestBody1

	return TestTransformationRequestBodyRequestBody{
		TestTransformationRequestBodyRequestBody1: &testTransformationRequestBodyRequestBody1,
		Type: typ,
	}
}

func CreateTestTransformationRequestBodyRequestBodyStr(str string) TestTransformationRequestBodyRequestBody {
	typ := TestTransformationRequestBodyRequestBodyTypeStr

	return TestTransformationRequestBodyRequestBody{
		Str:  &str,
		Type: typ,
	}
}

func (u *TestTransformationRequestBodyRequestBody) UnmarshalJSON(data []byte) error {

	testTransformationRequestBodyRequestBody1 := TestTransformationRequestBodyRequestBody1{}
	if err := utils.UnmarshalJSON(data, &testTransformationRequestBodyRequestBody1, "", true, true); err == nil {
		u.TestTransformationRequestBodyRequestBody1 = &testTransformationRequestBodyRequestBody1
		u.Type = TestTransformationRequestBodyRequestBodyTypeTestTransformationRequestBodyRequestBody1
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = TestTransformationRequestBodyRequestBodyTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TestTransformationRequestBodyRequestBody) MarshalJSON() ([]byte, error) {
	if u.TestTransformationRequestBodyRequestBody1 != nil {
		return utils.MarshalJSON(u.TestTransformationRequestBodyRequestBody1, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// TestTransformationRequestBodyRequestParsedQuery - JSON representation of the query params
type TestTransformationRequestBodyRequestParsedQuery struct {
}

// TestTransformationRequestBodyRequest - Request input to use for the transformation execution
type TestTransformationRequestBodyRequest struct {
	// Body of the request
	Body *TestTransformationRequestBodyRequestBody `json:"body,omitempty"`
	// Headers of the request
	Headers map[string]string `json:"headers"`
	// JSON representation of the query params
	ParsedQuery *TestTransformationRequestBodyRequestParsedQuery `json:"parsed_query,omitempty"`
	// Path of the request
	Path *string `json:"path,omitempty"`
	// String representation of the query params of the request
	Query *string `json:"query,omitempty"`
}

func (o *TestTransformationRequestBodyRequest) GetBody() *TestTransformationRequestBodyRequestBody {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *TestTransformationRequestBodyRequest) GetHeaders() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Headers
}

func (o *TestTransformationRequestBodyRequest) GetParsedQuery() *TestTransformationRequestBodyRequestParsedQuery {
	if o == nil {
		return nil
	}
	return o.ParsedQuery
}

func (o *TestTransformationRequestBodyRequest) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *TestTransformationRequestBodyRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type TestTransformationRequestBody struct {
	// JavaScript code to be executed
	Code *string `json:"code,omitempty"`
	// Key-value environment variables to be passed to the transformation
	Env     *TestTransformationRequestBodyEnv `json:"env,omitempty"`
	EventID *string                           `json:"event_id,omitempty"`
	// Request input to use for the transformation execution
	Request *TestTransformationRequestBodyRequest `json:"request,omitempty"`
	// Transformation ID
	TransformationID *string `json:"transformation_id,omitempty"`
	// ID of the connection to use for the execution `context`
	WebhookID *string `json:"webhook_id,omitempty"`
}

func (o *TestTransformationRequestBody) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *TestTransformationRequestBody) GetEnv() *TestTransformationRequestBodyEnv {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *TestTransformationRequestBody) GetEventID() *string {
	if o == nil {
		return nil
	}
	return o.EventID
}

func (o *TestTransformationRequestBody) GetRequest() *TestTransformationRequestBodyRequest {
	if o == nil {
		return nil
	}
	return o.Request
}

func (o *TestTransformationRequestBody) GetTransformationID() *string {
	if o == nil {
		return nil
	}
	return o.TransformationID
}

func (o *TestTransformationRequestBody) GetWebhookID() *string {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

type TestTransformationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Transformation run output
	TransformationExecutorOutput *shared.TransformationExecutorOutput
}

func (o *TestTransformationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TestTransformationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TestTransformationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TestTransformationResponse) GetTransformationExecutorOutput() *shared.TransformationExecutorOutput {
	if o == nil {
		return nil
	}
	return o.TransformationExecutorOutput
}

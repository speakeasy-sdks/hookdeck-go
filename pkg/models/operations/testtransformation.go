// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
	"net/http"
)

type TestTransformationRequestBodyRequestBodyType string

const (
	TestTransformationRequestBodyRequestBodyTypeMapOfany TestTransformationRequestBodyRequestBodyType = "mapOfany"
	TestTransformationRequestBodyRequestBodyTypeStr      TestTransformationRequestBodyRequestBodyType = "str"
)

type TestTransformationRequestBodyRequestBody struct {
	MapOfany map[string]interface{}
	Str      *string

	Type TestTransformationRequestBodyRequestBodyType
}

func CreateTestTransformationRequestBodyRequestBodyMapOfany(mapOfany map[string]interface{}) TestTransformationRequestBodyRequestBody {
	typ := TestTransformationRequestBodyRequestBodyTypeMapOfany

	return TestTransformationRequestBodyRequestBody{
		MapOfany: mapOfany,
		Type:     typ,
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

	mapOfany := map[string]interface{}{}
	if err := utils.UnmarshalJSON(data, &mapOfany, "", true, true); err == nil {
		u.MapOfany = mapOfany
		u.Type = TestTransformationRequestBodyRequestBodyTypeMapOfany
		return nil
	}

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = TestTransformationRequestBodyRequestBodyTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TestTransformationRequestBodyRequestBody) MarshalJSON() ([]byte, error) {
	if u.MapOfany != nil {
		return utils.MarshalJSON(u.MapOfany, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// TestTransformationRequestBodyRequest - Request input to use for the transformation execution
type TestTransformationRequestBodyRequest struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// Body of the request
	Body *TestTransformationRequestBodyRequestBody `json:"body,omitempty"`
	// Headers of the request
	Headers map[string]string `json:"headers"`
	// JSON representation of the query params
	ParsedQuery map[string]interface{} `json:"parsed_query,omitempty"`
	// Path of the request
	Path *string `json:"path,omitempty"`
	// String representation of the query params of the request
	Query *string `json:"query,omitempty"`
}

func (t TestTransformationRequestBodyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TestTransformationRequestBodyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TestTransformationRequestBodyRequest) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
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

func (o *TestTransformationRequestBodyRequest) GetParsedQuery() map[string]interface{} {
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
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// JavaScript code to be executed
	Code *string `json:"code,omitempty"`
	// Key-value environment variables to be passed to the transformation
	Env     map[string]interface{} `json:"env,omitempty"`
	EventID *string                `json:"event_id,omitempty"`
	// Request input to use for the transformation execution
	Request *TestTransformationRequestBodyRequest `json:"request,omitempty"`
	// Transformation ID
	TransformationID *string `json:"transformation_id,omitempty"`
	// ID of the connection to use for the execution `context`
	WebhookID *string `json:"webhook_id,omitempty"`
}

func (t TestTransformationRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TestTransformationRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TestTransformationRequestBody) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *TestTransformationRequestBody) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *TestTransformationRequestBody) GetEnv() map[string]interface{} {
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

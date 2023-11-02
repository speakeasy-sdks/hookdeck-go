// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
	"net/http"
)

type UpsertTransformationRequestBodyEnvType string

const (
	UpsertTransformationRequestBodyEnvTypeStr     UpsertTransformationRequestBodyEnvType = "str"
	UpsertTransformationRequestBodyEnvTypeFloat32 UpsertTransformationRequestBodyEnvType = "float32"
)

type UpsertTransformationRequestBodyEnv struct {
	Str     *string
	Float32 *float32

	Type UpsertTransformationRequestBodyEnvType
}

func CreateUpsertTransformationRequestBodyEnvStr(str string) UpsertTransformationRequestBodyEnv {
	typ := UpsertTransformationRequestBodyEnvTypeStr

	return UpsertTransformationRequestBodyEnv{
		Str:  &str,
		Type: typ,
	}
}

func CreateUpsertTransformationRequestBodyEnvFloat32(float32T float32) UpsertTransformationRequestBodyEnv {
	typ := UpsertTransformationRequestBodyEnvTypeFloat32

	return UpsertTransformationRequestBodyEnv{
		Float32: &float32T,
		Type:    typ,
	}
}

func (u *UpsertTransformationRequestBodyEnv) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UpsertTransformationRequestBodyEnvTypeStr
		return nil
	}

	float32Var := float32(0)
	if err := utils.UnmarshalJSON(data, &float32Var, "", true, true); err == nil {
		u.Float32 = &float32Var
		u.Type = UpsertTransformationRequestBodyEnvTypeFloat32
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpsertTransformationRequestBodyEnv) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Float32 != nil {
		return utils.MarshalJSON(u.Float32, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpsertTransformationRequestBody struct {
	// JavaScript code to be executed as string
	Code string `json:"code"`
	// Key-value environment variables to be passed to the transformation
	Env map[string]UpsertTransformationRequestBodyEnv `json:"env,omitempty"`
	// A unique, human-friendly name for the transformation
	Name string `json:"name"`
}

func (o *UpsertTransformationRequestBody) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *UpsertTransformationRequestBody) GetEnv() map[string]UpsertTransformationRequestBodyEnv {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *UpsertTransformationRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type UpsertTransformationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A single transformation
	Transformation *shared.Transformation
}

func (o *UpsertTransformationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertTransformationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertTransformationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertTransformationResponse) GetTransformation() *shared.Transformation {
	if o == nil {
		return nil
	}
	return o.Transformation
}

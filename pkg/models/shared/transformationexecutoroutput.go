// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type TransformationExecutorOutputRequestBody2 struct {
}

type TransformationExecutorOutputRequestBodyType string

const (
	TransformationExecutorOutputRequestBodyTypeStr                                      TransformationExecutorOutputRequestBodyType = "str"
	TransformationExecutorOutputRequestBodyTypeTransformationExecutorOutputRequestBody2 TransformationExecutorOutputRequestBodyType = "TransformationExecutorOutput_request_body_2"
)

type TransformationExecutorOutputRequestBody struct {
	Str                                      *string
	TransformationExecutorOutputRequestBody2 *TransformationExecutorOutputRequestBody2

	Type TransformationExecutorOutputRequestBodyType
}

func CreateTransformationExecutorOutputRequestBodyStr(str string) TransformationExecutorOutputRequestBody {
	typ := TransformationExecutorOutputRequestBodyTypeStr

	return TransformationExecutorOutputRequestBody{
		Str:  &str,
		Type: typ,
	}
}

func CreateTransformationExecutorOutputRequestBodyTransformationExecutorOutputRequestBody2(transformationExecutorOutputRequestBody2 TransformationExecutorOutputRequestBody2) TransformationExecutorOutputRequestBody {
	typ := TransformationExecutorOutputRequestBodyTypeTransformationExecutorOutputRequestBody2

	return TransformationExecutorOutputRequestBody{
		TransformationExecutorOutputRequestBody2: &transformationExecutorOutputRequestBody2,
		Type:                                     typ,
	}
}

func (u *TransformationExecutorOutputRequestBody) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = TransformationExecutorOutputRequestBodyTypeStr
		return nil
	}

	transformationExecutorOutputRequestBody2 := new(TransformationExecutorOutputRequestBody2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&transformationExecutorOutputRequestBody2); err == nil {
		u.TransformationExecutorOutputRequestBody2 = transformationExecutorOutputRequestBody2
		u.Type = TransformationExecutorOutputRequestBodyTypeTransformationExecutorOutputRequestBody2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TransformationExecutorOutputRequestBody) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.TransformationExecutorOutputRequestBody2 != nil {
		return json.Marshal(u.TransformationExecutorOutputRequestBody2)
	}

	return nil, nil
}

type TransformationExecutorOutputRequestHeadersType string

const (
	TransformationExecutorOutputRequestHeadersTypeStr      TransformationExecutorOutputRequestHeadersType = "str"
	TransformationExecutorOutputRequestHeadersTypeMapOfany TransformationExecutorOutputRequestHeadersType = "mapOfany"
)

type TransformationExecutorOutputRequestHeaders struct {
	Str      *string
	MapOfany map[string]interface{}

	Type TransformationExecutorOutputRequestHeadersType
}

func CreateTransformationExecutorOutputRequestHeadersStr(str string) TransformationExecutorOutputRequestHeaders {
	typ := TransformationExecutorOutputRequestHeadersTypeStr

	return TransformationExecutorOutputRequestHeaders{
		Str:  &str,
		Type: typ,
	}
}

func CreateTransformationExecutorOutputRequestHeadersMapOfany(mapOfany map[string]interface{}) TransformationExecutorOutputRequestHeaders {
	typ := TransformationExecutorOutputRequestHeadersTypeMapOfany

	return TransformationExecutorOutputRequestHeaders{
		MapOfany: mapOfany,
		Type:     typ,
	}
}

func (u *TransformationExecutorOutputRequestHeaders) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = TransformationExecutorOutputRequestHeadersTypeStr
		return nil
	}

	mapOfany := map[string]interface{}{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&mapOfany); err == nil {
		u.MapOfany = mapOfany
		u.Type = TransformationExecutorOutputRequestHeadersTypeMapOfany
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TransformationExecutorOutputRequestHeaders) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.MapOfany != nil {
		return json.Marshal(u.MapOfany)
	}

	return nil, nil
}

type TransformationExecutorOutputRequestParsedQuery2 struct {
}

type TransformationExecutorOutputRequestParsedQueryType string

const (
	TransformationExecutorOutputRequestParsedQueryTypeStr                                             TransformationExecutorOutputRequestParsedQueryType = "str"
	TransformationExecutorOutputRequestParsedQueryTypeTransformationExecutorOutputRequestParsedQuery2 TransformationExecutorOutputRequestParsedQueryType = "TransformationExecutorOutput_request_parsed_query_2"
)

type TransformationExecutorOutputRequestParsedQuery struct {
	Str                                             *string
	TransformationExecutorOutputRequestParsedQuery2 *TransformationExecutorOutputRequestParsedQuery2

	Type TransformationExecutorOutputRequestParsedQueryType
}

func CreateTransformationExecutorOutputRequestParsedQueryStr(str string) TransformationExecutorOutputRequestParsedQuery {
	typ := TransformationExecutorOutputRequestParsedQueryTypeStr

	return TransformationExecutorOutputRequestParsedQuery{
		Str:  &str,
		Type: typ,
	}
}

func CreateTransformationExecutorOutputRequestParsedQueryTransformationExecutorOutputRequestParsedQuery2(transformationExecutorOutputRequestParsedQuery2 TransformationExecutorOutputRequestParsedQuery2) TransformationExecutorOutputRequestParsedQuery {
	typ := TransformationExecutorOutputRequestParsedQueryTypeTransformationExecutorOutputRequestParsedQuery2

	return TransformationExecutorOutputRequestParsedQuery{
		TransformationExecutorOutputRequestParsedQuery2: &transformationExecutorOutputRequestParsedQuery2,
		Type: typ,
	}
}

func (u *TransformationExecutorOutputRequestParsedQuery) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = TransformationExecutorOutputRequestParsedQueryTypeStr
		return nil
	}

	transformationExecutorOutputRequestParsedQuery2 := new(TransformationExecutorOutputRequestParsedQuery2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&transformationExecutorOutputRequestParsedQuery2); err == nil {
		u.TransformationExecutorOutputRequestParsedQuery2 = transformationExecutorOutputRequestParsedQuery2
		u.Type = TransformationExecutorOutputRequestParsedQueryTypeTransformationExecutorOutputRequestParsedQuery2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TransformationExecutorOutputRequestParsedQuery) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.TransformationExecutorOutputRequestParsedQuery2 != nil {
		return json.Marshal(u.TransformationExecutorOutputRequestParsedQuery2)
	}

	return nil, nil
}

type TransformationExecutorOutputRequestQuery1 struct {
}

type TransformationExecutorOutputRequestQueryType string

const (
	TransformationExecutorOutputRequestQueryTypeTransformationExecutorOutputRequestQuery1 TransformationExecutorOutputRequestQueryType = "TransformationExecutorOutput_request_query_1"
	TransformationExecutorOutputRequestQueryTypeStr                                       TransformationExecutorOutputRequestQueryType = "str"
)

type TransformationExecutorOutputRequestQuery struct {
	TransformationExecutorOutputRequestQuery1 *TransformationExecutorOutputRequestQuery1
	Str                                       *string

	Type TransformationExecutorOutputRequestQueryType
}

func CreateTransformationExecutorOutputRequestQueryTransformationExecutorOutputRequestQuery1(transformationExecutorOutputRequestQuery1 TransformationExecutorOutputRequestQuery1) TransformationExecutorOutputRequestQuery {
	typ := TransformationExecutorOutputRequestQueryTypeTransformationExecutorOutputRequestQuery1

	return TransformationExecutorOutputRequestQuery{
		TransformationExecutorOutputRequestQuery1: &transformationExecutorOutputRequestQuery1,
		Type: typ,
	}
}

func CreateTransformationExecutorOutputRequestQueryStr(str string) TransformationExecutorOutputRequestQuery {
	typ := TransformationExecutorOutputRequestQueryTypeStr

	return TransformationExecutorOutputRequestQuery{
		Str:  &str,
		Type: typ,
	}
}

func (u *TransformationExecutorOutputRequestQuery) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	transformationExecutorOutputRequestQuery1 := new(TransformationExecutorOutputRequestQuery1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&transformationExecutorOutputRequestQuery1); err == nil {
		u.TransformationExecutorOutputRequestQuery1 = transformationExecutorOutputRequestQuery1
		u.Type = TransformationExecutorOutputRequestQueryTypeTransformationExecutorOutputRequestQuery1
		return nil
	}

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = TransformationExecutorOutputRequestQueryTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TransformationExecutorOutputRequestQuery) MarshalJSON() ([]byte, error) {
	if u.TransformationExecutorOutputRequestQuery1 != nil {
		return json.Marshal(u.TransformationExecutorOutputRequestQuery1)
	}

	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	return nil, nil
}

type TransformationExecutorOutputRequest struct {
	Body        *TransformationExecutorOutputRequestBody        `json:"body,omitempty"`
	Headers     *TransformationExecutorOutputRequestHeaders     `json:"headers,omitempty"`
	ParsedQuery *TransformationExecutorOutputRequestParsedQuery `json:"parsed_query,omitempty"`
	Path        string                                          `json:"path"`
	Query       *TransformationExecutorOutputRequestQuery       `json:"query,omitempty"`
}

// TransformationExecutorOutput - Transformation run output
type TransformationExecutorOutput struct {
	Console     []ConsoleLine `json:"console,omitempty"`
	ExecutionID *string       `json:"execution_id,omitempty"`
	// The minimum log level to open the issue on
	LogLevel         TransformationExecutionLogLevel      `json:"log_level"`
	Request          *TransformationExecutorOutputRequest `json:"request,omitempty"`
	RequestID        *string                              `json:"request_id,omitempty"`
	TransformationID *string                              `json:"transformation_id,omitempty"`
}

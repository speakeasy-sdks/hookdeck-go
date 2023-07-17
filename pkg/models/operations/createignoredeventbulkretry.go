// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type CreateIgnoredEventBulkRetryRequestBodyQueryCauseType string

const (
	CreateIgnoredEventBulkRetryRequestBodyQueryCauseTypeStr        CreateIgnoredEventBulkRetryRequestBodyQueryCauseType = "str"
	CreateIgnoredEventBulkRetryRequestBodyQueryCauseTypeArrayOfstr CreateIgnoredEventBulkRetryRequestBodyQueryCauseType = "arrayOfstr"
)

type CreateIgnoredEventBulkRetryRequestBodyQueryCause struct {
	Str        *string
	ArrayOfstr []string

	Type CreateIgnoredEventBulkRetryRequestBodyQueryCauseType
}

func CreateCreateIgnoredEventBulkRetryRequestBodyQueryCauseStr(str string) CreateIgnoredEventBulkRetryRequestBodyQueryCause {
	typ := CreateIgnoredEventBulkRetryRequestBodyQueryCauseTypeStr

	return CreateIgnoredEventBulkRetryRequestBodyQueryCause{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateIgnoredEventBulkRetryRequestBodyQueryCauseArrayOfstr(arrayOfstr []string) CreateIgnoredEventBulkRetryRequestBodyQueryCause {
	typ := CreateIgnoredEventBulkRetryRequestBodyQueryCauseTypeArrayOfstr

	return CreateIgnoredEventBulkRetryRequestBodyQueryCause{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *CreateIgnoredEventBulkRetryRequestBodyQueryCause) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = CreateIgnoredEventBulkRetryRequestBodyQueryCauseTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateIgnoredEventBulkRetryRequestBodyQueryCauseTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateIgnoredEventBulkRetryRequestBodyQueryCause) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

type CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDType string

const (
	CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDTypeStr        CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDType = "str"
	CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDTypeArrayOfstr CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDType = "arrayOfstr"
)

type CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDType
}

func CreateCreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDStr(str string) CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID {
	typ := CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDTypeStr

	return CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDArrayOfstr(arrayOfstr []string) CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID {
	typ := CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDTypeArrayOfstr

	return CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateIgnoredEventBulkRetryRequestBodyQueryWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// CreateIgnoredEventBulkRetryRequestBodyQuery - Filter by the bulk retry ignored event query object
type CreateIgnoredEventBulkRetryRequestBodyQuery struct {
	// The cause of the ignored event
	Cause *CreateIgnoredEventBulkRetryRequestBodyQueryCause `json:"cause,omitempty"`
	// The associated transformation ID (only applicable to the cause `TRANSFORMATION_FAILED`)
	TransformationID *string `json:"transformation_id,omitempty"`
	// Connection ID of the ignored event
	WebhookID *CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID `json:"webhook_id,omitempty"`
}

func (o *CreateIgnoredEventBulkRetryRequestBodyQuery) GetCause() *CreateIgnoredEventBulkRetryRequestBodyQueryCause {
	if o == nil {
		return nil
	}
	return o.Cause
}

func (o *CreateIgnoredEventBulkRetryRequestBodyQuery) GetTransformationID() *string {
	if o == nil {
		return nil
	}
	return o.TransformationID
}

func (o *CreateIgnoredEventBulkRetryRequestBodyQuery) GetWebhookID() *CreateIgnoredEventBulkRetryRequestBodyQueryWebhookID {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

type CreateIgnoredEventBulkRetryRequestBody struct {
	// Filter by the bulk retry ignored event query object
	Query *CreateIgnoredEventBulkRetryRequestBodyQuery `json:"query,omitempty"`
}

func (o *CreateIgnoredEventBulkRetryRequestBody) GetQuery() *CreateIgnoredEventBulkRetryRequestBodyQuery {
	if o == nil {
		return nil
	}
	return o.Query
}

type CreateIgnoredEventBulkRetryResponse struct {
	// A single ignored events bulk retry
	BatchOperation *shared.BatchOperation
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

func (o *CreateIgnoredEventBulkRetryResponse) GetBatchOperation() *shared.BatchOperation {
	if o == nil {
		return nil
	}
	return o.BatchOperation
}

func (o *CreateIgnoredEventBulkRetryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIgnoredEventBulkRetryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIgnoredEventBulkRetryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

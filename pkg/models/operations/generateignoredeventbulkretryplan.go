// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type GenerateIgnoredEventBulkRetryPlanQueryCauseType string

const (
	GenerateIgnoredEventBulkRetryPlanQueryCauseTypeStr        GenerateIgnoredEventBulkRetryPlanQueryCauseType = "str"
	GenerateIgnoredEventBulkRetryPlanQueryCauseTypeArrayOfstr GenerateIgnoredEventBulkRetryPlanQueryCauseType = "arrayOfstr"
)

type GenerateIgnoredEventBulkRetryPlanQueryCause struct {
	Str        *string
	ArrayOfstr []string

	Type GenerateIgnoredEventBulkRetryPlanQueryCauseType
}

func CreateGenerateIgnoredEventBulkRetryPlanQueryCauseStr(str string) GenerateIgnoredEventBulkRetryPlanQueryCause {
	typ := GenerateIgnoredEventBulkRetryPlanQueryCauseTypeStr

	return GenerateIgnoredEventBulkRetryPlanQueryCause{
		Str:  &str,
		Type: typ,
	}
}

func CreateGenerateIgnoredEventBulkRetryPlanQueryCauseArrayOfstr(arrayOfstr []string) GenerateIgnoredEventBulkRetryPlanQueryCause {
	typ := GenerateIgnoredEventBulkRetryPlanQueryCauseTypeArrayOfstr

	return GenerateIgnoredEventBulkRetryPlanQueryCause{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GenerateIgnoredEventBulkRetryPlanQueryCause) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryCauseTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryCauseTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GenerateIgnoredEventBulkRetryPlanQueryCause) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

type GenerateIgnoredEventBulkRetryPlanQueryWebhookIDType string

const (
	GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeStr        GenerateIgnoredEventBulkRetryPlanQueryWebhookIDType = "str"
	GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeArrayOfstr GenerateIgnoredEventBulkRetryPlanQueryWebhookIDType = "arrayOfstr"
)

type GenerateIgnoredEventBulkRetryPlanQueryWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type GenerateIgnoredEventBulkRetryPlanQueryWebhookIDType
}

func CreateGenerateIgnoredEventBulkRetryPlanQueryWebhookIDStr(str string) GenerateIgnoredEventBulkRetryPlanQueryWebhookID {
	typ := GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeStr

	return GenerateIgnoredEventBulkRetryPlanQueryWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGenerateIgnoredEventBulkRetryPlanQueryWebhookIDArrayOfstr(arrayOfstr []string) GenerateIgnoredEventBulkRetryPlanQueryWebhookID {
	typ := GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeArrayOfstr

	return GenerateIgnoredEventBulkRetryPlanQueryWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GenerateIgnoredEventBulkRetryPlanQueryWebhookID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GenerateIgnoredEventBulkRetryPlanQueryWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// GenerateIgnoredEventBulkRetryPlanQuery - Filter by the bulk retry ignored event query object
type GenerateIgnoredEventBulkRetryPlanQuery struct {
	// The cause of the ignored event
	Cause *GenerateIgnoredEventBulkRetryPlanQueryCause `queryParam:"name=cause"`
	// The associated transformation ID (only applicable to the cause `TRANSFORMATION_FAILED`)
	TransformationID *string `queryParam:"name=transformation_id"`
	// Connection ID of the ignored event
	WebhookID *GenerateIgnoredEventBulkRetryPlanQueryWebhookID `queryParam:"name=webhook_id"`
}

type GenerateIgnoredEventBulkRetryPlanRequest struct {
	// Filter by the bulk retry ignored event query object
	Query *GenerateIgnoredEventBulkRetryPlanQuery `queryParam:"style=form,explode=true,name=query"`
}

type GenerateIgnoredEventBulkRetryPlanResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	// Ignored events bulk retry plan
	BatchOperationPlan *shared.BatchOperationPlan
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
}

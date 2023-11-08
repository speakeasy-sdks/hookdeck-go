// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type CauseType string

const (
	CauseTypeStr        CauseType = "str"
	CauseTypeArrayOfstr CauseType = "arrayOfstr"
)

type Cause struct {
	Str        *string
	ArrayOfstr []string

	Type CauseType
}

func CreateCauseStr(str string) Cause {
	typ := CauseTypeStr

	return Cause{
		Str:  &str,
		Type: typ,
	}
}

func CreateCauseArrayOfstr(arrayOfstr []string) Cause {
	typ := CauseTypeArrayOfstr

	return Cause{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *Cause) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CauseTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CauseTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Cause) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateIgnoredEventBulkRetryWebhookIDType string

const (
	CreateIgnoredEventBulkRetryWebhookIDTypeStr        CreateIgnoredEventBulkRetryWebhookIDType = "str"
	CreateIgnoredEventBulkRetryWebhookIDTypeArrayOfstr CreateIgnoredEventBulkRetryWebhookIDType = "arrayOfstr"
)

type CreateIgnoredEventBulkRetryWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type CreateIgnoredEventBulkRetryWebhookIDType
}

func CreateCreateIgnoredEventBulkRetryWebhookIDStr(str string) CreateIgnoredEventBulkRetryWebhookID {
	typ := CreateIgnoredEventBulkRetryWebhookIDTypeStr

	return CreateIgnoredEventBulkRetryWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateIgnoredEventBulkRetryWebhookIDArrayOfstr(arrayOfstr []string) CreateIgnoredEventBulkRetryWebhookID {
	typ := CreateIgnoredEventBulkRetryWebhookIDTypeArrayOfstr

	return CreateIgnoredEventBulkRetryWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *CreateIgnoredEventBulkRetryWebhookID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateIgnoredEventBulkRetryWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateIgnoredEventBulkRetryWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateIgnoredEventBulkRetryWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// CreateIgnoredEventBulkRetryQuery - Filter by the bulk retry ignored event query object
type CreateIgnoredEventBulkRetryQuery struct {
	// The cause of the ignored event
	Cause *Cause `json:"cause,omitempty"`
	// The associated transformation ID (only applicable to the cause `TRANSFORMATION_FAILED`)
	TransformationID *string `json:"transformation_id,omitempty"`
	// Connection ID of the ignored event
	WebhookID *CreateIgnoredEventBulkRetryWebhookID `json:"webhook_id,omitempty"`
}

func (o *CreateIgnoredEventBulkRetryQuery) GetCause() *Cause {
	if o == nil {
		return nil
	}
	return o.Cause
}

func (o *CreateIgnoredEventBulkRetryQuery) GetTransformationID() *string {
	if o == nil {
		return nil
	}
	return o.TransformationID
}

func (o *CreateIgnoredEventBulkRetryQuery) GetWebhookID() *CreateIgnoredEventBulkRetryWebhookID {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

type CreateIgnoredEventBulkRetryRequestBody struct {
	// Filter by the bulk retry ignored event query object
	Query *CreateIgnoredEventBulkRetryQuery `json:"query,omitempty"`
}

func (o *CreateIgnoredEventBulkRetryRequestBody) GetQuery() *CreateIgnoredEventBulkRetryQuery {
	if o == nil {
		return nil
	}
	return o.Query
}

type CreateIgnoredEventBulkRetryResponse struct {
	// A single ignored events bulk retry
	BatchOperation *components.BatchOperation
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateIgnoredEventBulkRetryResponse) GetBatchOperation() *components.BatchOperation {
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

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
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

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryCauseTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryCauseTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GenerateIgnoredEventBulkRetryPlanQueryCause) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
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

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GenerateIgnoredEventBulkRetryPlanQueryWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GenerateIgnoredEventBulkRetryPlanQueryWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
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

func (o *GenerateIgnoredEventBulkRetryPlanQuery) GetCause() *GenerateIgnoredEventBulkRetryPlanQueryCause {
	if o == nil {
		return nil
	}
	return o.Cause
}

func (o *GenerateIgnoredEventBulkRetryPlanQuery) GetTransformationID() *string {
	if o == nil {
		return nil
	}
	return o.TransformationID
}

func (o *GenerateIgnoredEventBulkRetryPlanQuery) GetWebhookID() *GenerateIgnoredEventBulkRetryPlanQueryWebhookID {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

type GenerateIgnoredEventBulkRetryPlanRequest struct {
	// Filter by the bulk retry ignored event query object
	Query *GenerateIgnoredEventBulkRetryPlanQuery `queryParam:"style=form,explode=true,name=query"`
}

func (o *GenerateIgnoredEventBulkRetryPlanRequest) GetQuery() *GenerateIgnoredEventBulkRetryPlanQuery {
	if o == nil {
		return nil
	}
	return o.Query
}

type GenerateIgnoredEventBulkRetryPlanResponse struct {
	// Ignored events bulk retry plan
	BatchOperationPlan *shared.BatchOperationPlan
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GenerateIgnoredEventBulkRetryPlanResponse) GetBatchOperationPlan() *shared.BatchOperationPlan {
	if o == nil {
		return nil
	}
	return o.BatchOperationPlan
}

func (o *GenerateIgnoredEventBulkRetryPlanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateIgnoredEventBulkRetryPlanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateIgnoredEventBulkRetryPlanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

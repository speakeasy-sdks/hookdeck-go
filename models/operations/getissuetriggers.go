// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
	"time"
)

type GetIssueTriggersQueryParam2 string

const (
	GetIssueTriggersQueryParam2Asc  GetIssueTriggersQueryParam2 = "asc"
	GetIssueTriggersQueryParam2Desc GetIssueTriggersQueryParam2 = "desc"
)

func (e GetIssueTriggersQueryParam2) ToPointer() *GetIssueTriggersQueryParam2 {
	return &e
}

func (e *GetIssueTriggersQueryParam2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIssueTriggersQueryParam2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueTriggersQueryParam2: %v", v)
	}
}

// GetIssueTriggersQueryParam1 - Sort direction
type GetIssueTriggersQueryParam1 string

const (
	GetIssueTriggersQueryParam1Asc  GetIssueTriggersQueryParam1 = "asc"
	GetIssueTriggersQueryParam1Desc GetIssueTriggersQueryParam1 = "desc"
)

func (e GetIssueTriggersQueryParam1) ToPointer() *GetIssueTriggersQueryParam1 {
	return &e
}

func (e *GetIssueTriggersQueryParam1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIssueTriggersQueryParam1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueTriggersQueryParam1: %v", v)
	}
}

type GetIssueTriggersQueryParamDirType string

const (
	GetIssueTriggersQueryParamDirTypeGetIssueTriggersQueryParam1        GetIssueTriggersQueryParamDirType = "getIssueTriggers_queryParam_1"
	GetIssueTriggersQueryParamDirTypeArrayOfgetIssueTriggersQueryParam2 GetIssueTriggersQueryParamDirType = "arrayOfgetIssueTriggers_queryParam_2"
)

type GetIssueTriggersQueryParamDir struct {
	GetIssueTriggersQueryParam1        *GetIssueTriggersQueryParam1
	ArrayOfgetIssueTriggersQueryParam2 []GetIssueTriggersQueryParam2

	Type GetIssueTriggersQueryParamDirType
}

func CreateGetIssueTriggersQueryParamDirGetIssueTriggersQueryParam1(getIssueTriggersQueryParam1 GetIssueTriggersQueryParam1) GetIssueTriggersQueryParamDir {
	typ := GetIssueTriggersQueryParamDirTypeGetIssueTriggersQueryParam1

	return GetIssueTriggersQueryParamDir{
		GetIssueTriggersQueryParam1: &getIssueTriggersQueryParam1,
		Type:                        typ,
	}
}

func CreateGetIssueTriggersQueryParamDirArrayOfgetIssueTriggersQueryParam2(arrayOfgetIssueTriggersQueryParam2 []GetIssueTriggersQueryParam2) GetIssueTriggersQueryParamDir {
	typ := GetIssueTriggersQueryParamDirTypeArrayOfgetIssueTriggersQueryParam2

	return GetIssueTriggersQueryParamDir{
		ArrayOfgetIssueTriggersQueryParam2: arrayOfgetIssueTriggersQueryParam2,
		Type:                               typ,
	}
}

func (u *GetIssueTriggersQueryParamDir) UnmarshalJSON(data []byte) error {

	getIssueTriggersQueryParam1 := GetIssueTriggersQueryParam1("")
	if err := utils.UnmarshalJSON(data, &getIssueTriggersQueryParam1, "", true, true); err == nil {
		u.GetIssueTriggersQueryParam1 = &getIssueTriggersQueryParam1
		u.Type = GetIssueTriggersQueryParamDirTypeGetIssueTriggersQueryParam1
		return nil
	}

	arrayOfgetIssueTriggersQueryParam2 := []GetIssueTriggersQueryParam2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetIssueTriggersQueryParam2, "", true, true); err == nil {
		u.ArrayOfgetIssueTriggersQueryParam2 = arrayOfgetIssueTriggersQueryParam2
		u.Type = GetIssueTriggersQueryParamDirTypeArrayOfgetIssueTriggersQueryParam2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueTriggersQueryParamDir) MarshalJSON() ([]byte, error) {
	if u.GetIssueTriggersQueryParam1 != nil {
		return utils.MarshalJSON(u.GetIssueTriggersQueryParam1, "", true)
	}

	if u.ArrayOfgetIssueTriggersQueryParam2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetIssueTriggersQueryParam2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetIssueTriggersQueryParamIssueTriggers2 - Date when the issue trigger was disabled
type GetIssueTriggersQueryParamIssueTriggers2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetIssueTriggersQueryParamIssueTriggers2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIssueTriggersQueryParamIssueTriggers2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetIssueTriggersQueryParamIssueTriggers2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetIssueTriggersQueryParamIssueTriggers2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetIssueTriggersQueryParamIssueTriggers2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetIssueTriggersQueryParamIssueTriggers2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetIssueTriggersQueryParamIssueTriggers2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type DisabledAtType string

const (
	DisabledAtTypeDateTime                                 DisabledAtType = "date-time"
	DisabledAtTypeGetIssueTriggersQueryParamIssueTriggers2 DisabledAtType = "getIssueTriggers_queryParam_IssueTriggers_2"
)

type DisabledAt struct {
	DateTime                                 *time.Time
	GetIssueTriggersQueryParamIssueTriggers2 *GetIssueTriggersQueryParamIssueTriggers2

	Type DisabledAtType
}

func CreateDisabledAtDateTime(dateTime time.Time) DisabledAt {
	typ := DisabledAtTypeDateTime

	return DisabledAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateDisabledAtGetIssueTriggersQueryParamIssueTriggers2(getIssueTriggersQueryParamIssueTriggers2 GetIssueTriggersQueryParamIssueTriggers2) DisabledAt {
	typ := DisabledAtTypeGetIssueTriggersQueryParamIssueTriggers2

	return DisabledAt{
		GetIssueTriggersQueryParamIssueTriggers2: &getIssueTriggersQueryParamIssueTriggers2,
		Type:                                     typ,
	}
}

func (u *DisabledAt) UnmarshalJSON(data []byte) error {

	getIssueTriggersQueryParamIssueTriggers2 := GetIssueTriggersQueryParamIssueTriggers2{}
	if err := utils.UnmarshalJSON(data, &getIssueTriggersQueryParamIssueTriggers2, "", true, true); err == nil {
		u.GetIssueTriggersQueryParamIssueTriggers2 = &getIssueTriggersQueryParamIssueTriggers2
		u.Type = DisabledAtTypeGetIssueTriggersQueryParamIssueTriggers2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = DisabledAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DisabledAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetIssueTriggersQueryParamIssueTriggers2 != nil {
		return utils.MarshalJSON(u.GetIssueTriggersQueryParamIssueTriggers2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssueTriggersQueryParamIssueTriggersOrderBy2 string

const (
	GetIssueTriggersQueryParamIssueTriggersOrderBy2CreatedAt GetIssueTriggersQueryParamIssueTriggersOrderBy2 = "created_at"
	GetIssueTriggersQueryParamIssueTriggersOrderBy2Type      GetIssueTriggersQueryParamIssueTriggersOrderBy2 = "type"
)

func (e GetIssueTriggersQueryParamIssueTriggersOrderBy2) ToPointer() *GetIssueTriggersQueryParamIssueTriggersOrderBy2 {
	return &e
}

func (e *GetIssueTriggersQueryParamIssueTriggersOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "type":
		*e = GetIssueTriggersQueryParamIssueTriggersOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueTriggersQueryParamIssueTriggersOrderBy2: %v", v)
	}
}

// GetIssueTriggersQueryParamIssueTriggers1 - Sort key(s)
type GetIssueTriggersQueryParamIssueTriggers1 string

const (
	GetIssueTriggersQueryParamIssueTriggers1CreatedAt GetIssueTriggersQueryParamIssueTriggers1 = "created_at"
	GetIssueTriggersQueryParamIssueTriggers1Type      GetIssueTriggersQueryParamIssueTriggers1 = "type"
)

func (e GetIssueTriggersQueryParamIssueTriggers1) ToPointer() *GetIssueTriggersQueryParamIssueTriggers1 {
	return &e
}

func (e *GetIssueTriggersQueryParamIssueTriggers1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "type":
		*e = GetIssueTriggersQueryParamIssueTriggers1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueTriggersQueryParamIssueTriggers1: %v", v)
	}
}

type GetIssueTriggersQueryParamOrderByType string

const (
	GetIssueTriggersQueryParamOrderByTypeGetIssueTriggersQueryParamIssueTriggers1               GetIssueTriggersQueryParamOrderByType = "getIssueTriggers_queryParam_IssueTriggers_1"
	GetIssueTriggersQueryParamOrderByTypeArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2 GetIssueTriggersQueryParamOrderByType = "arrayOfgetIssueTriggers_queryParam_IssueTriggers_order_by_2"
)

type GetIssueTriggersQueryParamOrderBy struct {
	GetIssueTriggersQueryParamIssueTriggers1               *GetIssueTriggersQueryParamIssueTriggers1
	ArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2 []GetIssueTriggersQueryParamIssueTriggersOrderBy2

	Type GetIssueTriggersQueryParamOrderByType
}

func CreateGetIssueTriggersQueryParamOrderByGetIssueTriggersQueryParamIssueTriggers1(getIssueTriggersQueryParamIssueTriggers1 GetIssueTriggersQueryParamIssueTriggers1) GetIssueTriggersQueryParamOrderBy {
	typ := GetIssueTriggersQueryParamOrderByTypeGetIssueTriggersQueryParamIssueTriggers1

	return GetIssueTriggersQueryParamOrderBy{
		GetIssueTriggersQueryParamIssueTriggers1: &getIssueTriggersQueryParamIssueTriggers1,
		Type:                                     typ,
	}
}

func CreateGetIssueTriggersQueryParamOrderByArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2(arrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2 []GetIssueTriggersQueryParamIssueTriggersOrderBy2) GetIssueTriggersQueryParamOrderBy {
	typ := GetIssueTriggersQueryParamOrderByTypeArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2

	return GetIssueTriggersQueryParamOrderBy{
		ArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2: arrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2,
		Type: typ,
	}
}

func (u *GetIssueTriggersQueryParamOrderBy) UnmarshalJSON(data []byte) error {

	getIssueTriggersQueryParamIssueTriggers1 := GetIssueTriggersQueryParamIssueTriggers1("")
	if err := utils.UnmarshalJSON(data, &getIssueTriggersQueryParamIssueTriggers1, "", true, true); err == nil {
		u.GetIssueTriggersQueryParamIssueTriggers1 = &getIssueTriggersQueryParamIssueTriggers1
		u.Type = GetIssueTriggersQueryParamOrderByTypeGetIssueTriggersQueryParamIssueTriggers1
		return nil
	}

	arrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2 := []GetIssueTriggersQueryParamIssueTriggersOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2 = arrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2
		u.Type = GetIssueTriggersQueryParamOrderByTypeArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueTriggersQueryParamOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetIssueTriggersQueryParamIssueTriggers1 != nil {
		return utils.MarshalJSON(u.GetIssueTriggersQueryParamIssueTriggers1, "", true)
	}

	if u.ArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetIssueTriggersQueryParamIssueTriggersOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssueTriggersRequest struct {
	Dir        *GetIssueTriggersQueryParamDir     `queryParam:"style=form,explode=true,name=dir"`
	DisabledAt *DisabledAt                        `queryParam:"style=form,explode=true,name=disabled_at"`
	Limit      *int64                             `queryParam:"style=form,explode=true,name=limit"`
	Name       *string                            `queryParam:"style=form,explode=true,name=name"`
	Next       *string                            `queryParam:"style=form,explode=true,name=next"`
	OrderBy    *GetIssueTriggersQueryParamOrderBy `queryParam:"style=form,explode=true,name=order_by"`
	Prev       *string                            `queryParam:"style=form,explode=true,name=prev"`
	// Issue type
	Type *components.IssueType `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetIssueTriggersRequest) GetDir() *GetIssueTriggersQueryParamDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetIssueTriggersRequest) GetDisabledAt() *DisabledAt {
	if o == nil {
		return nil
	}
	return o.DisabledAt
}

func (o *GetIssueTriggersRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetIssueTriggersRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetIssueTriggersRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetIssueTriggersRequest) GetOrderBy() *GetIssueTriggersQueryParamOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetIssueTriggersRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

func (o *GetIssueTriggersRequest) GetType() *components.IssueType {
	if o == nil {
		return nil
	}
	return o.Type
}

type GetIssueTriggersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of issue triggers
	IssueTriggerPaginatedResult *components.IssueTriggerPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIssueTriggersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIssueTriggersResponse) GetIssueTriggerPaginatedResult() *components.IssueTriggerPaginatedResult {
	if o == nil {
		return nil
	}
	return o.IssueTriggerPaginatedResult
}

func (o *GetIssueTriggersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIssueTriggersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
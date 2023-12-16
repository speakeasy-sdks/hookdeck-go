// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type GetRequestIgnoredEventsQueryParam2 string

const (
	GetRequestIgnoredEventsQueryParam2Asc  GetRequestIgnoredEventsQueryParam2 = "asc"
	GetRequestIgnoredEventsQueryParam2Desc GetRequestIgnoredEventsQueryParam2 = "desc"
)

func (e GetRequestIgnoredEventsQueryParam2) ToPointer() *GetRequestIgnoredEventsQueryParam2 {
	return &e
}

func (e *GetRequestIgnoredEventsQueryParam2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRequestIgnoredEventsQueryParam2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRequestIgnoredEventsQueryParam2: %v", v)
	}
}

type GetRequestIgnoredEventsQueryParam1 string

const (
	GetRequestIgnoredEventsQueryParam1Asc  GetRequestIgnoredEventsQueryParam1 = "asc"
	GetRequestIgnoredEventsQueryParam1Desc GetRequestIgnoredEventsQueryParam1 = "desc"
)

func (e GetRequestIgnoredEventsQueryParam1) ToPointer() *GetRequestIgnoredEventsQueryParam1 {
	return &e
}

func (e *GetRequestIgnoredEventsQueryParam1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRequestIgnoredEventsQueryParam1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRequestIgnoredEventsQueryParam1: %v", v)
	}
}

type GetRequestIgnoredEventsQueryParamDirType string

const (
	GetRequestIgnoredEventsQueryParamDirTypeGetRequestIgnoredEventsQueryParam1        GetRequestIgnoredEventsQueryParamDirType = "getRequestIgnoredEvents_queryParam_1"
	GetRequestIgnoredEventsQueryParamDirTypeArrayOfgetRequestIgnoredEventsQueryParam2 GetRequestIgnoredEventsQueryParamDirType = "arrayOfgetRequestIgnoredEvents_queryParam_2"
)

// GetRequestIgnoredEventsQueryParamDir - Sort direction
type GetRequestIgnoredEventsQueryParamDir struct {
	GetRequestIgnoredEventsQueryParam1        *GetRequestIgnoredEventsQueryParam1
	ArrayOfgetRequestIgnoredEventsQueryParam2 []GetRequestIgnoredEventsQueryParam2

	Type GetRequestIgnoredEventsQueryParamDirType
}

func CreateGetRequestIgnoredEventsQueryParamDirGetRequestIgnoredEventsQueryParam1(getRequestIgnoredEventsQueryParam1 GetRequestIgnoredEventsQueryParam1) GetRequestIgnoredEventsQueryParamDir {
	typ := GetRequestIgnoredEventsQueryParamDirTypeGetRequestIgnoredEventsQueryParam1

	return GetRequestIgnoredEventsQueryParamDir{
		GetRequestIgnoredEventsQueryParam1: &getRequestIgnoredEventsQueryParam1,
		Type:                               typ,
	}
}

func CreateGetRequestIgnoredEventsQueryParamDirArrayOfgetRequestIgnoredEventsQueryParam2(arrayOfgetRequestIgnoredEventsQueryParam2 []GetRequestIgnoredEventsQueryParam2) GetRequestIgnoredEventsQueryParamDir {
	typ := GetRequestIgnoredEventsQueryParamDirTypeArrayOfgetRequestIgnoredEventsQueryParam2

	return GetRequestIgnoredEventsQueryParamDir{
		ArrayOfgetRequestIgnoredEventsQueryParam2: arrayOfgetRequestIgnoredEventsQueryParam2,
		Type: typ,
	}
}

func (u *GetRequestIgnoredEventsQueryParamDir) UnmarshalJSON(data []byte) error {

	getRequestIgnoredEventsQueryParam1 := GetRequestIgnoredEventsQueryParam1("")
	if err := utils.UnmarshalJSON(data, &getRequestIgnoredEventsQueryParam1, "", true, true); err == nil {
		u.GetRequestIgnoredEventsQueryParam1 = &getRequestIgnoredEventsQueryParam1
		u.Type = GetRequestIgnoredEventsQueryParamDirTypeGetRequestIgnoredEventsQueryParam1
		return nil
	}

	arrayOfgetRequestIgnoredEventsQueryParam2 := []GetRequestIgnoredEventsQueryParam2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRequestIgnoredEventsQueryParam2, "", true, true); err == nil {
		u.ArrayOfgetRequestIgnoredEventsQueryParam2 = arrayOfgetRequestIgnoredEventsQueryParam2
		u.Type = GetRequestIgnoredEventsQueryParamDirTypeArrayOfgetRequestIgnoredEventsQueryParam2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRequestIgnoredEventsQueryParamDir) MarshalJSON() ([]byte, error) {
	if u.GetRequestIgnoredEventsQueryParam1 != nil {
		return utils.MarshalJSON(u.GetRequestIgnoredEventsQueryParam1, "", true)
	}

	if u.ArrayOfgetRequestIgnoredEventsQueryParam2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRequestIgnoredEventsQueryParam2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRequestIgnoredEventsQueryParamIDType string

const (
	GetRequestIgnoredEventsQueryParamIDTypeStr        GetRequestIgnoredEventsQueryParamIDType = "str"
	GetRequestIgnoredEventsQueryParamIDTypeArrayOfstr GetRequestIgnoredEventsQueryParamIDType = "arrayOfstr"
)

// GetRequestIgnoredEventsQueryParamID - Filter by ignored events IDs
type GetRequestIgnoredEventsQueryParamID struct {
	Str        *string
	ArrayOfstr []string

	Type GetRequestIgnoredEventsQueryParamIDType
}

func CreateGetRequestIgnoredEventsQueryParamIDStr(str string) GetRequestIgnoredEventsQueryParamID {
	typ := GetRequestIgnoredEventsQueryParamIDTypeStr

	return GetRequestIgnoredEventsQueryParamID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRequestIgnoredEventsQueryParamIDArrayOfstr(arrayOfstr []string) GetRequestIgnoredEventsQueryParamID {
	typ := GetRequestIgnoredEventsQueryParamIDTypeArrayOfstr

	return GetRequestIgnoredEventsQueryParamID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRequestIgnoredEventsQueryParamID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetRequestIgnoredEventsQueryParamIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRequestIgnoredEventsQueryParamIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRequestIgnoredEventsQueryParamID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 string

const (
	GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2CreatedAt GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 = "created_at"
)

func (e GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2) ToPointer() *GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 {
	return &e
}

func (e *GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2: %v", v)
	}
}

type GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1 string

const (
	GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1CreatedAt GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1 = "created_at"
)

func (e GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1) ToPointer() *GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1 {
	return &e
}

func (e *GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1: %v", v)
	}
}

type GetRequestIgnoredEventsQueryParamOrderByType string

const (
	GetRequestIgnoredEventsQueryParamOrderByTypeGetRequestIgnoredEventsQueryParamRequestIgnoredEvents1        GetRequestIgnoredEventsQueryParamOrderByType = "getRequestIgnoredEvents_queryParam_RequestIgnoredEvents_1"
	GetRequestIgnoredEventsQueryParamOrderByTypeArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 GetRequestIgnoredEventsQueryParamOrderByType = "arrayOfgetRequestIgnoredEvents_queryParam_RequestIgnoredEvents_2"
)

// GetRequestIgnoredEventsQueryParamOrderBy - Sort key(s)
type GetRequestIgnoredEventsQueryParamOrderBy struct {
	GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1        *GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1
	ArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 []GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2

	Type GetRequestIgnoredEventsQueryParamOrderByType
}

func CreateGetRequestIgnoredEventsQueryParamOrderByGetRequestIgnoredEventsQueryParamRequestIgnoredEvents1(getRequestIgnoredEventsQueryParamRequestIgnoredEvents1 GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1) GetRequestIgnoredEventsQueryParamOrderBy {
	typ := GetRequestIgnoredEventsQueryParamOrderByTypeGetRequestIgnoredEventsQueryParamRequestIgnoredEvents1

	return GetRequestIgnoredEventsQueryParamOrderBy{
		GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1: &getRequestIgnoredEventsQueryParamRequestIgnoredEvents1,
		Type: typ,
	}
}

func CreateGetRequestIgnoredEventsQueryParamOrderByArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2(arrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 []GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2) GetRequestIgnoredEventsQueryParamOrderBy {
	typ := GetRequestIgnoredEventsQueryParamOrderByTypeArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2

	return GetRequestIgnoredEventsQueryParamOrderBy{
		ArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2: arrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2,
		Type: typ,
	}
}

func (u *GetRequestIgnoredEventsQueryParamOrderBy) UnmarshalJSON(data []byte) error {

	getRequestIgnoredEventsQueryParamRequestIgnoredEvents1 := GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1("")
	if err := utils.UnmarshalJSON(data, &getRequestIgnoredEventsQueryParamRequestIgnoredEvents1, "", true, true); err == nil {
		u.GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1 = &getRequestIgnoredEventsQueryParamRequestIgnoredEvents1
		u.Type = GetRequestIgnoredEventsQueryParamOrderByTypeGetRequestIgnoredEventsQueryParamRequestIgnoredEvents1
		return nil
	}

	arrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 := []GetRequestIgnoredEventsQueryParamRequestIgnoredEvents2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2, "", true, true); err == nil {
		u.ArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 = arrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2
		u.Type = GetRequestIgnoredEventsQueryParamOrderByTypeArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRequestIgnoredEventsQueryParamOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1 != nil {
		return utils.MarshalJSON(u.GetRequestIgnoredEventsQueryParamRequestIgnoredEvents1, "", true)
	}

	if u.ArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRequestIgnoredEventsQueryParamRequestIgnoredEvents2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRequestIgnoredEventsRequest struct {
	// Sort direction
	Dir             *GetRequestIgnoredEventsQueryParamDir `queryParam:"style=form,explode=true,name=dir"`
	IDPathParameter string                                `pathParam:"style=simple,explode=false,name=id"`
	// Filter by ignored events IDs
	IDQueryParameter *GetRequestIgnoredEventsQueryParamID `queryParam:"style=form,explode=true,name=id"`
	Limit            *int64                               `queryParam:"style=form,explode=true,name=limit"`
	Next             *string                              `queryParam:"style=form,explode=true,name=next"`
	// Sort key(s)
	OrderBy *GetRequestIgnoredEventsQueryParamOrderBy `queryParam:"style=form,explode=true,name=order_by"`
	Prev    *string                                   `queryParam:"style=form,explode=true,name=prev"`
}

func (o *GetRequestIgnoredEventsRequest) GetDir() *GetRequestIgnoredEventsQueryParamDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetRequestIgnoredEventsRequest) GetIDPathParameter() string {
	if o == nil {
		return ""
	}
	return o.IDPathParameter
}

func (o *GetRequestIgnoredEventsRequest) GetIDQueryParameter() *GetRequestIgnoredEventsQueryParamID {
	if o == nil {
		return nil
	}
	return o.IDQueryParameter
}

func (o *GetRequestIgnoredEventsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetRequestIgnoredEventsRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetRequestIgnoredEventsRequest) GetOrderBy() *GetRequestIgnoredEventsQueryParamOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetRequestIgnoredEventsRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type GetRequestIgnoredEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of ignored events
	IgnoredEventPaginatedResult *components.IgnoredEventPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRequestIgnoredEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRequestIgnoredEventsResponse) GetIgnoredEventPaginatedResult() *components.IgnoredEventPaginatedResult {
	if o == nil {
		return nil
	}
	return o.IgnoredEventPaginatedResult
}

func (o *GetRequestIgnoredEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRequestIgnoredEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

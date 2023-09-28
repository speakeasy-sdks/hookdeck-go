// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
	"net/http"
)

type GetRequestIgnoredEventsDir2 string

const (
	GetRequestIgnoredEventsDir2Asc  GetRequestIgnoredEventsDir2 = "asc"
	GetRequestIgnoredEventsDir2Desc GetRequestIgnoredEventsDir2 = "desc"
)

func (e GetRequestIgnoredEventsDir2) ToPointer() *GetRequestIgnoredEventsDir2 {
	return &e
}

func (e *GetRequestIgnoredEventsDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRequestIgnoredEventsDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRequestIgnoredEventsDir2: %v", v)
	}
}

// GetRequestIgnoredEventsDir1 - Sort direction
type GetRequestIgnoredEventsDir1 string

const (
	GetRequestIgnoredEventsDir1Asc  GetRequestIgnoredEventsDir1 = "asc"
	GetRequestIgnoredEventsDir1Desc GetRequestIgnoredEventsDir1 = "desc"
)

func (e GetRequestIgnoredEventsDir1) ToPointer() *GetRequestIgnoredEventsDir1 {
	return &e
}

func (e *GetRequestIgnoredEventsDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRequestIgnoredEventsDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRequestIgnoredEventsDir1: %v", v)
	}
}

type GetRequestIgnoredEventsDirType string

const (
	GetRequestIgnoredEventsDirTypeGetRequestIgnoredEventsDir1        GetRequestIgnoredEventsDirType = "getRequestIgnoredEventsDir_1"
	GetRequestIgnoredEventsDirTypeArrayOfgetRequestIgnoredEventsDir2 GetRequestIgnoredEventsDirType = "arrayOfgetRequestIgnoredEventsDir_2"
)

type GetRequestIgnoredEventsDir struct {
	GetRequestIgnoredEventsDir1        *GetRequestIgnoredEventsDir1
	ArrayOfgetRequestIgnoredEventsDir2 []GetRequestIgnoredEventsDir2

	Type GetRequestIgnoredEventsDirType
}

func CreateGetRequestIgnoredEventsDirGetRequestIgnoredEventsDir1(getRequestIgnoredEventsDir1 GetRequestIgnoredEventsDir1) GetRequestIgnoredEventsDir {
	typ := GetRequestIgnoredEventsDirTypeGetRequestIgnoredEventsDir1

	return GetRequestIgnoredEventsDir{
		GetRequestIgnoredEventsDir1: &getRequestIgnoredEventsDir1,
		Type:                        typ,
	}
}

func CreateGetRequestIgnoredEventsDirArrayOfgetRequestIgnoredEventsDir2(arrayOfgetRequestIgnoredEventsDir2 []GetRequestIgnoredEventsDir2) GetRequestIgnoredEventsDir {
	typ := GetRequestIgnoredEventsDirTypeArrayOfgetRequestIgnoredEventsDir2

	return GetRequestIgnoredEventsDir{
		ArrayOfgetRequestIgnoredEventsDir2: arrayOfgetRequestIgnoredEventsDir2,
		Type:                               typ,
	}
}

func (u *GetRequestIgnoredEventsDir) UnmarshalJSON(data []byte) error {

	getRequestIgnoredEventsDir1 := new(GetRequestIgnoredEventsDir1)
	if err := utils.UnmarshalJSON(data, &getRequestIgnoredEventsDir1, "", true, true); err == nil {
		u.GetRequestIgnoredEventsDir1 = getRequestIgnoredEventsDir1
		u.Type = GetRequestIgnoredEventsDirTypeGetRequestIgnoredEventsDir1
		return nil
	}

	arrayOfgetRequestIgnoredEventsDir2 := []GetRequestIgnoredEventsDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRequestIgnoredEventsDir2, "", true, true); err == nil {
		u.ArrayOfgetRequestIgnoredEventsDir2 = arrayOfgetRequestIgnoredEventsDir2
		u.Type = GetRequestIgnoredEventsDirTypeArrayOfgetRequestIgnoredEventsDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRequestIgnoredEventsDir) MarshalJSON() ([]byte, error) {
	if u.GetRequestIgnoredEventsDir1 != nil {
		return utils.MarshalJSON(u.GetRequestIgnoredEventsDir1, "", true)
	}

	if u.ArrayOfgetRequestIgnoredEventsDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRequestIgnoredEventsDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRequestIgnoredEventsIDType string

const (
	GetRequestIgnoredEventsIDTypeStr        GetRequestIgnoredEventsIDType = "str"
	GetRequestIgnoredEventsIDTypeArrayOfstr GetRequestIgnoredEventsIDType = "arrayOfstr"
)

type GetRequestIgnoredEventsID struct {
	Str        *string
	ArrayOfstr []string

	Type GetRequestIgnoredEventsIDType
}

func CreateGetRequestIgnoredEventsIDStr(str string) GetRequestIgnoredEventsID {
	typ := GetRequestIgnoredEventsIDTypeStr

	return GetRequestIgnoredEventsID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRequestIgnoredEventsIDArrayOfstr(arrayOfstr []string) GetRequestIgnoredEventsID {
	typ := GetRequestIgnoredEventsIDTypeArrayOfstr

	return GetRequestIgnoredEventsID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRequestIgnoredEventsID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRequestIgnoredEventsIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRequestIgnoredEventsIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRequestIgnoredEventsID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRequestIgnoredEventsOrderByType string

const (
	GetRequestIgnoredEventsOrderByTypeStr        GetRequestIgnoredEventsOrderByType = "str"
	GetRequestIgnoredEventsOrderByTypeArrayOfstr GetRequestIgnoredEventsOrderByType = "arrayOfstr"
)

type GetRequestIgnoredEventsOrderBy struct {
	Str        *string
	ArrayOfstr []string

	Type GetRequestIgnoredEventsOrderByType
}

func CreateGetRequestIgnoredEventsOrderByStr(str string) GetRequestIgnoredEventsOrderBy {
	typ := GetRequestIgnoredEventsOrderByTypeStr

	return GetRequestIgnoredEventsOrderBy{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRequestIgnoredEventsOrderByArrayOfstr(arrayOfstr []string) GetRequestIgnoredEventsOrderBy {
	typ := GetRequestIgnoredEventsOrderByTypeArrayOfstr

	return GetRequestIgnoredEventsOrderBy{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRequestIgnoredEventsOrderBy) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRequestIgnoredEventsOrderByTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRequestIgnoredEventsOrderByTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRequestIgnoredEventsOrderBy) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRequestIgnoredEventsRequest struct {
	Dir              *GetRequestIgnoredEventsDir     `queryParam:"style=form,explode=true,name=dir"`
	IDPathParameter  string                          `pathParam:"style=simple,explode=false,name=id"`
	IDQueryParameter *GetRequestIgnoredEventsID      `queryParam:"style=form,explode=true,name=id"`
	Limit            *int64                          `queryParam:"style=form,explode=true,name=limit"`
	Next             *string                         `queryParam:"style=form,explode=true,name=next"`
	OrderBy          *GetRequestIgnoredEventsOrderBy `queryParam:"style=form,explode=true,name=order_by"`
	Prev             *string                         `queryParam:"style=form,explode=true,name=prev"`
}

func (o *GetRequestIgnoredEventsRequest) GetDir() *GetRequestIgnoredEventsDir {
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

func (o *GetRequestIgnoredEventsRequest) GetIDQueryParameter() *GetRequestIgnoredEventsID {
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

func (o *GetRequestIgnoredEventsRequest) GetOrderBy() *GetRequestIgnoredEventsOrderBy {
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
	IgnoredEventPaginatedResult *shared.IgnoredEventPaginatedResult
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

func (o *GetRequestIgnoredEventsResponse) GetIgnoredEventPaginatedResult() *shared.IgnoredEventPaginatedResult {
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

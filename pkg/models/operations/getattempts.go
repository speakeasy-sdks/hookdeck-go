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

type GetAttemptsDir2 string

const (
	GetAttemptsDir2Asc  GetAttemptsDir2 = "asc"
	GetAttemptsDir2Desc GetAttemptsDir2 = "desc"
)

func (e GetAttemptsDir2) ToPointer() *GetAttemptsDir2 {
	return &e
}

func (e *GetAttemptsDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetAttemptsDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAttemptsDir2: %v", v)
	}
}

// GetAttemptsDir1 - Sort direction
type GetAttemptsDir1 string

const (
	GetAttemptsDir1Asc  GetAttemptsDir1 = "asc"
	GetAttemptsDir1Desc GetAttemptsDir1 = "desc"
)

func (e GetAttemptsDir1) ToPointer() *GetAttemptsDir1 {
	return &e
}

func (e *GetAttemptsDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetAttemptsDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAttemptsDir1: %v", v)
	}
}

type GetAttemptsDirType string

const (
	GetAttemptsDirTypeGetAttemptsDir1        GetAttemptsDirType = "getAttemptsDir_1"
	GetAttemptsDirTypeArrayOfgetAttemptsDir2 GetAttemptsDirType = "arrayOfgetAttemptsDir_2"
)

type GetAttemptsDir struct {
	GetAttemptsDir1        *GetAttemptsDir1
	ArrayOfgetAttemptsDir2 []GetAttemptsDir2

	Type GetAttemptsDirType
}

func CreateGetAttemptsDirGetAttemptsDir1(getAttemptsDir1 GetAttemptsDir1) GetAttemptsDir {
	typ := GetAttemptsDirTypeGetAttemptsDir1

	return GetAttemptsDir{
		GetAttemptsDir1: &getAttemptsDir1,
		Type:            typ,
	}
}

func CreateGetAttemptsDirArrayOfgetAttemptsDir2(arrayOfgetAttemptsDir2 []GetAttemptsDir2) GetAttemptsDir {
	typ := GetAttemptsDirTypeArrayOfgetAttemptsDir2

	return GetAttemptsDir{
		ArrayOfgetAttemptsDir2: arrayOfgetAttemptsDir2,
		Type:                   typ,
	}
}

func (u *GetAttemptsDir) UnmarshalJSON(data []byte) error {

	getAttemptsDir1 := new(GetAttemptsDir1)
	if err := utils.UnmarshalJSON(data, &getAttemptsDir1, "", true, true); err == nil {
		u.GetAttemptsDir1 = getAttemptsDir1
		u.Type = GetAttemptsDirTypeGetAttemptsDir1
		return nil
	}

	arrayOfgetAttemptsDir2 := []GetAttemptsDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetAttemptsDir2, "", true, true); err == nil {
		u.ArrayOfgetAttemptsDir2 = arrayOfgetAttemptsDir2
		u.Type = GetAttemptsDirTypeArrayOfgetAttemptsDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetAttemptsDir) MarshalJSON() ([]byte, error) {
	if u.GetAttemptsDir1 != nil {
		return utils.MarshalJSON(u.GetAttemptsDir1, "", true)
	}

	if u.ArrayOfgetAttemptsDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetAttemptsDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetAttemptsEventIDType string

const (
	GetAttemptsEventIDTypeStr        GetAttemptsEventIDType = "str"
	GetAttemptsEventIDTypeArrayOfstr GetAttemptsEventIDType = "arrayOfstr"
)

type GetAttemptsEventID struct {
	Str        *string
	ArrayOfstr []string

	Type GetAttemptsEventIDType
}

func CreateGetAttemptsEventIDStr(str string) GetAttemptsEventID {
	typ := GetAttemptsEventIDTypeStr

	return GetAttemptsEventID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetAttemptsEventIDArrayOfstr(arrayOfstr []string) GetAttemptsEventID {
	typ := GetAttemptsEventIDTypeArrayOfstr

	return GetAttemptsEventID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetAttemptsEventID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetAttemptsEventIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetAttemptsEventIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetAttemptsEventID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetAttemptsOrderByType string

const (
	GetAttemptsOrderByTypeStr        GetAttemptsOrderByType = "str"
	GetAttemptsOrderByTypeArrayOfstr GetAttemptsOrderByType = "arrayOfstr"
)

type GetAttemptsOrderBy struct {
	Str        *string
	ArrayOfstr []string

	Type GetAttemptsOrderByType
}

func CreateGetAttemptsOrderByStr(str string) GetAttemptsOrderBy {
	typ := GetAttemptsOrderByTypeStr

	return GetAttemptsOrderBy{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetAttemptsOrderByArrayOfstr(arrayOfstr []string) GetAttemptsOrderBy {
	typ := GetAttemptsOrderByTypeArrayOfstr

	return GetAttemptsOrderBy{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetAttemptsOrderBy) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetAttemptsOrderByTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetAttemptsOrderByTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetAttemptsOrderBy) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetAttemptsRequest struct {
	Dir     *GetAttemptsDir     `queryParam:"style=form,explode=true,name=dir"`
	EventID *GetAttemptsEventID `queryParam:"style=form,explode=true,name=event_id"`
	Limit   *int64              `queryParam:"style=form,explode=true,name=limit"`
	Next    *string             `queryParam:"style=form,explode=true,name=next"`
	OrderBy *GetAttemptsOrderBy `queryParam:"style=form,explode=true,name=order_by"`
	Prev    *string             `queryParam:"style=form,explode=true,name=prev"`
}

func (o *GetAttemptsRequest) GetDir() *GetAttemptsDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetAttemptsRequest) GetEventID() *GetAttemptsEventID {
	if o == nil {
		return nil
	}
	return o.EventID
}

func (o *GetAttemptsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetAttemptsRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetAttemptsRequest) GetOrderBy() *GetAttemptsOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetAttemptsRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type GetAttemptsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of attempts
	EventAttemptPaginatedResult *shared.EventAttemptPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAttemptsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAttemptsResponse) GetEventAttemptPaginatedResult() *shared.EventAttemptPaginatedResult {
	if o == nil {
		return nil
	}
	return o.EventAttemptPaginatedResult
}

func (o *GetAttemptsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAttemptsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

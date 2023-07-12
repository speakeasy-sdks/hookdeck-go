// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
	"time"
)

// GetIgnoredEventBulkRetriesCancelledAt2 - Filter by date the bulk retry was cancelled
type GetIgnoredEventBulkRetriesCancelledAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetIgnoredEventBulkRetriesCancelledAtType string

const (
	GetIgnoredEventBulkRetriesCancelledAtTypeDateTime                               GetIgnoredEventBulkRetriesCancelledAtType = "date-time"
	GetIgnoredEventBulkRetriesCancelledAtTypeGetIgnoredEventBulkRetriesCancelledAt2 GetIgnoredEventBulkRetriesCancelledAtType = "getIgnoredEventBulkRetriesCancelledAt_2"
)

type GetIgnoredEventBulkRetriesCancelledAt struct {
	DateTime                               *time.Time
	GetIgnoredEventBulkRetriesCancelledAt2 *GetIgnoredEventBulkRetriesCancelledAt2

	Type GetIgnoredEventBulkRetriesCancelledAtType
}

func CreateGetIgnoredEventBulkRetriesCancelledAtDateTime(dateTime time.Time) GetIgnoredEventBulkRetriesCancelledAt {
	typ := GetIgnoredEventBulkRetriesCancelledAtTypeDateTime

	return GetIgnoredEventBulkRetriesCancelledAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIgnoredEventBulkRetriesCancelledAtGetIgnoredEventBulkRetriesCancelledAt2(getIgnoredEventBulkRetriesCancelledAt2 GetIgnoredEventBulkRetriesCancelledAt2) GetIgnoredEventBulkRetriesCancelledAt {
	typ := GetIgnoredEventBulkRetriesCancelledAtTypeGetIgnoredEventBulkRetriesCancelledAt2

	return GetIgnoredEventBulkRetriesCancelledAt{
		GetIgnoredEventBulkRetriesCancelledAt2: &getIgnoredEventBulkRetriesCancelledAt2,
		Type:                                   typ,
	}
}

func (u *GetIgnoredEventBulkRetriesCancelledAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetIgnoredEventBulkRetriesCancelledAtTypeDateTime
		return nil
	}

	getIgnoredEventBulkRetriesCancelledAt2 := new(GetIgnoredEventBulkRetriesCancelledAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIgnoredEventBulkRetriesCancelledAt2); err == nil {
		u.GetIgnoredEventBulkRetriesCancelledAt2 = getIgnoredEventBulkRetriesCancelledAt2
		u.Type = GetIgnoredEventBulkRetriesCancelledAtTypeGetIgnoredEventBulkRetriesCancelledAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesCancelledAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetIgnoredEventBulkRetriesCancelledAt2 != nil {
		return json.Marshal(u.GetIgnoredEventBulkRetriesCancelledAt2)
	}

	return nil, nil
}

// GetIgnoredEventBulkRetriesCompletedAt2 - Filter by date the bulk retry completed
type GetIgnoredEventBulkRetriesCompletedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetIgnoredEventBulkRetriesCompletedAtType string

const (
	GetIgnoredEventBulkRetriesCompletedAtTypeDateTime                               GetIgnoredEventBulkRetriesCompletedAtType = "date-time"
	GetIgnoredEventBulkRetriesCompletedAtTypeGetIgnoredEventBulkRetriesCompletedAt2 GetIgnoredEventBulkRetriesCompletedAtType = "getIgnoredEventBulkRetriesCompletedAt_2"
)

type GetIgnoredEventBulkRetriesCompletedAt struct {
	DateTime                               *time.Time
	GetIgnoredEventBulkRetriesCompletedAt2 *GetIgnoredEventBulkRetriesCompletedAt2

	Type GetIgnoredEventBulkRetriesCompletedAtType
}

func CreateGetIgnoredEventBulkRetriesCompletedAtDateTime(dateTime time.Time) GetIgnoredEventBulkRetriesCompletedAt {
	typ := GetIgnoredEventBulkRetriesCompletedAtTypeDateTime

	return GetIgnoredEventBulkRetriesCompletedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIgnoredEventBulkRetriesCompletedAtGetIgnoredEventBulkRetriesCompletedAt2(getIgnoredEventBulkRetriesCompletedAt2 GetIgnoredEventBulkRetriesCompletedAt2) GetIgnoredEventBulkRetriesCompletedAt {
	typ := GetIgnoredEventBulkRetriesCompletedAtTypeGetIgnoredEventBulkRetriesCompletedAt2

	return GetIgnoredEventBulkRetriesCompletedAt{
		GetIgnoredEventBulkRetriesCompletedAt2: &getIgnoredEventBulkRetriesCompletedAt2,
		Type:                                   typ,
	}
}

func (u *GetIgnoredEventBulkRetriesCompletedAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetIgnoredEventBulkRetriesCompletedAtTypeDateTime
		return nil
	}

	getIgnoredEventBulkRetriesCompletedAt2 := new(GetIgnoredEventBulkRetriesCompletedAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIgnoredEventBulkRetriesCompletedAt2); err == nil {
		u.GetIgnoredEventBulkRetriesCompletedAt2 = getIgnoredEventBulkRetriesCompletedAt2
		u.Type = GetIgnoredEventBulkRetriesCompletedAtTypeGetIgnoredEventBulkRetriesCompletedAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesCompletedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetIgnoredEventBulkRetriesCompletedAt2 != nil {
		return json.Marshal(u.GetIgnoredEventBulkRetriesCompletedAt2)
	}

	return nil, nil
}

// GetIgnoredEventBulkRetriesCreatedAt2 - Filter by date the bulk retry was created
type GetIgnoredEventBulkRetriesCreatedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetIgnoredEventBulkRetriesCreatedAtType string

const (
	GetIgnoredEventBulkRetriesCreatedAtTypeDateTime                             GetIgnoredEventBulkRetriesCreatedAtType = "date-time"
	GetIgnoredEventBulkRetriesCreatedAtTypeGetIgnoredEventBulkRetriesCreatedAt2 GetIgnoredEventBulkRetriesCreatedAtType = "getIgnoredEventBulkRetriesCreatedAt_2"
)

type GetIgnoredEventBulkRetriesCreatedAt struct {
	DateTime                             *time.Time
	GetIgnoredEventBulkRetriesCreatedAt2 *GetIgnoredEventBulkRetriesCreatedAt2

	Type GetIgnoredEventBulkRetriesCreatedAtType
}

func CreateGetIgnoredEventBulkRetriesCreatedAtDateTime(dateTime time.Time) GetIgnoredEventBulkRetriesCreatedAt {
	typ := GetIgnoredEventBulkRetriesCreatedAtTypeDateTime

	return GetIgnoredEventBulkRetriesCreatedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIgnoredEventBulkRetriesCreatedAtGetIgnoredEventBulkRetriesCreatedAt2(getIgnoredEventBulkRetriesCreatedAt2 GetIgnoredEventBulkRetriesCreatedAt2) GetIgnoredEventBulkRetriesCreatedAt {
	typ := GetIgnoredEventBulkRetriesCreatedAtTypeGetIgnoredEventBulkRetriesCreatedAt2

	return GetIgnoredEventBulkRetriesCreatedAt{
		GetIgnoredEventBulkRetriesCreatedAt2: &getIgnoredEventBulkRetriesCreatedAt2,
		Type:                                 typ,
	}
}

func (u *GetIgnoredEventBulkRetriesCreatedAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetIgnoredEventBulkRetriesCreatedAtTypeDateTime
		return nil
	}

	getIgnoredEventBulkRetriesCreatedAt2 := new(GetIgnoredEventBulkRetriesCreatedAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIgnoredEventBulkRetriesCreatedAt2); err == nil {
		u.GetIgnoredEventBulkRetriesCreatedAt2 = getIgnoredEventBulkRetriesCreatedAt2
		u.Type = GetIgnoredEventBulkRetriesCreatedAtTypeGetIgnoredEventBulkRetriesCreatedAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesCreatedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetIgnoredEventBulkRetriesCreatedAt2 != nil {
		return json.Marshal(u.GetIgnoredEventBulkRetriesCreatedAt2)
	}

	return nil, nil
}

type GetIgnoredEventBulkRetriesDir2 string

const (
	GetIgnoredEventBulkRetriesDir2Asc  GetIgnoredEventBulkRetriesDir2 = "asc"
	GetIgnoredEventBulkRetriesDir2Desc GetIgnoredEventBulkRetriesDir2 = "desc"
)

func (e GetIgnoredEventBulkRetriesDir2) ToPointer() *GetIgnoredEventBulkRetriesDir2 {
	return &e
}

func (e *GetIgnoredEventBulkRetriesDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIgnoredEventBulkRetriesDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIgnoredEventBulkRetriesDir2: %v", v)
	}
}

// GetIgnoredEventBulkRetriesDir1 - Sort direction
type GetIgnoredEventBulkRetriesDir1 string

const (
	GetIgnoredEventBulkRetriesDir1Asc  GetIgnoredEventBulkRetriesDir1 = "asc"
	GetIgnoredEventBulkRetriesDir1Desc GetIgnoredEventBulkRetriesDir1 = "desc"
)

func (e GetIgnoredEventBulkRetriesDir1) ToPointer() *GetIgnoredEventBulkRetriesDir1 {
	return &e
}

func (e *GetIgnoredEventBulkRetriesDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIgnoredEventBulkRetriesDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIgnoredEventBulkRetriesDir1: %v", v)
	}
}

type GetIgnoredEventBulkRetriesDirType string

const (
	GetIgnoredEventBulkRetriesDirTypeGetIgnoredEventBulkRetriesDir1        GetIgnoredEventBulkRetriesDirType = "getIgnoredEventBulkRetriesDir_1"
	GetIgnoredEventBulkRetriesDirTypeArrayOfgetIgnoredEventBulkRetriesDir2 GetIgnoredEventBulkRetriesDirType = "arrayOfgetIgnoredEventBulkRetriesDir_2"
)

type GetIgnoredEventBulkRetriesDir struct {
	GetIgnoredEventBulkRetriesDir1        *GetIgnoredEventBulkRetriesDir1
	ArrayOfgetIgnoredEventBulkRetriesDir2 []GetIgnoredEventBulkRetriesDir2

	Type GetIgnoredEventBulkRetriesDirType
}

func CreateGetIgnoredEventBulkRetriesDirGetIgnoredEventBulkRetriesDir1(getIgnoredEventBulkRetriesDir1 GetIgnoredEventBulkRetriesDir1) GetIgnoredEventBulkRetriesDir {
	typ := GetIgnoredEventBulkRetriesDirTypeGetIgnoredEventBulkRetriesDir1

	return GetIgnoredEventBulkRetriesDir{
		GetIgnoredEventBulkRetriesDir1: &getIgnoredEventBulkRetriesDir1,
		Type:                           typ,
	}
}

func CreateGetIgnoredEventBulkRetriesDirArrayOfgetIgnoredEventBulkRetriesDir2(arrayOfgetIgnoredEventBulkRetriesDir2 []GetIgnoredEventBulkRetriesDir2) GetIgnoredEventBulkRetriesDir {
	typ := GetIgnoredEventBulkRetriesDirTypeArrayOfgetIgnoredEventBulkRetriesDir2

	return GetIgnoredEventBulkRetriesDir{
		ArrayOfgetIgnoredEventBulkRetriesDir2: arrayOfgetIgnoredEventBulkRetriesDir2,
		Type:                                  typ,
	}
}

func (u *GetIgnoredEventBulkRetriesDir) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getIgnoredEventBulkRetriesDir1 := new(GetIgnoredEventBulkRetriesDir1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIgnoredEventBulkRetriesDir1); err == nil {
		u.GetIgnoredEventBulkRetriesDir1 = getIgnoredEventBulkRetriesDir1
		u.Type = GetIgnoredEventBulkRetriesDirTypeGetIgnoredEventBulkRetriesDir1
		return nil
	}

	arrayOfgetIgnoredEventBulkRetriesDir2 := []GetIgnoredEventBulkRetriesDir2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetIgnoredEventBulkRetriesDir2); err == nil {
		u.ArrayOfgetIgnoredEventBulkRetriesDir2 = arrayOfgetIgnoredEventBulkRetriesDir2
		u.Type = GetIgnoredEventBulkRetriesDirTypeArrayOfgetIgnoredEventBulkRetriesDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesDir) MarshalJSON() ([]byte, error) {
	if u.GetIgnoredEventBulkRetriesDir1 != nil {
		return json.Marshal(u.GetIgnoredEventBulkRetriesDir1)
	}

	if u.ArrayOfgetIgnoredEventBulkRetriesDir2 != nil {
		return json.Marshal(u.ArrayOfgetIgnoredEventBulkRetriesDir2)
	}

	return nil, nil
}

type GetIgnoredEventBulkRetriesIDType string

const (
	GetIgnoredEventBulkRetriesIDTypeStr        GetIgnoredEventBulkRetriesIDType = "str"
	GetIgnoredEventBulkRetriesIDTypeArrayOfstr GetIgnoredEventBulkRetriesIDType = "arrayOfstr"
)

type GetIgnoredEventBulkRetriesID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIgnoredEventBulkRetriesIDType
}

func CreateGetIgnoredEventBulkRetriesIDStr(str string) GetIgnoredEventBulkRetriesID {
	typ := GetIgnoredEventBulkRetriesIDTypeStr

	return GetIgnoredEventBulkRetriesID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIgnoredEventBulkRetriesIDArrayOfstr(arrayOfstr []string) GetIgnoredEventBulkRetriesID {
	typ := GetIgnoredEventBulkRetriesIDTypeArrayOfstr

	return GetIgnoredEventBulkRetriesID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIgnoredEventBulkRetriesID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetIgnoredEventBulkRetriesIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIgnoredEventBulkRetriesIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

type GetIgnoredEventBulkRetriesOrderBy2 string

const (
	GetIgnoredEventBulkRetriesOrderBy2CreatedAt GetIgnoredEventBulkRetriesOrderBy2 = "created_at"
)

func (e GetIgnoredEventBulkRetriesOrderBy2) ToPointer() *GetIgnoredEventBulkRetriesOrderBy2 {
	return &e
}

func (e *GetIgnoredEventBulkRetriesOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetIgnoredEventBulkRetriesOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIgnoredEventBulkRetriesOrderBy2: %v", v)
	}
}

// GetIgnoredEventBulkRetriesOrderBy1 - Sort key(s)
type GetIgnoredEventBulkRetriesOrderBy1 string

const (
	GetIgnoredEventBulkRetriesOrderBy1CreatedAt GetIgnoredEventBulkRetriesOrderBy1 = "created_at"
)

func (e GetIgnoredEventBulkRetriesOrderBy1) ToPointer() *GetIgnoredEventBulkRetriesOrderBy1 {
	return &e
}

func (e *GetIgnoredEventBulkRetriesOrderBy1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetIgnoredEventBulkRetriesOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIgnoredEventBulkRetriesOrderBy1: %v", v)
	}
}

type GetIgnoredEventBulkRetriesOrderByType string

const (
	GetIgnoredEventBulkRetriesOrderByTypeGetIgnoredEventBulkRetriesOrderBy1        GetIgnoredEventBulkRetriesOrderByType = "getIgnoredEventBulkRetriesOrderBy_1"
	GetIgnoredEventBulkRetriesOrderByTypeArrayOfgetIgnoredEventBulkRetriesOrderBy2 GetIgnoredEventBulkRetriesOrderByType = "arrayOfgetIgnoredEventBulkRetriesOrderBy_2"
)

type GetIgnoredEventBulkRetriesOrderBy struct {
	GetIgnoredEventBulkRetriesOrderBy1        *GetIgnoredEventBulkRetriesOrderBy1
	ArrayOfgetIgnoredEventBulkRetriesOrderBy2 []GetIgnoredEventBulkRetriesOrderBy2

	Type GetIgnoredEventBulkRetriesOrderByType
}

func CreateGetIgnoredEventBulkRetriesOrderByGetIgnoredEventBulkRetriesOrderBy1(getIgnoredEventBulkRetriesOrderBy1 GetIgnoredEventBulkRetriesOrderBy1) GetIgnoredEventBulkRetriesOrderBy {
	typ := GetIgnoredEventBulkRetriesOrderByTypeGetIgnoredEventBulkRetriesOrderBy1

	return GetIgnoredEventBulkRetriesOrderBy{
		GetIgnoredEventBulkRetriesOrderBy1: &getIgnoredEventBulkRetriesOrderBy1,
		Type:                               typ,
	}
}

func CreateGetIgnoredEventBulkRetriesOrderByArrayOfgetIgnoredEventBulkRetriesOrderBy2(arrayOfgetIgnoredEventBulkRetriesOrderBy2 []GetIgnoredEventBulkRetriesOrderBy2) GetIgnoredEventBulkRetriesOrderBy {
	typ := GetIgnoredEventBulkRetriesOrderByTypeArrayOfgetIgnoredEventBulkRetriesOrderBy2

	return GetIgnoredEventBulkRetriesOrderBy{
		ArrayOfgetIgnoredEventBulkRetriesOrderBy2: arrayOfgetIgnoredEventBulkRetriesOrderBy2,
		Type: typ,
	}
}

func (u *GetIgnoredEventBulkRetriesOrderBy) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getIgnoredEventBulkRetriesOrderBy1 := new(GetIgnoredEventBulkRetriesOrderBy1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIgnoredEventBulkRetriesOrderBy1); err == nil {
		u.GetIgnoredEventBulkRetriesOrderBy1 = getIgnoredEventBulkRetriesOrderBy1
		u.Type = GetIgnoredEventBulkRetriesOrderByTypeGetIgnoredEventBulkRetriesOrderBy1
		return nil
	}

	arrayOfgetIgnoredEventBulkRetriesOrderBy2 := []GetIgnoredEventBulkRetriesOrderBy2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetIgnoredEventBulkRetriesOrderBy2); err == nil {
		u.ArrayOfgetIgnoredEventBulkRetriesOrderBy2 = arrayOfgetIgnoredEventBulkRetriesOrderBy2
		u.Type = GetIgnoredEventBulkRetriesOrderByTypeArrayOfgetIgnoredEventBulkRetriesOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetIgnoredEventBulkRetriesOrderBy1 != nil {
		return json.Marshal(u.GetIgnoredEventBulkRetriesOrderBy1)
	}

	if u.ArrayOfgetIgnoredEventBulkRetriesOrderBy2 != nil {
		return json.Marshal(u.ArrayOfgetIgnoredEventBulkRetriesOrderBy2)
	}

	return nil, nil
}

type GetIgnoredEventBulkRetriesQueryCauseType string

const (
	GetIgnoredEventBulkRetriesQueryCauseTypeStr        GetIgnoredEventBulkRetriesQueryCauseType = "str"
	GetIgnoredEventBulkRetriesQueryCauseTypeArrayOfstr GetIgnoredEventBulkRetriesQueryCauseType = "arrayOfstr"
)

type GetIgnoredEventBulkRetriesQueryCause struct {
	Str        *string
	ArrayOfstr []string

	Type GetIgnoredEventBulkRetriesQueryCauseType
}

func CreateGetIgnoredEventBulkRetriesQueryCauseStr(str string) GetIgnoredEventBulkRetriesQueryCause {
	typ := GetIgnoredEventBulkRetriesQueryCauseTypeStr

	return GetIgnoredEventBulkRetriesQueryCause{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIgnoredEventBulkRetriesQueryCauseArrayOfstr(arrayOfstr []string) GetIgnoredEventBulkRetriesQueryCause {
	typ := GetIgnoredEventBulkRetriesQueryCauseTypeArrayOfstr

	return GetIgnoredEventBulkRetriesQueryCause{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIgnoredEventBulkRetriesQueryCause) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetIgnoredEventBulkRetriesQueryCauseTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIgnoredEventBulkRetriesQueryCauseTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesQueryCause) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

type GetIgnoredEventBulkRetriesQueryWebhookIDType string

const (
	GetIgnoredEventBulkRetriesQueryWebhookIDTypeStr        GetIgnoredEventBulkRetriesQueryWebhookIDType = "str"
	GetIgnoredEventBulkRetriesQueryWebhookIDTypeArrayOfstr GetIgnoredEventBulkRetriesQueryWebhookIDType = "arrayOfstr"
)

type GetIgnoredEventBulkRetriesQueryWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIgnoredEventBulkRetriesQueryWebhookIDType
}

func CreateGetIgnoredEventBulkRetriesQueryWebhookIDStr(str string) GetIgnoredEventBulkRetriesQueryWebhookID {
	typ := GetIgnoredEventBulkRetriesQueryWebhookIDTypeStr

	return GetIgnoredEventBulkRetriesQueryWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIgnoredEventBulkRetriesQueryWebhookIDArrayOfstr(arrayOfstr []string) GetIgnoredEventBulkRetriesQueryWebhookID {
	typ := GetIgnoredEventBulkRetriesQueryWebhookIDTypeArrayOfstr

	return GetIgnoredEventBulkRetriesQueryWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIgnoredEventBulkRetriesQueryWebhookID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetIgnoredEventBulkRetriesQueryWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIgnoredEventBulkRetriesQueryWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIgnoredEventBulkRetriesQueryWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// GetIgnoredEventBulkRetriesQuery - Filter by the bulk retry ignored event query object
type GetIgnoredEventBulkRetriesQuery struct {
	// The cause of the ignored event
	Cause *GetIgnoredEventBulkRetriesQueryCause `queryParam:"name=cause"`
	// The associated transformation ID (only applicable to the cause `TRANSFORMATION_FAILED`)
	TransformationID *string `queryParam:"name=transformation_id"`
	// Connection ID of the ignored event
	WebhookID *GetIgnoredEventBulkRetriesQueryWebhookID `queryParam:"name=webhook_id"`
}

type GetIgnoredEventBulkRetriesRequest struct {
	CancelledAt *GetIgnoredEventBulkRetriesCancelledAt `queryParam:"style=form,explode=true,name=cancelled_at"`
	CompletedAt *GetIgnoredEventBulkRetriesCompletedAt `queryParam:"style=form,explode=true,name=completed_at"`
	CreatedAt   *GetIgnoredEventBulkRetriesCreatedAt   `queryParam:"style=form,explode=true,name=created_at"`
	Dir         *GetIgnoredEventBulkRetriesDir         `queryParam:"style=form,explode=true,name=dir"`
	ID          *GetIgnoredEventBulkRetriesID          `queryParam:"style=form,explode=true,name=id"`
	InProgress  *bool                                  `queryParam:"style=form,explode=true,name=in_progress"`
	Limit       *int64                                 `queryParam:"style=form,explode=true,name=limit"`
	Next        *string                                `queryParam:"style=form,explode=true,name=next"`
	OrderBy     *GetIgnoredEventBulkRetriesOrderBy     `queryParam:"style=form,explode=true,name=order_by"`
	Prev        *string                                `queryParam:"style=form,explode=true,name=prev"`
	// Filter by the bulk retry ignored event query object
	Query             *GetIgnoredEventBulkRetriesQuery `queryParam:"style=form,explode=true,name=query"`
	QueryPartialMatch *bool                            `queryParam:"style=form,explode=true,name=query_partial_match"`
}

type GetIgnoredEventBulkRetriesResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	// List of ignored events bulk retries
	BatchOperationPaginatedResult *shared.BatchOperationPaginatedResult
	ContentType                   string
	StatusCode                    int
	RawResponse                   *http.Response
}

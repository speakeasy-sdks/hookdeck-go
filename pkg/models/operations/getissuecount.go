// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hookdeck-go/pkg/models/shared"
	"net/http"
	"time"
)

type GetIssueCountAggregationKeysErrorCodeType string

const (
	GetIssueCountAggregationKeysErrorCodeTypeAttemptErrorCode        GetIssueCountAggregationKeysErrorCodeType = "AttemptErrorCode"
	GetIssueCountAggregationKeysErrorCodeTypeArrayOfAttemptErrorCode GetIssueCountAggregationKeysErrorCodeType = "arrayOfAttemptErrorCode"
)

type GetIssueCountAggregationKeysErrorCode struct {
	AttemptErrorCode        *shared.AttemptErrorCode
	ArrayOfAttemptErrorCode []shared.AttemptErrorCode

	Type GetIssueCountAggregationKeysErrorCodeType
}

func CreateGetIssueCountAggregationKeysErrorCodeAttemptErrorCode(attemptErrorCode shared.AttemptErrorCode) GetIssueCountAggregationKeysErrorCode {
	typ := GetIssueCountAggregationKeysErrorCodeTypeAttemptErrorCode

	return GetIssueCountAggregationKeysErrorCode{
		AttemptErrorCode: &attemptErrorCode,
		Type:             typ,
	}
}

func CreateGetIssueCountAggregationKeysErrorCodeArrayOfAttemptErrorCode(arrayOfAttemptErrorCode []shared.AttemptErrorCode) GetIssueCountAggregationKeysErrorCode {
	typ := GetIssueCountAggregationKeysErrorCodeTypeArrayOfAttemptErrorCode

	return GetIssueCountAggregationKeysErrorCode{
		ArrayOfAttemptErrorCode: arrayOfAttemptErrorCode,
		Type:                    typ,
	}
}

func (u *GetIssueCountAggregationKeysErrorCode) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	attemptErrorCode := new(shared.AttemptErrorCode)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&attemptErrorCode); err == nil {
		u.AttemptErrorCode = attemptErrorCode
		u.Type = GetIssueCountAggregationKeysErrorCodeTypeAttemptErrorCode
		return nil
	}

	arrayOfAttemptErrorCode := []shared.AttemptErrorCode{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfAttemptErrorCode); err == nil {
		u.ArrayOfAttemptErrorCode = arrayOfAttemptErrorCode
		u.Type = GetIssueCountAggregationKeysErrorCodeTypeArrayOfAttemptErrorCode
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountAggregationKeysErrorCode) MarshalJSON() ([]byte, error) {
	if u.AttemptErrorCode != nil {
		return json.Marshal(u.AttemptErrorCode)
	}

	if u.ArrayOfAttemptErrorCode != nil {
		return json.Marshal(u.ArrayOfAttemptErrorCode)
	}

	return nil, nil
}

type GetIssueCountAggregationKeysResponseStatusType string

const (
	GetIssueCountAggregationKeysResponseStatusTypeFloat32        GetIssueCountAggregationKeysResponseStatusType = "float32"
	GetIssueCountAggregationKeysResponseStatusTypeArrayOffloat32 GetIssueCountAggregationKeysResponseStatusType = "arrayOffloat32"
)

type GetIssueCountAggregationKeysResponseStatus struct {
	Float32        *float32
	ArrayOffloat32 []float32

	Type GetIssueCountAggregationKeysResponseStatusType
}

func CreateGetIssueCountAggregationKeysResponseStatusFloat32(float32T float32) GetIssueCountAggregationKeysResponseStatus {
	typ := GetIssueCountAggregationKeysResponseStatusTypeFloat32

	return GetIssueCountAggregationKeysResponseStatus{
		Float32: &float32T,
		Type:    typ,
	}
}

func CreateGetIssueCountAggregationKeysResponseStatusArrayOffloat32(arrayOffloat32 []float32) GetIssueCountAggregationKeysResponseStatus {
	typ := GetIssueCountAggregationKeysResponseStatusTypeArrayOffloat32

	return GetIssueCountAggregationKeysResponseStatus{
		ArrayOffloat32: arrayOffloat32,
		Type:           typ,
	}
}

func (u *GetIssueCountAggregationKeysResponseStatus) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	float32T := new(float32)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&float32T); err == nil {
		u.Float32 = float32T
		u.Type = GetIssueCountAggregationKeysResponseStatusTypeFloat32
		return nil
	}

	arrayOffloat32 := []float32{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOffloat32); err == nil {
		u.ArrayOffloat32 = arrayOffloat32
		u.Type = GetIssueCountAggregationKeysResponseStatusTypeArrayOffloat32
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountAggregationKeysResponseStatus) MarshalJSON() ([]byte, error) {
	if u.Float32 != nil {
		return json.Marshal(u.Float32)
	}

	if u.ArrayOffloat32 != nil {
		return json.Marshal(u.ArrayOffloat32)
	}

	return nil, nil
}

type GetIssueCountAggregationKeysWebhookIDType string

const (
	GetIssueCountAggregationKeysWebhookIDTypeStr        GetIssueCountAggregationKeysWebhookIDType = "str"
	GetIssueCountAggregationKeysWebhookIDTypeArrayOfstr GetIssueCountAggregationKeysWebhookIDType = "arrayOfstr"
)

type GetIssueCountAggregationKeysWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssueCountAggregationKeysWebhookIDType
}

func CreateGetIssueCountAggregationKeysWebhookIDStr(str string) GetIssueCountAggregationKeysWebhookID {
	typ := GetIssueCountAggregationKeysWebhookIDTypeStr

	return GetIssueCountAggregationKeysWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssueCountAggregationKeysWebhookIDArrayOfstr(arrayOfstr []string) GetIssueCountAggregationKeysWebhookID {
	typ := GetIssueCountAggregationKeysWebhookIDTypeArrayOfstr

	return GetIssueCountAggregationKeysWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssueCountAggregationKeysWebhookID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetIssueCountAggregationKeysWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssueCountAggregationKeysWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountAggregationKeysWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// GetIssueCountAggregationKeys - Filter by aggregation keys
type GetIssueCountAggregationKeys struct {
	ErrorCode      *GetIssueCountAggregationKeysErrorCode      `queryParam:"name=error_code"`
	ResponseStatus *GetIssueCountAggregationKeysResponseStatus `queryParam:"name=response_status"`
	WebhookID      *GetIssueCountAggregationKeysWebhookID      `queryParam:"name=webhook_id"`
}

// GetIssueCountCreatedAt2 - Filter by created dates
type GetIssueCountCreatedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetIssueCountCreatedAtType string

const (
	GetIssueCountCreatedAtTypeDateTime                GetIssueCountCreatedAtType = "date-time"
	GetIssueCountCreatedAtTypeGetIssueCountCreatedAt2 GetIssueCountCreatedAtType = "getIssueCountCreatedAt_2"
)

type GetIssueCountCreatedAt struct {
	DateTime                *time.Time
	GetIssueCountCreatedAt2 *GetIssueCountCreatedAt2

	Type GetIssueCountCreatedAtType
}

func CreateGetIssueCountCreatedAtDateTime(dateTime time.Time) GetIssueCountCreatedAt {
	typ := GetIssueCountCreatedAtTypeDateTime

	return GetIssueCountCreatedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssueCountCreatedAtGetIssueCountCreatedAt2(getIssueCountCreatedAt2 GetIssueCountCreatedAt2) GetIssueCountCreatedAt {
	typ := GetIssueCountCreatedAtTypeGetIssueCountCreatedAt2

	return GetIssueCountCreatedAt{
		GetIssueCountCreatedAt2: &getIssueCountCreatedAt2,
		Type:                    typ,
	}
}

func (u *GetIssueCountCreatedAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetIssueCountCreatedAtTypeDateTime
		return nil
	}

	getIssueCountCreatedAt2 := new(GetIssueCountCreatedAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountCreatedAt2); err == nil {
		u.GetIssueCountCreatedAt2 = getIssueCountCreatedAt2
		u.Type = GetIssueCountCreatedAtTypeGetIssueCountCreatedAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountCreatedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetIssueCountCreatedAt2 != nil {
		return json.Marshal(u.GetIssueCountCreatedAt2)
	}

	return nil, nil
}

type GetIssueCountDir2 string

const (
	GetIssueCountDir2Asc  GetIssueCountDir2 = "asc"
	GetIssueCountDir2Desc GetIssueCountDir2 = "desc"
)

func (e GetIssueCountDir2) ToPointer() *GetIssueCountDir2 {
	return &e
}

func (e *GetIssueCountDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIssueCountDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountDir2: %v", v)
	}
}

// GetIssueCountDir1 - Sort direction
type GetIssueCountDir1 string

const (
	GetIssueCountDir1Asc  GetIssueCountDir1 = "asc"
	GetIssueCountDir1Desc GetIssueCountDir1 = "desc"
)

func (e GetIssueCountDir1) ToPointer() *GetIssueCountDir1 {
	return &e
}

func (e *GetIssueCountDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIssueCountDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountDir1: %v", v)
	}
}

type GetIssueCountDirType string

const (
	GetIssueCountDirTypeGetIssueCountDir1        GetIssueCountDirType = "getIssueCountDir_1"
	GetIssueCountDirTypeArrayOfgetIssueCountDir2 GetIssueCountDirType = "arrayOfgetIssueCountDir_2"
)

type GetIssueCountDir struct {
	GetIssueCountDir1        *GetIssueCountDir1
	ArrayOfgetIssueCountDir2 []GetIssueCountDir2

	Type GetIssueCountDirType
}

func CreateGetIssueCountDirGetIssueCountDir1(getIssueCountDir1 GetIssueCountDir1) GetIssueCountDir {
	typ := GetIssueCountDirTypeGetIssueCountDir1

	return GetIssueCountDir{
		GetIssueCountDir1: &getIssueCountDir1,
		Type:              typ,
	}
}

func CreateGetIssueCountDirArrayOfgetIssueCountDir2(arrayOfgetIssueCountDir2 []GetIssueCountDir2) GetIssueCountDir {
	typ := GetIssueCountDirTypeArrayOfgetIssueCountDir2

	return GetIssueCountDir{
		ArrayOfgetIssueCountDir2: arrayOfgetIssueCountDir2,
		Type:                     typ,
	}
}

func (u *GetIssueCountDir) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getIssueCountDir1 := new(GetIssueCountDir1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountDir1); err == nil {
		u.GetIssueCountDir1 = getIssueCountDir1
		u.Type = GetIssueCountDirTypeGetIssueCountDir1
		return nil
	}

	arrayOfgetIssueCountDir2 := []GetIssueCountDir2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetIssueCountDir2); err == nil {
		u.ArrayOfgetIssueCountDir2 = arrayOfgetIssueCountDir2
		u.Type = GetIssueCountDirTypeArrayOfgetIssueCountDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountDir) MarshalJSON() ([]byte, error) {
	if u.GetIssueCountDir1 != nil {
		return json.Marshal(u.GetIssueCountDir1)
	}

	if u.ArrayOfgetIssueCountDir2 != nil {
		return json.Marshal(u.ArrayOfgetIssueCountDir2)
	}

	return nil, nil
}

// GetIssueCountDismissedAt2 - Filter by dismissed dates
type GetIssueCountDismissedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetIssueCountDismissedAtType string

const (
	GetIssueCountDismissedAtTypeDateTime                  GetIssueCountDismissedAtType = "date-time"
	GetIssueCountDismissedAtTypeGetIssueCountDismissedAt2 GetIssueCountDismissedAtType = "getIssueCountDismissedAt_2"
)

type GetIssueCountDismissedAt struct {
	DateTime                  *time.Time
	GetIssueCountDismissedAt2 *GetIssueCountDismissedAt2

	Type GetIssueCountDismissedAtType
}

func CreateGetIssueCountDismissedAtDateTime(dateTime time.Time) GetIssueCountDismissedAt {
	typ := GetIssueCountDismissedAtTypeDateTime

	return GetIssueCountDismissedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssueCountDismissedAtGetIssueCountDismissedAt2(getIssueCountDismissedAt2 GetIssueCountDismissedAt2) GetIssueCountDismissedAt {
	typ := GetIssueCountDismissedAtTypeGetIssueCountDismissedAt2

	return GetIssueCountDismissedAt{
		GetIssueCountDismissedAt2: &getIssueCountDismissedAt2,
		Type:                      typ,
	}
}

func (u *GetIssueCountDismissedAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetIssueCountDismissedAtTypeDateTime
		return nil
	}

	getIssueCountDismissedAt2 := new(GetIssueCountDismissedAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountDismissedAt2); err == nil {
		u.GetIssueCountDismissedAt2 = getIssueCountDismissedAt2
		u.Type = GetIssueCountDismissedAtTypeGetIssueCountDismissedAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountDismissedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetIssueCountDismissedAt2 != nil {
		return json.Marshal(u.GetIssueCountDismissedAt2)
	}

	return nil, nil
}

// GetIssueCountFirstSeenAt2 - Filter by first seen dates
type GetIssueCountFirstSeenAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetIssueCountFirstSeenAtType string

const (
	GetIssueCountFirstSeenAtTypeDateTime                  GetIssueCountFirstSeenAtType = "date-time"
	GetIssueCountFirstSeenAtTypeGetIssueCountFirstSeenAt2 GetIssueCountFirstSeenAtType = "getIssueCountFirstSeenAt_2"
)

type GetIssueCountFirstSeenAt struct {
	DateTime                  *time.Time
	GetIssueCountFirstSeenAt2 *GetIssueCountFirstSeenAt2

	Type GetIssueCountFirstSeenAtType
}

func CreateGetIssueCountFirstSeenAtDateTime(dateTime time.Time) GetIssueCountFirstSeenAt {
	typ := GetIssueCountFirstSeenAtTypeDateTime

	return GetIssueCountFirstSeenAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssueCountFirstSeenAtGetIssueCountFirstSeenAt2(getIssueCountFirstSeenAt2 GetIssueCountFirstSeenAt2) GetIssueCountFirstSeenAt {
	typ := GetIssueCountFirstSeenAtTypeGetIssueCountFirstSeenAt2

	return GetIssueCountFirstSeenAt{
		GetIssueCountFirstSeenAt2: &getIssueCountFirstSeenAt2,
		Type:                      typ,
	}
}

func (u *GetIssueCountFirstSeenAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetIssueCountFirstSeenAtTypeDateTime
		return nil
	}

	getIssueCountFirstSeenAt2 := new(GetIssueCountFirstSeenAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountFirstSeenAt2); err == nil {
		u.GetIssueCountFirstSeenAt2 = getIssueCountFirstSeenAt2
		u.Type = GetIssueCountFirstSeenAtTypeGetIssueCountFirstSeenAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountFirstSeenAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetIssueCountFirstSeenAt2 != nil {
		return json.Marshal(u.GetIssueCountFirstSeenAt2)
	}

	return nil, nil
}

type GetIssueCountIDType string

const (
	GetIssueCountIDTypeStr        GetIssueCountIDType = "str"
	GetIssueCountIDTypeArrayOfstr GetIssueCountIDType = "arrayOfstr"
)

type GetIssueCountID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssueCountIDType
}

func CreateGetIssueCountIDStr(str string) GetIssueCountID {
	typ := GetIssueCountIDTypeStr

	return GetIssueCountID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssueCountIDArrayOfstr(arrayOfstr []string) GetIssueCountID {
	typ := GetIssueCountIDTypeArrayOfstr

	return GetIssueCountID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssueCountID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetIssueCountIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssueCountIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

type GetIssueCountIssueTriggerIDType string

const (
	GetIssueCountIssueTriggerIDTypeStr        GetIssueCountIssueTriggerIDType = "str"
	GetIssueCountIssueTriggerIDTypeArrayOfstr GetIssueCountIssueTriggerIDType = "arrayOfstr"
)

type GetIssueCountIssueTriggerID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssueCountIssueTriggerIDType
}

func CreateGetIssueCountIssueTriggerIDStr(str string) GetIssueCountIssueTriggerID {
	typ := GetIssueCountIssueTriggerIDTypeStr

	return GetIssueCountIssueTriggerID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssueCountIssueTriggerIDArrayOfstr(arrayOfstr []string) GetIssueCountIssueTriggerID {
	typ := GetIssueCountIssueTriggerIDTypeArrayOfstr

	return GetIssueCountIssueTriggerID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssueCountIssueTriggerID) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetIssueCountIssueTriggerIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssueCountIssueTriggerIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountIssueTriggerID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// GetIssueCountLastSeenAt2 - Filter by last seen dates
type GetIssueCountLastSeenAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

type GetIssueCountLastSeenAtType string

const (
	GetIssueCountLastSeenAtTypeDateTime                 GetIssueCountLastSeenAtType = "date-time"
	GetIssueCountLastSeenAtTypeGetIssueCountLastSeenAt2 GetIssueCountLastSeenAtType = "getIssueCountLastSeenAt_2"
)

type GetIssueCountLastSeenAt struct {
	DateTime                 *time.Time
	GetIssueCountLastSeenAt2 *GetIssueCountLastSeenAt2

	Type GetIssueCountLastSeenAtType
}

func CreateGetIssueCountLastSeenAtDateTime(dateTime time.Time) GetIssueCountLastSeenAt {
	typ := GetIssueCountLastSeenAtTypeDateTime

	return GetIssueCountLastSeenAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssueCountLastSeenAtGetIssueCountLastSeenAt2(getIssueCountLastSeenAt2 GetIssueCountLastSeenAt2) GetIssueCountLastSeenAt {
	typ := GetIssueCountLastSeenAtTypeGetIssueCountLastSeenAt2

	return GetIssueCountLastSeenAt{
		GetIssueCountLastSeenAt2: &getIssueCountLastSeenAt2,
		Type:                     typ,
	}
}

func (u *GetIssueCountLastSeenAt) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = GetIssueCountLastSeenAtTypeDateTime
		return nil
	}

	getIssueCountLastSeenAt2 := new(GetIssueCountLastSeenAt2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountLastSeenAt2); err == nil {
		u.GetIssueCountLastSeenAt2 = getIssueCountLastSeenAt2
		u.Type = GetIssueCountLastSeenAtTypeGetIssueCountLastSeenAt2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountLastSeenAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.GetIssueCountLastSeenAt2 != nil {
		return json.Marshal(u.GetIssueCountLastSeenAt2)
	}

	return nil, nil
}

type GetIssueCountMergedWithType string

const (
	GetIssueCountMergedWithTypeStr        GetIssueCountMergedWithType = "str"
	GetIssueCountMergedWithTypeArrayOfstr GetIssueCountMergedWithType = "arrayOfstr"
)

type GetIssueCountMergedWith struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssueCountMergedWithType
}

func CreateGetIssueCountMergedWithStr(str string) GetIssueCountMergedWith {
	typ := GetIssueCountMergedWithTypeStr

	return GetIssueCountMergedWith{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssueCountMergedWithArrayOfstr(arrayOfstr []string) GetIssueCountMergedWith {
	typ := GetIssueCountMergedWithTypeArrayOfstr

	return GetIssueCountMergedWith{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssueCountMergedWith) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = GetIssueCountMergedWithTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssueCountMergedWithTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountMergedWith) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

type GetIssueCountOrderBy2 string

const (
	GetIssueCountOrderBy2CreatedAt   GetIssueCountOrderBy2 = "created_at"
	GetIssueCountOrderBy2FirstSeenAt GetIssueCountOrderBy2 = "first_seen_at"
	GetIssueCountOrderBy2LastSeenAt  GetIssueCountOrderBy2 = "last_seen_at"
	GetIssueCountOrderBy2OpenedAt    GetIssueCountOrderBy2 = "opened_at"
	GetIssueCountOrderBy2Status      GetIssueCountOrderBy2 = "status"
)

func (e GetIssueCountOrderBy2) ToPointer() *GetIssueCountOrderBy2 {
	return &e
}

func (e *GetIssueCountOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "first_seen_at":
		fallthrough
	case "last_seen_at":
		fallthrough
	case "opened_at":
		fallthrough
	case "status":
		*e = GetIssueCountOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountOrderBy2: %v", v)
	}
}

// GetIssueCountOrderBy1 - Sort key(s)
type GetIssueCountOrderBy1 string

const (
	GetIssueCountOrderBy1CreatedAt   GetIssueCountOrderBy1 = "created_at"
	GetIssueCountOrderBy1FirstSeenAt GetIssueCountOrderBy1 = "first_seen_at"
	GetIssueCountOrderBy1LastSeenAt  GetIssueCountOrderBy1 = "last_seen_at"
	GetIssueCountOrderBy1OpenedAt    GetIssueCountOrderBy1 = "opened_at"
	GetIssueCountOrderBy1Status      GetIssueCountOrderBy1 = "status"
)

func (e GetIssueCountOrderBy1) ToPointer() *GetIssueCountOrderBy1 {
	return &e
}

func (e *GetIssueCountOrderBy1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "first_seen_at":
		fallthrough
	case "last_seen_at":
		fallthrough
	case "opened_at":
		fallthrough
	case "status":
		*e = GetIssueCountOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountOrderBy1: %v", v)
	}
}

type GetIssueCountOrderByType string

const (
	GetIssueCountOrderByTypeGetIssueCountOrderBy1        GetIssueCountOrderByType = "getIssueCountOrderBy_1"
	GetIssueCountOrderByTypeArrayOfgetIssueCountOrderBy2 GetIssueCountOrderByType = "arrayOfgetIssueCountOrderBy_2"
)

type GetIssueCountOrderBy struct {
	GetIssueCountOrderBy1        *GetIssueCountOrderBy1
	ArrayOfgetIssueCountOrderBy2 []GetIssueCountOrderBy2

	Type GetIssueCountOrderByType
}

func CreateGetIssueCountOrderByGetIssueCountOrderBy1(getIssueCountOrderBy1 GetIssueCountOrderBy1) GetIssueCountOrderBy {
	typ := GetIssueCountOrderByTypeGetIssueCountOrderBy1

	return GetIssueCountOrderBy{
		GetIssueCountOrderBy1: &getIssueCountOrderBy1,
		Type:                  typ,
	}
}

func CreateGetIssueCountOrderByArrayOfgetIssueCountOrderBy2(arrayOfgetIssueCountOrderBy2 []GetIssueCountOrderBy2) GetIssueCountOrderBy {
	typ := GetIssueCountOrderByTypeArrayOfgetIssueCountOrderBy2

	return GetIssueCountOrderBy{
		ArrayOfgetIssueCountOrderBy2: arrayOfgetIssueCountOrderBy2,
		Type:                         typ,
	}
}

func (u *GetIssueCountOrderBy) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getIssueCountOrderBy1 := new(GetIssueCountOrderBy1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountOrderBy1); err == nil {
		u.GetIssueCountOrderBy1 = getIssueCountOrderBy1
		u.Type = GetIssueCountOrderByTypeGetIssueCountOrderBy1
		return nil
	}

	arrayOfgetIssueCountOrderBy2 := []GetIssueCountOrderBy2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetIssueCountOrderBy2); err == nil {
		u.ArrayOfgetIssueCountOrderBy2 = arrayOfgetIssueCountOrderBy2
		u.Type = GetIssueCountOrderByTypeArrayOfgetIssueCountOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetIssueCountOrderBy1 != nil {
		return json.Marshal(u.GetIssueCountOrderBy1)
	}

	if u.ArrayOfgetIssueCountOrderBy2 != nil {
		return json.Marshal(u.ArrayOfgetIssueCountOrderBy2)
	}

	return nil, nil
}

// GetIssueCountStatus2 - Issue status
type GetIssueCountStatus2 string

const (
	GetIssueCountStatus2Opened       GetIssueCountStatus2 = "OPENED"
	GetIssueCountStatus2Ignored      GetIssueCountStatus2 = "IGNORED"
	GetIssueCountStatus2Acknowledged GetIssueCountStatus2 = "ACKNOWLEDGED"
	GetIssueCountStatus2Resolved     GetIssueCountStatus2 = "RESOLVED"
)

func (e GetIssueCountStatus2) ToPointer() *GetIssueCountStatus2 {
	return &e
}

func (e *GetIssueCountStatus2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OPENED":
		fallthrough
	case "IGNORED":
		fallthrough
	case "ACKNOWLEDGED":
		fallthrough
	case "RESOLVED":
		*e = GetIssueCountStatus2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountStatus2: %v", v)
	}
}

// GetIssueCountStatus1 - Issue status
type GetIssueCountStatus1 string

const (
	GetIssueCountStatus1Opened       GetIssueCountStatus1 = "OPENED"
	GetIssueCountStatus1Ignored      GetIssueCountStatus1 = "IGNORED"
	GetIssueCountStatus1Acknowledged GetIssueCountStatus1 = "ACKNOWLEDGED"
	GetIssueCountStatus1Resolved     GetIssueCountStatus1 = "RESOLVED"
)

func (e GetIssueCountStatus1) ToPointer() *GetIssueCountStatus1 {
	return &e
}

func (e *GetIssueCountStatus1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OPENED":
		fallthrough
	case "IGNORED":
		fallthrough
	case "ACKNOWLEDGED":
		fallthrough
	case "RESOLVED":
		*e = GetIssueCountStatus1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountStatus1: %v", v)
	}
}

type GetIssueCountStatusType string

const (
	GetIssueCountStatusTypeGetIssueCountStatus1        GetIssueCountStatusType = "getIssueCountStatus_1"
	GetIssueCountStatusTypeArrayOfgetIssueCountStatus2 GetIssueCountStatusType = "arrayOfgetIssueCountStatus_2"
)

type GetIssueCountStatus struct {
	GetIssueCountStatus1        *GetIssueCountStatus1
	ArrayOfgetIssueCountStatus2 []GetIssueCountStatus2

	Type GetIssueCountStatusType
}

func CreateGetIssueCountStatusGetIssueCountStatus1(getIssueCountStatus1 GetIssueCountStatus1) GetIssueCountStatus {
	typ := GetIssueCountStatusTypeGetIssueCountStatus1

	return GetIssueCountStatus{
		GetIssueCountStatus1: &getIssueCountStatus1,
		Type:                 typ,
	}
}

func CreateGetIssueCountStatusArrayOfgetIssueCountStatus2(arrayOfgetIssueCountStatus2 []GetIssueCountStatus2) GetIssueCountStatus {
	typ := GetIssueCountStatusTypeArrayOfgetIssueCountStatus2

	return GetIssueCountStatus{
		ArrayOfgetIssueCountStatus2: arrayOfgetIssueCountStatus2,
		Type:                        typ,
	}
}

func (u *GetIssueCountStatus) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getIssueCountStatus1 := new(GetIssueCountStatus1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountStatus1); err == nil {
		u.GetIssueCountStatus1 = getIssueCountStatus1
		u.Type = GetIssueCountStatusTypeGetIssueCountStatus1
		return nil
	}

	arrayOfgetIssueCountStatus2 := []GetIssueCountStatus2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetIssueCountStatus2); err == nil {
		u.ArrayOfgetIssueCountStatus2 = arrayOfgetIssueCountStatus2
		u.Type = GetIssueCountStatusTypeArrayOfgetIssueCountStatus2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountStatus) MarshalJSON() ([]byte, error) {
	if u.GetIssueCountStatus1 != nil {
		return json.Marshal(u.GetIssueCountStatus1)
	}

	if u.ArrayOfgetIssueCountStatus2 != nil {
		return json.Marshal(u.ArrayOfgetIssueCountStatus2)
	}

	return nil, nil
}

// GetIssueCountType2 - Issue type
type GetIssueCountType2 string

const (
	GetIssueCountType2Delivery       GetIssueCountType2 = "delivery"
	GetIssueCountType2Transformation GetIssueCountType2 = "transformation"
	GetIssueCountType2Backpressure   GetIssueCountType2 = "backpressure"
)

func (e GetIssueCountType2) ToPointer() *GetIssueCountType2 {
	return &e
}

func (e *GetIssueCountType2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delivery":
		fallthrough
	case "transformation":
		fallthrough
	case "backpressure":
		*e = GetIssueCountType2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountType2: %v", v)
	}
}

// GetIssueCountType1 - Issue type
type GetIssueCountType1 string

const (
	GetIssueCountType1Delivery       GetIssueCountType1 = "delivery"
	GetIssueCountType1Transformation GetIssueCountType1 = "transformation"
	GetIssueCountType1Backpressure   GetIssueCountType1 = "backpressure"
)

func (e GetIssueCountType1) ToPointer() *GetIssueCountType1 {
	return &e
}

func (e *GetIssueCountType1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delivery":
		fallthrough
	case "transformation":
		fallthrough
	case "backpressure":
		*e = GetIssueCountType1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssueCountType1: %v", v)
	}
}

type GetIssueCountTypeType string

const (
	GetIssueCountTypeTypeGetIssueCountType1        GetIssueCountTypeType = "getIssueCountType_1"
	GetIssueCountTypeTypeArrayOfgetIssueCountType2 GetIssueCountTypeType = "arrayOfgetIssueCountType_2"
)

type GetIssueCountType struct {
	GetIssueCountType1        *GetIssueCountType1
	ArrayOfgetIssueCountType2 []GetIssueCountType2

	Type GetIssueCountTypeType
}

func CreateGetIssueCountTypeGetIssueCountType1(getIssueCountType1 GetIssueCountType1) GetIssueCountType {
	typ := GetIssueCountTypeTypeGetIssueCountType1

	return GetIssueCountType{
		GetIssueCountType1: &getIssueCountType1,
		Type:               typ,
	}
}

func CreateGetIssueCountTypeArrayOfgetIssueCountType2(arrayOfgetIssueCountType2 []GetIssueCountType2) GetIssueCountType {
	typ := GetIssueCountTypeTypeArrayOfgetIssueCountType2

	return GetIssueCountType{
		ArrayOfgetIssueCountType2: arrayOfgetIssueCountType2,
		Type:                      typ,
	}
}

func (u *GetIssueCountType) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	getIssueCountType1 := new(GetIssueCountType1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&getIssueCountType1); err == nil {
		u.GetIssueCountType1 = getIssueCountType1
		u.Type = GetIssueCountTypeTypeGetIssueCountType1
		return nil
	}

	arrayOfgetIssueCountType2 := []GetIssueCountType2{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfgetIssueCountType2); err == nil {
		u.ArrayOfgetIssueCountType2 = arrayOfgetIssueCountType2
		u.Type = GetIssueCountTypeTypeArrayOfgetIssueCountType2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssueCountType) MarshalJSON() ([]byte, error) {
	if u.GetIssueCountType1 != nil {
		return json.Marshal(u.GetIssueCountType1)
	}

	if u.ArrayOfgetIssueCountType2 != nil {
		return json.Marshal(u.ArrayOfgetIssueCountType2)
	}

	return nil, nil
}

type GetIssueCountRequest struct {
	// Filter by aggregation keys
	AggregationKeys *GetIssueCountAggregationKeys `queryParam:"style=form,explode=true,name=aggregation_keys"`
	CreatedAt       *GetIssueCountCreatedAt       `queryParam:"style=form,explode=true,name=created_at"`
	Dir             *GetIssueCountDir             `queryParam:"style=form,explode=true,name=dir"`
	DismissedAt     *GetIssueCountDismissedAt     `queryParam:"style=form,explode=true,name=dismissed_at"`
	FirstSeenAt     *GetIssueCountFirstSeenAt     `queryParam:"style=form,explode=true,name=first_seen_at"`
	ID              *GetIssueCountID              `queryParam:"style=form,explode=true,name=id"`
	IssueTriggerID  *GetIssueCountIssueTriggerID  `queryParam:"style=form,explode=true,name=issue_trigger_id"`
	LastSeenAt      *GetIssueCountLastSeenAt      `queryParam:"style=form,explode=true,name=last_seen_at"`
	Limit           *int64                        `queryParam:"style=form,explode=true,name=limit"`
	MergedWith      *GetIssueCountMergedWith      `queryParam:"style=form,explode=true,name=merged_with"`
	Next            *string                       `queryParam:"style=form,explode=true,name=next"`
	OrderBy         *GetIssueCountOrderBy         `queryParam:"style=form,explode=true,name=order_by"`
	Prev            *string                       `queryParam:"style=form,explode=true,name=prev"`
	Status          *GetIssueCountStatus          `queryParam:"style=form,explode=true,name=status"`
	Type            *GetIssueCountType            `queryParam:"style=form,explode=true,name=type"`
}

type GetIssueCountResponse struct {
	// Unprocessable Entity
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// Issue count
	IssueCount  *shared.IssueCount
	StatusCode  int
	RawResponse *http.Response
}

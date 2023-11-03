// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
	"net/http"
	"time"
)

type GetIssuesAggregationKeysErrorCodeType string

const (
	GetIssuesAggregationKeysErrorCodeTypeAttemptErrorCode        GetIssuesAggregationKeysErrorCodeType = "AttemptErrorCode"
	GetIssuesAggregationKeysErrorCodeTypeArrayOfAttemptErrorCode GetIssuesAggregationKeysErrorCodeType = "arrayOfAttemptErrorCode"
)

type GetIssuesAggregationKeysErrorCode struct {
	AttemptErrorCode        *shared.AttemptErrorCode
	ArrayOfAttemptErrorCode []shared.AttemptErrorCode

	Type GetIssuesAggregationKeysErrorCodeType
}

func CreateGetIssuesAggregationKeysErrorCodeAttemptErrorCode(attemptErrorCode shared.AttemptErrorCode) GetIssuesAggregationKeysErrorCode {
	typ := GetIssuesAggregationKeysErrorCodeTypeAttemptErrorCode

	return GetIssuesAggregationKeysErrorCode{
		AttemptErrorCode: &attemptErrorCode,
		Type:             typ,
	}
}

func CreateGetIssuesAggregationKeysErrorCodeArrayOfAttemptErrorCode(arrayOfAttemptErrorCode []shared.AttemptErrorCode) GetIssuesAggregationKeysErrorCode {
	typ := GetIssuesAggregationKeysErrorCodeTypeArrayOfAttemptErrorCode

	return GetIssuesAggregationKeysErrorCode{
		ArrayOfAttemptErrorCode: arrayOfAttemptErrorCode,
		Type:                    typ,
	}
}

func (u *GetIssuesAggregationKeysErrorCode) UnmarshalJSON(data []byte) error {

	attemptErrorCode := shared.AttemptErrorCode("")
	if err := utils.UnmarshalJSON(data, &attemptErrorCode, "", true, true); err == nil {
		u.AttemptErrorCode = &attemptErrorCode
		u.Type = GetIssuesAggregationKeysErrorCodeTypeAttemptErrorCode
		return nil
	}

	arrayOfAttemptErrorCode := []shared.AttemptErrorCode{}
	if err := utils.UnmarshalJSON(data, &arrayOfAttemptErrorCode, "", true, true); err == nil {
		u.ArrayOfAttemptErrorCode = arrayOfAttemptErrorCode
		u.Type = GetIssuesAggregationKeysErrorCodeTypeArrayOfAttemptErrorCode
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesAggregationKeysErrorCode) MarshalJSON() ([]byte, error) {
	if u.AttemptErrorCode != nil {
		return utils.MarshalJSON(u.AttemptErrorCode, "", true)
	}

	if u.ArrayOfAttemptErrorCode != nil {
		return utils.MarshalJSON(u.ArrayOfAttemptErrorCode, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesAggregationKeysResponseStatusType string

const (
	GetIssuesAggregationKeysResponseStatusTypeFloat32        GetIssuesAggregationKeysResponseStatusType = "float32"
	GetIssuesAggregationKeysResponseStatusTypeArrayOffloat32 GetIssuesAggregationKeysResponseStatusType = "arrayOffloat32"
)

type GetIssuesAggregationKeysResponseStatus struct {
	Float32        *float32
	ArrayOffloat32 []float32

	Type GetIssuesAggregationKeysResponseStatusType
}

func CreateGetIssuesAggregationKeysResponseStatusFloat32(float32T float32) GetIssuesAggregationKeysResponseStatus {
	typ := GetIssuesAggregationKeysResponseStatusTypeFloat32

	return GetIssuesAggregationKeysResponseStatus{
		Float32: &float32T,
		Type:    typ,
	}
}

func CreateGetIssuesAggregationKeysResponseStatusArrayOffloat32(arrayOffloat32 []float32) GetIssuesAggregationKeysResponseStatus {
	typ := GetIssuesAggregationKeysResponseStatusTypeArrayOffloat32

	return GetIssuesAggregationKeysResponseStatus{
		ArrayOffloat32: arrayOffloat32,
		Type:           typ,
	}
}

func (u *GetIssuesAggregationKeysResponseStatus) UnmarshalJSON(data []byte) error {

	float32Var := float32(0)
	if err := utils.UnmarshalJSON(data, &float32Var, "", true, true); err == nil {
		u.Float32 = &float32Var
		u.Type = GetIssuesAggregationKeysResponseStatusTypeFloat32
		return nil
	}

	arrayOffloat32 := []float32{}
	if err := utils.UnmarshalJSON(data, &arrayOffloat32, "", true, true); err == nil {
		u.ArrayOffloat32 = arrayOffloat32
		u.Type = GetIssuesAggregationKeysResponseStatusTypeArrayOffloat32
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesAggregationKeysResponseStatus) MarshalJSON() ([]byte, error) {
	if u.Float32 != nil {
		return utils.MarshalJSON(u.Float32, "", true)
	}

	if u.ArrayOffloat32 != nil {
		return utils.MarshalJSON(u.ArrayOffloat32, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesAggregationKeysWebhookIDType string

const (
	GetIssuesAggregationKeysWebhookIDTypeStr        GetIssuesAggregationKeysWebhookIDType = "str"
	GetIssuesAggregationKeysWebhookIDTypeArrayOfstr GetIssuesAggregationKeysWebhookIDType = "arrayOfstr"
)

type GetIssuesAggregationKeysWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssuesAggregationKeysWebhookIDType
}

func CreateGetIssuesAggregationKeysWebhookIDStr(str string) GetIssuesAggregationKeysWebhookID {
	typ := GetIssuesAggregationKeysWebhookIDTypeStr

	return GetIssuesAggregationKeysWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssuesAggregationKeysWebhookIDArrayOfstr(arrayOfstr []string) GetIssuesAggregationKeysWebhookID {
	typ := GetIssuesAggregationKeysWebhookIDTypeArrayOfstr

	return GetIssuesAggregationKeysWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssuesAggregationKeysWebhookID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetIssuesAggregationKeysWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssuesAggregationKeysWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesAggregationKeysWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetIssuesAggregationKeys - Filter by aggregation keys
type GetIssuesAggregationKeys struct {
	ErrorCode      *GetIssuesAggregationKeysErrorCode      `queryParam:"name=error_code"`
	ResponseStatus *GetIssuesAggregationKeysResponseStatus `queryParam:"name=response_status"`
	WebhookID      *GetIssuesAggregationKeysWebhookID      `queryParam:"name=webhook_id"`
}

func (o *GetIssuesAggregationKeys) GetErrorCode() *GetIssuesAggregationKeysErrorCode {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *GetIssuesAggregationKeys) GetResponseStatus() *GetIssuesAggregationKeysResponseStatus {
	if o == nil {
		return nil
	}
	return o.ResponseStatus
}

func (o *GetIssuesAggregationKeys) GetWebhookID() *GetIssuesAggregationKeysWebhookID {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

// GetIssuesCreatedAt2 - Filter by created dates
type GetIssuesCreatedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetIssuesCreatedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIssuesCreatedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetIssuesCreatedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetIssuesCreatedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetIssuesCreatedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetIssuesCreatedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetIssuesCreatedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetIssuesCreatedAtType string

const (
	GetIssuesCreatedAtTypeDateTime            GetIssuesCreatedAtType = "date-time"
	GetIssuesCreatedAtTypeGetIssuesCreatedAt2 GetIssuesCreatedAtType = "getIssuesCreatedAt_2"
)

type GetIssuesCreatedAt struct {
	DateTime            *time.Time
	GetIssuesCreatedAt2 *GetIssuesCreatedAt2

	Type GetIssuesCreatedAtType
}

func CreateGetIssuesCreatedAtDateTime(dateTime time.Time) GetIssuesCreatedAt {
	typ := GetIssuesCreatedAtTypeDateTime

	return GetIssuesCreatedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssuesCreatedAtGetIssuesCreatedAt2(getIssuesCreatedAt2 GetIssuesCreatedAt2) GetIssuesCreatedAt {
	typ := GetIssuesCreatedAtTypeGetIssuesCreatedAt2

	return GetIssuesCreatedAt{
		GetIssuesCreatedAt2: &getIssuesCreatedAt2,
		Type:                typ,
	}
}

func (u *GetIssuesCreatedAt) UnmarshalJSON(data []byte) error {

	getIssuesCreatedAt2 := GetIssuesCreatedAt2{}
	if err := utils.UnmarshalJSON(data, &getIssuesCreatedAt2, "", true, true); err == nil {
		u.GetIssuesCreatedAt2 = &getIssuesCreatedAt2
		u.Type = GetIssuesCreatedAtTypeGetIssuesCreatedAt2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetIssuesCreatedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesCreatedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetIssuesCreatedAt2 != nil {
		return utils.MarshalJSON(u.GetIssuesCreatedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesDir2 string

const (
	GetIssuesDir2Asc  GetIssuesDir2 = "asc"
	GetIssuesDir2Desc GetIssuesDir2 = "desc"
)

func (e GetIssuesDir2) ToPointer() *GetIssuesDir2 {
	return &e
}

func (e *GetIssuesDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIssuesDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesDir2: %v", v)
	}
}

// GetIssuesDir1 - Sort direction
type GetIssuesDir1 string

const (
	GetIssuesDir1Asc  GetIssuesDir1 = "asc"
	GetIssuesDir1Desc GetIssuesDir1 = "desc"
)

func (e GetIssuesDir1) ToPointer() *GetIssuesDir1 {
	return &e
}

func (e *GetIssuesDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetIssuesDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesDir1: %v", v)
	}
}

type GetIssuesDirType string

const (
	GetIssuesDirTypeGetIssuesDir1        GetIssuesDirType = "getIssuesDir_1"
	GetIssuesDirTypeArrayOfgetIssuesDir2 GetIssuesDirType = "arrayOfgetIssuesDir_2"
)

type GetIssuesDir struct {
	GetIssuesDir1        *GetIssuesDir1
	ArrayOfgetIssuesDir2 []GetIssuesDir2

	Type GetIssuesDirType
}

func CreateGetIssuesDirGetIssuesDir1(getIssuesDir1 GetIssuesDir1) GetIssuesDir {
	typ := GetIssuesDirTypeGetIssuesDir1

	return GetIssuesDir{
		GetIssuesDir1: &getIssuesDir1,
		Type:          typ,
	}
}

func CreateGetIssuesDirArrayOfgetIssuesDir2(arrayOfgetIssuesDir2 []GetIssuesDir2) GetIssuesDir {
	typ := GetIssuesDirTypeArrayOfgetIssuesDir2

	return GetIssuesDir{
		ArrayOfgetIssuesDir2: arrayOfgetIssuesDir2,
		Type:                 typ,
	}
}

func (u *GetIssuesDir) UnmarshalJSON(data []byte) error {

	getIssuesDir1 := GetIssuesDir1("")
	if err := utils.UnmarshalJSON(data, &getIssuesDir1, "", true, true); err == nil {
		u.GetIssuesDir1 = &getIssuesDir1
		u.Type = GetIssuesDirTypeGetIssuesDir1
		return nil
	}

	arrayOfgetIssuesDir2 := []GetIssuesDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetIssuesDir2, "", true, true); err == nil {
		u.ArrayOfgetIssuesDir2 = arrayOfgetIssuesDir2
		u.Type = GetIssuesDirTypeArrayOfgetIssuesDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesDir) MarshalJSON() ([]byte, error) {
	if u.GetIssuesDir1 != nil {
		return utils.MarshalJSON(u.GetIssuesDir1, "", true)
	}

	if u.ArrayOfgetIssuesDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetIssuesDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetIssuesDismissedAt2 - Filter by dismissed dates
type GetIssuesDismissedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetIssuesDismissedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIssuesDismissedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetIssuesDismissedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetIssuesDismissedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetIssuesDismissedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetIssuesDismissedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetIssuesDismissedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetIssuesDismissedAtType string

const (
	GetIssuesDismissedAtTypeDateTime              GetIssuesDismissedAtType = "date-time"
	GetIssuesDismissedAtTypeGetIssuesDismissedAt2 GetIssuesDismissedAtType = "getIssuesDismissedAt_2"
)

type GetIssuesDismissedAt struct {
	DateTime              *time.Time
	GetIssuesDismissedAt2 *GetIssuesDismissedAt2

	Type GetIssuesDismissedAtType
}

func CreateGetIssuesDismissedAtDateTime(dateTime time.Time) GetIssuesDismissedAt {
	typ := GetIssuesDismissedAtTypeDateTime

	return GetIssuesDismissedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssuesDismissedAtGetIssuesDismissedAt2(getIssuesDismissedAt2 GetIssuesDismissedAt2) GetIssuesDismissedAt {
	typ := GetIssuesDismissedAtTypeGetIssuesDismissedAt2

	return GetIssuesDismissedAt{
		GetIssuesDismissedAt2: &getIssuesDismissedAt2,
		Type:                  typ,
	}
}

func (u *GetIssuesDismissedAt) UnmarshalJSON(data []byte) error {

	getIssuesDismissedAt2 := GetIssuesDismissedAt2{}
	if err := utils.UnmarshalJSON(data, &getIssuesDismissedAt2, "", true, true); err == nil {
		u.GetIssuesDismissedAt2 = &getIssuesDismissedAt2
		u.Type = GetIssuesDismissedAtTypeGetIssuesDismissedAt2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetIssuesDismissedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesDismissedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetIssuesDismissedAt2 != nil {
		return utils.MarshalJSON(u.GetIssuesDismissedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetIssuesFirstSeenAt2 - Filter by first seen dates
type GetIssuesFirstSeenAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetIssuesFirstSeenAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIssuesFirstSeenAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetIssuesFirstSeenAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetIssuesFirstSeenAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetIssuesFirstSeenAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetIssuesFirstSeenAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetIssuesFirstSeenAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetIssuesFirstSeenAtType string

const (
	GetIssuesFirstSeenAtTypeDateTime              GetIssuesFirstSeenAtType = "date-time"
	GetIssuesFirstSeenAtTypeGetIssuesFirstSeenAt2 GetIssuesFirstSeenAtType = "getIssuesFirstSeenAt_2"
)

type GetIssuesFirstSeenAt struct {
	DateTime              *time.Time
	GetIssuesFirstSeenAt2 *GetIssuesFirstSeenAt2

	Type GetIssuesFirstSeenAtType
}

func CreateGetIssuesFirstSeenAtDateTime(dateTime time.Time) GetIssuesFirstSeenAt {
	typ := GetIssuesFirstSeenAtTypeDateTime

	return GetIssuesFirstSeenAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssuesFirstSeenAtGetIssuesFirstSeenAt2(getIssuesFirstSeenAt2 GetIssuesFirstSeenAt2) GetIssuesFirstSeenAt {
	typ := GetIssuesFirstSeenAtTypeGetIssuesFirstSeenAt2

	return GetIssuesFirstSeenAt{
		GetIssuesFirstSeenAt2: &getIssuesFirstSeenAt2,
		Type:                  typ,
	}
}

func (u *GetIssuesFirstSeenAt) UnmarshalJSON(data []byte) error {

	getIssuesFirstSeenAt2 := GetIssuesFirstSeenAt2{}
	if err := utils.UnmarshalJSON(data, &getIssuesFirstSeenAt2, "", true, true); err == nil {
		u.GetIssuesFirstSeenAt2 = &getIssuesFirstSeenAt2
		u.Type = GetIssuesFirstSeenAtTypeGetIssuesFirstSeenAt2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetIssuesFirstSeenAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesFirstSeenAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetIssuesFirstSeenAt2 != nil {
		return utils.MarshalJSON(u.GetIssuesFirstSeenAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesIDType string

const (
	GetIssuesIDTypeStr        GetIssuesIDType = "str"
	GetIssuesIDTypeArrayOfstr GetIssuesIDType = "arrayOfstr"
)

type GetIssuesID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssuesIDType
}

func CreateGetIssuesIDStr(str string) GetIssuesID {
	typ := GetIssuesIDTypeStr

	return GetIssuesID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssuesIDArrayOfstr(arrayOfstr []string) GetIssuesID {
	typ := GetIssuesIDTypeArrayOfstr

	return GetIssuesID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssuesID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetIssuesIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssuesIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesIssueTriggerIDType string

const (
	GetIssuesIssueTriggerIDTypeStr        GetIssuesIssueTriggerIDType = "str"
	GetIssuesIssueTriggerIDTypeArrayOfstr GetIssuesIssueTriggerIDType = "arrayOfstr"
)

type GetIssuesIssueTriggerID struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssuesIssueTriggerIDType
}

func CreateGetIssuesIssueTriggerIDStr(str string) GetIssuesIssueTriggerID {
	typ := GetIssuesIssueTriggerIDTypeStr

	return GetIssuesIssueTriggerID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssuesIssueTriggerIDArrayOfstr(arrayOfstr []string) GetIssuesIssueTriggerID {
	typ := GetIssuesIssueTriggerIDTypeArrayOfstr

	return GetIssuesIssueTriggerID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssuesIssueTriggerID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetIssuesIssueTriggerIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssuesIssueTriggerIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesIssueTriggerID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetIssuesLastSeenAt2 - Filter by last seen dates
type GetIssuesLastSeenAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetIssuesLastSeenAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIssuesLastSeenAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetIssuesLastSeenAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetIssuesLastSeenAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetIssuesLastSeenAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetIssuesLastSeenAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetIssuesLastSeenAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetIssuesLastSeenAtType string

const (
	GetIssuesLastSeenAtTypeDateTime             GetIssuesLastSeenAtType = "date-time"
	GetIssuesLastSeenAtTypeGetIssuesLastSeenAt2 GetIssuesLastSeenAtType = "getIssuesLastSeenAt_2"
)

type GetIssuesLastSeenAt struct {
	DateTime             *time.Time
	GetIssuesLastSeenAt2 *GetIssuesLastSeenAt2

	Type GetIssuesLastSeenAtType
}

func CreateGetIssuesLastSeenAtDateTime(dateTime time.Time) GetIssuesLastSeenAt {
	typ := GetIssuesLastSeenAtTypeDateTime

	return GetIssuesLastSeenAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetIssuesLastSeenAtGetIssuesLastSeenAt2(getIssuesLastSeenAt2 GetIssuesLastSeenAt2) GetIssuesLastSeenAt {
	typ := GetIssuesLastSeenAtTypeGetIssuesLastSeenAt2

	return GetIssuesLastSeenAt{
		GetIssuesLastSeenAt2: &getIssuesLastSeenAt2,
		Type:                 typ,
	}
}

func (u *GetIssuesLastSeenAt) UnmarshalJSON(data []byte) error {

	getIssuesLastSeenAt2 := GetIssuesLastSeenAt2{}
	if err := utils.UnmarshalJSON(data, &getIssuesLastSeenAt2, "", true, true); err == nil {
		u.GetIssuesLastSeenAt2 = &getIssuesLastSeenAt2
		u.Type = GetIssuesLastSeenAtTypeGetIssuesLastSeenAt2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetIssuesLastSeenAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesLastSeenAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetIssuesLastSeenAt2 != nil {
		return utils.MarshalJSON(u.GetIssuesLastSeenAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesMergedWithType string

const (
	GetIssuesMergedWithTypeStr        GetIssuesMergedWithType = "str"
	GetIssuesMergedWithTypeArrayOfstr GetIssuesMergedWithType = "arrayOfstr"
)

type GetIssuesMergedWith struct {
	Str        *string
	ArrayOfstr []string

	Type GetIssuesMergedWithType
}

func CreateGetIssuesMergedWithStr(str string) GetIssuesMergedWith {
	typ := GetIssuesMergedWithTypeStr

	return GetIssuesMergedWith{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetIssuesMergedWithArrayOfstr(arrayOfstr []string) GetIssuesMergedWith {
	typ := GetIssuesMergedWithTypeArrayOfstr

	return GetIssuesMergedWith{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetIssuesMergedWith) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetIssuesMergedWithTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetIssuesMergedWithTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesMergedWith) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesOrderBy2 string

const (
	GetIssuesOrderBy2CreatedAt   GetIssuesOrderBy2 = "created_at"
	GetIssuesOrderBy2FirstSeenAt GetIssuesOrderBy2 = "first_seen_at"
	GetIssuesOrderBy2LastSeenAt  GetIssuesOrderBy2 = "last_seen_at"
	GetIssuesOrderBy2OpenedAt    GetIssuesOrderBy2 = "opened_at"
	GetIssuesOrderBy2Status      GetIssuesOrderBy2 = "status"
)

func (e GetIssuesOrderBy2) ToPointer() *GetIssuesOrderBy2 {
	return &e
}

func (e *GetIssuesOrderBy2) UnmarshalJSON(data []byte) error {
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
		*e = GetIssuesOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesOrderBy2: %v", v)
	}
}

// GetIssuesOrderBy1 - Sort key(s)
type GetIssuesOrderBy1 string

const (
	GetIssuesOrderBy1CreatedAt   GetIssuesOrderBy1 = "created_at"
	GetIssuesOrderBy1FirstSeenAt GetIssuesOrderBy1 = "first_seen_at"
	GetIssuesOrderBy1LastSeenAt  GetIssuesOrderBy1 = "last_seen_at"
	GetIssuesOrderBy1OpenedAt    GetIssuesOrderBy1 = "opened_at"
	GetIssuesOrderBy1Status      GetIssuesOrderBy1 = "status"
)

func (e GetIssuesOrderBy1) ToPointer() *GetIssuesOrderBy1 {
	return &e
}

func (e *GetIssuesOrderBy1) UnmarshalJSON(data []byte) error {
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
		*e = GetIssuesOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesOrderBy1: %v", v)
	}
}

type GetIssuesOrderByType string

const (
	GetIssuesOrderByTypeGetIssuesOrderBy1        GetIssuesOrderByType = "getIssuesOrderBy_1"
	GetIssuesOrderByTypeArrayOfgetIssuesOrderBy2 GetIssuesOrderByType = "arrayOfgetIssuesOrderBy_2"
)

type GetIssuesOrderBy struct {
	GetIssuesOrderBy1        *GetIssuesOrderBy1
	ArrayOfgetIssuesOrderBy2 []GetIssuesOrderBy2

	Type GetIssuesOrderByType
}

func CreateGetIssuesOrderByGetIssuesOrderBy1(getIssuesOrderBy1 GetIssuesOrderBy1) GetIssuesOrderBy {
	typ := GetIssuesOrderByTypeGetIssuesOrderBy1

	return GetIssuesOrderBy{
		GetIssuesOrderBy1: &getIssuesOrderBy1,
		Type:              typ,
	}
}

func CreateGetIssuesOrderByArrayOfgetIssuesOrderBy2(arrayOfgetIssuesOrderBy2 []GetIssuesOrderBy2) GetIssuesOrderBy {
	typ := GetIssuesOrderByTypeArrayOfgetIssuesOrderBy2

	return GetIssuesOrderBy{
		ArrayOfgetIssuesOrderBy2: arrayOfgetIssuesOrderBy2,
		Type:                     typ,
	}
}

func (u *GetIssuesOrderBy) UnmarshalJSON(data []byte) error {

	getIssuesOrderBy1 := GetIssuesOrderBy1("")
	if err := utils.UnmarshalJSON(data, &getIssuesOrderBy1, "", true, true); err == nil {
		u.GetIssuesOrderBy1 = &getIssuesOrderBy1
		u.Type = GetIssuesOrderByTypeGetIssuesOrderBy1
		return nil
	}

	arrayOfgetIssuesOrderBy2 := []GetIssuesOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetIssuesOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetIssuesOrderBy2 = arrayOfgetIssuesOrderBy2
		u.Type = GetIssuesOrderByTypeArrayOfgetIssuesOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetIssuesOrderBy1 != nil {
		return utils.MarshalJSON(u.GetIssuesOrderBy1, "", true)
	}

	if u.ArrayOfgetIssuesOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetIssuesOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetIssuesStatus2 - Issue status
type GetIssuesStatus2 string

const (
	GetIssuesStatus2Opened       GetIssuesStatus2 = "OPENED"
	GetIssuesStatus2Ignored      GetIssuesStatus2 = "IGNORED"
	GetIssuesStatus2Acknowledged GetIssuesStatus2 = "ACKNOWLEDGED"
	GetIssuesStatus2Resolved     GetIssuesStatus2 = "RESOLVED"
)

func (e GetIssuesStatus2) ToPointer() *GetIssuesStatus2 {
	return &e
}

func (e *GetIssuesStatus2) UnmarshalJSON(data []byte) error {
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
		*e = GetIssuesStatus2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesStatus2: %v", v)
	}
}

// GetIssuesStatus1 - Issue status
type GetIssuesStatus1 string

const (
	GetIssuesStatus1Opened       GetIssuesStatus1 = "OPENED"
	GetIssuesStatus1Ignored      GetIssuesStatus1 = "IGNORED"
	GetIssuesStatus1Acknowledged GetIssuesStatus1 = "ACKNOWLEDGED"
	GetIssuesStatus1Resolved     GetIssuesStatus1 = "RESOLVED"
)

func (e GetIssuesStatus1) ToPointer() *GetIssuesStatus1 {
	return &e
}

func (e *GetIssuesStatus1) UnmarshalJSON(data []byte) error {
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
		*e = GetIssuesStatus1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesStatus1: %v", v)
	}
}

type GetIssuesStatusType string

const (
	GetIssuesStatusTypeGetIssuesStatus1        GetIssuesStatusType = "getIssuesStatus_1"
	GetIssuesStatusTypeArrayOfgetIssuesStatus2 GetIssuesStatusType = "arrayOfgetIssuesStatus_2"
)

type GetIssuesStatus struct {
	GetIssuesStatus1        *GetIssuesStatus1
	ArrayOfgetIssuesStatus2 []GetIssuesStatus2

	Type GetIssuesStatusType
}

func CreateGetIssuesStatusGetIssuesStatus1(getIssuesStatus1 GetIssuesStatus1) GetIssuesStatus {
	typ := GetIssuesStatusTypeGetIssuesStatus1

	return GetIssuesStatus{
		GetIssuesStatus1: &getIssuesStatus1,
		Type:             typ,
	}
}

func CreateGetIssuesStatusArrayOfgetIssuesStatus2(arrayOfgetIssuesStatus2 []GetIssuesStatus2) GetIssuesStatus {
	typ := GetIssuesStatusTypeArrayOfgetIssuesStatus2

	return GetIssuesStatus{
		ArrayOfgetIssuesStatus2: arrayOfgetIssuesStatus2,
		Type:                    typ,
	}
}

func (u *GetIssuesStatus) UnmarshalJSON(data []byte) error {

	getIssuesStatus1 := GetIssuesStatus1("")
	if err := utils.UnmarshalJSON(data, &getIssuesStatus1, "", true, true); err == nil {
		u.GetIssuesStatus1 = &getIssuesStatus1
		u.Type = GetIssuesStatusTypeGetIssuesStatus1
		return nil
	}

	arrayOfgetIssuesStatus2 := []GetIssuesStatus2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetIssuesStatus2, "", true, true); err == nil {
		u.ArrayOfgetIssuesStatus2 = arrayOfgetIssuesStatus2
		u.Type = GetIssuesStatusTypeArrayOfgetIssuesStatus2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesStatus) MarshalJSON() ([]byte, error) {
	if u.GetIssuesStatus1 != nil {
		return utils.MarshalJSON(u.GetIssuesStatus1, "", true)
	}

	if u.ArrayOfgetIssuesStatus2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetIssuesStatus2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetIssuesType2 - Issue type
type GetIssuesType2 string

const (
	GetIssuesType2Delivery       GetIssuesType2 = "delivery"
	GetIssuesType2Transformation GetIssuesType2 = "transformation"
	GetIssuesType2Backpressure   GetIssuesType2 = "backpressure"
)

func (e GetIssuesType2) ToPointer() *GetIssuesType2 {
	return &e
}

func (e *GetIssuesType2) UnmarshalJSON(data []byte) error {
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
		*e = GetIssuesType2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesType2: %v", v)
	}
}

// GetIssuesType1 - Issue type
type GetIssuesType1 string

const (
	GetIssuesType1Delivery       GetIssuesType1 = "delivery"
	GetIssuesType1Transformation GetIssuesType1 = "transformation"
	GetIssuesType1Backpressure   GetIssuesType1 = "backpressure"
)

func (e GetIssuesType1) ToPointer() *GetIssuesType1 {
	return &e
}

func (e *GetIssuesType1) UnmarshalJSON(data []byte) error {
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
		*e = GetIssuesType1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetIssuesType1: %v", v)
	}
}

type GetIssuesTypeType string

const (
	GetIssuesTypeTypeGetIssuesType1        GetIssuesTypeType = "getIssuesType_1"
	GetIssuesTypeTypeArrayOfgetIssuesType2 GetIssuesTypeType = "arrayOfgetIssuesType_2"
)

type GetIssuesType struct {
	GetIssuesType1        *GetIssuesType1
	ArrayOfgetIssuesType2 []GetIssuesType2

	Type GetIssuesTypeType
}

func CreateGetIssuesTypeGetIssuesType1(getIssuesType1 GetIssuesType1) GetIssuesType {
	typ := GetIssuesTypeTypeGetIssuesType1

	return GetIssuesType{
		GetIssuesType1: &getIssuesType1,
		Type:           typ,
	}
}

func CreateGetIssuesTypeArrayOfgetIssuesType2(arrayOfgetIssuesType2 []GetIssuesType2) GetIssuesType {
	typ := GetIssuesTypeTypeArrayOfgetIssuesType2

	return GetIssuesType{
		ArrayOfgetIssuesType2: arrayOfgetIssuesType2,
		Type:                  typ,
	}
}

func (u *GetIssuesType) UnmarshalJSON(data []byte) error {

	getIssuesType1 := GetIssuesType1("")
	if err := utils.UnmarshalJSON(data, &getIssuesType1, "", true, true); err == nil {
		u.GetIssuesType1 = &getIssuesType1
		u.Type = GetIssuesTypeTypeGetIssuesType1
		return nil
	}

	arrayOfgetIssuesType2 := []GetIssuesType2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetIssuesType2, "", true, true); err == nil {
		u.ArrayOfgetIssuesType2 = arrayOfgetIssuesType2
		u.Type = GetIssuesTypeTypeArrayOfgetIssuesType2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetIssuesType) MarshalJSON() ([]byte, error) {
	if u.GetIssuesType1 != nil {
		return utils.MarshalJSON(u.GetIssuesType1, "", true)
	}

	if u.ArrayOfgetIssuesType2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetIssuesType2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetIssuesRequest struct {
	// Filter by aggregation keys
	AggregationKeys *GetIssuesAggregationKeys `queryParam:"style=form,explode=true,name=aggregation_keys"`
	CreatedAt       *GetIssuesCreatedAt       `queryParam:"style=form,explode=true,name=created_at"`
	Dir             *GetIssuesDir             `queryParam:"style=form,explode=true,name=dir"`
	DismissedAt     *GetIssuesDismissedAt     `queryParam:"style=form,explode=true,name=dismissed_at"`
	FirstSeenAt     *GetIssuesFirstSeenAt     `queryParam:"style=form,explode=true,name=first_seen_at"`
	ID              *GetIssuesID              `queryParam:"style=form,explode=true,name=id"`
	IssueTriggerID  *GetIssuesIssueTriggerID  `queryParam:"style=form,explode=true,name=issue_trigger_id"`
	LastSeenAt      *GetIssuesLastSeenAt      `queryParam:"style=form,explode=true,name=last_seen_at"`
	Limit           *int64                    `queryParam:"style=form,explode=true,name=limit"`
	MergedWith      *GetIssuesMergedWith      `queryParam:"style=form,explode=true,name=merged_with"`
	Next            *string                   `queryParam:"style=form,explode=true,name=next"`
	OrderBy         *GetIssuesOrderBy         `queryParam:"style=form,explode=true,name=order_by"`
	Prev            *string                   `queryParam:"style=form,explode=true,name=prev"`
	Status          *GetIssuesStatus          `queryParam:"style=form,explode=true,name=status"`
	Type            *GetIssuesType            `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetIssuesRequest) GetAggregationKeys() *GetIssuesAggregationKeys {
	if o == nil {
		return nil
	}
	return o.AggregationKeys
}

func (o *GetIssuesRequest) GetCreatedAt() *GetIssuesCreatedAt {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetIssuesRequest) GetDir() *GetIssuesDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetIssuesRequest) GetDismissedAt() *GetIssuesDismissedAt {
	if o == nil {
		return nil
	}
	return o.DismissedAt
}

func (o *GetIssuesRequest) GetFirstSeenAt() *GetIssuesFirstSeenAt {
	if o == nil {
		return nil
	}
	return o.FirstSeenAt
}

func (o *GetIssuesRequest) GetID() *GetIssuesID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetIssuesRequest) GetIssueTriggerID() *GetIssuesIssueTriggerID {
	if o == nil {
		return nil
	}
	return o.IssueTriggerID
}

func (o *GetIssuesRequest) GetLastSeenAt() *GetIssuesLastSeenAt {
	if o == nil {
		return nil
	}
	return o.LastSeenAt
}

func (o *GetIssuesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetIssuesRequest) GetMergedWith() *GetIssuesMergedWith {
	if o == nil {
		return nil
	}
	return o.MergedWith
}

func (o *GetIssuesRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetIssuesRequest) GetOrderBy() *GetIssuesOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetIssuesRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

func (o *GetIssuesRequest) GetStatus() *GetIssuesStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetIssuesRequest) GetType() *GetIssuesType {
	if o == nil {
		return nil
	}
	return o.Type
}

type GetIssuesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of issues
	IssueWithDataPaginatedResult *shared.IssueWithDataPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIssuesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIssuesResponse) GetIssueWithDataPaginatedResult() *shared.IssueWithDataPaginatedResult {
	if o == nil {
		return nil
	}
	return o.IssueWithDataPaginatedResult
}

func (o *GetIssuesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIssuesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

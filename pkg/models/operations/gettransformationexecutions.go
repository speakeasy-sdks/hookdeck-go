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

// GetTransformationExecutionsCreatedAt2 - ISO date of the transformation's execution
type GetTransformationExecutionsCreatedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetTransformationExecutionsCreatedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTransformationExecutionsCreatedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetTransformationExecutionsCreatedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetTransformationExecutionsCreatedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetTransformationExecutionsCreatedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetTransformationExecutionsCreatedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetTransformationExecutionsCreatedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetTransformationExecutionsCreatedAtType string

const (
	GetTransformationExecutionsCreatedAtTypeDateTime                              GetTransformationExecutionsCreatedAtType = "date-time"
	GetTransformationExecutionsCreatedAtTypeGetTransformationExecutionsCreatedAt2 GetTransformationExecutionsCreatedAtType = "getTransformationExecutionsCreatedAt_2"
)

type GetTransformationExecutionsCreatedAt struct {
	DateTime                              *time.Time
	GetTransformationExecutionsCreatedAt2 *GetTransformationExecutionsCreatedAt2

	Type GetTransformationExecutionsCreatedAtType
}

func CreateGetTransformationExecutionsCreatedAtDateTime(dateTime time.Time) GetTransformationExecutionsCreatedAt {
	typ := GetTransformationExecutionsCreatedAtTypeDateTime

	return GetTransformationExecutionsCreatedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetTransformationExecutionsCreatedAtGetTransformationExecutionsCreatedAt2(getTransformationExecutionsCreatedAt2 GetTransformationExecutionsCreatedAt2) GetTransformationExecutionsCreatedAt {
	typ := GetTransformationExecutionsCreatedAtTypeGetTransformationExecutionsCreatedAt2

	return GetTransformationExecutionsCreatedAt{
		GetTransformationExecutionsCreatedAt2: &getTransformationExecutionsCreatedAt2,
		Type:                                  typ,
	}
}

func (u *GetTransformationExecutionsCreatedAt) UnmarshalJSON(data []byte) error {

	getTransformationExecutionsCreatedAt2 := new(GetTransformationExecutionsCreatedAt2)
	if err := utils.UnmarshalJSON(data, &getTransformationExecutionsCreatedAt2, "", true, true); err == nil {
		u.GetTransformationExecutionsCreatedAt2 = getTransformationExecutionsCreatedAt2
		u.Type = GetTransformationExecutionsCreatedAtTypeGetTransformationExecutionsCreatedAt2
		return nil
	}

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = GetTransformationExecutionsCreatedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetTransformationExecutionsCreatedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetTransformationExecutionsCreatedAt2 != nil {
		return utils.MarshalJSON(u.GetTransformationExecutionsCreatedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetTransformationExecutionsDir2 string

const (
	GetTransformationExecutionsDir2Asc  GetTransformationExecutionsDir2 = "asc"
	GetTransformationExecutionsDir2Desc GetTransformationExecutionsDir2 = "desc"
)

func (e GetTransformationExecutionsDir2) ToPointer() *GetTransformationExecutionsDir2 {
	return &e
}

func (e *GetTransformationExecutionsDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetTransformationExecutionsDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTransformationExecutionsDir2: %v", v)
	}
}

// GetTransformationExecutionsDir1 - Sort direction
type GetTransformationExecutionsDir1 string

const (
	GetTransformationExecutionsDir1Asc  GetTransformationExecutionsDir1 = "asc"
	GetTransformationExecutionsDir1Desc GetTransformationExecutionsDir1 = "desc"
)

func (e GetTransformationExecutionsDir1) ToPointer() *GetTransformationExecutionsDir1 {
	return &e
}

func (e *GetTransformationExecutionsDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetTransformationExecutionsDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTransformationExecutionsDir1: %v", v)
	}
}

type GetTransformationExecutionsDirType string

const (
	GetTransformationExecutionsDirTypeGetTransformationExecutionsDir1        GetTransformationExecutionsDirType = "getTransformationExecutionsDir_1"
	GetTransformationExecutionsDirTypeArrayOfgetTransformationExecutionsDir2 GetTransformationExecutionsDirType = "arrayOfgetTransformationExecutionsDir_2"
)

type GetTransformationExecutionsDir struct {
	GetTransformationExecutionsDir1        *GetTransformationExecutionsDir1
	ArrayOfgetTransformationExecutionsDir2 []GetTransformationExecutionsDir2

	Type GetTransformationExecutionsDirType
}

func CreateGetTransformationExecutionsDirGetTransformationExecutionsDir1(getTransformationExecutionsDir1 GetTransformationExecutionsDir1) GetTransformationExecutionsDir {
	typ := GetTransformationExecutionsDirTypeGetTransformationExecutionsDir1

	return GetTransformationExecutionsDir{
		GetTransformationExecutionsDir1: &getTransformationExecutionsDir1,
		Type:                            typ,
	}
}

func CreateGetTransformationExecutionsDirArrayOfgetTransformationExecutionsDir2(arrayOfgetTransformationExecutionsDir2 []GetTransformationExecutionsDir2) GetTransformationExecutionsDir {
	typ := GetTransformationExecutionsDirTypeArrayOfgetTransformationExecutionsDir2

	return GetTransformationExecutionsDir{
		ArrayOfgetTransformationExecutionsDir2: arrayOfgetTransformationExecutionsDir2,
		Type:                                   typ,
	}
}

func (u *GetTransformationExecutionsDir) UnmarshalJSON(data []byte) error {

	getTransformationExecutionsDir1 := new(GetTransformationExecutionsDir1)
	if err := utils.UnmarshalJSON(data, &getTransformationExecutionsDir1, "", true, true); err == nil {
		u.GetTransformationExecutionsDir1 = getTransformationExecutionsDir1
		u.Type = GetTransformationExecutionsDirTypeGetTransformationExecutionsDir1
		return nil
	}

	arrayOfgetTransformationExecutionsDir2 := []GetTransformationExecutionsDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetTransformationExecutionsDir2, "", true, true); err == nil {
		u.ArrayOfgetTransformationExecutionsDir2 = arrayOfgetTransformationExecutionsDir2
		u.Type = GetTransformationExecutionsDirTypeArrayOfgetTransformationExecutionsDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetTransformationExecutionsDir) MarshalJSON() ([]byte, error) {
	if u.GetTransformationExecutionsDir1 != nil {
		return utils.MarshalJSON(u.GetTransformationExecutionsDir1, "", true)
	}

	if u.ArrayOfgetTransformationExecutionsDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetTransformationExecutionsDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetTransformationExecutionsIssueIDType string

const (
	GetTransformationExecutionsIssueIDTypeStr        GetTransformationExecutionsIssueIDType = "str"
	GetTransformationExecutionsIssueIDTypeArrayOfstr GetTransformationExecutionsIssueIDType = "arrayOfstr"
)

type GetTransformationExecutionsIssueID struct {
	Str        *string
	ArrayOfstr []string

	Type GetTransformationExecutionsIssueIDType
}

func CreateGetTransformationExecutionsIssueIDStr(str string) GetTransformationExecutionsIssueID {
	typ := GetTransformationExecutionsIssueIDTypeStr

	return GetTransformationExecutionsIssueID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetTransformationExecutionsIssueIDArrayOfstr(arrayOfstr []string) GetTransformationExecutionsIssueID {
	typ := GetTransformationExecutionsIssueIDTypeArrayOfstr

	return GetTransformationExecutionsIssueID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetTransformationExecutionsIssueID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetTransformationExecutionsIssueIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetTransformationExecutionsIssueIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetTransformationExecutionsIssueID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetTransformationExecutionsLogLevel2 string

const (
	GetTransformationExecutionsLogLevel2LessThanNilGreaterThan GetTransformationExecutionsLogLevel2 = "<nil>"
	GetTransformationExecutionsLogLevel2Debug                  GetTransformationExecutionsLogLevel2 = "debug"
	GetTransformationExecutionsLogLevel2Info                   GetTransformationExecutionsLogLevel2 = "info"
	GetTransformationExecutionsLogLevel2Warn                   GetTransformationExecutionsLogLevel2 = "warn"
	GetTransformationExecutionsLogLevel2Error                  GetTransformationExecutionsLogLevel2 = "error"
	GetTransformationExecutionsLogLevel2Fatal                  GetTransformationExecutionsLogLevel2 = "fatal"
)

func (e GetTransformationExecutionsLogLevel2) ToPointer() *GetTransformationExecutionsLogLevel2 {
	return &e
}

func (e *GetTransformationExecutionsLogLevel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "<nil>":
		fallthrough
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		fallthrough
	case "fatal":
		*e = GetTransformationExecutionsLogLevel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTransformationExecutionsLogLevel2: %v", v)
	}
}

// GetTransformationExecutionsLogLevel1 - Log level of the execution
type GetTransformationExecutionsLogLevel1 string

const (
	GetTransformationExecutionsLogLevel1LessThanNilGreaterThan GetTransformationExecutionsLogLevel1 = "<nil>"
	GetTransformationExecutionsLogLevel1Debug                  GetTransformationExecutionsLogLevel1 = "debug"
	GetTransformationExecutionsLogLevel1Info                   GetTransformationExecutionsLogLevel1 = "info"
	GetTransformationExecutionsLogLevel1Warn                   GetTransformationExecutionsLogLevel1 = "warn"
	GetTransformationExecutionsLogLevel1Error                  GetTransformationExecutionsLogLevel1 = "error"
	GetTransformationExecutionsLogLevel1Fatal                  GetTransformationExecutionsLogLevel1 = "fatal"
)

func (e GetTransformationExecutionsLogLevel1) ToPointer() *GetTransformationExecutionsLogLevel1 {
	return &e
}

func (e *GetTransformationExecutionsLogLevel1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "<nil>":
		fallthrough
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		fallthrough
	case "fatal":
		*e = GetTransformationExecutionsLogLevel1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTransformationExecutionsLogLevel1: %v", v)
	}
}

type GetTransformationExecutionsLogLevelType string

const (
	GetTransformationExecutionsLogLevelTypeGetTransformationExecutionsLogLevel1        GetTransformationExecutionsLogLevelType = "getTransformationExecutionsLogLevel_1"
	GetTransformationExecutionsLogLevelTypeArrayOfgetTransformationExecutionsLogLevel2 GetTransformationExecutionsLogLevelType = "arrayOfgetTransformationExecutionsLogLevel_2"
)

type GetTransformationExecutionsLogLevel struct {
	GetTransformationExecutionsLogLevel1        *GetTransformationExecutionsLogLevel1
	ArrayOfgetTransformationExecutionsLogLevel2 []GetTransformationExecutionsLogLevel2

	Type GetTransformationExecutionsLogLevelType
}

func CreateGetTransformationExecutionsLogLevelGetTransformationExecutionsLogLevel1(getTransformationExecutionsLogLevel1 GetTransformationExecutionsLogLevel1) GetTransformationExecutionsLogLevel {
	typ := GetTransformationExecutionsLogLevelTypeGetTransformationExecutionsLogLevel1

	return GetTransformationExecutionsLogLevel{
		GetTransformationExecutionsLogLevel1: &getTransformationExecutionsLogLevel1,
		Type:                                 typ,
	}
}

func CreateGetTransformationExecutionsLogLevelArrayOfgetTransformationExecutionsLogLevel2(arrayOfgetTransformationExecutionsLogLevel2 []GetTransformationExecutionsLogLevel2) GetTransformationExecutionsLogLevel {
	typ := GetTransformationExecutionsLogLevelTypeArrayOfgetTransformationExecutionsLogLevel2

	return GetTransformationExecutionsLogLevel{
		ArrayOfgetTransformationExecutionsLogLevel2: arrayOfgetTransformationExecutionsLogLevel2,
		Type: typ,
	}
}

func (u *GetTransformationExecutionsLogLevel) UnmarshalJSON(data []byte) error {

	getTransformationExecutionsLogLevel1 := new(GetTransformationExecutionsLogLevel1)
	if err := utils.UnmarshalJSON(data, &getTransformationExecutionsLogLevel1, "", true, true); err == nil {
		u.GetTransformationExecutionsLogLevel1 = getTransformationExecutionsLogLevel1
		u.Type = GetTransformationExecutionsLogLevelTypeGetTransformationExecutionsLogLevel1
		return nil
	}

	arrayOfgetTransformationExecutionsLogLevel2 := []GetTransformationExecutionsLogLevel2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetTransformationExecutionsLogLevel2, "", true, true); err == nil {
		u.ArrayOfgetTransformationExecutionsLogLevel2 = arrayOfgetTransformationExecutionsLogLevel2
		u.Type = GetTransformationExecutionsLogLevelTypeArrayOfgetTransformationExecutionsLogLevel2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetTransformationExecutionsLogLevel) MarshalJSON() ([]byte, error) {
	if u.GetTransformationExecutionsLogLevel1 != nil {
		return utils.MarshalJSON(u.GetTransformationExecutionsLogLevel1, "", true)
	}

	if u.ArrayOfgetTransformationExecutionsLogLevel2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetTransformationExecutionsLogLevel2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetTransformationExecutionsOrderByType string

const (
	GetTransformationExecutionsOrderByTypeStr        GetTransformationExecutionsOrderByType = "str"
	GetTransformationExecutionsOrderByTypeArrayOfstr GetTransformationExecutionsOrderByType = "arrayOfstr"
)

type GetTransformationExecutionsOrderBy struct {
	Str        *string
	ArrayOfstr []string

	Type GetTransformationExecutionsOrderByType
}

func CreateGetTransformationExecutionsOrderByStr(str string) GetTransformationExecutionsOrderBy {
	typ := GetTransformationExecutionsOrderByTypeStr

	return GetTransformationExecutionsOrderBy{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetTransformationExecutionsOrderByArrayOfstr(arrayOfstr []string) GetTransformationExecutionsOrderBy {
	typ := GetTransformationExecutionsOrderByTypeArrayOfstr

	return GetTransformationExecutionsOrderBy{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetTransformationExecutionsOrderBy) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetTransformationExecutionsOrderByTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetTransformationExecutionsOrderByTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetTransformationExecutionsOrderBy) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetTransformationExecutionsWebhookIDType string

const (
	GetTransformationExecutionsWebhookIDTypeStr        GetTransformationExecutionsWebhookIDType = "str"
	GetTransformationExecutionsWebhookIDTypeArrayOfstr GetTransformationExecutionsWebhookIDType = "arrayOfstr"
)

type GetTransformationExecutionsWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type GetTransformationExecutionsWebhookIDType
}

func CreateGetTransformationExecutionsWebhookIDStr(str string) GetTransformationExecutionsWebhookID {
	typ := GetTransformationExecutionsWebhookIDTypeStr

	return GetTransformationExecutionsWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetTransformationExecutionsWebhookIDArrayOfstr(arrayOfstr []string) GetTransformationExecutionsWebhookID {
	typ := GetTransformationExecutionsWebhookIDTypeArrayOfstr

	return GetTransformationExecutionsWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetTransformationExecutionsWebhookID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetTransformationExecutionsWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetTransformationExecutionsWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetTransformationExecutionsWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetTransformationExecutionsRequest struct {
	CreatedAt *GetTransformationExecutionsCreatedAt `queryParam:"style=form,explode=true,name=created_at"`
	Dir       *GetTransformationExecutionsDir       `queryParam:"style=form,explode=true,name=dir"`
	ID        string                                `pathParam:"style=simple,explode=false,name=id"`
	IssueID   *GetTransformationExecutionsIssueID   `queryParam:"style=form,explode=true,name=issue_id"`
	Limit     *int64                                `queryParam:"style=form,explode=true,name=limit"`
	LogLevel  *GetTransformationExecutionsLogLevel  `queryParam:"style=form,explode=true,name=log_level"`
	Next      *string                               `queryParam:"style=form,explode=true,name=next"`
	OrderBy   *GetTransformationExecutionsOrderBy   `queryParam:"style=form,explode=true,name=order_by"`
	Prev      *string                               `queryParam:"style=form,explode=true,name=prev"`
	WebhookID *GetTransformationExecutionsWebhookID `queryParam:"style=form,explode=true,name=webhook_id"`
}

func (o *GetTransformationExecutionsRequest) GetCreatedAt() *GetTransformationExecutionsCreatedAt {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetTransformationExecutionsRequest) GetDir() *GetTransformationExecutionsDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetTransformationExecutionsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetTransformationExecutionsRequest) GetIssueID() *GetTransformationExecutionsIssueID {
	if o == nil {
		return nil
	}
	return o.IssueID
}

func (o *GetTransformationExecutionsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetTransformationExecutionsRequest) GetLogLevel() *GetTransformationExecutionsLogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *GetTransformationExecutionsRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetTransformationExecutionsRequest) GetOrderBy() *GetTransformationExecutionsOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetTransformationExecutionsRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

func (o *GetTransformationExecutionsRequest) GetWebhookID() *GetTransformationExecutionsWebhookID {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

type GetTransformationExecutionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// List of transformation executions
	TransformationExecutionPaginatedResult *shared.TransformationExecutionPaginatedResult
}

func (o *GetTransformationExecutionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransformationExecutionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransformationExecutionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransformationExecutionsResponse) GetTransformationExecutionPaginatedResult() *shared.TransformationExecutionPaginatedResult {
	if o == nil {
		return nil
	}
	return o.TransformationExecutionPaginatedResult
}

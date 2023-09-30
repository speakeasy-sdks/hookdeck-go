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

type GetBookmarksDir2 string

const (
	GetBookmarksDir2Asc  GetBookmarksDir2 = "asc"
	GetBookmarksDir2Desc GetBookmarksDir2 = "desc"
)

func (e GetBookmarksDir2) ToPointer() *GetBookmarksDir2 {
	return &e
}

func (e *GetBookmarksDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetBookmarksDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetBookmarksDir2: %v", v)
	}
}

// GetBookmarksDir1 - Sort direction
type GetBookmarksDir1 string

const (
	GetBookmarksDir1Asc  GetBookmarksDir1 = "asc"
	GetBookmarksDir1Desc GetBookmarksDir1 = "desc"
)

func (e GetBookmarksDir1) ToPointer() *GetBookmarksDir1 {
	return &e
}

func (e *GetBookmarksDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetBookmarksDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetBookmarksDir1: %v", v)
	}
}

type GetBookmarksDirType string

const (
	GetBookmarksDirTypeGetBookmarksDir1        GetBookmarksDirType = "getBookmarksDir_1"
	GetBookmarksDirTypeArrayOfgetBookmarksDir2 GetBookmarksDirType = "arrayOfgetBookmarksDir_2"
)

type GetBookmarksDir struct {
	GetBookmarksDir1        *GetBookmarksDir1
	ArrayOfgetBookmarksDir2 []GetBookmarksDir2

	Type GetBookmarksDirType
}

func CreateGetBookmarksDirGetBookmarksDir1(getBookmarksDir1 GetBookmarksDir1) GetBookmarksDir {
	typ := GetBookmarksDirTypeGetBookmarksDir1

	return GetBookmarksDir{
		GetBookmarksDir1: &getBookmarksDir1,
		Type:             typ,
	}
}

func CreateGetBookmarksDirArrayOfgetBookmarksDir2(arrayOfgetBookmarksDir2 []GetBookmarksDir2) GetBookmarksDir {
	typ := GetBookmarksDirTypeArrayOfgetBookmarksDir2

	return GetBookmarksDir{
		ArrayOfgetBookmarksDir2: arrayOfgetBookmarksDir2,
		Type:                    typ,
	}
}

func (u *GetBookmarksDir) UnmarshalJSON(data []byte) error {

	getBookmarksDir1 := new(GetBookmarksDir1)
	if err := utils.UnmarshalJSON(data, &getBookmarksDir1, "", true, true); err == nil {
		u.GetBookmarksDir1 = getBookmarksDir1
		u.Type = GetBookmarksDirTypeGetBookmarksDir1
		return nil
	}

	arrayOfgetBookmarksDir2 := []GetBookmarksDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetBookmarksDir2, "", true, true); err == nil {
		u.ArrayOfgetBookmarksDir2 = arrayOfgetBookmarksDir2
		u.Type = GetBookmarksDirTypeArrayOfgetBookmarksDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksDir) MarshalJSON() ([]byte, error) {
	if u.GetBookmarksDir1 != nil {
		return utils.MarshalJSON(u.GetBookmarksDir1, "", true)
	}

	if u.ArrayOfgetBookmarksDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetBookmarksDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetBookmarksEventDataIDType string

const (
	GetBookmarksEventDataIDTypeStr        GetBookmarksEventDataIDType = "str"
	GetBookmarksEventDataIDTypeArrayOfstr GetBookmarksEventDataIDType = "arrayOfstr"
)

type GetBookmarksEventDataID struct {
	Str        *string
	ArrayOfstr []string

	Type GetBookmarksEventDataIDType
}

func CreateGetBookmarksEventDataIDStr(str string) GetBookmarksEventDataID {
	typ := GetBookmarksEventDataIDTypeStr

	return GetBookmarksEventDataID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetBookmarksEventDataIDArrayOfstr(arrayOfstr []string) GetBookmarksEventDataID {
	typ := GetBookmarksEventDataIDTypeArrayOfstr

	return GetBookmarksEventDataID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetBookmarksEventDataID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetBookmarksEventDataIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetBookmarksEventDataIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksEventDataID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetBookmarksIDType string

const (
	GetBookmarksIDTypeStr        GetBookmarksIDType = "str"
	GetBookmarksIDTypeArrayOfstr GetBookmarksIDType = "arrayOfstr"
)

type GetBookmarksID struct {
	Str        *string
	ArrayOfstr []string

	Type GetBookmarksIDType
}

func CreateGetBookmarksIDStr(str string) GetBookmarksID {
	typ := GetBookmarksIDTypeStr

	return GetBookmarksID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetBookmarksIDArrayOfstr(arrayOfstr []string) GetBookmarksID {
	typ := GetBookmarksIDTypeArrayOfstr

	return GetBookmarksID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetBookmarksID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetBookmarksIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetBookmarksIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetBookmarksLabelType string

const (
	GetBookmarksLabelTypeStr        GetBookmarksLabelType = "str"
	GetBookmarksLabelTypeArrayOfstr GetBookmarksLabelType = "arrayOfstr"
)

type GetBookmarksLabel struct {
	Str        *string
	ArrayOfstr []string

	Type GetBookmarksLabelType
}

func CreateGetBookmarksLabelStr(str string) GetBookmarksLabel {
	typ := GetBookmarksLabelTypeStr

	return GetBookmarksLabel{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetBookmarksLabelArrayOfstr(arrayOfstr []string) GetBookmarksLabel {
	typ := GetBookmarksLabelTypeArrayOfstr

	return GetBookmarksLabel{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetBookmarksLabel) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetBookmarksLabelTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetBookmarksLabelTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksLabel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetBookmarksLastUsedAt2 - Filter by last used date
type GetBookmarksLastUsedAt2 struct {
	Any *bool `queryParam:"name=any"`
	// Last used date
	Gt *time.Time `queryParam:"name=gt"`
	// Last used date
	Gte *time.Time `queryParam:"name=gte"`
	// Last used date
	Le *time.Time `queryParam:"name=le"`
	// Last used date
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetBookmarksLastUsedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetBookmarksLastUsedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetBookmarksLastUsedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetBookmarksLastUsedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetBookmarksLastUsedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetBookmarksLastUsedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetBookmarksLastUsedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetBookmarksLastUsedAtType string

const (
	GetBookmarksLastUsedAtTypeDateTime                GetBookmarksLastUsedAtType = "date-time"
	GetBookmarksLastUsedAtTypeGetBookmarksLastUsedAt2 GetBookmarksLastUsedAtType = "getBookmarksLastUsedAt_2"
)

type GetBookmarksLastUsedAt struct {
	DateTime                *time.Time
	GetBookmarksLastUsedAt2 *GetBookmarksLastUsedAt2

	Type GetBookmarksLastUsedAtType
}

func CreateGetBookmarksLastUsedAtDateTime(dateTime time.Time) GetBookmarksLastUsedAt {
	typ := GetBookmarksLastUsedAtTypeDateTime

	return GetBookmarksLastUsedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetBookmarksLastUsedAtGetBookmarksLastUsedAt2(getBookmarksLastUsedAt2 GetBookmarksLastUsedAt2) GetBookmarksLastUsedAt {
	typ := GetBookmarksLastUsedAtTypeGetBookmarksLastUsedAt2

	return GetBookmarksLastUsedAt{
		GetBookmarksLastUsedAt2: &getBookmarksLastUsedAt2,
		Type:                    typ,
	}
}

func (u *GetBookmarksLastUsedAt) UnmarshalJSON(data []byte) error {

	getBookmarksLastUsedAt2 := new(GetBookmarksLastUsedAt2)
	if err := utils.UnmarshalJSON(data, &getBookmarksLastUsedAt2, "", true, true); err == nil {
		u.GetBookmarksLastUsedAt2 = getBookmarksLastUsedAt2
		u.Type = GetBookmarksLastUsedAtTypeGetBookmarksLastUsedAt2
		return nil
	}

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = GetBookmarksLastUsedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksLastUsedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetBookmarksLastUsedAt2 != nil {
		return utils.MarshalJSON(u.GetBookmarksLastUsedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetBookmarksNameType string

const (
	GetBookmarksNameTypeStr        GetBookmarksNameType = "str"
	GetBookmarksNameTypeArrayOfstr GetBookmarksNameType = "arrayOfstr"
)

type GetBookmarksName struct {
	Str        *string
	ArrayOfstr []string

	Type GetBookmarksNameType
}

func CreateGetBookmarksNameStr(str string) GetBookmarksName {
	typ := GetBookmarksNameTypeStr

	return GetBookmarksName{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetBookmarksNameArrayOfstr(arrayOfstr []string) GetBookmarksName {
	typ := GetBookmarksNameTypeArrayOfstr

	return GetBookmarksName{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetBookmarksName) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetBookmarksNameTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetBookmarksNameTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksName) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetBookmarksOrderByType string

const (
	GetBookmarksOrderByTypeStr        GetBookmarksOrderByType = "str"
	GetBookmarksOrderByTypeArrayOfstr GetBookmarksOrderByType = "arrayOfstr"
)

type GetBookmarksOrderBy struct {
	Str        *string
	ArrayOfstr []string

	Type GetBookmarksOrderByType
}

func CreateGetBookmarksOrderByStr(str string) GetBookmarksOrderBy {
	typ := GetBookmarksOrderByTypeStr

	return GetBookmarksOrderBy{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetBookmarksOrderByArrayOfstr(arrayOfstr []string) GetBookmarksOrderBy {
	typ := GetBookmarksOrderByTypeArrayOfstr

	return GetBookmarksOrderBy{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetBookmarksOrderBy) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetBookmarksOrderByTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetBookmarksOrderByTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksOrderBy) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetBookmarksWebhookIDType string

const (
	GetBookmarksWebhookIDTypeStr        GetBookmarksWebhookIDType = "str"
	GetBookmarksWebhookIDTypeArrayOfstr GetBookmarksWebhookIDType = "arrayOfstr"
)

type GetBookmarksWebhookID struct {
	Str        *string
	ArrayOfstr []string

	Type GetBookmarksWebhookIDType
}

func CreateGetBookmarksWebhookIDStr(str string) GetBookmarksWebhookID {
	typ := GetBookmarksWebhookIDTypeStr

	return GetBookmarksWebhookID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetBookmarksWebhookIDArrayOfstr(arrayOfstr []string) GetBookmarksWebhookID {
	typ := GetBookmarksWebhookIDTypeArrayOfstr

	return GetBookmarksWebhookID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetBookmarksWebhookID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetBookmarksWebhookIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetBookmarksWebhookIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetBookmarksWebhookID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetBookmarksRequest struct {
	Dir         *GetBookmarksDir         `queryParam:"style=form,explode=true,name=dir"`
	EventDataID *GetBookmarksEventDataID `queryParam:"style=form,explode=true,name=event_data_id"`
	ID          *GetBookmarksID          `queryParam:"style=form,explode=true,name=id"`
	Label       *GetBookmarksLabel       `queryParam:"style=form,explode=true,name=label"`
	LastUsedAt  *GetBookmarksLastUsedAt  `queryParam:"style=form,explode=true,name=last_used_at"`
	Limit       *int64                   `queryParam:"style=form,explode=true,name=limit"`
	Name        *GetBookmarksName        `queryParam:"style=form,explode=true,name=name"`
	Next        *string                  `queryParam:"style=form,explode=true,name=next"`
	OrderBy     *GetBookmarksOrderBy     `queryParam:"style=form,explode=true,name=order_by"`
	Prev        *string                  `queryParam:"style=form,explode=true,name=prev"`
	WebhookID   *GetBookmarksWebhookID   `queryParam:"style=form,explode=true,name=webhook_id"`
}

func (o *GetBookmarksRequest) GetDir() *GetBookmarksDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetBookmarksRequest) GetEventDataID() *GetBookmarksEventDataID {
	if o == nil {
		return nil
	}
	return o.EventDataID
}

func (o *GetBookmarksRequest) GetID() *GetBookmarksID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetBookmarksRequest) GetLabel() *GetBookmarksLabel {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *GetBookmarksRequest) GetLastUsedAt() *GetBookmarksLastUsedAt {
	if o == nil {
		return nil
	}
	return o.LastUsedAt
}

func (o *GetBookmarksRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetBookmarksRequest) GetName() *GetBookmarksName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetBookmarksRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetBookmarksRequest) GetOrderBy() *GetBookmarksOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetBookmarksRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

func (o *GetBookmarksRequest) GetWebhookID() *GetBookmarksWebhookID {
	if o == nil {
		return nil
	}
	return o.WebhookID
}

type GetBookmarksResponse struct {
	// List of bookmarks
	BookmarkPaginatedResult *shared.BookmarkPaginatedResult
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBookmarksResponse) GetBookmarkPaginatedResult() *shared.BookmarkPaginatedResult {
	if o == nil {
		return nil
	}
	return o.BookmarkPaginatedResult
}

func (o *GetBookmarksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBookmarksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBookmarksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

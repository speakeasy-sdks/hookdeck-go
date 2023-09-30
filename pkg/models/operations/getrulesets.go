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

// GetRulesetsArchivedAt2 - Date the ruleset was archived
type GetRulesetsArchivedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetRulesetsArchivedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRulesetsArchivedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetRulesetsArchivedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetRulesetsArchivedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetRulesetsArchivedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetRulesetsArchivedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetRulesetsArchivedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetRulesetsArchivedAtType string

const (
	GetRulesetsArchivedAtTypeDateTime               GetRulesetsArchivedAtType = "date-time"
	GetRulesetsArchivedAtTypeGetRulesetsArchivedAt2 GetRulesetsArchivedAtType = "getRulesetsArchivedAt_2"
)

type GetRulesetsArchivedAt struct {
	DateTime               *time.Time
	GetRulesetsArchivedAt2 *GetRulesetsArchivedAt2

	Type GetRulesetsArchivedAtType
}

func CreateGetRulesetsArchivedAtDateTime(dateTime time.Time) GetRulesetsArchivedAt {
	typ := GetRulesetsArchivedAtTypeDateTime

	return GetRulesetsArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetRulesetsArchivedAtGetRulesetsArchivedAt2(getRulesetsArchivedAt2 GetRulesetsArchivedAt2) GetRulesetsArchivedAt {
	typ := GetRulesetsArchivedAtTypeGetRulesetsArchivedAt2

	return GetRulesetsArchivedAt{
		GetRulesetsArchivedAt2: &getRulesetsArchivedAt2,
		Type:                   typ,
	}
}

func (u *GetRulesetsArchivedAt) UnmarshalJSON(data []byte) error {

	getRulesetsArchivedAt2 := new(GetRulesetsArchivedAt2)
	if err := utils.UnmarshalJSON(data, &getRulesetsArchivedAt2, "", true, true); err == nil {
		u.GetRulesetsArchivedAt2 = getRulesetsArchivedAt2
		u.Type = GetRulesetsArchivedAtTypeGetRulesetsArchivedAt2
		return nil
	}

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = GetRulesetsArchivedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetRulesetsArchivedAt2 != nil {
		return utils.MarshalJSON(u.GetRulesetsArchivedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsDir2 string

const (
	GetRulesetsDir2Asc  GetRulesetsDir2 = "asc"
	GetRulesetsDir2Desc GetRulesetsDir2 = "desc"
)

func (e GetRulesetsDir2) ToPointer() *GetRulesetsDir2 {
	return &e
}

func (e *GetRulesetsDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRulesetsDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsDir2: %v", v)
	}
}

// GetRulesetsDir1 - Sort direction
type GetRulesetsDir1 string

const (
	GetRulesetsDir1Asc  GetRulesetsDir1 = "asc"
	GetRulesetsDir1Desc GetRulesetsDir1 = "desc"
)

func (e GetRulesetsDir1) ToPointer() *GetRulesetsDir1 {
	return &e
}

func (e *GetRulesetsDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRulesetsDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsDir1: %v", v)
	}
}

type GetRulesetsDirType string

const (
	GetRulesetsDirTypeGetRulesetsDir1        GetRulesetsDirType = "getRulesetsDir_1"
	GetRulesetsDirTypeArrayOfgetRulesetsDir2 GetRulesetsDirType = "arrayOfgetRulesetsDir_2"
)

type GetRulesetsDir struct {
	GetRulesetsDir1        *GetRulesetsDir1
	ArrayOfgetRulesetsDir2 []GetRulesetsDir2

	Type GetRulesetsDirType
}

func CreateGetRulesetsDirGetRulesetsDir1(getRulesetsDir1 GetRulesetsDir1) GetRulesetsDir {
	typ := GetRulesetsDirTypeGetRulesetsDir1

	return GetRulesetsDir{
		GetRulesetsDir1: &getRulesetsDir1,
		Type:            typ,
	}
}

func CreateGetRulesetsDirArrayOfgetRulesetsDir2(arrayOfgetRulesetsDir2 []GetRulesetsDir2) GetRulesetsDir {
	typ := GetRulesetsDirTypeArrayOfgetRulesetsDir2

	return GetRulesetsDir{
		ArrayOfgetRulesetsDir2: arrayOfgetRulesetsDir2,
		Type:                   typ,
	}
}

func (u *GetRulesetsDir) UnmarshalJSON(data []byte) error {

	getRulesetsDir1 := new(GetRulesetsDir1)
	if err := utils.UnmarshalJSON(data, &getRulesetsDir1, "", true, true); err == nil {
		u.GetRulesetsDir1 = getRulesetsDir1
		u.Type = GetRulesetsDirTypeGetRulesetsDir1
		return nil
	}

	arrayOfgetRulesetsDir2 := []GetRulesetsDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRulesetsDir2, "", true, true); err == nil {
		u.ArrayOfgetRulesetsDir2 = arrayOfgetRulesetsDir2
		u.Type = GetRulesetsDirTypeArrayOfgetRulesetsDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsDir) MarshalJSON() ([]byte, error) {
	if u.GetRulesetsDir1 != nil {
		return utils.MarshalJSON(u.GetRulesetsDir1, "", true)
	}

	if u.ArrayOfgetRulesetsDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRulesetsDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsIDType string

const (
	GetRulesetsIDTypeStr        GetRulesetsIDType = "str"
	GetRulesetsIDTypeArrayOfstr GetRulesetsIDType = "arrayOfstr"
)

type GetRulesetsID struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsIDType
}

func CreateGetRulesetsIDStr(str string) GetRulesetsID {
	typ := GetRulesetsIDTypeStr

	return GetRulesetsID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsIDArrayOfstr(arrayOfstr []string) GetRulesetsID {
	typ := GetRulesetsIDTypeArrayOfstr

	return GetRulesetsID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsNameContainsType string

const (
	GetRulesetsNameContainsTypeStr        GetRulesetsNameContainsType = "str"
	GetRulesetsNameContainsTypeArrayOfstr GetRulesetsNameContainsType = "arrayOfstr"
)

type GetRulesetsNameContains struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsNameContainsType
}

func CreateGetRulesetsNameContainsStr(str string) GetRulesetsNameContains {
	typ := GetRulesetsNameContainsTypeStr

	return GetRulesetsNameContains{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsNameContainsArrayOfstr(arrayOfstr []string) GetRulesetsNameContains {
	typ := GetRulesetsNameContainsTypeArrayOfstr

	return GetRulesetsNameContains{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsNameContains) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsNameContainsTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsNameContainsTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsNameContains) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsNameGtType string

const (
	GetRulesetsNameGtTypeStr        GetRulesetsNameGtType = "str"
	GetRulesetsNameGtTypeArrayOfstr GetRulesetsNameGtType = "arrayOfstr"
)

type GetRulesetsNameGt struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsNameGtType
}

func CreateGetRulesetsNameGtStr(str string) GetRulesetsNameGt {
	typ := GetRulesetsNameGtTypeStr

	return GetRulesetsNameGt{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsNameGtArrayOfstr(arrayOfstr []string) GetRulesetsNameGt {
	typ := GetRulesetsNameGtTypeArrayOfstr

	return GetRulesetsNameGt{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsNameGt) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsNameGtTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsNameGtTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsNameGt) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsNameGteType string

const (
	GetRulesetsNameGteTypeStr        GetRulesetsNameGteType = "str"
	GetRulesetsNameGteTypeArrayOfstr GetRulesetsNameGteType = "arrayOfstr"
)

type GetRulesetsNameGte struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsNameGteType
}

func CreateGetRulesetsNameGteStr(str string) GetRulesetsNameGte {
	typ := GetRulesetsNameGteTypeStr

	return GetRulesetsNameGte{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsNameGteArrayOfstr(arrayOfstr []string) GetRulesetsNameGte {
	typ := GetRulesetsNameGteTypeArrayOfstr

	return GetRulesetsNameGte{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsNameGte) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsNameGteTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsNameGteTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsNameGte) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsNameLeType string

const (
	GetRulesetsNameLeTypeStr        GetRulesetsNameLeType = "str"
	GetRulesetsNameLeTypeArrayOfstr GetRulesetsNameLeType = "arrayOfstr"
)

type GetRulesetsNameLe struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsNameLeType
}

func CreateGetRulesetsNameLeStr(str string) GetRulesetsNameLe {
	typ := GetRulesetsNameLeTypeStr

	return GetRulesetsNameLe{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsNameLeArrayOfstr(arrayOfstr []string) GetRulesetsNameLe {
	typ := GetRulesetsNameLeTypeArrayOfstr

	return GetRulesetsNameLe{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsNameLe) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsNameLeTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsNameLeTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsNameLe) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsNameLteType string

const (
	GetRulesetsNameLteTypeStr        GetRulesetsNameLteType = "str"
	GetRulesetsNameLteTypeArrayOfstr GetRulesetsNameLteType = "arrayOfstr"
)

type GetRulesetsNameLte struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsNameLteType
}

func CreateGetRulesetsNameLteStr(str string) GetRulesetsNameLte {
	typ := GetRulesetsNameLteTypeStr

	return GetRulesetsNameLte{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsNameLteArrayOfstr(arrayOfstr []string) GetRulesetsNameLte {
	typ := GetRulesetsNameLteTypeArrayOfstr

	return GetRulesetsNameLte{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsNameLte) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsNameLteTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsNameLteTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsNameLte) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsName struct {
	Any      *bool                    `queryParam:"name=any"`
	Contains *GetRulesetsNameContains `queryParam:"name=contains"`
	Gt       *GetRulesetsNameGt       `queryParam:"name=gt"`
	Gte      *GetRulesetsNameGte      `queryParam:"name=gte"`
	Le       *GetRulesetsNameLe       `queryParam:"name=le"`
	Lte      *GetRulesetsNameLte      `queryParam:"name=lte"`
}

func (o *GetRulesetsName) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetRulesetsName) GetContains() *GetRulesetsNameContains {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *GetRulesetsName) GetGt() *GetRulesetsNameGt {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetRulesetsName) GetGte() *GetRulesetsNameGte {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetRulesetsName) GetLe() *GetRulesetsNameLe {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetRulesetsName) GetLte() *GetRulesetsNameLte {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetRulesetsOrderByType string

const (
	GetRulesetsOrderByTypeStr        GetRulesetsOrderByType = "str"
	GetRulesetsOrderByTypeArrayOfstr GetRulesetsOrderByType = "arrayOfstr"
)

type GetRulesetsOrderBy struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsOrderByType
}

func CreateGetRulesetsOrderByStr(str string) GetRulesetsOrderBy {
	typ := GetRulesetsOrderByTypeStr

	return GetRulesetsOrderBy{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsOrderByArrayOfstr(arrayOfstr []string) GetRulesetsOrderBy {
	typ := GetRulesetsOrderByTypeArrayOfstr

	return GetRulesetsOrderBy{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsOrderBy) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetRulesetsOrderByTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsOrderByTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsOrderBy) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsRequest struct {
	Archived   *bool                  `queryParam:"style=form,explode=true,name=archived"`
	ArchivedAt *GetRulesetsArchivedAt `queryParam:"style=form,explode=true,name=archived_at"`
	Dir        *GetRulesetsDir        `queryParam:"style=form,explode=true,name=dir"`
	ID         *GetRulesetsID         `queryParam:"style=form,explode=true,name=id"`
	Limit      *int64                 `queryParam:"style=form,explode=true,name=limit"`
	Name       *GetRulesetsName       `queryParam:"style=form,explode=true,name=name"`
	Next       *string                `queryParam:"style=form,explode=true,name=next"`
	OrderBy    *GetRulesetsOrderBy    `queryParam:"style=form,explode=true,name=order_by"`
	Prev       *string                `queryParam:"style=form,explode=true,name=prev"`
}

func (o *GetRulesetsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetRulesetsRequest) GetArchivedAt() *GetRulesetsArchivedAt {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetRulesetsRequest) GetDir() *GetRulesetsDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetRulesetsRequest) GetID() *GetRulesetsID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetRulesetsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetRulesetsRequest) GetName() *GetRulesetsName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetRulesetsRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetRulesetsRequest) GetOrderBy() *GetRulesetsOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetRulesetsRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type GetRulesetsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of rulesets
	RulesetPaginatedResult *shared.RulesetPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRulesetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRulesetsResponse) GetRulesetPaginatedResult() *shared.RulesetPaginatedResult {
	if o == nil {
		return nil
	}
	return o.RulesetPaginatedResult
}

func (o *GetRulesetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRulesetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

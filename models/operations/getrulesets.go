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

type GetRulesetsQueryParam2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetRulesetsQueryParam2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRulesetsQueryParam2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetRulesetsQueryParam2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetRulesetsQueryParam2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetRulesetsQueryParam2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetRulesetsQueryParam2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetRulesetsQueryParam2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetRulesetsQueryParamArchivedAtType string

const (
	GetRulesetsQueryParamArchivedAtTypeDateTime               GetRulesetsQueryParamArchivedAtType = "date-time"
	GetRulesetsQueryParamArchivedAtTypeGetRulesetsQueryParam2 GetRulesetsQueryParamArchivedAtType = "getRulesets_queryParam_2"
)

// GetRulesetsQueryParamArchivedAt - Date the ruleset was archived
type GetRulesetsQueryParamArchivedAt struct {
	DateTime               *time.Time
	GetRulesetsQueryParam2 *GetRulesetsQueryParam2

	Type GetRulesetsQueryParamArchivedAtType
}

func CreateGetRulesetsQueryParamArchivedAtDateTime(dateTime time.Time) GetRulesetsQueryParamArchivedAt {
	typ := GetRulesetsQueryParamArchivedAtTypeDateTime

	return GetRulesetsQueryParamArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetRulesetsQueryParamArchivedAtGetRulesetsQueryParam2(getRulesetsQueryParam2 GetRulesetsQueryParam2) GetRulesetsQueryParamArchivedAt {
	typ := GetRulesetsQueryParamArchivedAtTypeGetRulesetsQueryParam2

	return GetRulesetsQueryParamArchivedAt{
		GetRulesetsQueryParam2: &getRulesetsQueryParam2,
		Type:                   typ,
	}
}

func (u *GetRulesetsQueryParamArchivedAt) UnmarshalJSON(data []byte) error {

	getRulesetsQueryParam2 := GetRulesetsQueryParam2{}
	if err := utils.UnmarshalJSON(data, &getRulesetsQueryParam2, "", true, true); err == nil {
		u.GetRulesetsQueryParam2 = &getRulesetsQueryParam2
		u.Type = GetRulesetsQueryParamArchivedAtTypeGetRulesetsQueryParam2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetRulesetsQueryParamArchivedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsQueryParamArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetRulesetsQueryParam2 != nil {
		return utils.MarshalJSON(u.GetRulesetsQueryParam2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsQueryParamRulesets2 string

const (
	GetRulesetsQueryParamRulesets2Asc  GetRulesetsQueryParamRulesets2 = "asc"
	GetRulesetsQueryParamRulesets2Desc GetRulesetsQueryParamRulesets2 = "desc"
)

func (e GetRulesetsQueryParamRulesets2) ToPointer() *GetRulesetsQueryParamRulesets2 {
	return &e
}

func (e *GetRulesetsQueryParamRulesets2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRulesetsQueryParamRulesets2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsQueryParamRulesets2: %v", v)
	}
}

type GetRulesetsQueryParam1 string

const (
	GetRulesetsQueryParam1Asc  GetRulesetsQueryParam1 = "asc"
	GetRulesetsQueryParam1Desc GetRulesetsQueryParam1 = "desc"
)

func (e GetRulesetsQueryParam1) ToPointer() *GetRulesetsQueryParam1 {
	return &e
}

func (e *GetRulesetsQueryParam1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetRulesetsQueryParam1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsQueryParam1: %v", v)
	}
}

type GetRulesetsQueryParamDirType string

const (
	GetRulesetsQueryParamDirTypeGetRulesetsQueryParam1                GetRulesetsQueryParamDirType = "getRulesets_queryParam_1"
	GetRulesetsQueryParamDirTypeArrayOfgetRulesetsQueryParamRulesets2 GetRulesetsQueryParamDirType = "arrayOfgetRulesets_queryParam_Rulesets_2"
)

// GetRulesetsQueryParamDir - Sort direction
type GetRulesetsQueryParamDir struct {
	GetRulesetsQueryParam1                *GetRulesetsQueryParam1
	ArrayOfgetRulesetsQueryParamRulesets2 []GetRulesetsQueryParamRulesets2

	Type GetRulesetsQueryParamDirType
}

func CreateGetRulesetsQueryParamDirGetRulesetsQueryParam1(getRulesetsQueryParam1 GetRulesetsQueryParam1) GetRulesetsQueryParamDir {
	typ := GetRulesetsQueryParamDirTypeGetRulesetsQueryParam1

	return GetRulesetsQueryParamDir{
		GetRulesetsQueryParam1: &getRulesetsQueryParam1,
		Type:                   typ,
	}
}

func CreateGetRulesetsQueryParamDirArrayOfgetRulesetsQueryParamRulesets2(arrayOfgetRulesetsQueryParamRulesets2 []GetRulesetsQueryParamRulesets2) GetRulesetsQueryParamDir {
	typ := GetRulesetsQueryParamDirTypeArrayOfgetRulesetsQueryParamRulesets2

	return GetRulesetsQueryParamDir{
		ArrayOfgetRulesetsQueryParamRulesets2: arrayOfgetRulesetsQueryParamRulesets2,
		Type:                                  typ,
	}
}

func (u *GetRulesetsQueryParamDir) UnmarshalJSON(data []byte) error {

	getRulesetsQueryParam1 := GetRulesetsQueryParam1("")
	if err := utils.UnmarshalJSON(data, &getRulesetsQueryParam1, "", true, true); err == nil {
		u.GetRulesetsQueryParam1 = &getRulesetsQueryParam1
		u.Type = GetRulesetsQueryParamDirTypeGetRulesetsQueryParam1
		return nil
	}

	arrayOfgetRulesetsQueryParamRulesets2 := []GetRulesetsQueryParamRulesets2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRulesetsQueryParamRulesets2, "", true, true); err == nil {
		u.ArrayOfgetRulesetsQueryParamRulesets2 = arrayOfgetRulesetsQueryParamRulesets2
		u.Type = GetRulesetsQueryParamDirTypeArrayOfgetRulesetsQueryParamRulesets2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsQueryParamDir) MarshalJSON() ([]byte, error) {
	if u.GetRulesetsQueryParam1 != nil {
		return utils.MarshalJSON(u.GetRulesetsQueryParam1, "", true)
	}

	if u.ArrayOfgetRulesetsQueryParamRulesets2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRulesetsQueryParamRulesets2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsQueryParamIDType string

const (
	GetRulesetsQueryParamIDTypeStr        GetRulesetsQueryParamIDType = "str"
	GetRulesetsQueryParamIDTypeArrayOfstr GetRulesetsQueryParamIDType = "arrayOfstr"
)

// GetRulesetsQueryParamID - Filter by ruleset IDs
type GetRulesetsQueryParamID struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsQueryParamIDType
}

func CreateGetRulesetsQueryParamIDStr(str string) GetRulesetsQueryParamID {
	typ := GetRulesetsQueryParamIDTypeStr

	return GetRulesetsQueryParamID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsQueryParamIDArrayOfstr(arrayOfstr []string) GetRulesetsQueryParamID {
	typ := GetRulesetsQueryParamIDTypeArrayOfstr

	return GetRulesetsQueryParamID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsQueryParamID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetRulesetsQueryParamIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsQueryParamIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsQueryParamID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ContainsType string

const (
	ContainsTypeStr        ContainsType = "str"
	ContainsTypeArrayOfstr ContainsType = "arrayOfstr"
)

type Contains struct {
	Str        *string
	ArrayOfstr []string

	Type ContainsType
}

func CreateContainsStr(str string) Contains {
	typ := ContainsTypeStr

	return Contains{
		Str:  &str,
		Type: typ,
	}
}

func CreateContainsArrayOfstr(arrayOfstr []string) Contains {
	typ := ContainsTypeArrayOfstr

	return Contains{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *Contains) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ContainsTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = ContainsTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Contains) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GtType string

const (
	GtTypeStr        GtType = "str"
	GtTypeArrayOfstr GtType = "arrayOfstr"
)

type Gt struct {
	Str        *string
	ArrayOfstr []string

	Type GtType
}

func CreateGtStr(str string) Gt {
	typ := GtTypeStr

	return Gt{
		Str:  &str,
		Type: typ,
	}
}

func CreateGtArrayOfstr(arrayOfstr []string) Gt {
	typ := GtTypeArrayOfstr

	return Gt{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *Gt) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GtTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GtTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Gt) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GteType string

const (
	GteTypeStr        GteType = "str"
	GteTypeArrayOfstr GteType = "arrayOfstr"
)

type Gte struct {
	Str        *string
	ArrayOfstr []string

	Type GteType
}

func CreateGteStr(str string) Gte {
	typ := GteTypeStr

	return Gte{
		Str:  &str,
		Type: typ,
	}
}

func CreateGteArrayOfstr(arrayOfstr []string) Gte {
	typ := GteTypeArrayOfstr

	return Gte{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *Gte) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GteTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GteTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Gte) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type LeType string

const (
	LeTypeStr        LeType = "str"
	LeTypeArrayOfstr LeType = "arrayOfstr"
)

type Le struct {
	Str        *string
	ArrayOfstr []string

	Type LeType
}

func CreateLeStr(str string) Le {
	typ := LeTypeStr

	return Le{
		Str:  &str,
		Type: typ,
	}
}

func CreateLeArrayOfstr(arrayOfstr []string) Le {
	typ := LeTypeArrayOfstr

	return Le{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *Le) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = LeTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = LeTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Le) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type LteType string

const (
	LteTypeStr        LteType = "str"
	LteTypeArrayOfstr LteType = "arrayOfstr"
)

type Lte struct {
	Str        *string
	ArrayOfstr []string

	Type LteType
}

func CreateLteStr(str string) Lte {
	typ := LteTypeStr

	return Lte{
		Str:  &str,
		Type: typ,
	}
}

func CreateLteArrayOfstr(arrayOfstr []string) Lte {
	typ := LteTypeArrayOfstr

	return Lte{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *Lte) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = LteTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = LteTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Lte) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsQueryParamRulesetsName2 struct {
	Any      *bool     `queryParam:"name=any"`
	Contains *Contains `queryParam:"name=contains"`
	Gt       *Gt       `queryParam:"name=gt"`
	Gte      *Gte      `queryParam:"name=gte"`
	Le       *Le       `queryParam:"name=le"`
	Lte      *Lte      `queryParam:"name=lte"`
}

func (o *GetRulesetsQueryParamRulesetsName2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetRulesetsQueryParamRulesetsName2) GetContains() *Contains {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *GetRulesetsQueryParamRulesetsName2) GetGt() *Gt {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetRulesetsQueryParamRulesetsName2) GetGte() *Gte {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetRulesetsQueryParamRulesetsName2) GetLe() *Le {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetRulesetsQueryParamRulesetsName2) GetLte() *Lte {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetRulesetsQueryParamRulesets1Type string

const (
	GetRulesetsQueryParamRulesets1TypeStr        GetRulesetsQueryParamRulesets1Type = "str"
	GetRulesetsQueryParamRulesets1TypeArrayOfstr GetRulesetsQueryParamRulesets1Type = "arrayOfstr"
)

type GetRulesetsQueryParamRulesets1 struct {
	Str        *string
	ArrayOfstr []string

	Type GetRulesetsQueryParamRulesets1Type
}

func CreateGetRulesetsQueryParamRulesets1Str(str string) GetRulesetsQueryParamRulesets1 {
	typ := GetRulesetsQueryParamRulesets1TypeStr

	return GetRulesetsQueryParamRulesets1{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetRulesetsQueryParamRulesets1ArrayOfstr(arrayOfstr []string) GetRulesetsQueryParamRulesets1 {
	typ := GetRulesetsQueryParamRulesets1TypeArrayOfstr

	return GetRulesetsQueryParamRulesets1{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetRulesetsQueryParamRulesets1) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetRulesetsQueryParamRulesets1TypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetRulesetsQueryParamRulesets1TypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsQueryParamRulesets1) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsQueryParamNameType string

const (
	GetRulesetsQueryParamNameTypeGetRulesetsQueryParamRulesets1     GetRulesetsQueryParamNameType = "getRulesets_queryParam_Rulesets_1"
	GetRulesetsQueryParamNameTypeGetRulesetsQueryParamRulesetsName2 GetRulesetsQueryParamNameType = "getRulesets_queryParam_Rulesets_name_2"
)

// GetRulesetsQueryParamName - The ruleset name
type GetRulesetsQueryParamName struct {
	GetRulesetsQueryParamRulesets1     *GetRulesetsQueryParamRulesets1
	GetRulesetsQueryParamRulesetsName2 *GetRulesetsQueryParamRulesetsName2

	Type GetRulesetsQueryParamNameType
}

func CreateGetRulesetsQueryParamNameGetRulesetsQueryParamRulesets1(getRulesetsQueryParamRulesets1 GetRulesetsQueryParamRulesets1) GetRulesetsQueryParamName {
	typ := GetRulesetsQueryParamNameTypeGetRulesetsQueryParamRulesets1

	return GetRulesetsQueryParamName{
		GetRulesetsQueryParamRulesets1: &getRulesetsQueryParamRulesets1,
		Type:                           typ,
	}
}

func CreateGetRulesetsQueryParamNameGetRulesetsQueryParamRulesetsName2(getRulesetsQueryParamRulesetsName2 GetRulesetsQueryParamRulesetsName2) GetRulesetsQueryParamName {
	typ := GetRulesetsQueryParamNameTypeGetRulesetsQueryParamRulesetsName2

	return GetRulesetsQueryParamName{
		GetRulesetsQueryParamRulesetsName2: &getRulesetsQueryParamRulesetsName2,
		Type:                               typ,
	}
}

func (u *GetRulesetsQueryParamName) UnmarshalJSON(data []byte) error {

	getRulesetsQueryParamRulesetsName2 := GetRulesetsQueryParamRulesetsName2{}
	if err := utils.UnmarshalJSON(data, &getRulesetsQueryParamRulesetsName2, "", true, true); err == nil {
		u.GetRulesetsQueryParamRulesetsName2 = &getRulesetsQueryParamRulesetsName2
		u.Type = GetRulesetsQueryParamNameTypeGetRulesetsQueryParamRulesetsName2
		return nil
	}

	getRulesetsQueryParamRulesets1 := GetRulesetsQueryParamRulesets1{}
	if err := utils.UnmarshalJSON(data, &getRulesetsQueryParamRulesets1, "", true, true); err == nil {
		u.GetRulesetsQueryParamRulesets1 = &getRulesetsQueryParamRulesets1
		u.Type = GetRulesetsQueryParamNameTypeGetRulesetsQueryParamRulesets1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsQueryParamName) MarshalJSON() ([]byte, error) {
	if u.GetRulesetsQueryParamRulesets1 != nil {
		return utils.MarshalJSON(u.GetRulesetsQueryParamRulesets1, "", true)
	}

	if u.GetRulesetsQueryParamRulesetsName2 != nil {
		return utils.MarshalJSON(u.GetRulesetsQueryParamRulesetsName2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsQueryParamRulesetsOrderBy2 string

const (
	GetRulesetsQueryParamRulesetsOrderBy2CreatedAt GetRulesetsQueryParamRulesetsOrderBy2 = "created_at"
)

func (e GetRulesetsQueryParamRulesetsOrderBy2) ToPointer() *GetRulesetsQueryParamRulesetsOrderBy2 {
	return &e
}

func (e *GetRulesetsQueryParamRulesetsOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetRulesetsQueryParamRulesetsOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsQueryParamRulesetsOrderBy2: %v", v)
	}
}

type GetRulesetsQueryParamRulesetsOrderBy1 string

const (
	GetRulesetsQueryParamRulesetsOrderBy1CreatedAt GetRulesetsQueryParamRulesetsOrderBy1 = "created_at"
)

func (e GetRulesetsQueryParamRulesetsOrderBy1) ToPointer() *GetRulesetsQueryParamRulesetsOrderBy1 {
	return &e
}

func (e *GetRulesetsQueryParamRulesetsOrderBy1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetRulesetsQueryParamRulesetsOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRulesetsQueryParamRulesetsOrderBy1: %v", v)
	}
}

type GetRulesetsQueryParamOrderByType string

const (
	GetRulesetsQueryParamOrderByTypeGetRulesetsQueryParamRulesetsOrderBy1        GetRulesetsQueryParamOrderByType = "getRulesets_queryParam_Rulesets_order_by_1"
	GetRulesetsQueryParamOrderByTypeArrayOfgetRulesetsQueryParamRulesetsOrderBy2 GetRulesetsQueryParamOrderByType = "arrayOfgetRulesets_queryParam_Rulesets_order_by_2"
)

// GetRulesetsQueryParamOrderBy - Sort key(s)
type GetRulesetsQueryParamOrderBy struct {
	GetRulesetsQueryParamRulesetsOrderBy1        *GetRulesetsQueryParamRulesetsOrderBy1
	ArrayOfgetRulesetsQueryParamRulesetsOrderBy2 []GetRulesetsQueryParamRulesetsOrderBy2

	Type GetRulesetsQueryParamOrderByType
}

func CreateGetRulesetsQueryParamOrderByGetRulesetsQueryParamRulesetsOrderBy1(getRulesetsQueryParamRulesetsOrderBy1 GetRulesetsQueryParamRulesetsOrderBy1) GetRulesetsQueryParamOrderBy {
	typ := GetRulesetsQueryParamOrderByTypeGetRulesetsQueryParamRulesetsOrderBy1

	return GetRulesetsQueryParamOrderBy{
		GetRulesetsQueryParamRulesetsOrderBy1: &getRulesetsQueryParamRulesetsOrderBy1,
		Type:                                  typ,
	}
}

func CreateGetRulesetsQueryParamOrderByArrayOfgetRulesetsQueryParamRulesetsOrderBy2(arrayOfgetRulesetsQueryParamRulesetsOrderBy2 []GetRulesetsQueryParamRulesetsOrderBy2) GetRulesetsQueryParamOrderBy {
	typ := GetRulesetsQueryParamOrderByTypeArrayOfgetRulesetsQueryParamRulesetsOrderBy2

	return GetRulesetsQueryParamOrderBy{
		ArrayOfgetRulesetsQueryParamRulesetsOrderBy2: arrayOfgetRulesetsQueryParamRulesetsOrderBy2,
		Type: typ,
	}
}

func (u *GetRulesetsQueryParamOrderBy) UnmarshalJSON(data []byte) error {

	getRulesetsQueryParamRulesetsOrderBy1 := GetRulesetsQueryParamRulesetsOrderBy1("")
	if err := utils.UnmarshalJSON(data, &getRulesetsQueryParamRulesetsOrderBy1, "", true, true); err == nil {
		u.GetRulesetsQueryParamRulesetsOrderBy1 = &getRulesetsQueryParamRulesetsOrderBy1
		u.Type = GetRulesetsQueryParamOrderByTypeGetRulesetsQueryParamRulesetsOrderBy1
		return nil
	}

	arrayOfgetRulesetsQueryParamRulesetsOrderBy2 := []GetRulesetsQueryParamRulesetsOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetRulesetsQueryParamRulesetsOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetRulesetsQueryParamRulesetsOrderBy2 = arrayOfgetRulesetsQueryParamRulesetsOrderBy2
		u.Type = GetRulesetsQueryParamOrderByTypeArrayOfgetRulesetsQueryParamRulesetsOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetRulesetsQueryParamOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetRulesetsQueryParamRulesetsOrderBy1 != nil {
		return utils.MarshalJSON(u.GetRulesetsQueryParamRulesetsOrderBy1, "", true)
	}

	if u.ArrayOfgetRulesetsQueryParamRulesetsOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetRulesetsQueryParamRulesetsOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetRulesetsRequest struct {
	Archived *bool `queryParam:"style=form,explode=true,name=archived"`
	// Date the ruleset was archived
	ArchivedAt *GetRulesetsQueryParamArchivedAt `queryParam:"style=form,explode=true,name=archived_at"`
	// Sort direction
	Dir *GetRulesetsQueryParamDir `queryParam:"style=form,explode=true,name=dir"`
	// Filter by ruleset IDs
	ID    *GetRulesetsQueryParamID `queryParam:"style=form,explode=true,name=id"`
	Limit *int64                   `queryParam:"style=form,explode=true,name=limit"`
	// The ruleset name
	Name *GetRulesetsQueryParamName `queryParam:"style=form,explode=true,name=name"`
	Next *string                    `queryParam:"style=form,explode=true,name=next"`
	// Sort key(s)
	OrderBy *GetRulesetsQueryParamOrderBy `queryParam:"style=form,explode=true,name=order_by"`
	Prev    *string                       `queryParam:"style=form,explode=true,name=prev"`
}

func (o *GetRulesetsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetRulesetsRequest) GetArchivedAt() *GetRulesetsQueryParamArchivedAt {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetRulesetsRequest) GetDir() *GetRulesetsQueryParamDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetRulesetsRequest) GetID() *GetRulesetsQueryParamID {
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

func (o *GetRulesetsRequest) GetName() *GetRulesetsQueryParamName {
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

func (o *GetRulesetsRequest) GetOrderBy() *GetRulesetsQueryParamOrderBy {
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
	RulesetPaginatedResult *components.RulesetPaginatedResult
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

func (o *GetRulesetsResponse) GetRulesetPaginatedResult() *components.RulesetPaginatedResult {
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

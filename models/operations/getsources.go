// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/v2/internal/utils"
	"github.com/speakeasy-sdks/hookdeck-go/v2/models/components"
	"net/http"
	"time"
)

type GetSourcesQueryParam2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetSourcesQueryParam2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetSourcesQueryParam2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetSourcesQueryParam2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetSourcesQueryParam2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetSourcesQueryParam2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetSourcesQueryParam2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetSourcesQueryParam2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetSourcesQueryParamArchivedAtType string

const (
	GetSourcesQueryParamArchivedAtTypeDateTime              GetSourcesQueryParamArchivedAtType = "date-time"
	GetSourcesQueryParamArchivedAtTypeGetSourcesQueryParam2 GetSourcesQueryParamArchivedAtType = "getSources_queryParam_2"
)

// GetSourcesQueryParamArchivedAt - Date the source was archived
type GetSourcesQueryParamArchivedAt struct {
	DateTime              *time.Time
	GetSourcesQueryParam2 *GetSourcesQueryParam2

	Type GetSourcesQueryParamArchivedAtType
}

func CreateGetSourcesQueryParamArchivedAtDateTime(dateTime time.Time) GetSourcesQueryParamArchivedAt {
	typ := GetSourcesQueryParamArchivedAtTypeDateTime

	return GetSourcesQueryParamArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetSourcesQueryParamArchivedAtGetSourcesQueryParam2(getSourcesQueryParam2 GetSourcesQueryParam2) GetSourcesQueryParamArchivedAt {
	typ := GetSourcesQueryParamArchivedAtTypeGetSourcesQueryParam2

	return GetSourcesQueryParamArchivedAt{
		GetSourcesQueryParam2: &getSourcesQueryParam2,
		Type:                  typ,
	}
}

func (u *GetSourcesQueryParamArchivedAt) UnmarshalJSON(data []byte) error {

	getSourcesQueryParam2 := GetSourcesQueryParam2{}
	if err := utils.UnmarshalJSON(data, &getSourcesQueryParam2, "", true, true); err == nil {
		u.GetSourcesQueryParam2 = &getSourcesQueryParam2
		u.Type = GetSourcesQueryParamArchivedAtTypeGetSourcesQueryParam2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetSourcesQueryParamArchivedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesQueryParamArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetSourcesQueryParam2 != nil {
		return utils.MarshalJSON(u.GetSourcesQueryParam2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetSourcesQueryParamSources2 string

const (
	GetSourcesQueryParamSources2Asc  GetSourcesQueryParamSources2 = "asc"
	GetSourcesQueryParamSources2Desc GetSourcesQueryParamSources2 = "desc"
)

func (e GetSourcesQueryParamSources2) ToPointer() *GetSourcesQueryParamSources2 {
	return &e
}

func (e *GetSourcesQueryParamSources2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetSourcesQueryParamSources2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesQueryParamSources2: %v", v)
	}
}

type GetSourcesQueryParam1 string

const (
	GetSourcesQueryParam1Asc  GetSourcesQueryParam1 = "asc"
	GetSourcesQueryParam1Desc GetSourcesQueryParam1 = "desc"
)

func (e GetSourcesQueryParam1) ToPointer() *GetSourcesQueryParam1 {
	return &e
}

func (e *GetSourcesQueryParam1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetSourcesQueryParam1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesQueryParam1: %v", v)
	}
}

type GetSourcesQueryParamDirType string

const (
	GetSourcesQueryParamDirTypeGetSourcesQueryParam1               GetSourcesQueryParamDirType = "getSources_queryParam_1"
	GetSourcesQueryParamDirTypeArrayOfgetSourcesQueryParamSources2 GetSourcesQueryParamDirType = "arrayOfgetSources_queryParam_Sources_2"
)

// GetSourcesQueryParamDir - Sort direction
type GetSourcesQueryParamDir struct {
	GetSourcesQueryParam1               *GetSourcesQueryParam1
	ArrayOfgetSourcesQueryParamSources2 []GetSourcesQueryParamSources2

	Type GetSourcesQueryParamDirType
}

func CreateGetSourcesQueryParamDirGetSourcesQueryParam1(getSourcesQueryParam1 GetSourcesQueryParam1) GetSourcesQueryParamDir {
	typ := GetSourcesQueryParamDirTypeGetSourcesQueryParam1

	return GetSourcesQueryParamDir{
		GetSourcesQueryParam1: &getSourcesQueryParam1,
		Type:                  typ,
	}
}

func CreateGetSourcesQueryParamDirArrayOfgetSourcesQueryParamSources2(arrayOfgetSourcesQueryParamSources2 []GetSourcesQueryParamSources2) GetSourcesQueryParamDir {
	typ := GetSourcesQueryParamDirTypeArrayOfgetSourcesQueryParamSources2

	return GetSourcesQueryParamDir{
		ArrayOfgetSourcesQueryParamSources2: arrayOfgetSourcesQueryParamSources2,
		Type:                                typ,
	}
}

func (u *GetSourcesQueryParamDir) UnmarshalJSON(data []byte) error {

	getSourcesQueryParam1 := GetSourcesQueryParam1("")
	if err := utils.UnmarshalJSON(data, &getSourcesQueryParam1, "", true, true); err == nil {
		u.GetSourcesQueryParam1 = &getSourcesQueryParam1
		u.Type = GetSourcesQueryParamDirTypeGetSourcesQueryParam1
		return nil
	}

	arrayOfgetSourcesQueryParamSources2 := []GetSourcesQueryParamSources2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetSourcesQueryParamSources2, "", true, true); err == nil {
		u.ArrayOfgetSourcesQueryParamSources2 = arrayOfgetSourcesQueryParamSources2
		u.Type = GetSourcesQueryParamDirTypeArrayOfgetSourcesQueryParamSources2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesQueryParamDir) MarshalJSON() ([]byte, error) {
	if u.GetSourcesQueryParam1 != nil {
		return utils.MarshalJSON(u.GetSourcesQueryParam1, "", true)
	}

	if u.ArrayOfgetSourcesQueryParamSources2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetSourcesQueryParamSources2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetSourcesQueryParamIDType string

const (
	GetSourcesQueryParamIDTypeStr        GetSourcesQueryParamIDType = "str"
	GetSourcesQueryParamIDTypeArrayOfstr GetSourcesQueryParamIDType = "arrayOfstr"
)

// GetSourcesQueryParamID - Filter by source IDs
type GetSourcesQueryParamID struct {
	Str        *string
	ArrayOfstr []string

	Type GetSourcesQueryParamIDType
}

func CreateGetSourcesQueryParamIDStr(str string) GetSourcesQueryParamID {
	typ := GetSourcesQueryParamIDTypeStr

	return GetSourcesQueryParamID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetSourcesQueryParamIDArrayOfstr(arrayOfstr []string) GetSourcesQueryParamID {
	typ := GetSourcesQueryParamIDTypeArrayOfstr

	return GetSourcesQueryParamID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetSourcesQueryParamID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetSourcesQueryParamIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetSourcesQueryParamIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesQueryParamID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetSourcesQueryParamSourcesIntegrationID2 struct {
	Any *bool `queryParam:"name=any"`
}

func (o *GetSourcesQueryParamSourcesIntegrationID2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

type IntegrationIDType string

const (
	IntegrationIDTypeStr                                       IntegrationIDType = "str"
	IntegrationIDTypeGetSourcesQueryParamSourcesIntegrationID2 IntegrationIDType = "getSources_queryParam_Sources_integration_id_2"
)

// IntegrationID - Filter by integration IDs
type IntegrationID struct {
	Str                                       *string
	GetSourcesQueryParamSourcesIntegrationID2 *GetSourcesQueryParamSourcesIntegrationID2

	Type IntegrationIDType
}

func CreateIntegrationIDStr(str string) IntegrationID {
	typ := IntegrationIDTypeStr

	return IntegrationID{
		Str:  &str,
		Type: typ,
	}
}

func CreateIntegrationIDGetSourcesQueryParamSourcesIntegrationID2(getSourcesQueryParamSourcesIntegrationID2 GetSourcesQueryParamSourcesIntegrationID2) IntegrationID {
	typ := IntegrationIDTypeGetSourcesQueryParamSourcesIntegrationID2

	return IntegrationID{
		GetSourcesQueryParamSourcesIntegrationID2: &getSourcesQueryParamSourcesIntegrationID2,
		Type: typ,
	}
}

func (u *IntegrationID) UnmarshalJSON(data []byte) error {

	getSourcesQueryParamSourcesIntegrationID2 := GetSourcesQueryParamSourcesIntegrationID2{}
	if err := utils.UnmarshalJSON(data, &getSourcesQueryParamSourcesIntegrationID2, "", true, true); err == nil {
		u.GetSourcesQueryParamSourcesIntegrationID2 = &getSourcesQueryParamSourcesIntegrationID2
		u.Type = IntegrationIDTypeGetSourcesQueryParamSourcesIntegrationID2
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = IntegrationIDTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u IntegrationID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.GetSourcesQueryParamSourcesIntegrationID2 != nil {
		return utils.MarshalJSON(u.GetSourcesQueryParamSourcesIntegrationID2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetSourcesQueryParamSourcesName2 struct {
	Any      *bool   `queryParam:"name=any"`
	Contains *string `queryParam:"name=contains"`
	Gt       *string `queryParam:"name=gt"`
	Gte      *string `queryParam:"name=gte"`
	Le       *string `queryParam:"name=le"`
	Lte      *string `queryParam:"name=lte"`
}

func (o *GetSourcesQueryParamSourcesName2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetSourcesQueryParamSourcesName2) GetContains() *string {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *GetSourcesQueryParamSourcesName2) GetGt() *string {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetSourcesQueryParamSourcesName2) GetGte() *string {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetSourcesQueryParamSourcesName2) GetLe() *string {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetSourcesQueryParamSourcesName2) GetLte() *string {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetSourcesQueryParamNameType string

const (
	GetSourcesQueryParamNameTypeStr                              GetSourcesQueryParamNameType = "str"
	GetSourcesQueryParamNameTypeGetSourcesQueryParamSourcesName2 GetSourcesQueryParamNameType = "getSources_queryParam_Sources_name_2"
)

// GetSourcesQueryParamName - The source name
type GetSourcesQueryParamName struct {
	Str                              *string
	GetSourcesQueryParamSourcesName2 *GetSourcesQueryParamSourcesName2

	Type GetSourcesQueryParamNameType
}

func CreateGetSourcesQueryParamNameStr(str string) GetSourcesQueryParamName {
	typ := GetSourcesQueryParamNameTypeStr

	return GetSourcesQueryParamName{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetSourcesQueryParamNameGetSourcesQueryParamSourcesName2(getSourcesQueryParamSourcesName2 GetSourcesQueryParamSourcesName2) GetSourcesQueryParamName {
	typ := GetSourcesQueryParamNameTypeGetSourcesQueryParamSourcesName2

	return GetSourcesQueryParamName{
		GetSourcesQueryParamSourcesName2: &getSourcesQueryParamSourcesName2,
		Type:                             typ,
	}
}

func (u *GetSourcesQueryParamName) UnmarshalJSON(data []byte) error {

	getSourcesQueryParamSourcesName2 := GetSourcesQueryParamSourcesName2{}
	if err := utils.UnmarshalJSON(data, &getSourcesQueryParamSourcesName2, "", true, true); err == nil {
		u.GetSourcesQueryParamSourcesName2 = &getSourcesQueryParamSourcesName2
		u.Type = GetSourcesQueryParamNameTypeGetSourcesQueryParamSourcesName2
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetSourcesQueryParamNameTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesQueryParamName) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.GetSourcesQueryParamSourcesName2 != nil {
		return utils.MarshalJSON(u.GetSourcesQueryParamSourcesName2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetSourcesQueryParamSourcesOrderBy2 string

const (
	GetSourcesQueryParamSourcesOrderBy2CreatedAt GetSourcesQueryParamSourcesOrderBy2 = "created_at"
)

func (e GetSourcesQueryParamSourcesOrderBy2) ToPointer() *GetSourcesQueryParamSourcesOrderBy2 {
	return &e
}

func (e *GetSourcesQueryParamSourcesOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetSourcesQueryParamSourcesOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesQueryParamSourcesOrderBy2: %v", v)
	}
}

type GetSourcesQueryParamSources1 string

const (
	GetSourcesQueryParamSources1CreatedAt GetSourcesQueryParamSources1 = "created_at"
)

func (e GetSourcesQueryParamSources1) ToPointer() *GetSourcesQueryParamSources1 {
	return &e
}

func (e *GetSourcesQueryParamSources1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetSourcesQueryParamSources1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSourcesQueryParamSources1: %v", v)
	}
}

type GetSourcesQueryParamOrderByType string

const (
	GetSourcesQueryParamOrderByTypeGetSourcesQueryParamSources1               GetSourcesQueryParamOrderByType = "getSources_queryParam_Sources_1"
	GetSourcesQueryParamOrderByTypeArrayOfgetSourcesQueryParamSourcesOrderBy2 GetSourcesQueryParamOrderByType = "arrayOfgetSources_queryParam_Sources_order_by_2"
)

// GetSourcesQueryParamOrderBy - Sort key(s)
type GetSourcesQueryParamOrderBy struct {
	GetSourcesQueryParamSources1               *GetSourcesQueryParamSources1
	ArrayOfgetSourcesQueryParamSourcesOrderBy2 []GetSourcesQueryParamSourcesOrderBy2

	Type GetSourcesQueryParamOrderByType
}

func CreateGetSourcesQueryParamOrderByGetSourcesQueryParamSources1(getSourcesQueryParamSources1 GetSourcesQueryParamSources1) GetSourcesQueryParamOrderBy {
	typ := GetSourcesQueryParamOrderByTypeGetSourcesQueryParamSources1

	return GetSourcesQueryParamOrderBy{
		GetSourcesQueryParamSources1: &getSourcesQueryParamSources1,
		Type:                         typ,
	}
}

func CreateGetSourcesQueryParamOrderByArrayOfgetSourcesQueryParamSourcesOrderBy2(arrayOfgetSourcesQueryParamSourcesOrderBy2 []GetSourcesQueryParamSourcesOrderBy2) GetSourcesQueryParamOrderBy {
	typ := GetSourcesQueryParamOrderByTypeArrayOfgetSourcesQueryParamSourcesOrderBy2

	return GetSourcesQueryParamOrderBy{
		ArrayOfgetSourcesQueryParamSourcesOrderBy2: arrayOfgetSourcesQueryParamSourcesOrderBy2,
		Type: typ,
	}
}

func (u *GetSourcesQueryParamOrderBy) UnmarshalJSON(data []byte) error {

	getSourcesQueryParamSources1 := GetSourcesQueryParamSources1("")
	if err := utils.UnmarshalJSON(data, &getSourcesQueryParamSources1, "", true, true); err == nil {
		u.GetSourcesQueryParamSources1 = &getSourcesQueryParamSources1
		u.Type = GetSourcesQueryParamOrderByTypeGetSourcesQueryParamSources1
		return nil
	}

	arrayOfgetSourcesQueryParamSourcesOrderBy2 := []GetSourcesQueryParamSourcesOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetSourcesQueryParamSourcesOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetSourcesQueryParamSourcesOrderBy2 = arrayOfgetSourcesQueryParamSourcesOrderBy2
		u.Type = GetSourcesQueryParamOrderByTypeArrayOfgetSourcesQueryParamSourcesOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetSourcesQueryParamOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetSourcesQueryParamSources1 != nil {
		return utils.MarshalJSON(u.GetSourcesQueryParamSources1, "", true)
	}

	if u.ArrayOfgetSourcesQueryParamSourcesOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetSourcesQueryParamSourcesOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetSourcesRequest struct {
	Archived *bool `queryParam:"style=form,explode=true,name=archived"`
	// Date the source was archived
	ArchivedAt *GetSourcesQueryParamArchivedAt `queryParam:"style=form,explode=true,name=archived_at"`
	// Sort direction
	Dir *GetSourcesQueryParamDir `queryParam:"style=form,explode=true,name=dir"`
	// Filter by source IDs
	ID *GetSourcesQueryParamID `queryParam:"style=form,explode=true,name=id"`
	// Filter by integration IDs
	IntegrationID *IntegrationID `queryParam:"style=form,explode=true,name=integration_id"`
	Limit         *int64         `queryParam:"style=form,explode=true,name=limit"`
	// The source name
	Name *GetSourcesQueryParamName `queryParam:"style=form,explode=true,name=name"`
	Next *string                   `queryParam:"style=form,explode=true,name=next"`
	// Sort key(s)
	OrderBy *GetSourcesQueryParamOrderBy `queryParam:"style=form,explode=true,name=order_by"`
	Prev    *string                      `queryParam:"style=form,explode=true,name=prev"`
}

func (o *GetSourcesRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetSourcesRequest) GetArchivedAt() *GetSourcesQueryParamArchivedAt {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetSourcesRequest) GetDir() *GetSourcesQueryParamDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetSourcesRequest) GetID() *GetSourcesQueryParamID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetSourcesRequest) GetIntegrationID() *IntegrationID {
	if o == nil {
		return nil
	}
	return o.IntegrationID
}

func (o *GetSourcesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetSourcesRequest) GetName() *GetSourcesQueryParamName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetSourcesRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetSourcesRequest) GetOrderBy() *GetSourcesQueryParamOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetSourcesRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

type GetSourcesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of sources
	SourcePaginatedResult *components.SourcePaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSourcesResponse) GetSourcePaginatedResult() *components.SourcePaginatedResult {
	if o == nil {
		return nil
	}
	return o.SourcePaginatedResult
}

func (o *GetSourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

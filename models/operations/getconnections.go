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

type GetConnectionsQueryParam2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetConnectionsQueryParam2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetConnectionsQueryParam2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetConnectionsQueryParam2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetConnectionsQueryParam2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetConnectionsQueryParam2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetConnectionsQueryParam2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetConnectionsQueryParam2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type ArchivedAtType string

const (
	ArchivedAtTypeDateTime                  ArchivedAtType = "date-time"
	ArchivedAtTypeGetConnectionsQueryParam2 ArchivedAtType = "getConnections_queryParam_2"
)

// ArchivedAt - Date the connection was archived
type ArchivedAt struct {
	DateTime                  *time.Time
	GetConnectionsQueryParam2 *GetConnectionsQueryParam2

	Type ArchivedAtType
}

func CreateArchivedAtDateTime(dateTime time.Time) ArchivedAt {
	typ := ArchivedAtTypeDateTime

	return ArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateArchivedAtGetConnectionsQueryParam2(getConnectionsQueryParam2 GetConnectionsQueryParam2) ArchivedAt {
	typ := ArchivedAtTypeGetConnectionsQueryParam2

	return ArchivedAt{
		GetConnectionsQueryParam2: &getConnectionsQueryParam2,
		Type:                      typ,
	}
}

func (u *ArchivedAt) UnmarshalJSON(data []byte) error {

	getConnectionsQueryParam2 := GetConnectionsQueryParam2{}
	if err := utils.UnmarshalJSON(data, &getConnectionsQueryParam2, "", true, true); err == nil {
		u.GetConnectionsQueryParam2 = &getConnectionsQueryParam2
		u.Type = ArchivedAtTypeGetConnectionsQueryParam2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = ArchivedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetConnectionsQueryParam2 != nil {
		return utils.MarshalJSON(u.GetConnectionsQueryParam2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsQueryParamDestinationIDType string

const (
	GetConnectionsQueryParamDestinationIDTypeStr        GetConnectionsQueryParamDestinationIDType = "str"
	GetConnectionsQueryParamDestinationIDTypeArrayOfstr GetConnectionsQueryParamDestinationIDType = "arrayOfstr"
)

// GetConnectionsQueryParamDestinationID - Filter by associated destination IDs
type GetConnectionsQueryParamDestinationID struct {
	Str        *string
	ArrayOfstr []string

	Type GetConnectionsQueryParamDestinationIDType
}

func CreateGetConnectionsQueryParamDestinationIDStr(str string) GetConnectionsQueryParamDestinationID {
	typ := GetConnectionsQueryParamDestinationIDTypeStr

	return GetConnectionsQueryParamDestinationID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetConnectionsQueryParamDestinationIDArrayOfstr(arrayOfstr []string) GetConnectionsQueryParamDestinationID {
	typ := GetConnectionsQueryParamDestinationIDTypeArrayOfstr

	return GetConnectionsQueryParamDestinationID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetConnectionsQueryParamDestinationID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetConnectionsQueryParamDestinationIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetConnectionsQueryParamDestinationIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsQueryParamDestinationID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsQueryParamConnections2 string

const (
	GetConnectionsQueryParamConnections2Asc  GetConnectionsQueryParamConnections2 = "asc"
	GetConnectionsQueryParamConnections2Desc GetConnectionsQueryParamConnections2 = "desc"
)

func (e GetConnectionsQueryParamConnections2) ToPointer() *GetConnectionsQueryParamConnections2 {
	return &e
}

func (e *GetConnectionsQueryParamConnections2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetConnectionsQueryParamConnections2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsQueryParamConnections2: %v", v)
	}
}

type GetConnectionsQueryParam1 string

const (
	GetConnectionsQueryParam1Asc  GetConnectionsQueryParam1 = "asc"
	GetConnectionsQueryParam1Desc GetConnectionsQueryParam1 = "desc"
)

func (e GetConnectionsQueryParam1) ToPointer() *GetConnectionsQueryParam1 {
	return &e
}

func (e *GetConnectionsQueryParam1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetConnectionsQueryParam1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsQueryParam1: %v", v)
	}
}

type GetConnectionsQueryParamDirType string

const (
	GetConnectionsQueryParamDirTypeGetConnectionsQueryParam1                   GetConnectionsQueryParamDirType = "getConnections_queryParam_1"
	GetConnectionsQueryParamDirTypeArrayOfgetConnectionsQueryParamConnections2 GetConnectionsQueryParamDirType = "arrayOfgetConnections_queryParam_Connections_2"
)

// GetConnectionsQueryParamDir - Sort direction
type GetConnectionsQueryParamDir struct {
	GetConnectionsQueryParam1                   *GetConnectionsQueryParam1
	ArrayOfgetConnectionsQueryParamConnections2 []GetConnectionsQueryParamConnections2

	Type GetConnectionsQueryParamDirType
}

func CreateGetConnectionsQueryParamDirGetConnectionsQueryParam1(getConnectionsQueryParam1 GetConnectionsQueryParam1) GetConnectionsQueryParamDir {
	typ := GetConnectionsQueryParamDirTypeGetConnectionsQueryParam1

	return GetConnectionsQueryParamDir{
		GetConnectionsQueryParam1: &getConnectionsQueryParam1,
		Type:                      typ,
	}
}

func CreateGetConnectionsQueryParamDirArrayOfgetConnectionsQueryParamConnections2(arrayOfgetConnectionsQueryParamConnections2 []GetConnectionsQueryParamConnections2) GetConnectionsQueryParamDir {
	typ := GetConnectionsQueryParamDirTypeArrayOfgetConnectionsQueryParamConnections2

	return GetConnectionsQueryParamDir{
		ArrayOfgetConnectionsQueryParamConnections2: arrayOfgetConnectionsQueryParamConnections2,
		Type: typ,
	}
}

func (u *GetConnectionsQueryParamDir) UnmarshalJSON(data []byte) error {

	getConnectionsQueryParam1 := GetConnectionsQueryParam1("")
	if err := utils.UnmarshalJSON(data, &getConnectionsQueryParam1, "", true, true); err == nil {
		u.GetConnectionsQueryParam1 = &getConnectionsQueryParam1
		u.Type = GetConnectionsQueryParamDirTypeGetConnectionsQueryParam1
		return nil
	}

	arrayOfgetConnectionsQueryParamConnections2 := []GetConnectionsQueryParamConnections2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetConnectionsQueryParamConnections2, "", true, true); err == nil {
		u.ArrayOfgetConnectionsQueryParamConnections2 = arrayOfgetConnectionsQueryParamConnections2
		u.Type = GetConnectionsQueryParamDirTypeArrayOfgetConnectionsQueryParamConnections2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsQueryParamDir) MarshalJSON() ([]byte, error) {
	if u.GetConnectionsQueryParam1 != nil {
		return utils.MarshalJSON(u.GetConnectionsQueryParam1, "", true)
	}

	if u.ArrayOfgetConnectionsQueryParamConnections2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetConnectionsQueryParamConnections2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsQueryParamIDType string

const (
	GetConnectionsQueryParamIDTypeStr        GetConnectionsQueryParamIDType = "str"
	GetConnectionsQueryParamIDTypeArrayOfstr GetConnectionsQueryParamIDType = "arrayOfstr"
)

// GetConnectionsQueryParamID - Filter by connection IDs
type GetConnectionsQueryParamID struct {
	Str        *string
	ArrayOfstr []string

	Type GetConnectionsQueryParamIDType
}

func CreateGetConnectionsQueryParamIDStr(str string) GetConnectionsQueryParamID {
	typ := GetConnectionsQueryParamIDTypeStr

	return GetConnectionsQueryParamID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetConnectionsQueryParamIDArrayOfstr(arrayOfstr []string) GetConnectionsQueryParamID {
	typ := GetConnectionsQueryParamIDTypeArrayOfstr

	return GetConnectionsQueryParamID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetConnectionsQueryParamID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetConnectionsQueryParamIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetConnectionsQueryParamIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsQueryParamID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsQueryParamConnectionsName2 struct {
	Any      *bool   `queryParam:"name=any"`
	Contains *string `queryParam:"name=contains"`
	Gt       *string `queryParam:"name=gt"`
	Gte      *string `queryParam:"name=gte"`
	Le       *string `queryParam:"name=le"`
	Lte      *string `queryParam:"name=lte"`
}

func (o *GetConnectionsQueryParamConnectionsName2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetConnectionsQueryParamConnectionsName2) GetContains() *string {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *GetConnectionsQueryParamConnectionsName2) GetGt() *string {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetConnectionsQueryParamConnectionsName2) GetGte() *string {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetConnectionsQueryParamConnectionsName2) GetLe() *string {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetConnectionsQueryParamConnectionsName2) GetLte() *string {
	if o == nil {
		return nil
	}
	return o.Lte
}

type QueryParamNameType string

const (
	QueryParamNameTypeStr                                      QueryParamNameType = "str"
	QueryParamNameTypeGetConnectionsQueryParamConnectionsName2 QueryParamNameType = "getConnections_queryParam_Connections_name_2"
)

// QueryParamName - Filter by connection name
type QueryParamName struct {
	Str                                      *string
	GetConnectionsQueryParamConnectionsName2 *GetConnectionsQueryParamConnectionsName2

	Type QueryParamNameType
}

func CreateQueryParamNameStr(str string) QueryParamName {
	typ := QueryParamNameTypeStr

	return QueryParamName{
		Str:  &str,
		Type: typ,
	}
}

func CreateQueryParamNameGetConnectionsQueryParamConnectionsName2(getConnectionsQueryParamConnectionsName2 GetConnectionsQueryParamConnectionsName2) QueryParamName {
	typ := QueryParamNameTypeGetConnectionsQueryParamConnectionsName2

	return QueryParamName{
		GetConnectionsQueryParamConnectionsName2: &getConnectionsQueryParamConnectionsName2,
		Type:                                     typ,
	}
}

func (u *QueryParamName) UnmarshalJSON(data []byte) error {

	getConnectionsQueryParamConnectionsName2 := GetConnectionsQueryParamConnectionsName2{}
	if err := utils.UnmarshalJSON(data, &getConnectionsQueryParamConnectionsName2, "", true, true); err == nil {
		u.GetConnectionsQueryParamConnectionsName2 = &getConnectionsQueryParamConnectionsName2
		u.Type = QueryParamNameTypeGetConnectionsQueryParamConnectionsName2
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = QueryParamNameTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u QueryParamName) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.GetConnectionsQueryParamConnectionsName2 != nil {
		return utils.MarshalJSON(u.GetConnectionsQueryParamConnectionsName2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsQueryParamConnectionsOrderBy2 string

const (
	GetConnectionsQueryParamConnectionsOrderBy2CreatedAt GetConnectionsQueryParamConnectionsOrderBy2 = "created_at"
)

func (e GetConnectionsQueryParamConnectionsOrderBy2) ToPointer() *GetConnectionsQueryParamConnectionsOrderBy2 {
	return &e
}

func (e *GetConnectionsQueryParamConnectionsOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetConnectionsQueryParamConnectionsOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsQueryParamConnectionsOrderBy2: %v", v)
	}
}

type GetConnectionsQueryParamConnections1 string

const (
	GetConnectionsQueryParamConnections1CreatedAt GetConnectionsQueryParamConnections1 = "created_at"
)

func (e GetConnectionsQueryParamConnections1) ToPointer() *GetConnectionsQueryParamConnections1 {
	return &e
}

func (e *GetConnectionsQueryParamConnections1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetConnectionsQueryParamConnections1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsQueryParamConnections1: %v", v)
	}
}

type GetConnectionsQueryParamOrderByType string

const (
	GetConnectionsQueryParamOrderByTypeGetConnectionsQueryParamConnections1               GetConnectionsQueryParamOrderByType = "getConnections_queryParam_Connections_1"
	GetConnectionsQueryParamOrderByTypeArrayOfgetConnectionsQueryParamConnectionsOrderBy2 GetConnectionsQueryParamOrderByType = "arrayOfgetConnections_queryParam_Connections_order_by_2"
)

// GetConnectionsQueryParamOrderBy - Sort key(s)
type GetConnectionsQueryParamOrderBy struct {
	GetConnectionsQueryParamConnections1               *GetConnectionsQueryParamConnections1
	ArrayOfgetConnectionsQueryParamConnectionsOrderBy2 []GetConnectionsQueryParamConnectionsOrderBy2

	Type GetConnectionsQueryParamOrderByType
}

func CreateGetConnectionsQueryParamOrderByGetConnectionsQueryParamConnections1(getConnectionsQueryParamConnections1 GetConnectionsQueryParamConnections1) GetConnectionsQueryParamOrderBy {
	typ := GetConnectionsQueryParamOrderByTypeGetConnectionsQueryParamConnections1

	return GetConnectionsQueryParamOrderBy{
		GetConnectionsQueryParamConnections1: &getConnectionsQueryParamConnections1,
		Type:                                 typ,
	}
}

func CreateGetConnectionsQueryParamOrderByArrayOfgetConnectionsQueryParamConnectionsOrderBy2(arrayOfgetConnectionsQueryParamConnectionsOrderBy2 []GetConnectionsQueryParamConnectionsOrderBy2) GetConnectionsQueryParamOrderBy {
	typ := GetConnectionsQueryParamOrderByTypeArrayOfgetConnectionsQueryParamConnectionsOrderBy2

	return GetConnectionsQueryParamOrderBy{
		ArrayOfgetConnectionsQueryParamConnectionsOrderBy2: arrayOfgetConnectionsQueryParamConnectionsOrderBy2,
		Type: typ,
	}
}

func (u *GetConnectionsQueryParamOrderBy) UnmarshalJSON(data []byte) error {

	getConnectionsQueryParamConnections1 := GetConnectionsQueryParamConnections1("")
	if err := utils.UnmarshalJSON(data, &getConnectionsQueryParamConnections1, "", true, true); err == nil {
		u.GetConnectionsQueryParamConnections1 = &getConnectionsQueryParamConnections1
		u.Type = GetConnectionsQueryParamOrderByTypeGetConnectionsQueryParamConnections1
		return nil
	}

	arrayOfgetConnectionsQueryParamConnectionsOrderBy2 := []GetConnectionsQueryParamConnectionsOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetConnectionsQueryParamConnectionsOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetConnectionsQueryParamConnectionsOrderBy2 = arrayOfgetConnectionsQueryParamConnectionsOrderBy2
		u.Type = GetConnectionsQueryParamOrderByTypeArrayOfgetConnectionsQueryParamConnectionsOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsQueryParamOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetConnectionsQueryParamConnections1 != nil {
		return utils.MarshalJSON(u.GetConnectionsQueryParamConnections1, "", true)
	}

	if u.ArrayOfgetConnectionsQueryParamConnectionsOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetConnectionsQueryParamConnectionsOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsQueryParamConnectionsPausedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetConnectionsQueryParamConnectionsPausedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetConnectionsQueryParamConnectionsPausedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetConnectionsQueryParamConnectionsPausedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetConnectionsQueryParamConnectionsPausedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetConnectionsQueryParamConnectionsPausedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetConnectionsQueryParamConnectionsPausedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetConnectionsQueryParamConnectionsPausedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type PausedAtType string

const (
	PausedAtTypeDateTime                                     PausedAtType = "date-time"
	PausedAtTypeGetConnectionsQueryParamConnectionsPausedAt2 PausedAtType = "getConnections_queryParam_Connections_paused_at_2"
)

// PausedAt - Date the connection was paused
type PausedAt struct {
	DateTime                                     *time.Time
	GetConnectionsQueryParamConnectionsPausedAt2 *GetConnectionsQueryParamConnectionsPausedAt2

	Type PausedAtType
}

func CreatePausedAtDateTime(dateTime time.Time) PausedAt {
	typ := PausedAtTypeDateTime

	return PausedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreatePausedAtGetConnectionsQueryParamConnectionsPausedAt2(getConnectionsQueryParamConnectionsPausedAt2 GetConnectionsQueryParamConnectionsPausedAt2) PausedAt {
	typ := PausedAtTypeGetConnectionsQueryParamConnectionsPausedAt2

	return PausedAt{
		GetConnectionsQueryParamConnectionsPausedAt2: &getConnectionsQueryParamConnectionsPausedAt2,
		Type: typ,
	}
}

func (u *PausedAt) UnmarshalJSON(data []byte) error {

	getConnectionsQueryParamConnectionsPausedAt2 := GetConnectionsQueryParamConnectionsPausedAt2{}
	if err := utils.UnmarshalJSON(data, &getConnectionsQueryParamConnectionsPausedAt2, "", true, true); err == nil {
		u.GetConnectionsQueryParamConnectionsPausedAt2 = &getConnectionsQueryParamConnectionsPausedAt2
		u.Type = PausedAtTypeGetConnectionsQueryParamConnectionsPausedAt2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = PausedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PausedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetConnectionsQueryParamConnectionsPausedAt2 != nil {
		return utils.MarshalJSON(u.GetConnectionsQueryParamConnectionsPausedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsQueryParamSourceIDType string

const (
	GetConnectionsQueryParamSourceIDTypeStr        GetConnectionsQueryParamSourceIDType = "str"
	GetConnectionsQueryParamSourceIDTypeArrayOfstr GetConnectionsQueryParamSourceIDType = "arrayOfstr"
)

// GetConnectionsQueryParamSourceID - Filter by associated source IDs
type GetConnectionsQueryParamSourceID struct {
	Str        *string
	ArrayOfstr []string

	Type GetConnectionsQueryParamSourceIDType
}

func CreateGetConnectionsQueryParamSourceIDStr(str string) GetConnectionsQueryParamSourceID {
	typ := GetConnectionsQueryParamSourceIDTypeStr

	return GetConnectionsQueryParamSourceID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetConnectionsQueryParamSourceIDArrayOfstr(arrayOfstr []string) GetConnectionsQueryParamSourceID {
	typ := GetConnectionsQueryParamSourceIDTypeArrayOfstr

	return GetConnectionsQueryParamSourceID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetConnectionsQueryParamSourceID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetConnectionsQueryParamSourceIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetConnectionsQueryParamSourceIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsQueryParamSourceID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsRequest struct {
	Archived *bool `queryParam:"style=form,explode=true,name=archived"`
	// Date the connection was archived
	ArchivedAt *ArchivedAt `queryParam:"style=form,explode=true,name=archived_at"`
	// Filter by associated destination IDs
	DestinationID *GetConnectionsQueryParamDestinationID `queryParam:"style=form,explode=true,name=destination_id"`
	// Sort direction
	Dir      *GetConnectionsQueryParamDir `queryParam:"style=form,explode=true,name=dir"`
	FullName *string                      `queryParam:"style=form,explode=true,name=full_name"`
	// Filter by connection IDs
	ID *GetConnectionsQueryParamID `queryParam:"style=form,explode=true,name=id"`
	// Filter by connection name
	Name *QueryParamName `queryParam:"style=form,explode=true,name=name"`
	// Sort key(s)
	OrderBy *GetConnectionsQueryParamOrderBy `queryParam:"style=form,explode=true,name=order_by"`
	// Date the connection was paused
	PausedAt *PausedAt `queryParam:"style=form,explode=true,name=paused_at"`
	// Filter by associated source IDs
	SourceID *GetConnectionsQueryParamSourceID `queryParam:"style=form,explode=true,name=source_id"`
}

func (o *GetConnectionsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetConnectionsRequest) GetArchivedAt() *ArchivedAt {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetConnectionsRequest) GetDestinationID() *GetConnectionsQueryParamDestinationID {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *GetConnectionsRequest) GetDir() *GetConnectionsQueryParamDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetConnectionsRequest) GetFullName() *string {
	if o == nil {
		return nil
	}
	return o.FullName
}

func (o *GetConnectionsRequest) GetID() *GetConnectionsQueryParamID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetConnectionsRequest) GetName() *QueryParamName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetConnectionsRequest) GetOrderBy() *GetConnectionsQueryParamOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetConnectionsRequest) GetPausedAt() *PausedAt {
	if o == nil {
		return nil
	}
	return o.PausedAt
}

func (o *GetConnectionsRequest) GetSourceID() *GetConnectionsQueryParamSourceID {
	if o == nil {
		return nil
	}
	return o.SourceID
}

type GetConnectionsResponse struct {
	// List of connections
	ConnectionPaginatedResult *components.ConnectionPaginatedResult
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConnectionsResponse) GetConnectionPaginatedResult() *components.ConnectionPaginatedResult {
	if o == nil {
		return nil
	}
	return o.ConnectionPaginatedResult
}

func (o *GetConnectionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

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

// GetConnectionsArchivedAt2 - Date the connection was archived
type GetConnectionsArchivedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetConnectionsArchivedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetConnectionsArchivedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetConnectionsArchivedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetConnectionsArchivedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetConnectionsArchivedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetConnectionsArchivedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetConnectionsArchivedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetConnectionsArchivedAtType string

const (
	GetConnectionsArchivedAtTypeDateTime                  GetConnectionsArchivedAtType = "date-time"
	GetConnectionsArchivedAtTypeGetConnectionsArchivedAt2 GetConnectionsArchivedAtType = "getConnectionsArchivedAt_2"
)

type GetConnectionsArchivedAt struct {
	DateTime                  *time.Time
	GetConnectionsArchivedAt2 *GetConnectionsArchivedAt2

	Type GetConnectionsArchivedAtType
}

func CreateGetConnectionsArchivedAtDateTime(dateTime time.Time) GetConnectionsArchivedAt {
	typ := GetConnectionsArchivedAtTypeDateTime

	return GetConnectionsArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetConnectionsArchivedAtGetConnectionsArchivedAt2(getConnectionsArchivedAt2 GetConnectionsArchivedAt2) GetConnectionsArchivedAt {
	typ := GetConnectionsArchivedAtTypeGetConnectionsArchivedAt2

	return GetConnectionsArchivedAt{
		GetConnectionsArchivedAt2: &getConnectionsArchivedAt2,
		Type:                      typ,
	}
}

func (u *GetConnectionsArchivedAt) UnmarshalJSON(data []byte) error {

	getConnectionsArchivedAt2 := GetConnectionsArchivedAt2{}
	if err := utils.UnmarshalJSON(data, &getConnectionsArchivedAt2, "", true, true); err == nil {
		u.GetConnectionsArchivedAt2 = &getConnectionsArchivedAt2
		u.Type = GetConnectionsArchivedAtTypeGetConnectionsArchivedAt2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetConnectionsArchivedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetConnectionsArchivedAt2 != nil {
		return utils.MarshalJSON(u.GetConnectionsArchivedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsDestinationIDType string

const (
	GetConnectionsDestinationIDTypeStr        GetConnectionsDestinationIDType = "str"
	GetConnectionsDestinationIDTypeArrayOfstr GetConnectionsDestinationIDType = "arrayOfstr"
)

type GetConnectionsDestinationID struct {
	Str        *string
	ArrayOfstr []string

	Type GetConnectionsDestinationIDType
}

func CreateGetConnectionsDestinationIDStr(str string) GetConnectionsDestinationID {
	typ := GetConnectionsDestinationIDTypeStr

	return GetConnectionsDestinationID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetConnectionsDestinationIDArrayOfstr(arrayOfstr []string) GetConnectionsDestinationID {
	typ := GetConnectionsDestinationIDTypeArrayOfstr

	return GetConnectionsDestinationID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetConnectionsDestinationID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetConnectionsDestinationIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetConnectionsDestinationIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsDestinationID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsDir2 string

const (
	GetConnectionsDir2Asc  GetConnectionsDir2 = "asc"
	GetConnectionsDir2Desc GetConnectionsDir2 = "desc"
)

func (e GetConnectionsDir2) ToPointer() *GetConnectionsDir2 {
	return &e
}

func (e *GetConnectionsDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetConnectionsDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsDir2: %v", v)
	}
}

// GetConnectionsDir1 - Sort direction
type GetConnectionsDir1 string

const (
	GetConnectionsDir1Asc  GetConnectionsDir1 = "asc"
	GetConnectionsDir1Desc GetConnectionsDir1 = "desc"
)

func (e GetConnectionsDir1) ToPointer() *GetConnectionsDir1 {
	return &e
}

func (e *GetConnectionsDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetConnectionsDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsDir1: %v", v)
	}
}

type GetConnectionsDirType string

const (
	GetConnectionsDirTypeGetConnectionsDir1        GetConnectionsDirType = "getConnectionsDir_1"
	GetConnectionsDirTypeArrayOfgetConnectionsDir2 GetConnectionsDirType = "arrayOfgetConnectionsDir_2"
)

type GetConnectionsDir struct {
	GetConnectionsDir1        *GetConnectionsDir1
	ArrayOfgetConnectionsDir2 []GetConnectionsDir2

	Type GetConnectionsDirType
}

func CreateGetConnectionsDirGetConnectionsDir1(getConnectionsDir1 GetConnectionsDir1) GetConnectionsDir {
	typ := GetConnectionsDirTypeGetConnectionsDir1

	return GetConnectionsDir{
		GetConnectionsDir1: &getConnectionsDir1,
		Type:               typ,
	}
}

func CreateGetConnectionsDirArrayOfgetConnectionsDir2(arrayOfgetConnectionsDir2 []GetConnectionsDir2) GetConnectionsDir {
	typ := GetConnectionsDirTypeArrayOfgetConnectionsDir2

	return GetConnectionsDir{
		ArrayOfgetConnectionsDir2: arrayOfgetConnectionsDir2,
		Type:                      typ,
	}
}

func (u *GetConnectionsDir) UnmarshalJSON(data []byte) error {

	getConnectionsDir1 := GetConnectionsDir1("")
	if err := utils.UnmarshalJSON(data, &getConnectionsDir1, "", true, true); err == nil {
		u.GetConnectionsDir1 = &getConnectionsDir1
		u.Type = GetConnectionsDirTypeGetConnectionsDir1
		return nil
	}

	arrayOfgetConnectionsDir2 := []GetConnectionsDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetConnectionsDir2, "", true, true); err == nil {
		u.ArrayOfgetConnectionsDir2 = arrayOfgetConnectionsDir2
		u.Type = GetConnectionsDirTypeArrayOfgetConnectionsDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsDir) MarshalJSON() ([]byte, error) {
	if u.GetConnectionsDir1 != nil {
		return utils.MarshalJSON(u.GetConnectionsDir1, "", true)
	}

	if u.ArrayOfgetConnectionsDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetConnectionsDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsIDType string

const (
	GetConnectionsIDTypeStr        GetConnectionsIDType = "str"
	GetConnectionsIDTypeArrayOfstr GetConnectionsIDType = "arrayOfstr"
)

type GetConnectionsID struct {
	Str        *string
	ArrayOfstr []string

	Type GetConnectionsIDType
}

func CreateGetConnectionsIDStr(str string) GetConnectionsID {
	typ := GetConnectionsIDTypeStr

	return GetConnectionsID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetConnectionsIDArrayOfstr(arrayOfstr []string) GetConnectionsID {
	typ := GetConnectionsIDTypeArrayOfstr

	return GetConnectionsID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetConnectionsID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetConnectionsIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetConnectionsIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetConnectionsName2 - Filter by connection name
type GetConnectionsName2 struct {
	Any      *bool   `queryParam:"name=any"`
	Contains *string `queryParam:"name=contains"`
	Gt       *string `queryParam:"name=gt"`
	Gte      *string `queryParam:"name=gte"`
	Le       *string `queryParam:"name=le"`
	Lte      *string `queryParam:"name=lte"`
}

func (o *GetConnectionsName2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetConnectionsName2) GetContains() *string {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *GetConnectionsName2) GetGt() *string {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetConnectionsName2) GetGte() *string {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetConnectionsName2) GetLe() *string {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetConnectionsName2) GetLte() *string {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetConnectionsNameType string

const (
	GetConnectionsNameTypeStr                 GetConnectionsNameType = "str"
	GetConnectionsNameTypeGetConnectionsName2 GetConnectionsNameType = "getConnectionsName_2"
)

type GetConnectionsName struct {
	Str                 *string
	GetConnectionsName2 *GetConnectionsName2

	Type GetConnectionsNameType
}

func CreateGetConnectionsNameStr(str string) GetConnectionsName {
	typ := GetConnectionsNameTypeStr

	return GetConnectionsName{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetConnectionsNameGetConnectionsName2(getConnectionsName2 GetConnectionsName2) GetConnectionsName {
	typ := GetConnectionsNameTypeGetConnectionsName2

	return GetConnectionsName{
		GetConnectionsName2: &getConnectionsName2,
		Type:                typ,
	}
}

func (u *GetConnectionsName) UnmarshalJSON(data []byte) error {

	getConnectionsName2 := GetConnectionsName2{}
	if err := utils.UnmarshalJSON(data, &getConnectionsName2, "", true, true); err == nil {
		u.GetConnectionsName2 = &getConnectionsName2
		u.Type = GetConnectionsNameTypeGetConnectionsName2
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetConnectionsNameTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsName) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.GetConnectionsName2 != nil {
		return utils.MarshalJSON(u.GetConnectionsName2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsOrderBy2 string

const (
	GetConnectionsOrderBy2CreatedAt GetConnectionsOrderBy2 = "created_at"
)

func (e GetConnectionsOrderBy2) ToPointer() *GetConnectionsOrderBy2 {
	return &e
}

func (e *GetConnectionsOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetConnectionsOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsOrderBy2: %v", v)
	}
}

// GetConnectionsOrderBy1 - Sort key(s)
type GetConnectionsOrderBy1 string

const (
	GetConnectionsOrderBy1CreatedAt GetConnectionsOrderBy1 = "created_at"
)

func (e GetConnectionsOrderBy1) ToPointer() *GetConnectionsOrderBy1 {
	return &e
}

func (e *GetConnectionsOrderBy1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetConnectionsOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConnectionsOrderBy1: %v", v)
	}
}

type GetConnectionsOrderByType string

const (
	GetConnectionsOrderByTypeGetConnectionsOrderBy1        GetConnectionsOrderByType = "getConnectionsOrderBy_1"
	GetConnectionsOrderByTypeArrayOfgetConnectionsOrderBy2 GetConnectionsOrderByType = "arrayOfgetConnectionsOrderBy_2"
)

type GetConnectionsOrderBy struct {
	GetConnectionsOrderBy1        *GetConnectionsOrderBy1
	ArrayOfgetConnectionsOrderBy2 []GetConnectionsOrderBy2

	Type GetConnectionsOrderByType
}

func CreateGetConnectionsOrderByGetConnectionsOrderBy1(getConnectionsOrderBy1 GetConnectionsOrderBy1) GetConnectionsOrderBy {
	typ := GetConnectionsOrderByTypeGetConnectionsOrderBy1

	return GetConnectionsOrderBy{
		GetConnectionsOrderBy1: &getConnectionsOrderBy1,
		Type:                   typ,
	}
}

func CreateGetConnectionsOrderByArrayOfgetConnectionsOrderBy2(arrayOfgetConnectionsOrderBy2 []GetConnectionsOrderBy2) GetConnectionsOrderBy {
	typ := GetConnectionsOrderByTypeArrayOfgetConnectionsOrderBy2

	return GetConnectionsOrderBy{
		ArrayOfgetConnectionsOrderBy2: arrayOfgetConnectionsOrderBy2,
		Type:                          typ,
	}
}

func (u *GetConnectionsOrderBy) UnmarshalJSON(data []byte) error {

	getConnectionsOrderBy1 := GetConnectionsOrderBy1("")
	if err := utils.UnmarshalJSON(data, &getConnectionsOrderBy1, "", true, true); err == nil {
		u.GetConnectionsOrderBy1 = &getConnectionsOrderBy1
		u.Type = GetConnectionsOrderByTypeGetConnectionsOrderBy1
		return nil
	}

	arrayOfgetConnectionsOrderBy2 := []GetConnectionsOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetConnectionsOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetConnectionsOrderBy2 = arrayOfgetConnectionsOrderBy2
		u.Type = GetConnectionsOrderByTypeArrayOfgetConnectionsOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetConnectionsOrderBy1 != nil {
		return utils.MarshalJSON(u.GetConnectionsOrderBy1, "", true)
	}

	if u.ArrayOfgetConnectionsOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetConnectionsOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetConnectionsPausedAt2 - Date the connection was paused
type GetConnectionsPausedAt2 struct {
	Any *bool      `queryParam:"name=any"`
	Gt  *time.Time `queryParam:"name=gt"`
	Gte *time.Time `queryParam:"name=gte"`
	Le  *time.Time `queryParam:"name=le"`
	Lte *time.Time `queryParam:"name=lte"`
}

func (g GetConnectionsPausedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetConnectionsPausedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetConnectionsPausedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetConnectionsPausedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetConnectionsPausedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetConnectionsPausedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetConnectionsPausedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetConnectionsPausedAtType string

const (
	GetConnectionsPausedAtTypeDateTime                GetConnectionsPausedAtType = "date-time"
	GetConnectionsPausedAtTypeGetConnectionsPausedAt2 GetConnectionsPausedAtType = "getConnectionsPausedAt_2"
)

type GetConnectionsPausedAt struct {
	DateTime                *time.Time
	GetConnectionsPausedAt2 *GetConnectionsPausedAt2

	Type GetConnectionsPausedAtType
}

func CreateGetConnectionsPausedAtDateTime(dateTime time.Time) GetConnectionsPausedAt {
	typ := GetConnectionsPausedAtTypeDateTime

	return GetConnectionsPausedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetConnectionsPausedAtGetConnectionsPausedAt2(getConnectionsPausedAt2 GetConnectionsPausedAt2) GetConnectionsPausedAt {
	typ := GetConnectionsPausedAtTypeGetConnectionsPausedAt2

	return GetConnectionsPausedAt{
		GetConnectionsPausedAt2: &getConnectionsPausedAt2,
		Type:                    typ,
	}
}

func (u *GetConnectionsPausedAt) UnmarshalJSON(data []byte) error {

	getConnectionsPausedAt2 := GetConnectionsPausedAt2{}
	if err := utils.UnmarshalJSON(data, &getConnectionsPausedAt2, "", true, true); err == nil {
		u.GetConnectionsPausedAt2 = &getConnectionsPausedAt2
		u.Type = GetConnectionsPausedAtTypeGetConnectionsPausedAt2
		return nil
	}

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = GetConnectionsPausedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsPausedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetConnectionsPausedAt2 != nil {
		return utils.MarshalJSON(u.GetConnectionsPausedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsSourceIDType string

const (
	GetConnectionsSourceIDTypeStr        GetConnectionsSourceIDType = "str"
	GetConnectionsSourceIDTypeArrayOfstr GetConnectionsSourceIDType = "arrayOfstr"
)

type GetConnectionsSourceID struct {
	Str        *string
	ArrayOfstr []string

	Type GetConnectionsSourceIDType
}

func CreateGetConnectionsSourceIDStr(str string) GetConnectionsSourceID {
	typ := GetConnectionsSourceIDTypeStr

	return GetConnectionsSourceID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetConnectionsSourceIDArrayOfstr(arrayOfstr []string) GetConnectionsSourceID {
	typ := GetConnectionsSourceIDTypeArrayOfstr

	return GetConnectionsSourceID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetConnectionsSourceID) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = GetConnectionsSourceIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetConnectionsSourceIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetConnectionsSourceID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetConnectionsRequest struct {
	Archived      *bool                        `queryParam:"style=form,explode=true,name=archived"`
	ArchivedAt    *GetConnectionsArchivedAt    `queryParam:"style=form,explode=true,name=archived_at"`
	DestinationID *GetConnectionsDestinationID `queryParam:"style=form,explode=true,name=destination_id"`
	Dir           *GetConnectionsDir           `queryParam:"style=form,explode=true,name=dir"`
	FullName      *string                      `queryParam:"style=form,explode=true,name=full_name"`
	ID            *GetConnectionsID            `queryParam:"style=form,explode=true,name=id"`
	Name          *GetConnectionsName          `queryParam:"style=form,explode=true,name=name"`
	OrderBy       *GetConnectionsOrderBy       `queryParam:"style=form,explode=true,name=order_by"`
	PausedAt      *GetConnectionsPausedAt      `queryParam:"style=form,explode=true,name=paused_at"`
	SourceID      *GetConnectionsSourceID      `queryParam:"style=form,explode=true,name=source_id"`
}

func (o *GetConnectionsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetConnectionsRequest) GetArchivedAt() *GetConnectionsArchivedAt {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetConnectionsRequest) GetDestinationID() *GetConnectionsDestinationID {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *GetConnectionsRequest) GetDir() *GetConnectionsDir {
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

func (o *GetConnectionsRequest) GetID() *GetConnectionsID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetConnectionsRequest) GetName() *GetConnectionsName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetConnectionsRequest) GetOrderBy() *GetConnectionsOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetConnectionsRequest) GetPausedAt() *GetConnectionsPausedAt {
	if o == nil {
		return nil
	}
	return o.PausedAt
}

func (o *GetConnectionsRequest) GetSourceID() *GetConnectionsSourceID {
	if o == nil {
		return nil
	}
	return o.SourceID
}

type GetConnectionsResponse struct {
	// List of connections
	ConnectionPaginatedResult *shared.ConnectionPaginatedResult
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConnectionsResponse) GetConnectionPaginatedResult() *shared.ConnectionPaginatedResult {
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

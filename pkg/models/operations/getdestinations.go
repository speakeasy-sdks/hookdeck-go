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

// GetDestinationsArchivedAt2 - Date the destination was archived
type GetDestinationsArchivedAt2 struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Any                  *bool                  `queryParam:"name=any"`
	Gt                   *time.Time             `queryParam:"name=gt"`
	Gte                  *time.Time             `queryParam:"name=gte"`
	Le                   *time.Time             `queryParam:"name=le"`
	Lte                  *time.Time             `queryParam:"name=lte"`
}

func (g GetDestinationsArchivedAt2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDestinationsArchivedAt2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDestinationsArchivedAt2) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GetDestinationsArchivedAt2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetDestinationsArchivedAt2) GetGt() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetDestinationsArchivedAt2) GetGte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetDestinationsArchivedAt2) GetLe() *time.Time {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetDestinationsArchivedAt2) GetLte() *time.Time {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetDestinationsArchivedAtType string

const (
	GetDestinationsArchivedAtTypeDateTime                   GetDestinationsArchivedAtType = "date-time"
	GetDestinationsArchivedAtTypeGetDestinationsArchivedAt2 GetDestinationsArchivedAtType = "getDestinationsArchivedAt_2"
)

type GetDestinationsArchivedAt struct {
	DateTime                   *time.Time
	GetDestinationsArchivedAt2 *GetDestinationsArchivedAt2

	Type GetDestinationsArchivedAtType
}

func CreateGetDestinationsArchivedAtDateTime(dateTime time.Time) GetDestinationsArchivedAt {
	typ := GetDestinationsArchivedAtTypeDateTime

	return GetDestinationsArchivedAt{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateGetDestinationsArchivedAtGetDestinationsArchivedAt2(getDestinationsArchivedAt2 GetDestinationsArchivedAt2) GetDestinationsArchivedAt {
	typ := GetDestinationsArchivedAtTypeGetDestinationsArchivedAt2

	return GetDestinationsArchivedAt{
		GetDestinationsArchivedAt2: &getDestinationsArchivedAt2,
		Type:                       typ,
	}
}

func (u *GetDestinationsArchivedAt) UnmarshalJSON(data []byte) error {

	getDestinationsArchivedAt2 := new(GetDestinationsArchivedAt2)
	if err := utils.UnmarshalJSON(data, &getDestinationsArchivedAt2, "", true, true); err == nil {
		u.GetDestinationsArchivedAt2 = getDestinationsArchivedAt2
		u.Type = GetDestinationsArchivedAtTypeGetDestinationsArchivedAt2
		return nil
	}

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = GetDestinationsArchivedAtTypeDateTime
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDestinationsArchivedAt) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.GetDestinationsArchivedAt2 != nil {
		return utils.MarshalJSON(u.GetDestinationsArchivedAt2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetDestinationsCliPath2 - Path for the CLI destination
type GetDestinationsCliPath2 struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Any                  *bool                  `queryParam:"name=any"`
}

func (g GetDestinationsCliPath2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDestinationsCliPath2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDestinationsCliPath2) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GetDestinationsCliPath2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

type GetDestinationsCliPathType string

const (
	GetDestinationsCliPathTypeStr                     GetDestinationsCliPathType = "str"
	GetDestinationsCliPathTypeGetDestinationsCliPath2 GetDestinationsCliPathType = "getDestinationsCliPath_2"
	GetDestinationsCliPathTypeArrayOfstr              GetDestinationsCliPathType = "arrayOfstr"
)

type GetDestinationsCliPath struct {
	Str                     *string
	GetDestinationsCliPath2 *GetDestinationsCliPath2
	ArrayOfstr              []string

	Type GetDestinationsCliPathType
}

func CreateGetDestinationsCliPathStr(str string) GetDestinationsCliPath {
	typ := GetDestinationsCliPathTypeStr

	return GetDestinationsCliPath{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetDestinationsCliPathGetDestinationsCliPath2(getDestinationsCliPath2 GetDestinationsCliPath2) GetDestinationsCliPath {
	typ := GetDestinationsCliPathTypeGetDestinationsCliPath2

	return GetDestinationsCliPath{
		GetDestinationsCliPath2: &getDestinationsCliPath2,
		Type:                    typ,
	}
}

func CreateGetDestinationsCliPathArrayOfstr(arrayOfstr []string) GetDestinationsCliPath {
	typ := GetDestinationsCliPathTypeArrayOfstr

	return GetDestinationsCliPath{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetDestinationsCliPath) UnmarshalJSON(data []byte) error {

	getDestinationsCliPath2 := new(GetDestinationsCliPath2)
	if err := utils.UnmarshalJSON(data, &getDestinationsCliPath2, "", true, true); err == nil {
		u.GetDestinationsCliPath2 = getDestinationsCliPath2
		u.Type = GetDestinationsCliPathTypeGetDestinationsCliPath2
		return nil
	}

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetDestinationsCliPathTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetDestinationsCliPathTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDestinationsCliPath) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.GetDestinationsCliPath2 != nil {
		return utils.MarshalJSON(u.GetDestinationsCliPath2, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetDestinationsDir2 string

const (
	GetDestinationsDir2Asc  GetDestinationsDir2 = "asc"
	GetDestinationsDir2Desc GetDestinationsDir2 = "desc"
)

func (e GetDestinationsDir2) ToPointer() *GetDestinationsDir2 {
	return &e
}

func (e *GetDestinationsDir2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetDestinationsDir2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationsDir2: %v", v)
	}
}

// GetDestinationsDir1 - Sort direction
type GetDestinationsDir1 string

const (
	GetDestinationsDir1Asc  GetDestinationsDir1 = "asc"
	GetDestinationsDir1Desc GetDestinationsDir1 = "desc"
)

func (e GetDestinationsDir1) ToPointer() *GetDestinationsDir1 {
	return &e
}

func (e *GetDestinationsDir1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetDestinationsDir1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationsDir1: %v", v)
	}
}

type GetDestinationsDirType string

const (
	GetDestinationsDirTypeGetDestinationsDir1        GetDestinationsDirType = "getDestinationsDir_1"
	GetDestinationsDirTypeArrayOfgetDestinationsDir2 GetDestinationsDirType = "arrayOfgetDestinationsDir_2"
)

type GetDestinationsDir struct {
	GetDestinationsDir1        *GetDestinationsDir1
	ArrayOfgetDestinationsDir2 []GetDestinationsDir2

	Type GetDestinationsDirType
}

func CreateGetDestinationsDirGetDestinationsDir1(getDestinationsDir1 GetDestinationsDir1) GetDestinationsDir {
	typ := GetDestinationsDirTypeGetDestinationsDir1

	return GetDestinationsDir{
		GetDestinationsDir1: &getDestinationsDir1,
		Type:                typ,
	}
}

func CreateGetDestinationsDirArrayOfgetDestinationsDir2(arrayOfgetDestinationsDir2 []GetDestinationsDir2) GetDestinationsDir {
	typ := GetDestinationsDirTypeArrayOfgetDestinationsDir2

	return GetDestinationsDir{
		ArrayOfgetDestinationsDir2: arrayOfgetDestinationsDir2,
		Type:                       typ,
	}
}

func (u *GetDestinationsDir) UnmarshalJSON(data []byte) error {

	getDestinationsDir1 := new(GetDestinationsDir1)
	if err := utils.UnmarshalJSON(data, &getDestinationsDir1, "", true, true); err == nil {
		u.GetDestinationsDir1 = getDestinationsDir1
		u.Type = GetDestinationsDirTypeGetDestinationsDir1
		return nil
	}

	arrayOfgetDestinationsDir2 := []GetDestinationsDir2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetDestinationsDir2, "", true, true); err == nil {
		u.ArrayOfgetDestinationsDir2 = arrayOfgetDestinationsDir2
		u.Type = GetDestinationsDirTypeArrayOfgetDestinationsDir2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDestinationsDir) MarshalJSON() ([]byte, error) {
	if u.GetDestinationsDir1 != nil {
		return utils.MarshalJSON(u.GetDestinationsDir1, "", true)
	}

	if u.ArrayOfgetDestinationsDir2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetDestinationsDir2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetDestinationsIDType string

const (
	GetDestinationsIDTypeStr        GetDestinationsIDType = "str"
	GetDestinationsIDTypeArrayOfstr GetDestinationsIDType = "arrayOfstr"
)

type GetDestinationsID struct {
	Str        *string
	ArrayOfstr []string

	Type GetDestinationsIDType
}

func CreateGetDestinationsIDStr(str string) GetDestinationsID {
	typ := GetDestinationsIDTypeStr

	return GetDestinationsID{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetDestinationsIDArrayOfstr(arrayOfstr []string) GetDestinationsID {
	typ := GetDestinationsIDTypeArrayOfstr

	return GetDestinationsID{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetDestinationsID) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetDestinationsIDTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetDestinationsIDTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDestinationsID) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// GetDestinationsName2 - The destination name
type GetDestinationsName2 struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Any                  *bool                  `queryParam:"name=any"`
	Contains             *string                `queryParam:"name=contains"`
	Gt                   *string                `queryParam:"name=gt"`
	Gte                  *string                `queryParam:"name=gte"`
	Le                   *string                `queryParam:"name=le"`
	Lte                  *string                `queryParam:"name=lte"`
}

func (g GetDestinationsName2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetDestinationsName2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDestinationsName2) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *GetDestinationsName2) GetAny() *bool {
	if o == nil {
		return nil
	}
	return o.Any
}

func (o *GetDestinationsName2) GetContains() *string {
	if o == nil {
		return nil
	}
	return o.Contains
}

func (o *GetDestinationsName2) GetGt() *string {
	if o == nil {
		return nil
	}
	return o.Gt
}

func (o *GetDestinationsName2) GetGte() *string {
	if o == nil {
		return nil
	}
	return o.Gte
}

func (o *GetDestinationsName2) GetLe() *string {
	if o == nil {
		return nil
	}
	return o.Le
}

func (o *GetDestinationsName2) GetLte() *string {
	if o == nil {
		return nil
	}
	return o.Lte
}

type GetDestinationsNameType string

const (
	GetDestinationsNameTypeStr                  GetDestinationsNameType = "str"
	GetDestinationsNameTypeGetDestinationsName2 GetDestinationsNameType = "getDestinationsName_2"
)

type GetDestinationsName struct {
	Str                  *string
	GetDestinationsName2 *GetDestinationsName2

	Type GetDestinationsNameType
}

func CreateGetDestinationsNameStr(str string) GetDestinationsName {
	typ := GetDestinationsNameTypeStr

	return GetDestinationsName{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetDestinationsNameGetDestinationsName2(getDestinationsName2 GetDestinationsName2) GetDestinationsName {
	typ := GetDestinationsNameTypeGetDestinationsName2

	return GetDestinationsName{
		GetDestinationsName2: &getDestinationsName2,
		Type:                 typ,
	}
}

func (u *GetDestinationsName) UnmarshalJSON(data []byte) error {

	getDestinationsName2 := new(GetDestinationsName2)
	if err := utils.UnmarshalJSON(data, &getDestinationsName2, "", true, true); err == nil {
		u.GetDestinationsName2 = getDestinationsName2
		u.Type = GetDestinationsNameTypeGetDestinationsName2
		return nil
	}

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetDestinationsNameTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDestinationsName) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.GetDestinationsName2 != nil {
		return utils.MarshalJSON(u.GetDestinationsName2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetDestinationsOrderBy2 string

const (
	GetDestinationsOrderBy2CreatedAt GetDestinationsOrderBy2 = "created_at"
)

func (e GetDestinationsOrderBy2) ToPointer() *GetDestinationsOrderBy2 {
	return &e
}

func (e *GetDestinationsOrderBy2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetDestinationsOrderBy2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationsOrderBy2: %v", v)
	}
}

// GetDestinationsOrderBy1 - Sort key(s)
type GetDestinationsOrderBy1 string

const (
	GetDestinationsOrderBy1CreatedAt GetDestinationsOrderBy1 = "created_at"
)

func (e GetDestinationsOrderBy1) ToPointer() *GetDestinationsOrderBy1 {
	return &e
}

func (e *GetDestinationsOrderBy1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		*e = GetDestinationsOrderBy1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationsOrderBy1: %v", v)
	}
}

type GetDestinationsOrderByType string

const (
	GetDestinationsOrderByTypeGetDestinationsOrderBy1        GetDestinationsOrderByType = "getDestinationsOrderBy_1"
	GetDestinationsOrderByTypeArrayOfgetDestinationsOrderBy2 GetDestinationsOrderByType = "arrayOfgetDestinationsOrderBy_2"
)

type GetDestinationsOrderBy struct {
	GetDestinationsOrderBy1        *GetDestinationsOrderBy1
	ArrayOfgetDestinationsOrderBy2 []GetDestinationsOrderBy2

	Type GetDestinationsOrderByType
}

func CreateGetDestinationsOrderByGetDestinationsOrderBy1(getDestinationsOrderBy1 GetDestinationsOrderBy1) GetDestinationsOrderBy {
	typ := GetDestinationsOrderByTypeGetDestinationsOrderBy1

	return GetDestinationsOrderBy{
		GetDestinationsOrderBy1: &getDestinationsOrderBy1,
		Type:                    typ,
	}
}

func CreateGetDestinationsOrderByArrayOfgetDestinationsOrderBy2(arrayOfgetDestinationsOrderBy2 []GetDestinationsOrderBy2) GetDestinationsOrderBy {
	typ := GetDestinationsOrderByTypeArrayOfgetDestinationsOrderBy2

	return GetDestinationsOrderBy{
		ArrayOfgetDestinationsOrderBy2: arrayOfgetDestinationsOrderBy2,
		Type:                           typ,
	}
}

func (u *GetDestinationsOrderBy) UnmarshalJSON(data []byte) error {

	getDestinationsOrderBy1 := new(GetDestinationsOrderBy1)
	if err := utils.UnmarshalJSON(data, &getDestinationsOrderBy1, "", true, true); err == nil {
		u.GetDestinationsOrderBy1 = getDestinationsOrderBy1
		u.Type = GetDestinationsOrderByTypeGetDestinationsOrderBy1
		return nil
	}

	arrayOfgetDestinationsOrderBy2 := []GetDestinationsOrderBy2{}
	if err := utils.UnmarshalJSON(data, &arrayOfgetDestinationsOrderBy2, "", true, true); err == nil {
		u.ArrayOfgetDestinationsOrderBy2 = arrayOfgetDestinationsOrderBy2
		u.Type = GetDestinationsOrderByTypeArrayOfgetDestinationsOrderBy2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDestinationsOrderBy) MarshalJSON() ([]byte, error) {
	if u.GetDestinationsOrderBy1 != nil {
		return utils.MarshalJSON(u.GetDestinationsOrderBy1, "", true)
	}

	if u.ArrayOfgetDestinationsOrderBy2 != nil {
		return utils.MarshalJSON(u.ArrayOfgetDestinationsOrderBy2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetDestinationsURLType string

const (
	GetDestinationsURLTypeStr        GetDestinationsURLType = "str"
	GetDestinationsURLTypeArrayOfstr GetDestinationsURLType = "arrayOfstr"
)

type GetDestinationsURL struct {
	Str        *string
	ArrayOfstr []string

	Type GetDestinationsURLType
}

func CreateGetDestinationsURLStr(str string) GetDestinationsURL {
	typ := GetDestinationsURLTypeStr

	return GetDestinationsURL{
		Str:  &str,
		Type: typ,
	}
}

func CreateGetDestinationsURLArrayOfstr(arrayOfstr []string) GetDestinationsURL {
	typ := GetDestinationsURLTypeArrayOfstr

	return GetDestinationsURL{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *GetDestinationsURL) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = GetDestinationsURLTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = GetDestinationsURLTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetDestinationsURL) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetDestinationsRequest struct {
	Archived   *bool                      `queryParam:"style=form,explode=true,name=archived"`
	ArchivedAt *GetDestinationsArchivedAt `queryParam:"style=form,explode=true,name=archived_at"`
	CliPath    *GetDestinationsCliPath    `queryParam:"style=form,explode=true,name=cli_path"`
	Dir        *GetDestinationsDir        `queryParam:"style=form,explode=true,name=dir"`
	ID         *GetDestinationsID         `queryParam:"style=form,explode=true,name=id"`
	Limit      *int64                     `queryParam:"style=form,explode=true,name=limit"`
	Name       *GetDestinationsName       `queryParam:"style=form,explode=true,name=name"`
	Next       *string                    `queryParam:"style=form,explode=true,name=next"`
	OrderBy    *GetDestinationsOrderBy    `queryParam:"style=form,explode=true,name=order_by"`
	Prev       *string                    `queryParam:"style=form,explode=true,name=prev"`
	URL        *GetDestinationsURL        `queryParam:"style=form,explode=true,name=url"`
}

func (o *GetDestinationsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetDestinationsRequest) GetArchivedAt() *GetDestinationsArchivedAt {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetDestinationsRequest) GetCliPath() *GetDestinationsCliPath {
	if o == nil {
		return nil
	}
	return o.CliPath
}

func (o *GetDestinationsRequest) GetDir() *GetDestinationsDir {
	if o == nil {
		return nil
	}
	return o.Dir
}

func (o *GetDestinationsRequest) GetID() *GetDestinationsID {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetDestinationsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetDestinationsRequest) GetName() *GetDestinationsName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetDestinationsRequest) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetDestinationsRequest) GetOrderBy() *GetDestinationsOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *GetDestinationsRequest) GetPrev() *string {
	if o == nil {
		return nil
	}
	return o.Prev
}

func (o *GetDestinationsRequest) GetURL() *GetDestinationsURL {
	if o == nil {
		return nil
	}
	return o.URL
}

type GetDestinationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of destinations
	DestinationPaginatedResult *shared.DestinationPaginatedResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDestinationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDestinationsResponse) GetDestinationPaginatedResult() *shared.DestinationPaginatedResult {
	if o == nil {
		return nil
	}
	return o.DestinationPaginatedResult
}

func (o *GetDestinationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDestinationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

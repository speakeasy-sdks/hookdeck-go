// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type EventDataBody2 struct {
}

type EventDataBodyType string

const (
	EventDataBodyTypeStr            EventDataBodyType = "str"
	EventDataBodyTypeEventDataBody2 EventDataBodyType = "Event_data_body_2"
	EventDataBodyTypeArrayOfany     EventDataBodyType = "arrayOfany"
)

type EventDataBody struct {
	Str            *string
	EventDataBody2 *EventDataBody2
	ArrayOfany     []interface{}

	Type EventDataBodyType
}

func CreateEventDataBodyStr(str string) EventDataBody {
	typ := EventDataBodyTypeStr

	return EventDataBody{
		Str:  &str,
		Type: typ,
	}
}

func CreateEventDataBodyEventDataBody2(eventDataBody2 EventDataBody2) EventDataBody {
	typ := EventDataBodyTypeEventDataBody2

	return EventDataBody{
		EventDataBody2: &eventDataBody2,
		Type:           typ,
	}
}

func CreateEventDataBodyArrayOfany(arrayOfany []interface{}) EventDataBody {
	typ := EventDataBodyTypeArrayOfany

	return EventDataBody{
		ArrayOfany: arrayOfany,
		Type:       typ,
	}
}

func (u *EventDataBody) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = EventDataBodyTypeStr
		return nil
	}

	eventDataBody2 := new(EventDataBody2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&eventDataBody2); err == nil {
		u.EventDataBody2 = eventDataBody2
		u.Type = EventDataBodyTypeEventDataBody2
		return nil
	}

	arrayOfany := []interface{}{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfany); err == nil {
		u.ArrayOfany = arrayOfany
		u.Type = EventDataBodyTypeArrayOfany
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EventDataBody) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.EventDataBody2 != nil {
		return json.Marshal(u.EventDataBody2)
	}

	if u.ArrayOfany != nil {
		return json.Marshal(u.ArrayOfany)
	}

	return nil, nil
}

type EventDataHeadersType string

const (
	EventDataHeadersTypeStr      EventDataHeadersType = "str"
	EventDataHeadersTypeMapOfstr EventDataHeadersType = "mapOfstr"
)

type EventDataHeaders struct {
	Str      *string
	MapOfstr map[string]string

	Type EventDataHeadersType
}

func CreateEventDataHeadersStr(str string) EventDataHeaders {
	typ := EventDataHeadersTypeStr

	return EventDataHeaders{
		Str:  &str,
		Type: typ,
	}
}

func CreateEventDataHeadersMapOfstr(mapOfstr map[string]string) EventDataHeaders {
	typ := EventDataHeadersTypeMapOfstr

	return EventDataHeaders{
		MapOfstr: mapOfstr,
		Type:     typ,
	}
}

func (u *EventDataHeaders) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = EventDataHeadersTypeStr
		return nil
	}

	mapOfstr := map[string]string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&mapOfstr); err == nil {
		u.MapOfstr = mapOfstr
		u.Type = EventDataHeadersTypeMapOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EventDataHeaders) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.MapOfstr != nil {
		return json.Marshal(u.MapOfstr)
	}

	return nil, nil
}

type EventDataParsedQuery2 struct {
}

type EventDataParsedQueryType string

const (
	EventDataParsedQueryTypeStr                   EventDataParsedQueryType = "str"
	EventDataParsedQueryTypeEventDataParsedQuery2 EventDataParsedQueryType = "Event_data_parsed_query_2"
)

type EventDataParsedQuery struct {
	Str                   *string
	EventDataParsedQuery2 *EventDataParsedQuery2

	Type EventDataParsedQueryType
}

func CreateEventDataParsedQueryStr(str string) EventDataParsedQuery {
	typ := EventDataParsedQueryTypeStr

	return EventDataParsedQuery{
		Str:  &str,
		Type: typ,
	}
}

func CreateEventDataParsedQueryEventDataParsedQuery2(eventDataParsedQuery2 EventDataParsedQuery2) EventDataParsedQuery {
	typ := EventDataParsedQueryTypeEventDataParsedQuery2

	return EventDataParsedQuery{
		EventDataParsedQuery2: &eventDataParsedQuery2,
		Type:                  typ,
	}
}

func (u *EventDataParsedQuery) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = EventDataParsedQueryTypeStr
		return nil
	}

	eventDataParsedQuery2 := new(EventDataParsedQuery2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&eventDataParsedQuery2); err == nil {
		u.EventDataParsedQuery2 = eventDataParsedQuery2
		u.Type = EventDataParsedQueryTypeEventDataParsedQuery2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EventDataParsedQuery) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.EventDataParsedQuery2 != nil {
		return json.Marshal(u.EventDataParsedQuery2)
	}

	return nil, nil
}

type EventData struct {
	Body           *EventDataBody        `json:"body,omitempty"`
	Headers        *EventDataHeaders     `json:"headers,omitempty"`
	IsLargePayload *bool                 `json:"is_large_payload,omitempty"`
	ParsedQuery    *EventDataParsedQuery `json:"parsed_query,omitempty"`
	Path           string                `json:"path"`
	Query          *string               `json:"query,omitempty"`
}

func (o *EventData) GetBody() *EventDataBody {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *EventData) GetHeaders() *EventDataHeaders {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *EventData) GetIsLargePayload() *bool {
	if o == nil {
		return nil
	}
	return o.IsLargePayload
}

func (o *EventData) GetParsedQuery() *EventDataParsedQuery {
	if o == nil {
		return nil
	}
	return o.ParsedQuery
}

func (o *EventData) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *EventData) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

// Event - A single event
type Event struct {
	// Number of delivery attempts made
	Attempts int64 `json:"attempts"`
	// ID of the CLI the event is sent to
	CliID *string `json:"cli_id,omitempty"`
	// Date the event was created
	CreatedAt   time.Time  `json:"created_at"`
	CreatedAtID *string    `json:"created_at_id,omitempty"`
	Data        *EventData `json:"data,omitempty"`
	// ID of the associated destination
	DestinationID string `json:"destination_id"`
	// ID of the request data
	EventDataID string `json:"event_data_id"`
	// ID of the event
	ID string `json:"id"`
	// Date of the most recently attempted retry
	LastAttemptAt   *time.Time `json:"last_attempt_at,omitempty"`
	LastAttemptAtID *string    `json:"last_attempt_at_id,omitempty"`
	// Date of the next scheduled retry
	NextAttemptAt *time.Time `json:"next_attempt_at,omitempty"`
	// ID of the request that created the event
	RequestID string `json:"request_id"`
	// Event status
	ResponseStatus *int64 `json:"response_status,omitempty"`
	// ID of the associated source
	SourceID string      `json:"source_id"`
	Status   EventStatus `json:"status"`
	// Date of the latest successful attempt
	SuccessfulAt *time.Time `json:"successful_at,omitempty"`
	// ID of the workspace
	TeamID string `json:"team_id"`
	// Date the event was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// ID of the associated connection
	WebhookID string `json:"webhook_id"`
}

func (o *Event) GetAttempts() int64 {
	if o == nil {
		return 0
	}
	return o.Attempts
}

func (o *Event) GetCliID() *string {
	if o == nil {
		return nil
	}
	return o.CliID
}

func (o *Event) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Event) GetCreatedAtID() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAtID
}

func (o *Event) GetData() *EventData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Event) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *Event) GetEventDataID() string {
	if o == nil {
		return ""
	}
	return o.EventDataID
}

func (o *Event) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Event) GetLastAttemptAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastAttemptAt
}

func (o *Event) GetLastAttemptAtID() *string {
	if o == nil {
		return nil
	}
	return o.LastAttemptAtID
}

func (o *Event) GetNextAttemptAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.NextAttemptAt
}

func (o *Event) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

func (o *Event) GetResponseStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.ResponseStatus
}

func (o *Event) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *Event) GetStatus() EventStatus {
	if o == nil {
		return EventStatus("")
	}
	return o.Status
}

func (o *Event) GetSuccessfulAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.SuccessfulAt
}

func (o *Event) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *Event) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Event) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

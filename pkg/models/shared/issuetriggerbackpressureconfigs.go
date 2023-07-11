// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type IssueTriggerBackpressureConfigsDestinationsType string

const (
	IssueTriggerBackpressureConfigsDestinationsTypeStr        IssueTriggerBackpressureConfigsDestinationsType = "str"
	IssueTriggerBackpressureConfigsDestinationsTypeArrayOfstr IssueTriggerBackpressureConfigsDestinationsType = "arrayOfstr"
)

type IssueTriggerBackpressureConfigsDestinations struct {
	Str        *string
	ArrayOfstr []string

	Type IssueTriggerBackpressureConfigsDestinationsType
}

func CreateIssueTriggerBackpressureConfigsDestinationsStr(str string) IssueTriggerBackpressureConfigsDestinations {
	typ := IssueTriggerBackpressureConfigsDestinationsTypeStr

	return IssueTriggerBackpressureConfigsDestinations{
		Str:  &str,
		Type: typ,
	}
}

func CreateIssueTriggerBackpressureConfigsDestinationsArrayOfstr(arrayOfstr []string) IssueTriggerBackpressureConfigsDestinations {
	typ := IssueTriggerBackpressureConfigsDestinationsTypeArrayOfstr

	return IssueTriggerBackpressureConfigsDestinations{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *IssueTriggerBackpressureConfigsDestinations) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = IssueTriggerBackpressureConfigsDestinationsTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = IssueTriggerBackpressureConfigsDestinationsTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u IssueTriggerBackpressureConfigsDestinations) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// IssueTriggerBackpressureConfigs - Configurations for a 'Backpressure' issue trigger
type IssueTriggerBackpressureConfigs struct {
	// The minimum delay (backpressure) to open the issue for min of 1 minute (60000) and max of 1 day (86400000)
	Delay int64 `json:"delay"`
	// A pattern to match on the destination name or array of destination IDs. Use `*` as wildcard.
	Destinations IssueTriggerBackpressureConfigsDestinations `json:"destinations"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/v2/internal/utils"
)

type DestinationsType string

const (
	DestinationsTypeStr        DestinationsType = "str"
	DestinationsTypeArrayOfstr DestinationsType = "arrayOfstr"
)

// Destinations - A pattern to match on the destination name or array of destination IDs. Use `*` as wildcard.
type Destinations struct {
	Str        *string
	ArrayOfstr []string

	Type DestinationsType
}

func CreateDestinationsStr(str string) Destinations {
	typ := DestinationsTypeStr

	return Destinations{
		Str:  &str,
		Type: typ,
	}
}

func CreateDestinationsArrayOfstr(arrayOfstr []string) Destinations {
	typ := DestinationsTypeArrayOfstr

	return Destinations{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *Destinations) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DestinationsTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = DestinationsTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Destinations) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// IssueTriggerBackpressureConfigs - Configurations for a 'Backpressure' issue trigger
type IssueTriggerBackpressureConfigs struct {
	// The minimum delay (backpressure) to open the issue for min of 1 minute (60000) and max of 1 day (86400000)
	Delay int64 `json:"delay"`
	// A pattern to match on the destination name or array of destination IDs. Use `*` as wildcard.
	Destinations Destinations `json:"destinations"`
}

func (o *IssueTriggerBackpressureConfigs) GetDelay() int64 {
	if o == nil {
		return 0
	}
	return o.Delay
}

func (o *IssueTriggerBackpressureConfigs) GetDestinations() Destinations {
	if o == nil {
		return Destinations{}
	}
	return o.Destinations
}

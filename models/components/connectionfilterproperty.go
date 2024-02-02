// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/v2/internal/utils"
)

type Four struct {
}

type ConnectionFilterPropertyType string

const (
	ConnectionFilterPropertyTypeStr     ConnectionFilterPropertyType = "str"
	ConnectionFilterPropertyTypeFloat32 ConnectionFilterPropertyType = "float32"
	ConnectionFilterPropertyTypeBoolean ConnectionFilterPropertyType = "boolean"
	ConnectionFilterPropertyTypeFour    ConnectionFilterPropertyType = "4"
)

// ConnectionFilterProperty - JSON using our filter syntax to filter on request headers
type ConnectionFilterProperty struct {
	Str     *string
	Float32 *float32
	Boolean *bool
	Four    *Four

	Type ConnectionFilterPropertyType
}

func CreateConnectionFilterPropertyStr(str string) ConnectionFilterProperty {
	typ := ConnectionFilterPropertyTypeStr

	return ConnectionFilterProperty{
		Str:  &str,
		Type: typ,
	}
}

func CreateConnectionFilterPropertyFloat32(float32T float32) ConnectionFilterProperty {
	typ := ConnectionFilterPropertyTypeFloat32

	return ConnectionFilterProperty{
		Float32: &float32T,
		Type:    typ,
	}
}

func CreateConnectionFilterPropertyBoolean(boolean bool) ConnectionFilterProperty {
	typ := ConnectionFilterPropertyTypeBoolean

	return ConnectionFilterProperty{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateConnectionFilterPropertyFour(four Four) ConnectionFilterProperty {
	typ := ConnectionFilterPropertyTypeFour

	return ConnectionFilterProperty{
		Four: &four,
		Type: typ,
	}
}

func (u *ConnectionFilterProperty) UnmarshalJSON(data []byte) error {

	four := Four{}
	if err := utils.UnmarshalJSON(data, &four, "", true, true); err == nil {
		u.Four = &four
		u.Type = ConnectionFilterPropertyTypeFour
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ConnectionFilterPropertyTypeStr
		return nil
	}

	float32Var := float32(0)
	if err := utils.UnmarshalJSON(data, &float32Var, "", true, true); err == nil {
		u.Float32 = &float32Var
		u.Type = ConnectionFilterPropertyTypeFloat32
		return nil
	}

	boolean := false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = ConnectionFilterPropertyTypeBoolean
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConnectionFilterProperty) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Float32 != nil {
		return utils.MarshalJSON(u.Float32, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Four != nil {
		return utils.MarshalJSON(u.Four, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

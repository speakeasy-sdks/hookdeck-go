// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

// ConnectionFilterProperty4 - JSON using our filter syntax to filter on request headers
type ConnectionFilterProperty4 struct {
}

type ConnectionFilterPropertyType string

const (
	ConnectionFilterPropertyTypeStr                       ConnectionFilterPropertyType = "str"
	ConnectionFilterPropertyTypeFloat32                   ConnectionFilterPropertyType = "float32"
	ConnectionFilterPropertyTypeBoolean                   ConnectionFilterPropertyType = "boolean"
	ConnectionFilterPropertyTypeConnectionFilterProperty4 ConnectionFilterPropertyType = "ConnectionFilterProperty_4"
)

type ConnectionFilterProperty struct {
	Str                       *string
	Float32                   *float32
	Boolean                   *bool
	ConnectionFilterProperty4 *ConnectionFilterProperty4

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

func CreateConnectionFilterPropertyConnectionFilterProperty4(connectionFilterProperty4 ConnectionFilterProperty4) ConnectionFilterProperty {
	typ := ConnectionFilterPropertyTypeConnectionFilterProperty4

	return ConnectionFilterProperty{
		ConnectionFilterProperty4: &connectionFilterProperty4,
		Type:                      typ,
	}
}

func (u *ConnectionFilterProperty) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = ConnectionFilterPropertyTypeStr
		return nil
	}

	float32Var := new(float32)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&float32Var); err == nil {
		u.Float32 = float32Var
		u.Type = ConnectionFilterPropertyTypeFloat32
		return nil
	}

	boolean := new(bool)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&boolean); err == nil {
		u.Boolean = boolean
		u.Type = ConnectionFilterPropertyTypeBoolean
		return nil
	}

	connectionFilterProperty4 := new(ConnectionFilterProperty4)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&connectionFilterProperty4); err == nil {
		u.ConnectionFilterProperty4 = connectionFilterProperty4
		u.Type = ConnectionFilterPropertyTypeConnectionFilterProperty4
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConnectionFilterProperty) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.Float32 != nil {
		return json.Marshal(u.Float32)
	}

	if u.Boolean != nil {
		return json.Marshal(u.Boolean)
	}

	if u.ConnectionFilterProperty4 != nil {
		return json.Marshal(u.ConnectionFilterProperty4)
	}

	return nil, nil
}

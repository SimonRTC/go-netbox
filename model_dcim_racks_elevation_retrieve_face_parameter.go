/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.7 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// DcimRacksElevationRetrieveFaceParameter the model 'DcimRacksElevationRetrieveFaceParameter'
type DcimRacksElevationRetrieveFaceParameter string

// List of dcim_racks_elevation_retrieve_face_parameter
const (
	DCIMRACKSELEVATIONRETRIEVEFACEPARAMETER_FRONT DcimRacksElevationRetrieveFaceParameter = "front"
	DCIMRACKSELEVATIONRETRIEVEFACEPARAMETER_REAR  DcimRacksElevationRetrieveFaceParameter = "rear"
)

// All allowed values of DcimRacksElevationRetrieveFaceParameter enum
var AllowedDcimRacksElevationRetrieveFaceParameterEnumValues = []DcimRacksElevationRetrieveFaceParameter{
	"front",
	"rear",
}

func (v *DcimRacksElevationRetrieveFaceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimRacksElevationRetrieveFaceParameter(value)
	for _, existing := range AllowedDcimRacksElevationRetrieveFaceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimRacksElevationRetrieveFaceParameter", value)
}

// NewDcimRacksElevationRetrieveFaceParameterFromValue returns a pointer to a valid DcimRacksElevationRetrieveFaceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimRacksElevationRetrieveFaceParameterFromValue(v string) (*DcimRacksElevationRetrieveFaceParameter, error) {
	ev := DcimRacksElevationRetrieveFaceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimRacksElevationRetrieveFaceParameter: valid values are %v", v, AllowedDcimRacksElevationRetrieveFaceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimRacksElevationRetrieveFaceParameter) IsValid() bool {
	for _, existing := range AllowedDcimRacksElevationRetrieveFaceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_racks_elevation_retrieve_face_parameter value
func (v DcimRacksElevationRetrieveFaceParameter) Ptr() *DcimRacksElevationRetrieveFaceParameter {
	return &v
}

type NullableDcimRacksElevationRetrieveFaceParameter struct {
	value *DcimRacksElevationRetrieveFaceParameter
	isSet bool
}

func (v NullableDcimRacksElevationRetrieveFaceParameter) Get() *DcimRacksElevationRetrieveFaceParameter {
	return v.value
}

func (v *NullableDcimRacksElevationRetrieveFaceParameter) Set(val *DcimRacksElevationRetrieveFaceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimRacksElevationRetrieveFaceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimRacksElevationRetrieveFaceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimRacksElevationRetrieveFaceParameter(val *DcimRacksElevationRetrieveFaceParameter) *NullableDcimRacksElevationRetrieveFaceParameter {
	return &NullableDcimRacksElevationRetrieveFaceParameter{value: val, isSet: true}
}

func (v NullableDcimRacksElevationRetrieveFaceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimRacksElevationRetrieveFaceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

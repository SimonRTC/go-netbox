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

// RackAirflowLabel the model 'RackAirflowLabel'
type RackAirflowLabel string

// List of Rack_airflow_label
const (
	RACKAIRFLOWLABEL_FRONT_TO_REAR RackAirflowLabel = "Front to rear"
	RACKAIRFLOWLABEL_REAR_TO_FRONT RackAirflowLabel = "Rear to front"
)

// All allowed values of RackAirflowLabel enum
var AllowedRackAirflowLabelEnumValues = []RackAirflowLabel{
	"Front to rear",
	"Rear to front",
}

func (v *RackAirflowLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackAirflowLabel(value)
	for _, existing := range AllowedRackAirflowLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackAirflowLabel", value)
}

// NewRackAirflowLabelFromValue returns a pointer to a valid RackAirflowLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackAirflowLabelFromValue(v string) (*RackAirflowLabel, error) {
	ev := RackAirflowLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackAirflowLabel: valid values are %v", v, AllowedRackAirflowLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackAirflowLabel) IsValid() bool {
	for _, existing := range AllowedRackAirflowLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_airflow_label value
func (v RackAirflowLabel) Ptr() *RackAirflowLabel {
	return &v
}

type NullableRackAirflowLabel struct {
	value *RackAirflowLabel
	isSet bool
}

func (v NullableRackAirflowLabel) Get() *RackAirflowLabel {
	return v.value
}

func (v *NullableRackAirflowLabel) Set(val *RackAirflowLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRackAirflowLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRackAirflowLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackAirflowLabel(val *RackAirflowLabel) *NullableRackAirflowLabel {
	return &NullableRackAirflowLabel{value: val, isSet: true}
}

func (v NullableRackAirflowLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackAirflowLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
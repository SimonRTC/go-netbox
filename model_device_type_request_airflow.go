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

// DeviceTypeRequestAirflow * `front-to-rear` - Front to rear * `rear-to-front` - Rear to front * `left-to-right` - Left to right * `right-to-left` - Right to left * `side-to-rear` - Side to rear * `rear-to-side` - Rear to side * `bottom-to-top` - Bottom to top * `top-to-bottom` - Top to bottom * `passive` - Passive * `mixed` - Mixed
type DeviceTypeRequestAirflow string

// List of DeviceTypeRequest_airflow
const (
	DEVICETYPEREQUESTAIRFLOW_FRONT_TO_REAR DeviceTypeRequestAirflow = "front-to-rear"
	DEVICETYPEREQUESTAIRFLOW_REAR_TO_FRONT DeviceTypeRequestAirflow = "rear-to-front"
	DEVICETYPEREQUESTAIRFLOW_LEFT_TO_RIGHT DeviceTypeRequestAirflow = "left-to-right"
	DEVICETYPEREQUESTAIRFLOW_RIGHT_TO_LEFT DeviceTypeRequestAirflow = "right-to-left"
	DEVICETYPEREQUESTAIRFLOW_SIDE_TO_REAR  DeviceTypeRequestAirflow = "side-to-rear"
	DEVICETYPEREQUESTAIRFLOW_REAR_TO_SIDE  DeviceTypeRequestAirflow = "rear-to-side"
	DEVICETYPEREQUESTAIRFLOW_BOTTOM_TO_TOP DeviceTypeRequestAirflow = "bottom-to-top"
	DEVICETYPEREQUESTAIRFLOW_TOP_TO_BOTTOM DeviceTypeRequestAirflow = "top-to-bottom"
	DEVICETYPEREQUESTAIRFLOW_PASSIVE       DeviceTypeRequestAirflow = "passive"
	DEVICETYPEREQUESTAIRFLOW_MIXED         DeviceTypeRequestAirflow = "mixed"
	DEVICETYPEREQUESTAIRFLOW_EMPTY         DeviceTypeRequestAirflow = ""
)

// All allowed values of DeviceTypeRequestAirflow enum
var AllowedDeviceTypeRequestAirflowEnumValues = []DeviceTypeRequestAirflow{
	"front-to-rear",
	"rear-to-front",
	"left-to-right",
	"right-to-left",
	"side-to-rear",
	"rear-to-side",
	"bottom-to-top",
	"top-to-bottom",
	"passive",
	"mixed",
	"",
}

func (v *DeviceTypeRequestAirflow) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceTypeRequestAirflow(value)
	for _, existing := range AllowedDeviceTypeRequestAirflowEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceTypeRequestAirflow", value)
}

// NewDeviceTypeRequestAirflowFromValue returns a pointer to a valid DeviceTypeRequestAirflow
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeRequestAirflowFromValue(v string) (*DeviceTypeRequestAirflow, error) {
	ev := DeviceTypeRequestAirflow(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceTypeRequestAirflow: valid values are %v", v, AllowedDeviceTypeRequestAirflowEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceTypeRequestAirflow) IsValid() bool {
	for _, existing := range AllowedDeviceTypeRequestAirflowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceTypeRequest_airflow value
func (v DeviceTypeRequestAirflow) Ptr() *DeviceTypeRequestAirflow {
	return &v
}

type NullableDeviceTypeRequestAirflow struct {
	value *DeviceTypeRequestAirflow
	isSet bool
}

func (v NullableDeviceTypeRequestAirflow) Get() *DeviceTypeRequestAirflow {
	return v.value
}

func (v *NullableDeviceTypeRequestAirflow) Set(val *DeviceTypeRequestAirflow) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeRequestAirflow) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeRequestAirflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeRequestAirflow(val *DeviceTypeRequestAirflow) *NullableDeviceTypeRequestAirflow {
	return &NullableDeviceTypeRequestAirflow{value: val, isSet: true}
}

func (v NullableDeviceTypeRequestAirflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeRequestAirflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

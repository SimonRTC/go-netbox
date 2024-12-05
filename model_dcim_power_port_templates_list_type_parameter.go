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

// DcimPowerPortTemplatesListTypeParameter the model 'DcimPowerPortTemplatesListTypeParameter'
type DcimPowerPortTemplatesListTypeParameter string

// List of dcim_power_port_templates_list_type_parameter
const (
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_CALIFORNIA_STYLE  DcimPowerPortTemplatesListTypeParameter = "California Style"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_DC                DcimPowerPortTemplatesListTypeParameter = "DC"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_IEC_60309         DcimPowerPortTemplatesListTypeParameter = "IEC 60309"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_IEC_60320         DcimPowerPortTemplatesListTypeParameter = "IEC 60320"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_IEC_60906_1       DcimPowerPortTemplatesListTypeParameter = "IEC 60906-1"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_INTERNATIONAL_ITA DcimPowerPortTemplatesListTypeParameter = "International/ITA"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_MOLEX             DcimPowerPortTemplatesListTypeParameter = "Molex"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_NEMA__LOCKING     DcimPowerPortTemplatesListTypeParameter = "NEMA (Locking)"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_NEMA__NON_LOCKING DcimPowerPortTemplatesListTypeParameter = "NEMA (Non-locking)"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_OTHER             DcimPowerPortTemplatesListTypeParameter = "Other"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_PROPRIETARY       DcimPowerPortTemplatesListTypeParameter = "Proprietary"
	DCIMPOWERPORTTEMPLATESLISTTYPEPARAMETER_USB               DcimPowerPortTemplatesListTypeParameter = "USB"
)

// All allowed values of DcimPowerPortTemplatesListTypeParameter enum
var AllowedDcimPowerPortTemplatesListTypeParameterEnumValues = []DcimPowerPortTemplatesListTypeParameter{
	"California Style",
	"DC",
	"IEC 60309",
	"IEC 60320",
	"IEC 60906-1",
	"International/ITA",
	"Molex",
	"NEMA (Locking)",
	"NEMA (Non-locking)",
	"Other",
	"Proprietary",
	"USB",
}

func (v *DcimPowerPortTemplatesListTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimPowerPortTemplatesListTypeParameter(value)
	for _, existing := range AllowedDcimPowerPortTemplatesListTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimPowerPortTemplatesListTypeParameter", value)
}

// NewDcimPowerPortTemplatesListTypeParameterFromValue returns a pointer to a valid DcimPowerPortTemplatesListTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimPowerPortTemplatesListTypeParameterFromValue(v string) (*DcimPowerPortTemplatesListTypeParameter, error) {
	ev := DcimPowerPortTemplatesListTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimPowerPortTemplatesListTypeParameter: valid values are %v", v, AllowedDcimPowerPortTemplatesListTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimPowerPortTemplatesListTypeParameter) IsValid() bool {
	for _, existing := range AllowedDcimPowerPortTemplatesListTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_power_port_templates_list_type_parameter value
func (v DcimPowerPortTemplatesListTypeParameter) Ptr() *DcimPowerPortTemplatesListTypeParameter {
	return &v
}

type NullableDcimPowerPortTemplatesListTypeParameter struct {
	value *DcimPowerPortTemplatesListTypeParameter
	isSet bool
}

func (v NullableDcimPowerPortTemplatesListTypeParameter) Get() *DcimPowerPortTemplatesListTypeParameter {
	return v.value
}

func (v *NullableDcimPowerPortTemplatesListTypeParameter) Set(val *DcimPowerPortTemplatesListTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimPowerPortTemplatesListTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimPowerPortTemplatesListTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimPowerPortTemplatesListTypeParameter(val *DcimPowerPortTemplatesListTypeParameter) *NullableDcimPowerPortTemplatesListTypeParameter {
	return &NullableDcimPowerPortTemplatesListTypeParameter{value: val, isSet: true}
}

func (v NullableDcimPowerPortTemplatesListTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimPowerPortTemplatesListTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

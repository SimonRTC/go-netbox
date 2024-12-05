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

// ExtrasCustomFieldChoiceSetsListBaseChoicesParameter the model 'ExtrasCustomFieldChoiceSetsListBaseChoicesParameter'
type ExtrasCustomFieldChoiceSetsListBaseChoicesParameter string

// List of extras_custom_field_choice_sets_list_base_choices_parameter
const (
	EXTRASCUSTOMFIELDCHOICESETSLISTBASECHOICESPARAMETER_IATA      ExtrasCustomFieldChoiceSetsListBaseChoicesParameter = "IATA"
	EXTRASCUSTOMFIELDCHOICESETSLISTBASECHOICESPARAMETER_ISO_3166  ExtrasCustomFieldChoiceSetsListBaseChoicesParameter = "ISO_3166"
	EXTRASCUSTOMFIELDCHOICESETSLISTBASECHOICESPARAMETER_UN_LOCODE ExtrasCustomFieldChoiceSetsListBaseChoicesParameter = "UN_LOCODE"
)

// All allowed values of ExtrasCustomFieldChoiceSetsListBaseChoicesParameter enum
var AllowedExtrasCustomFieldChoiceSetsListBaseChoicesParameterEnumValues = []ExtrasCustomFieldChoiceSetsListBaseChoicesParameter{
	"IATA",
	"ISO_3166",
	"UN_LOCODE",
}

func (v *ExtrasCustomFieldChoiceSetsListBaseChoicesParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtrasCustomFieldChoiceSetsListBaseChoicesParameter(value)
	for _, existing := range AllowedExtrasCustomFieldChoiceSetsListBaseChoicesParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtrasCustomFieldChoiceSetsListBaseChoicesParameter", value)
}

// NewExtrasCustomFieldChoiceSetsListBaseChoicesParameterFromValue returns a pointer to a valid ExtrasCustomFieldChoiceSetsListBaseChoicesParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtrasCustomFieldChoiceSetsListBaseChoicesParameterFromValue(v string) (*ExtrasCustomFieldChoiceSetsListBaseChoicesParameter, error) {
	ev := ExtrasCustomFieldChoiceSetsListBaseChoicesParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtrasCustomFieldChoiceSetsListBaseChoicesParameter: valid values are %v", v, AllowedExtrasCustomFieldChoiceSetsListBaseChoicesParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtrasCustomFieldChoiceSetsListBaseChoicesParameter) IsValid() bool {
	for _, existing := range AllowedExtrasCustomFieldChoiceSetsListBaseChoicesParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to extras_custom_field_choice_sets_list_base_choices_parameter value
func (v ExtrasCustomFieldChoiceSetsListBaseChoicesParameter) Ptr() *ExtrasCustomFieldChoiceSetsListBaseChoicesParameter {
	return &v
}

type NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter struct {
	value *ExtrasCustomFieldChoiceSetsListBaseChoicesParameter
	isSet bool
}

func (v NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter) Get() *ExtrasCustomFieldChoiceSetsListBaseChoicesParameter {
	return v.value
}

func (v *NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter) Set(val *ExtrasCustomFieldChoiceSetsListBaseChoicesParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter(val *ExtrasCustomFieldChoiceSetsListBaseChoicesParameter) *NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter {
	return &NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter{value: val, isSet: true}
}

func (v NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtrasCustomFieldChoiceSetsListBaseChoicesParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.3 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the InterfaceTemplatePoeType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceTemplatePoeType{}

// InterfaceTemplatePoeType struct for InterfaceTemplatePoeType
type InterfaceTemplatePoeType struct {
	Value                *InterfacePoeTypeValue `json:"value,omitempty"`
	Label                *InterfacePoeTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceTemplatePoeType InterfaceTemplatePoeType

// NewInterfaceTemplatePoeType instantiates a new InterfaceTemplatePoeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceTemplatePoeType() *InterfaceTemplatePoeType {
	this := InterfaceTemplatePoeType{}
	return &this
}

// NewInterfaceTemplatePoeTypeWithDefaults instantiates a new InterfaceTemplatePoeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTemplatePoeTypeWithDefaults() *InterfaceTemplatePoeType {
	this := InterfaceTemplatePoeType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceTemplatePoeType) GetValue() InterfacePoeTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret InterfacePoeTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplatePoeType) GetValueOk() (*InterfacePoeTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceTemplatePoeType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given InterfacePoeTypeValue and assigns it to the Value field.
func (o *InterfaceTemplatePoeType) SetValue(v InterfacePoeTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceTemplatePoeType) GetLabel() InterfacePoeTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret InterfacePoeTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplatePoeType) GetLabelOk() (*InterfacePoeTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceTemplatePoeType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given InterfacePoeTypeLabel and assigns it to the Label field.
func (o *InterfaceTemplatePoeType) SetLabel(v InterfacePoeTypeLabel) {
	o.Label = &v
}

func (o InterfaceTemplatePoeType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceTemplatePoeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceTemplatePoeType) UnmarshalJSON(data []byte) (err error) {
	varInterfaceTemplatePoeType := _InterfaceTemplatePoeType{}

	err = json.Unmarshal(data, &varInterfaceTemplatePoeType)

	if err != nil {
		return err
	}

	*o = InterfaceTemplatePoeType(varInterfaceTemplatePoeType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceTemplatePoeType struct {
	value *InterfaceTemplatePoeType
	isSet bool
}

func (v NullableInterfaceTemplatePoeType) Get() *InterfaceTemplatePoeType {
	return v.value
}

func (v *NullableInterfaceTemplatePoeType) Set(val *InterfaceTemplatePoeType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplatePoeType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplatePoeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplatePoeType(val *InterfaceTemplatePoeType) *NullableInterfaceTemplatePoeType {
	return &NullableInterfaceTemplatePoeType{value: val, isSet: true}
}

func (v NullableInterfaceTemplatePoeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplatePoeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
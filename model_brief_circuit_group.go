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

// checks if the BriefCircuitGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefCircuitGroup{}

// BriefCircuitGroup Adds support for custom fields and tags.
type BriefCircuitGroup struct {
	Id                   int32  `json:"id"`
	Url                  string `json:"url"`
	Display              string `json:"display"`
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _BriefCircuitGroup BriefCircuitGroup

// NewBriefCircuitGroup instantiates a new BriefCircuitGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefCircuitGroup(id int32, url string, display string, name string) *BriefCircuitGroup {
	this := BriefCircuitGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewBriefCircuitGroupWithDefaults instantiates a new BriefCircuitGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefCircuitGroupWithDefaults() *BriefCircuitGroup {
	this := BriefCircuitGroup{}
	return &this
}

// GetId returns the Id field value
func (o *BriefCircuitGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefCircuitGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefCircuitGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefCircuitGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefCircuitGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefCircuitGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefCircuitGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefCircuitGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefCircuitGroup) SetName(v string) {
	o.Name = v
}

func (o BriefCircuitGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefCircuitGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefCircuitGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBriefCircuitGroup := _BriefCircuitGroup{}

	err = json.Unmarshal(data, &varBriefCircuitGroup)

	if err != nil {
		return err
	}

	*o = BriefCircuitGroup(varBriefCircuitGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefCircuitGroup struct {
	value *BriefCircuitGroup
	isSet bool
}

func (v NullableBriefCircuitGroup) Get() *BriefCircuitGroup {
	return v.value
}

func (v *NullableBriefCircuitGroup) Set(val *BriefCircuitGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCircuitGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCircuitGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCircuitGroup(val *BriefCircuitGroup) *NullableBriefCircuitGroup {
	return &NullableBriefCircuitGroup{value: val, isSet: true}
}

func (v NullableBriefCircuitGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCircuitGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

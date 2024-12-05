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

// checks if the BriefVRF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVRF{}

// BriefVRF Adds support for custom fields and tags.
type BriefVRF struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Name    string `json:"name"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd                   NullableString `json:"rd,omitempty"`
	Description          *string        `json:"description,omitempty"`
	PrefixCount          *int64         `json:"prefix_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefVRF BriefVRF

// NewBriefVRF instantiates a new BriefVRF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVRF(id int32, url string, display string, name string) *BriefVRF {
	this := BriefVRF{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewBriefVRFWithDefaults instantiates a new BriefVRF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVRFWithDefaults() *BriefVRF {
	this := BriefVRF{}
	return &this
}

// GetId returns the Id field value
func (o *BriefVRF) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefVRF) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefVRF) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefVRF) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefVRF) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefVRF) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefVRF) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefVRF) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefVRF) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefVRF) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefVRF) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefVRF) SetName(v string) {
	o.Name = v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BriefVRF) GetRd() string {
	if o == nil || IsNil(o.Rd.Get()) {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BriefVRF) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *BriefVRF) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *BriefVRF) SetRd(v string) {
	o.Rd.Set(&v)
}

// SetRdNil sets the value for Rd to be an explicit nil
func (o *BriefVRF) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *BriefVRF) UnsetRd() {
	o.Rd.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVRF) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVRF) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVRF) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVRF) SetDescription(v string) {
	o.Description = &v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *BriefVRF) GetPrefixCount() int64 {
	if o == nil || IsNil(o.PrefixCount) {
		var ret int64
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVRF) GetPrefixCountOk() (*int64, bool) {
	if o == nil || IsNil(o.PrefixCount) {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *BriefVRF) HasPrefixCount() bool {
	if o != nil && !IsNil(o.PrefixCount) {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int64 and assigns it to the PrefixCount field.
func (o *BriefVRF) SetPrefixCount(v int64) {
	o.PrefixCount = &v
}

func (o BriefVRF) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVRF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if o.Rd.IsSet() {
		toSerialize["rd"] = o.Rd.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PrefixCount) {
		toSerialize["prefix_count"] = o.PrefixCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefVRF) UnmarshalJSON(data []byte) (err error) {
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

	varBriefVRF := _BriefVRF{}

	err = json.Unmarshal(data, &varBriefVRF)

	if err != nil {
		return err
	}

	*o = BriefVRF(varBriefVRF)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rd")
		delete(additionalProperties, "description")
		delete(additionalProperties, "prefix_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVRF struct {
	value *BriefVRF
	isSet bool
}

func (v NullableBriefVRF) Get() *BriefVRF {
	return v.value
}

func (v *NullableBriefVRF) Set(val *BriefVRF) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVRF) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVRF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVRF(val *BriefVRF) *NullableBriefVRF {
	return &NullableBriefVRF{value: val, isSet: true}
}

func (v NullableBriefVRF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVRF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

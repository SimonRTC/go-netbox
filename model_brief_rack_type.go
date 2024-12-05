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

// checks if the BriefRackType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefRackType{}

// BriefRackType Adds support for custom fields and tags.
type BriefRackType struct {
	Id                   int32             `json:"id"`
	Url                  string            `json:"url"`
	Display              string            `json:"display"`
	Manufacturer         BriefManufacturer `json:"manufacturer"`
	Model                string            `json:"model"`
	Slug                 string            `json:"slug"`
	Description          *string           `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefRackType BriefRackType

// NewBriefRackType instantiates a new BriefRackType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefRackType(id int32, url string, display string, manufacturer BriefManufacturer, model string, slug string) *BriefRackType {
	this := BriefRackType{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	return &this
}

// NewBriefRackTypeWithDefaults instantiates a new BriefRackType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefRackTypeWithDefaults() *BriefRackType {
	this := BriefRackType{}
	return &this
}

// GetId returns the Id field value
func (o *BriefRackType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefRackType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefRackType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefRackType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefRackType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefRackType) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefRackType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefRackType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefRackType) SetDisplay(v string) {
	o.Display = v
}

// GetManufacturer returns the Manufacturer field value
func (o *BriefRackType) GetManufacturer() BriefManufacturer {
	if o == nil {
		var ret BriefManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *BriefRackType) GetManufacturerOk() (*BriefManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *BriefRackType) SetManufacturer(v BriefManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *BriefRackType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BriefRackType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BriefRackType) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *BriefRackType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefRackType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefRackType) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefRackType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefRackType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefRackType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefRackType) SetDescription(v string) {
	o.Description = &v
}

func (o BriefRackType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefRackType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefRackType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"manufacturer",
		"model",
		"slug",
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

	varBriefRackType := _BriefRackType{}

	err = json.Unmarshal(data, &varBriefRackType)

	if err != nil {
		return err
	}

	*o = BriefRackType(varBriefRackType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefRackType struct {
	value *BriefRackType
	isSet bool
}

func (v NullableBriefRackType) Get() *BriefRackType {
	return v.value
}

func (v *NullableBriefRackType) Set(val *BriefRackType) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefRackType) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefRackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefRackType(val *BriefRackType) *NullableBriefRackType {
	return &NullableBriefRackType{value: val, isSet: true}
}

func (v NullableBriefRackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefRackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

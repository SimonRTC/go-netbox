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

// checks if the BriefProviderAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefProviderAccount{}

// BriefProviderAccount Adds support for custom fields and tags.
type BriefProviderAccount struct {
	Id                   int32   `json:"id"`
	Url                  string  `json:"url"`
	Display              string  `json:"display"`
	Name                 *string `json:"name,omitempty"`
	Account              string  `json:"account"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefProviderAccount BriefProviderAccount

// NewBriefProviderAccount instantiates a new BriefProviderAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefProviderAccount(id int32, url string, display string, account string) *BriefProviderAccount {
	this := BriefProviderAccount{}
	this.Id = id
	this.Url = url
	this.Display = display
	var name string = ""
	this.Name = &name
	this.Account = account
	return &this
}

// NewBriefProviderAccountWithDefaults instantiates a new BriefProviderAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefProviderAccountWithDefaults() *BriefProviderAccount {
	this := BriefProviderAccount{}
	var name string = ""
	this.Name = &name
	return &this
}

// GetId returns the Id field value
func (o *BriefProviderAccount) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefProviderAccount) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefProviderAccount) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefProviderAccount) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefProviderAccount) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefProviderAccount) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefProviderAccount) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefProviderAccount) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefProviderAccount) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BriefProviderAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefProviderAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BriefProviderAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BriefProviderAccount) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value
func (o *BriefProviderAccount) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *BriefProviderAccount) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *BriefProviderAccount) SetAccount(v string) {
	o.Account = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefProviderAccount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefProviderAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefProviderAccount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefProviderAccount) SetDescription(v string) {
	o.Description = &v
}

func (o BriefProviderAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefProviderAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["account"] = o.Account
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefProviderAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"account",
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

	varBriefProviderAccount := _BriefProviderAccount{}

	err = json.Unmarshal(data, &varBriefProviderAccount)

	if err != nil {
		return err
	}

	*o = BriefProviderAccount(varBriefProviderAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "account")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefProviderAccount struct {
	value *BriefProviderAccount
	isSet bool
}

func (v NullableBriefProviderAccount) Get() *BriefProviderAccount {
	return v.value
}

func (v *NullableBriefProviderAccount) Set(val *BriefProviderAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefProviderAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefProviderAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefProviderAccount(val *BriefProviderAccount) *NullableBriefProviderAccount {
	return &NullableBriefProviderAccount{value: val, isSet: true}
}

func (v NullableBriefProviderAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefProviderAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

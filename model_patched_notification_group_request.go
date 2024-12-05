/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.7 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedNotificationGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedNotificationGroupRequest{}

// PatchedNotificationGroupRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedNotificationGroupRequest struct {
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	Groups               []int32 `json:"groups,omitempty"`
	Users                []int32 `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedNotificationGroupRequest PatchedNotificationGroupRequest

// NewPatchedNotificationGroupRequest instantiates a new PatchedNotificationGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNotificationGroupRequest() *PatchedNotificationGroupRequest {
	this := PatchedNotificationGroupRequest{}
	return &this
}

// NewPatchedNotificationGroupRequestWithDefaults instantiates a new PatchedNotificationGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNotificationGroupRequestWithDefaults() *PatchedNotificationGroupRequest {
	this := PatchedNotificationGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedNotificationGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedNotificationGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedNotificationGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedNotificationGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedNotificationGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedNotificationGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PatchedNotificationGroupRequest) GetGroups() []int32 {
	if o == nil || IsNil(o.Groups) {
		var ret []int32
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationGroupRequest) GetGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PatchedNotificationGroupRequest) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *PatchedNotificationGroupRequest) SetGroups(v []int32) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PatchedNotificationGroupRequest) GetUsers() []int32 {
	if o == nil || IsNil(o.Users) {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationGroupRequest) GetUsersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PatchedNotificationGroupRequest) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *PatchedNotificationGroupRequest) SetUsers(v []int32) {
	o.Users = v
}

func (o PatchedNotificationGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedNotificationGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedNotificationGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedNotificationGroupRequest := _PatchedNotificationGroupRequest{}

	err = json.Unmarshal(data, &varPatchedNotificationGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedNotificationGroupRequest(varPatchedNotificationGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedNotificationGroupRequest struct {
	value *PatchedNotificationGroupRequest
	isSet bool
}

func (v NullablePatchedNotificationGroupRequest) Get() *PatchedNotificationGroupRequest {
	return v.value
}

func (v *NullablePatchedNotificationGroupRequest) Set(val *PatchedNotificationGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNotificationGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNotificationGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNotificationGroupRequest(val *PatchedNotificationGroupRequest) *NullablePatchedNotificationGroupRequest {
	return &NullablePatchedNotificationGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedNotificationGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNotificationGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the GenericObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericObject{}

// GenericObject Minimal representation of some generic object identified by ContentType and PK.
type GenericObject struct {
	ObjectType           string      `json:"object_type"`
	ObjectId             int32       `json:"object_id"`
	Object               interface{} `json:"object"`
	AdditionalProperties map[string]interface{}
}

type _GenericObject GenericObject

// NewGenericObject instantiates a new GenericObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericObject(objectType string, objectId int32, object interface{}) *GenericObject {
	this := GenericObject{}
	this.ObjectType = objectType
	this.ObjectId = objectId
	this.Object = object
	return &this
}

// NewGenericObjectWithDefaults instantiates a new GenericObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericObjectWithDefaults() *GenericObject {
	this := GenericObject{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *GenericObject) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *GenericObject) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *GenericObject) SetObjectType(v string) {
	o.ObjectType = v
}

// GetObjectId returns the ObjectId field value
func (o *GenericObject) GetObjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *GenericObject) GetObjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *GenericObject) SetObjectId(v int32) {
	o.ObjectId = v
}

// GetObject returns the Object field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GenericObject) GetObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenericObject) GetObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GenericObject) SetObject(v interface{}) {
	o.Object = v
}

func (o GenericObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["object_id"] = o.ObjectId
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenericObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_type",
		"object_id",
		"object",
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

	varGenericObject := _GenericObject{}

	err = json.Unmarshal(data, &varGenericObject)

	if err != nil {
		return err
	}

	*o = GenericObject(varGenericObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericObject struct {
	value *GenericObject
	isSet bool
}

func (v NullableGenericObject) Get() *GenericObject {
	return v.value
}

func (v *NullableGenericObject) Set(val *GenericObject) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericObject) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericObject(val *GenericObject) *NullableGenericObject {
	return &NullableGenericObject{value: val, isSet: true}
}

func (v NullableGenericObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

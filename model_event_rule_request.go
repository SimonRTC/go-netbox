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

// checks if the EventRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventRuleRequest{}

// EventRuleRequest Adds support for custom fields and tags.
type EventRuleRequest struct {
	ObjectTypes []string `json:"object_types"`
	Name        string   `json:"name"`
	Enabled     *bool    `json:"enabled,omitempty"`
	// The types of event which will trigger this rule.
	EventTypes []EventRuleEventTypesInner `json:"event_types"`
	// A set of conditions which determine whether the event will be generated.
	Conditions           interface{}              `json:"conditions,omitempty"`
	ActionType           EventRuleActionTypeValue `json:"action_type"`
	ActionObjectType     string                   `json:"action_object_type"`
	ActionObjectId       NullableInt64            `json:"action_object_id,omitempty"`
	Description          *string                  `json:"description,omitempty"`
	CustomFields         map[string]interface{}   `json:"custom_fields,omitempty"`
	Tags                 []NestedTagRequest       `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventRuleRequest EventRuleRequest

// NewEventRuleRequest instantiates a new EventRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventRuleRequest(objectTypes []string, name string, eventTypes []EventRuleEventTypesInner, actionType EventRuleActionTypeValue, actionObjectType string) *EventRuleRequest {
	this := EventRuleRequest{}
	this.ObjectTypes = objectTypes
	this.Name = name
	this.EventTypes = eventTypes
	this.ActionType = actionType
	this.ActionObjectType = actionObjectType
	return &this
}

// NewEventRuleRequestWithDefaults instantiates a new EventRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventRuleRequestWithDefaults() *EventRuleRequest {
	this := EventRuleRequest{}
	return &this
}

// GetObjectTypes returns the ObjectTypes field value
func (o *EventRuleRequest) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *EventRuleRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value
func (o *EventRuleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventRuleRequest) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EventRuleRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EventRuleRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EventRuleRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventTypes returns the EventTypes field value
func (o *EventRuleRequest) GetEventTypes() []EventRuleEventTypesInner {
	if o == nil {
		var ret []EventRuleEventTypesInner
		return ret
	}

	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetEventTypesOk() ([]EventRuleEventTypesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventTypes, true
}

// SetEventTypes sets field value
func (o *EventRuleRequest) SetEventTypes(v []EventRuleEventTypesInner) {
	o.EventTypes = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventRuleRequest) GetConditions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventRuleRequest) GetConditionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *EventRuleRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given interface{} and assigns it to the Conditions field.
func (o *EventRuleRequest) SetConditions(v interface{}) {
	o.Conditions = v
}

// GetActionType returns the ActionType field value
func (o *EventRuleRequest) GetActionType() EventRuleActionTypeValue {
	if o == nil {
		var ret EventRuleActionTypeValue
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetActionTypeOk() (*EventRuleActionTypeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *EventRuleRequest) SetActionType(v EventRuleActionTypeValue) {
	o.ActionType = v
}

// GetActionObjectType returns the ActionObjectType field value
func (o *EventRuleRequest) GetActionObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionObjectType
}

// GetActionObjectTypeOk returns a tuple with the ActionObjectType field value
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetActionObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionObjectType, true
}

// SetActionObjectType sets field value
func (o *EventRuleRequest) SetActionObjectType(v string) {
	o.ActionObjectType = v
}

// GetActionObjectId returns the ActionObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventRuleRequest) GetActionObjectId() int64 {
	if o == nil || IsNil(o.ActionObjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ActionObjectId.Get()
}

// GetActionObjectIdOk returns a tuple with the ActionObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventRuleRequest) GetActionObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionObjectId.Get(), o.ActionObjectId.IsSet()
}

// HasActionObjectId returns a boolean if a field has been set.
func (o *EventRuleRequest) HasActionObjectId() bool {
	if o != nil && o.ActionObjectId.IsSet() {
		return true
	}

	return false
}

// SetActionObjectId gets a reference to the given NullableInt64 and assigns it to the ActionObjectId field.
func (o *EventRuleRequest) SetActionObjectId(v int64) {
	o.ActionObjectId.Set(&v)
}

// SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil
func (o *EventRuleRequest) SetActionObjectIdNil() {
	o.ActionObjectId.Set(nil)
}

// UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
func (o *EventRuleRequest) UnsetActionObjectId() {
	o.ActionObjectId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventRuleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventRuleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *EventRuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *EventRuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *EventRuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventRuleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRuleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventRuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *EventRuleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o EventRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_types"] = o.ObjectTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["event_types"] = o.EventTypes
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	toSerialize["action_type"] = o.ActionType
	toSerialize["action_object_type"] = o.ActionObjectType
	if o.ActionObjectId.IsSet() {
		toSerialize["action_object_id"] = o.ActionObjectId.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventRuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_types",
		"name",
		"event_types",
		"action_type",
		"action_object_type",
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

	varEventRuleRequest := _EventRuleRequest{}

	err = json.Unmarshal(data, &varEventRuleRequest)

	if err != nil {
		return err
	}

	*o = EventRuleRequest(varEventRuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "event_types")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "action_type")
		delete(additionalProperties, "action_object_type")
		delete(additionalProperties, "action_object_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventRuleRequest struct {
	value *EventRuleRequest
	isSet bool
}

func (v NullableEventRuleRequest) Get() *EventRuleRequest {
	return v.value
}

func (v *NullableEventRuleRequest) Set(val *EventRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRuleRequest(val *EventRuleRequest) *NullableEventRuleRequest {
	return &NullableEventRuleRequest{value: val, isSet: true}
}

func (v NullableEventRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

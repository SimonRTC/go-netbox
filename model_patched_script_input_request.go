/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.7 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the PatchedScriptInputRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedScriptInputRequest{}

// PatchedScriptInputRequest struct for PatchedScriptInputRequest
type PatchedScriptInputRequest struct {
	Data                 interface{}   `json:"data,omitempty"`
	Commit               *bool         `json:"commit,omitempty"`
	ScheduleAt           NullableTime  `json:"schedule_at,omitempty"`
	Interval             NullableInt32 `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedScriptInputRequest PatchedScriptInputRequest

// NewPatchedScriptInputRequest instantiates a new PatchedScriptInputRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedScriptInputRequest() *PatchedScriptInputRequest {
	this := PatchedScriptInputRequest{}
	return &this
}

// NewPatchedScriptInputRequestWithDefaults instantiates a new PatchedScriptInputRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedScriptInputRequestWithDefaults() *PatchedScriptInputRequest {
	this := PatchedScriptInputRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedScriptInputRequest) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedScriptInputRequest) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PatchedScriptInputRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *PatchedScriptInputRequest) SetData(v interface{}) {
	o.Data = v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *PatchedScriptInputRequest) GetCommit() bool {
	if o == nil || IsNil(o.Commit) {
		var ret bool
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedScriptInputRequest) GetCommitOk() (*bool, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *PatchedScriptInputRequest) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given bool and assigns it to the Commit field.
func (o *PatchedScriptInputRequest) SetCommit(v bool) {
	o.Commit = &v
}

// GetScheduleAt returns the ScheduleAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedScriptInputRequest) GetScheduleAt() time.Time {
	if o == nil || IsNil(o.ScheduleAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleAt.Get()
}

// GetScheduleAtOk returns a tuple with the ScheduleAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedScriptInputRequest) GetScheduleAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleAt.Get(), o.ScheduleAt.IsSet()
}

// HasScheduleAt returns a boolean if a field has been set.
func (o *PatchedScriptInputRequest) HasScheduleAt() bool {
	if o != nil && o.ScheduleAt.IsSet() {
		return true
	}

	return false
}

// SetScheduleAt gets a reference to the given NullableTime and assigns it to the ScheduleAt field.
func (o *PatchedScriptInputRequest) SetScheduleAt(v time.Time) {
	o.ScheduleAt.Set(&v)
}

// SetScheduleAtNil sets the value for ScheduleAt to be an explicit nil
func (o *PatchedScriptInputRequest) SetScheduleAtNil() {
	o.ScheduleAt.Set(nil)
}

// UnsetScheduleAt ensures that no value is present for ScheduleAt, not even an explicit nil
func (o *PatchedScriptInputRequest) UnsetScheduleAt() {
	o.ScheduleAt.Unset()
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedScriptInputRequest) GetInterval() int32 {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret int32
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedScriptInputRequest) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *PatchedScriptInputRequest) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableInt32 and assigns it to the Interval field.
func (o *PatchedScriptInputRequest) SetInterval(v int32) {
	o.Interval.Set(&v)
}

// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *PatchedScriptInputRequest) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *PatchedScriptInputRequest) UnsetInterval() {
	o.Interval.Unset()
}

func (o PatchedScriptInputRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedScriptInputRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if o.ScheduleAt.IsSet() {
		toSerialize["schedule_at"] = o.ScheduleAt.Get()
	}
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedScriptInputRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedScriptInputRequest := _PatchedScriptInputRequest{}

	err = json.Unmarshal(data, &varPatchedScriptInputRequest)

	if err != nil {
		return err
	}

	*o = PatchedScriptInputRequest(varPatchedScriptInputRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "commit")
		delete(additionalProperties, "schedule_at")
		delete(additionalProperties, "interval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedScriptInputRequest struct {
	value *PatchedScriptInputRequest
	isSet bool
}

func (v NullablePatchedScriptInputRequest) Get() *PatchedScriptInputRequest {
	return v.value
}

func (v *NullablePatchedScriptInputRequest) Set(val *PatchedScriptInputRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedScriptInputRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedScriptInputRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedScriptInputRequest(val *PatchedScriptInputRequest) *NullablePatchedScriptInputRequest {
	return &NullablePatchedScriptInputRequest{value: val, isSet: true}
}

func (v NullablePatchedScriptInputRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedScriptInputRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
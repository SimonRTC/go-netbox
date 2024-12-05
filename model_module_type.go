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
	"time"
)

// checks if the ModuleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleType{}

// ModuleType Adds support for custom fields and tags.
type ModuleType struct {
	Id           int32             `json:"id"`
	Url          string            `json:"url"`
	DisplayUrl   string            `json:"display_url"`
	Display      string            `json:"display"`
	Manufacturer BriefManufacturer `json:"manufacturer"`
	Model        string            `json:"model"`
	// Discrete part number (optional)
	PartNumber           *string                      `json:"part_number,omitempty"`
	Airflow              NullableModuleTypeAirflow    `json:"airflow,omitempty"`
	Weight               NullableFloat64              `json:"weight,omitempty"`
	WeightUnit           NullableDeviceTypeWeightUnit `json:"weight_unit,omitempty"`
	Description          *string                      `json:"description,omitempty"`
	Comments             *string                      `json:"comments,omitempty"`
	Tags                 []NestedTag                  `json:"tags,omitempty"`
	CustomFields         map[string]interface{}       `json:"custom_fields,omitempty"`
	Created              NullableTime                 `json:"created"`
	LastUpdated          NullableTime                 `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _ModuleType ModuleType

// NewModuleType instantiates a new ModuleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleType(id int32, url string, displayUrl string, display string, manufacturer BriefManufacturer, model string, created NullableTime, lastUpdated NullableTime) *ModuleType {
	this := ModuleType{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Manufacturer = manufacturer
	this.Model = model
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewModuleTypeWithDefaults instantiates a new ModuleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleTypeWithDefaults() *ModuleType {
	this := ModuleType{}
	return &this
}

// GetId returns the Id field value
func (o *ModuleType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModuleType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModuleType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ModuleType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ModuleType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ModuleType) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *ModuleType) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *ModuleType) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *ModuleType) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *ModuleType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ModuleType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ModuleType) SetDisplay(v string) {
	o.Display = v
}

// GetManufacturer returns the Manufacturer field value
func (o *ModuleType) GetManufacturer() BriefManufacturer {
	if o == nil {
		var ret BriefManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *ModuleType) GetManufacturerOk() (*BriefManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *ModuleType) SetManufacturer(v BriefManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *ModuleType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ModuleType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ModuleType) SetModel(v string) {
	o.Model = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *ModuleType) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleType) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *ModuleType) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *ModuleType) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetAirflow returns the Airflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleType) GetAirflow() ModuleTypeAirflow {
	if o == nil || IsNil(o.Airflow.Get()) {
		var ret ModuleTypeAirflow
		return ret
	}
	return *o.Airflow.Get()
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleType) GetAirflowOk() (*ModuleTypeAirflow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Airflow.Get(), o.Airflow.IsSet()
}

// HasAirflow returns a boolean if a field has been set.
func (o *ModuleType) HasAirflow() bool {
	if o != nil && o.Airflow.IsSet() {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given NullableModuleTypeAirflow and assigns it to the Airflow field.
func (o *ModuleType) SetAirflow(v ModuleTypeAirflow) {
	o.Airflow.Set(&v)
}

// SetAirflowNil sets the value for Airflow to be an explicit nil
func (o *ModuleType) SetAirflowNil() {
	o.Airflow.Set(nil)
}

// UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
func (o *ModuleType) UnsetAirflow() {
	o.Airflow.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleType) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleType) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *ModuleType) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *ModuleType) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *ModuleType) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *ModuleType) UnsetWeight() {
	o.Weight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleType) GetWeightUnit() DeviceTypeWeightUnit {
	if o == nil || IsNil(o.WeightUnit.Get()) {
		var ret DeviceTypeWeightUnit
		return ret
	}
	return *o.WeightUnit.Get()
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleType) GetWeightUnitOk() (*DeviceTypeWeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightUnit.Get(), o.WeightUnit.IsSet()
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *ModuleType) HasWeightUnit() bool {
	if o != nil && o.WeightUnit.IsSet() {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given NullableDeviceTypeWeightUnit and assigns it to the WeightUnit field.
func (o *ModuleType) SetWeightUnit(v DeviceTypeWeightUnit) {
	o.WeightUnit.Set(&v)
}

// SetWeightUnitNil sets the value for WeightUnit to be an explicit nil
func (o *ModuleType) SetWeightUnitNil() {
	o.WeightUnit.Set(nil)
}

// UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
func (o *ModuleType) UnsetWeightUnit() {
	o.WeightUnit.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModuleType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModuleType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModuleType) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ModuleType) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleType) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ModuleType) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ModuleType) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModuleType) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleType) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModuleType) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ModuleType) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ModuleType) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleType) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ModuleType) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ModuleType) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ModuleType) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleType) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ModuleType) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ModuleType) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleType) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ModuleType) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ModuleType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if o.Airflow.IsSet() {
		toSerialize["airflow"] = o.Airflow.Get()
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.WeightUnit.IsSet() {
		toSerialize["weight_unit"] = o.WeightUnit.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"manufacturer",
		"model",
		"created",
		"last_updated",
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

	varModuleType := _ModuleType{}

	err = json.Unmarshal(data, &varModuleType)

	if err != nil {
		return err
	}

	*o = ModuleType(varModuleType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "part_number")
		delete(additionalProperties, "airflow")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleType struct {
	value *ModuleType
	isSet bool
}

func (v NullableModuleType) Get() *ModuleType {
	return v.value
}

func (v *NullableModuleType) Set(val *ModuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleType(val *ModuleType) *NullableModuleType {
	return &NullableModuleType{value: val, isSet: true}
}

func (v NullableModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

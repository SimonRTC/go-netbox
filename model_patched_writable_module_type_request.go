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

// checks if the PatchedWritableModuleTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableModuleTypeRequest{}

// PatchedWritableModuleTypeRequest Adds support for custom fields and tags.
type PatchedWritableModuleTypeRequest struct {
	Manufacturer *BriefManufacturerRequest `json:"manufacturer,omitempty"`
	Model        *string                   `json:"model,omitempty"`
	// Discrete part number (optional)
	PartNumber           *string                    `json:"part_number,omitempty"`
	Airflow              *ModuleTypeAirflowValue    `json:"airflow,omitempty"`
	Weight               NullableFloat64            `json:"weight,omitempty"`
	WeightUnit           *DeviceTypeWeightUnitValue `json:"weight_unit,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	Comments             *string                    `json:"comments,omitempty"`
	Tags                 []NestedTagRequest         `json:"tags,omitempty"`
	CustomFields         map[string]interface{}     `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableModuleTypeRequest PatchedWritableModuleTypeRequest

// NewPatchedWritableModuleTypeRequest instantiates a new PatchedWritableModuleTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableModuleTypeRequest() *PatchedWritableModuleTypeRequest {
	this := PatchedWritableModuleTypeRequest{}
	return &this
}

// NewPatchedWritableModuleTypeRequestWithDefaults instantiates a new PatchedWritableModuleTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableModuleTypeRequestWithDefaults() *PatchedWritableModuleTypeRequest {
	this := PatchedWritableModuleTypeRequest{}
	return &this
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil || IsNil(o.Manufacturer) {
		var ret BriefManufacturerRequest
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given BriefManufacturerRequest and assigns it to the Manufacturer field.
func (o *PatchedWritableModuleTypeRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *PatchedWritableModuleTypeRequest) SetModel(v string) {
	o.Model = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *PatchedWritableModuleTypeRequest) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetAirflow returns the Airflow field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetAirflow() ModuleTypeAirflowValue {
	if o == nil || IsNil(o.Airflow) {
		var ret ModuleTypeAirflowValue
		return ret
	}
	return *o.Airflow
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetAirflowOk() (*ModuleTypeAirflowValue, bool) {
	if o == nil || IsNil(o.Airflow) {
		return nil, false
	}
	return o.Airflow, true
}

// HasAirflow returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasAirflow() bool {
	if o != nil && !IsNil(o.Airflow) {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given ModuleTypeAirflowValue and assigns it to the Airflow field.
func (o *PatchedWritableModuleTypeRequest) SetAirflow(v ModuleTypeAirflowValue) {
	o.Airflow = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableModuleTypeRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableModuleTypeRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *PatchedWritableModuleTypeRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *PatchedWritableModuleTypeRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *PatchedWritableModuleTypeRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetWeightUnit() DeviceTypeWeightUnitValue {
	if o == nil || IsNil(o.WeightUnit) {
		var ret DeviceTypeWeightUnitValue
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetWeightUnitOk() (*DeviceTypeWeightUnitValue, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given DeviceTypeWeightUnitValue and assigns it to the WeightUnit field.
func (o *PatchedWritableModuleTypeRequest) SetWeightUnit(v DeviceTypeWeightUnitValue) {
	o.WeightUnit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableModuleTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableModuleTypeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableModuleTypeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableModuleTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableModuleTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableModuleTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableModuleTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableModuleTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableModuleTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if !IsNil(o.Airflow) {
		toSerialize["airflow"] = o.Airflow
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if !IsNil(o.WeightUnit) {
		toSerialize["weight_unit"] = o.WeightUnit
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableModuleTypeRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableModuleTypeRequest := _PatchedWritableModuleTypeRequest{}

	err = json.Unmarshal(data, &varPatchedWritableModuleTypeRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableModuleTypeRequest(varPatchedWritableModuleTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableModuleTypeRequest struct {
	value *PatchedWritableModuleTypeRequest
	isSet bool
}

func (v NullablePatchedWritableModuleTypeRequest) Get() *PatchedWritableModuleTypeRequest {
	return v.value
}

func (v *NullablePatchedWritableModuleTypeRequest) Set(val *PatchedWritableModuleTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableModuleTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableModuleTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableModuleTypeRequest(val *PatchedWritableModuleTypeRequest) *NullablePatchedWritableModuleTypeRequest {
	return &NullablePatchedWritableModuleTypeRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableModuleTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableModuleTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

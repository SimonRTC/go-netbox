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

// checks if the IPRangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPRangeRequest{}

// IPRangeRequest Adds support for custom fields and tags.
type IPRangeRequest struct {
	StartAddress string                     `json:"start_address"`
	EndAddress   string                     `json:"end_address"`
	Vrf          NullableBriefVRFRequest    `json:"vrf,omitempty"`
	Tenant       NullableBriefTenantRequest `json:"tenant,omitempty"`
	Status       *IPRangeStatusValue        `json:"status,omitempty"`
	Role         NullableBriefRoleRequest   `json:"role,omitempty"`
	Description  *string                    `json:"description,omitempty"`
	Comments     *string                    `json:"comments,omitempty"`
	Tags         []NestedTagRequest         `json:"tags,omitempty"`
	CustomFields map[string]interface{}     `json:"custom_fields,omitempty"`
	// Treat as fully utilized
	MarkUtilized         *bool `json:"mark_utilized,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPRangeRequest IPRangeRequest

// NewIPRangeRequest instantiates a new IPRangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPRangeRequest(startAddress string, endAddress string) *IPRangeRequest {
	this := IPRangeRequest{}
	this.StartAddress = startAddress
	this.EndAddress = endAddress
	return &this
}

// NewIPRangeRequestWithDefaults instantiates a new IPRangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPRangeRequestWithDefaults() *IPRangeRequest {
	this := IPRangeRequest{}
	return &this
}

// GetStartAddress returns the StartAddress field value
func (o *IPRangeRequest) GetStartAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartAddress
}

// GetStartAddressOk returns a tuple with the StartAddress field value
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetStartAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAddress, true
}

// SetStartAddress sets field value
func (o *IPRangeRequest) SetStartAddress(v string) {
	o.StartAddress = v
}

// GetEndAddress returns the EndAddress field value
func (o *IPRangeRequest) GetEndAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetEndAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndAddress, true
}

// SetEndAddress sets field value
func (o *IPRangeRequest) SetEndAddress(v string) {
	o.EndAddress = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPRangeRequest) GetVrf() BriefVRFRequest {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret BriefVRFRequest
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPRangeRequest) GetVrfOk() (*BriefVRFRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *IPRangeRequest) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableBriefVRFRequest and assigns it to the Vrf field.
func (o *IPRangeRequest) SetVrf(v BriefVRFRequest) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *IPRangeRequest) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *IPRangeRequest) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPRangeRequest) GetTenant() BriefTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPRangeRequest) GetTenantOk() (*BriefTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *IPRangeRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenantRequest and assigns it to the Tenant field.
func (o *IPRangeRequest) SetTenant(v BriefTenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *IPRangeRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *IPRangeRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IPRangeRequest) GetStatus() IPRangeStatusValue {
	if o == nil || IsNil(o.Status) {
		var ret IPRangeStatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetStatusOk() (*IPRangeStatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IPRangeRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IPRangeStatusValue and assigns it to the Status field.
func (o *IPRangeRequest) SetStatus(v IPRangeStatusValue) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPRangeRequest) GetRole() BriefRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPRangeRequest) GetRoleOk() (*BriefRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *IPRangeRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefRoleRequest and assigns it to the Role field.
func (o *IPRangeRequest) SetRole(v BriefRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *IPRangeRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *IPRangeRequest) UnsetRole() {
	o.Role.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPRangeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPRangeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPRangeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPRangeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPRangeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPRangeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPRangeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPRangeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *IPRangeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPRangeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPRangeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IPRangeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetMarkUtilized returns the MarkUtilized field value if set, zero value otherwise.
func (o *IPRangeRequest) GetMarkUtilized() bool {
	if o == nil || IsNil(o.MarkUtilized) {
		var ret bool
		return ret
	}
	return *o.MarkUtilized
}

// GetMarkUtilizedOk returns a tuple with the MarkUtilized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeRequest) GetMarkUtilizedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkUtilized) {
		return nil, false
	}
	return o.MarkUtilized, true
}

// HasMarkUtilized returns a boolean if a field has been set.
func (o *IPRangeRequest) HasMarkUtilized() bool {
	if o != nil && !IsNil(o.MarkUtilized) {
		return true
	}

	return false
}

// SetMarkUtilized gets a reference to the given bool and assigns it to the MarkUtilized field.
func (o *IPRangeRequest) SetMarkUtilized(v bool) {
	o.MarkUtilized = &v
}

func (o IPRangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPRangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_address"] = o.StartAddress
	toSerialize["end_address"] = o.EndAddress
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
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
	if !IsNil(o.MarkUtilized) {
		toSerialize["mark_utilized"] = o.MarkUtilized
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPRangeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start_address",
		"end_address",
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

	varIPRangeRequest := _IPRangeRequest{}

	err = json.Unmarshal(data, &varIPRangeRequest)

	if err != nil {
		return err
	}

	*o = IPRangeRequest(varIPRangeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "start_address")
		delete(additionalProperties, "end_address")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "mark_utilized")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPRangeRequest struct {
	value *IPRangeRequest
	isSet bool
}

func (v NullableIPRangeRequest) Get() *IPRangeRequest {
	return v.value
}

func (v *NullableIPRangeRequest) Set(val *IPRangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIPRangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIPRangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPRangeRequest(val *IPRangeRequest) *NullableIPRangeRequest {
	return &NullableIPRangeRequest{value: val, isSet: true}
}

func (v NullableIPRangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPRangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

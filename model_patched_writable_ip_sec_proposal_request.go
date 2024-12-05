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

// checks if the PatchedWritableIPSecProposalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableIPSecProposalRequest{}

// PatchedWritableIPSecProposalRequest Adds support for custom fields and tags.
type PatchedWritableIPSecProposalRequest struct {
	Name                    *string         `json:"name,omitempty"`
	Description             *string         `json:"description,omitempty"`
	EncryptionAlgorithm     *Encryption     `json:"encryption_algorithm,omitempty"`
	AuthenticationAlgorithm *Authentication `json:"authentication_algorithm,omitempty"`
	// Security association lifetime (seconds)
	SaLifetimeSeconds NullableInt32 `json:"sa_lifetime_seconds,omitempty"`
	// Security association lifetime (in kilobytes)
	SaLifetimeData       NullableInt32          `json:"sa_lifetime_data,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableIPSecProposalRequest PatchedWritableIPSecProposalRequest

// NewPatchedWritableIPSecProposalRequest instantiates a new PatchedWritableIPSecProposalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableIPSecProposalRequest() *PatchedWritableIPSecProposalRequest {
	this := PatchedWritableIPSecProposalRequest{}
	return &this
}

// NewPatchedWritableIPSecProposalRequestWithDefaults instantiates a new PatchedWritableIPSecProposalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableIPSecProposalRequestWithDefaults() *PatchedWritableIPSecProposalRequest {
	this := PatchedWritableIPSecProposalRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProposalRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProposalRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableIPSecProposalRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProposalRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProposalRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableIPSecProposalRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProposalRequest) GetEncryptionAlgorithm() Encryption {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		var ret Encryption
		return ret
	}
	return *o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProposalRequest) GetEncryptionAlgorithmOk() (*Encryption, bool) {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		return nil, false
	}
	return o.EncryptionAlgorithm, true
}

// HasEncryptionAlgorithm returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.EncryptionAlgorithm) {
		return true
	}

	return false
}

// SetEncryptionAlgorithm gets a reference to the given Encryption and assigns it to the EncryptionAlgorithm field.
func (o *PatchedWritableIPSecProposalRequest) SetEncryptionAlgorithm(v Encryption) {
	o.EncryptionAlgorithm = &v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProposalRequest) GetAuthenticationAlgorithm() Authentication {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		var ret Authentication
		return ret
	}
	return *o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProposalRequest) GetAuthenticationAlgorithmOk() (*Authentication, bool) {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		return nil, false
	}
	return o.AuthenticationAlgorithm, true
}

// HasAuthenticationAlgorithm returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasAuthenticationAlgorithm() bool {
	if o != nil && !IsNil(o.AuthenticationAlgorithm) {
		return true
	}

	return false
}

// SetAuthenticationAlgorithm gets a reference to the given Authentication and assigns it to the AuthenticationAlgorithm field.
func (o *PatchedWritableIPSecProposalRequest) SetAuthenticationAlgorithm(v Authentication) {
	o.AuthenticationAlgorithm = &v
}

// GetSaLifetimeSeconds returns the SaLifetimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeSeconds() int32 {
	if o == nil || IsNil(o.SaLifetimeSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetimeSeconds.Get()
}

// GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeSeconds.Get(), o.SaLifetimeSeconds.IsSet()
}

// HasSaLifetimeSeconds returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasSaLifetimeSeconds() bool {
	if o != nil && o.SaLifetimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeSeconds gets a reference to the given NullableInt32 and assigns it to the SaLifetimeSeconds field.
func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeSeconds(v int32) {
	o.SaLifetimeSeconds.Set(&v)
}

// SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil
func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeSecondsNil() {
	o.SaLifetimeSeconds.Set(nil)
}

// UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
func (o *PatchedWritableIPSecProposalRequest) UnsetSaLifetimeSeconds() {
	o.SaLifetimeSeconds.Unset()
}

// GetSaLifetimeData returns the SaLifetimeData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeData() int32 {
	if o == nil || IsNil(o.SaLifetimeData.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetimeData.Get()
}

// GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeDataOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetimeData.Get(), o.SaLifetimeData.IsSet()
}

// HasSaLifetimeData returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasSaLifetimeData() bool {
	if o != nil && o.SaLifetimeData.IsSet() {
		return true
	}

	return false
}

// SetSaLifetimeData gets a reference to the given NullableInt32 and assigns it to the SaLifetimeData field.
func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeData(v int32) {
	o.SaLifetimeData.Set(&v)
}

// SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil
func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeDataNil() {
	o.SaLifetimeData.Set(nil)
}

// UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
func (o *PatchedWritableIPSecProposalRequest) UnsetSaLifetimeData() {
	o.SaLifetimeData.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProposalRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProposalRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableIPSecProposalRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProposalRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProposalRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableIPSecProposalRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProposalRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProposalRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProposalRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableIPSecProposalRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableIPSecProposalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableIPSecProposalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EncryptionAlgorithm) {
		toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	}
	if !IsNil(o.AuthenticationAlgorithm) {
		toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	}
	if o.SaLifetimeSeconds.IsSet() {
		toSerialize["sa_lifetime_seconds"] = o.SaLifetimeSeconds.Get()
	}
	if o.SaLifetimeData.IsSet() {
		toSerialize["sa_lifetime_data"] = o.SaLifetimeData.Get()
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

func (o *PatchedWritableIPSecProposalRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableIPSecProposalRequest := _PatchedWritableIPSecProposalRequest{}

	err = json.Unmarshal(data, &varPatchedWritableIPSecProposalRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableIPSecProposalRequest(varPatchedWritableIPSecProposalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "sa_lifetime_seconds")
		delete(additionalProperties, "sa_lifetime_data")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableIPSecProposalRequest struct {
	value *PatchedWritableIPSecProposalRequest
	isSet bool
}

func (v NullablePatchedWritableIPSecProposalRequest) Get() *PatchedWritableIPSecProposalRequest {
	return v.value
}

func (v *NullablePatchedWritableIPSecProposalRequest) Set(val *PatchedWritableIPSecProposalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIPSecProposalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIPSecProposalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIPSecProposalRequest(val *PatchedWritableIPSecProposalRequest) *NullablePatchedWritableIPSecProposalRequest {
	return &NullablePatchedWritableIPSecProposalRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableIPSecProposalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIPSecProposalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

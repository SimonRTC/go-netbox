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

// IKEProposalAuthenticationAlgorithmLabel the model 'IKEProposalAuthenticationAlgorithmLabel'
type IKEProposalAuthenticationAlgorithmLabel string

// List of IKEProposal_authentication_algorithm_label
const (
	IKEPROPOSALAUTHENTICATIONALGORITHMLABEL_SHA_1_HMAC   IKEProposalAuthenticationAlgorithmLabel = "SHA-1 HMAC"
	IKEPROPOSALAUTHENTICATIONALGORITHMLABEL_SHA_256_HMAC IKEProposalAuthenticationAlgorithmLabel = "SHA-256 HMAC"
	IKEPROPOSALAUTHENTICATIONALGORITHMLABEL_SHA_384_HMAC IKEProposalAuthenticationAlgorithmLabel = "SHA-384 HMAC"
	IKEPROPOSALAUTHENTICATIONALGORITHMLABEL_SHA_512_HMAC IKEProposalAuthenticationAlgorithmLabel = "SHA-512 HMAC"
	IKEPROPOSALAUTHENTICATIONALGORITHMLABEL_MD5_HMAC     IKEProposalAuthenticationAlgorithmLabel = "MD5 HMAC"
)

// All allowed values of IKEProposalAuthenticationAlgorithmLabel enum
var AllowedIKEProposalAuthenticationAlgorithmLabelEnumValues = []IKEProposalAuthenticationAlgorithmLabel{
	"SHA-1 HMAC",
	"SHA-256 HMAC",
	"SHA-384 HMAC",
	"SHA-512 HMAC",
	"MD5 HMAC",
}

func (v *IKEProposalAuthenticationAlgorithmLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEProposalAuthenticationAlgorithmLabel(value)
	for _, existing := range AllowedIKEProposalAuthenticationAlgorithmLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEProposalAuthenticationAlgorithmLabel", value)
}

// NewIKEProposalAuthenticationAlgorithmLabelFromValue returns a pointer to a valid IKEProposalAuthenticationAlgorithmLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEProposalAuthenticationAlgorithmLabelFromValue(v string) (*IKEProposalAuthenticationAlgorithmLabel, error) {
	ev := IKEProposalAuthenticationAlgorithmLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEProposalAuthenticationAlgorithmLabel: valid values are %v", v, AllowedIKEProposalAuthenticationAlgorithmLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEProposalAuthenticationAlgorithmLabel) IsValid() bool {
	for _, existing := range AllowedIKEProposalAuthenticationAlgorithmLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEProposal_authentication_algorithm_label value
func (v IKEProposalAuthenticationAlgorithmLabel) Ptr() *IKEProposalAuthenticationAlgorithmLabel {
	return &v
}

type NullableIKEProposalAuthenticationAlgorithmLabel struct {
	value *IKEProposalAuthenticationAlgorithmLabel
	isSet bool
}

func (v NullableIKEProposalAuthenticationAlgorithmLabel) Get() *IKEProposalAuthenticationAlgorithmLabel {
	return v.value
}

func (v *NullableIKEProposalAuthenticationAlgorithmLabel) Set(val *IKEProposalAuthenticationAlgorithmLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalAuthenticationAlgorithmLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalAuthenticationAlgorithmLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalAuthenticationAlgorithmLabel(val *IKEProposalAuthenticationAlgorithmLabel) *NullableIKEProposalAuthenticationAlgorithmLabel {
	return &NullableIKEProposalAuthenticationAlgorithmLabel{value: val, isSet: true}
}

func (v NullableIKEProposalAuthenticationAlgorithmLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalAuthenticationAlgorithmLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

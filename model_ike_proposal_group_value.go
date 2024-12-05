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

// IKEProposalGroupValue * `1` - Group 1 * `2` - Group 2 * `5` - Group 5 * `14` - Group 14 * `15` - Group 15 * `16` - Group 16 * `17` - Group 17 * `18` - Group 18 * `19` - Group 19 * `20` - Group 20 * `21` - Group 21 * `22` - Group 22 * `23` - Group 23 * `24` - Group 24 * `25` - Group 25 * `26` - Group 26 * `27` - Group 27 * `28` - Group 28 * `29` - Group 29 * `30` - Group 30 * `31` - Group 31 * `32` - Group 32 * `33` - Group 33 * `34` - Group 34
type IKEProposalGroupValue int32

// List of IKEProposal_group_value
const (
	IKEPROPOSALGROUPVALUE__1  IKEProposalGroupValue = 1
	IKEPROPOSALGROUPVALUE__2  IKEProposalGroupValue = 2
	IKEPROPOSALGROUPVALUE__5  IKEProposalGroupValue = 5
	IKEPROPOSALGROUPVALUE__14 IKEProposalGroupValue = 14
	IKEPROPOSALGROUPVALUE__15 IKEProposalGroupValue = 15
	IKEPROPOSALGROUPVALUE__16 IKEProposalGroupValue = 16
	IKEPROPOSALGROUPVALUE__17 IKEProposalGroupValue = 17
	IKEPROPOSALGROUPVALUE__18 IKEProposalGroupValue = 18
	IKEPROPOSALGROUPVALUE__19 IKEProposalGroupValue = 19
	IKEPROPOSALGROUPVALUE__20 IKEProposalGroupValue = 20
	IKEPROPOSALGROUPVALUE__21 IKEProposalGroupValue = 21
	IKEPROPOSALGROUPVALUE__22 IKEProposalGroupValue = 22
	IKEPROPOSALGROUPVALUE__23 IKEProposalGroupValue = 23
	IKEPROPOSALGROUPVALUE__24 IKEProposalGroupValue = 24
	IKEPROPOSALGROUPVALUE__25 IKEProposalGroupValue = 25
	IKEPROPOSALGROUPVALUE__26 IKEProposalGroupValue = 26
	IKEPROPOSALGROUPVALUE__27 IKEProposalGroupValue = 27
	IKEPROPOSALGROUPVALUE__28 IKEProposalGroupValue = 28
	IKEPROPOSALGROUPVALUE__29 IKEProposalGroupValue = 29
	IKEPROPOSALGROUPVALUE__30 IKEProposalGroupValue = 30
	IKEPROPOSALGROUPVALUE__31 IKEProposalGroupValue = 31
	IKEPROPOSALGROUPVALUE__32 IKEProposalGroupValue = 32
	IKEPROPOSALGROUPVALUE__33 IKEProposalGroupValue = 33
	IKEPROPOSALGROUPVALUE__34 IKEProposalGroupValue = 34
)

// All allowed values of IKEProposalGroupValue enum
var AllowedIKEProposalGroupValueEnumValues = []IKEProposalGroupValue{
	1,
	2,
	5,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
}

func (v *IKEProposalGroupValue) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEProposalGroupValue(value)
	for _, existing := range AllowedIKEProposalGroupValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEProposalGroupValue", value)
}

// NewIKEProposalGroupValueFromValue returns a pointer to a valid IKEProposalGroupValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEProposalGroupValueFromValue(v int32) (*IKEProposalGroupValue, error) {
	ev := IKEProposalGroupValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEProposalGroupValue: valid values are %v", v, AllowedIKEProposalGroupValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEProposalGroupValue) IsValid() bool {
	for _, existing := range AllowedIKEProposalGroupValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEProposal_group_value value
func (v IKEProposalGroupValue) Ptr() *IKEProposalGroupValue {
	return &v
}

type NullableIKEProposalGroupValue struct {
	value *IKEProposalGroupValue
	isSet bool
}

func (v NullableIKEProposalGroupValue) Get() *IKEProposalGroupValue {
	return v.value
}

func (v *NullableIKEProposalGroupValue) Set(val *IKEProposalGroupValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalGroupValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalGroupValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalGroupValue(val *IKEProposalGroupValue) *NullableIKEProposalGroupValue {
	return &NullableIKEProposalGroupValue{value: val, isSet: true}
}

func (v NullableIKEProposalGroupValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalGroupValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

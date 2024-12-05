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

// PatchedWritableRackRequestFormFactor * `2-post-frame` - 2-post frame * `4-post-frame` - 4-post frame * `4-post-cabinet` - 4-post cabinet * `wall-frame` - Wall-mounted frame * `wall-frame-vertical` - Wall-mounted frame (vertical) * `wall-cabinet` - Wall-mounted cabinet * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
type PatchedWritableRackRequestFormFactor string

// List of PatchedWritableRackRequest_form_factor
const (
	PATCHEDWRITABLERACKREQUESTFORMFACTOR__2_POST_FRAME         PatchedWritableRackRequestFormFactor = "2-post-frame"
	PATCHEDWRITABLERACKREQUESTFORMFACTOR__4_POST_FRAME         PatchedWritableRackRequestFormFactor = "4-post-frame"
	PATCHEDWRITABLERACKREQUESTFORMFACTOR__4_POST_CABINET       PatchedWritableRackRequestFormFactor = "4-post-cabinet"
	PATCHEDWRITABLERACKREQUESTFORMFACTOR_WALL_FRAME            PatchedWritableRackRequestFormFactor = "wall-frame"
	PATCHEDWRITABLERACKREQUESTFORMFACTOR_WALL_FRAME_VERTICAL   PatchedWritableRackRequestFormFactor = "wall-frame-vertical"
	PATCHEDWRITABLERACKREQUESTFORMFACTOR_WALL_CABINET          PatchedWritableRackRequestFormFactor = "wall-cabinet"
	PATCHEDWRITABLERACKREQUESTFORMFACTOR_WALL_CABINET_VERTICAL PatchedWritableRackRequestFormFactor = "wall-cabinet-vertical"
	PATCHEDWRITABLERACKREQUESTFORMFACTOR_EMPTY                 PatchedWritableRackRequestFormFactor = ""
)

// All allowed values of PatchedWritableRackRequestFormFactor enum
var AllowedPatchedWritableRackRequestFormFactorEnumValues = []PatchedWritableRackRequestFormFactor{
	"2-post-frame",
	"4-post-frame",
	"4-post-cabinet",
	"wall-frame",
	"wall-frame-vertical",
	"wall-cabinet",
	"wall-cabinet-vertical",
	"",
}

func (v *PatchedWritableRackRequestFormFactor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableRackRequestFormFactor(value)
	for _, existing := range AllowedPatchedWritableRackRequestFormFactorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableRackRequestFormFactor", value)
}

// NewPatchedWritableRackRequestFormFactorFromValue returns a pointer to a valid PatchedWritableRackRequestFormFactor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableRackRequestFormFactorFromValue(v string) (*PatchedWritableRackRequestFormFactor, error) {
	ev := PatchedWritableRackRequestFormFactor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableRackRequestFormFactor: valid values are %v", v, AllowedPatchedWritableRackRequestFormFactorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableRackRequestFormFactor) IsValid() bool {
	for _, existing := range AllowedPatchedWritableRackRequestFormFactorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableRackRequest_form_factor value
func (v PatchedWritableRackRequestFormFactor) Ptr() *PatchedWritableRackRequestFormFactor {
	return &v
}

type NullablePatchedWritableRackRequestFormFactor struct {
	value *PatchedWritableRackRequestFormFactor
	isSet bool
}

func (v NullablePatchedWritableRackRequestFormFactor) Get() *PatchedWritableRackRequestFormFactor {
	return v.value
}

func (v *NullablePatchedWritableRackRequestFormFactor) Set(val *PatchedWritableRackRequestFormFactor) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequestFormFactor) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequestFormFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequestFormFactor(val *PatchedWritableRackRequestFormFactor) *NullablePatchedWritableRackRequestFormFactor {
	return &NullablePatchedWritableRackRequestFormFactor{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequestFormFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequestFormFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

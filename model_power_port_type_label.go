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

// PowerPortTypeLabel the model 'PowerPortTypeLabel'
type PowerPortTypeLabel string

// List of PowerPort_type_label
const (
	POWERPORTTYPELABEL_C6                          PowerPortTypeLabel = "C6"
	POWERPORTTYPELABEL_C8                          PowerPortTypeLabel = "C8"
	POWERPORTTYPELABEL_C14                         PowerPortTypeLabel = "C14"
	POWERPORTTYPELABEL_C16                         PowerPortTypeLabel = "C16"
	POWERPORTTYPELABEL_C20                         PowerPortTypeLabel = "C20"
	POWERPORTTYPELABEL_C22                         PowerPortTypeLabel = "C22"
	POWERPORTTYPELABEL_PNE_4_H                     PowerPortTypeLabel = "P+N+E 4H"
	POWERPORTTYPELABEL_PNE_6_H                     PowerPortTypeLabel = "P+N+E 6H"
	POWERPORTTYPELABEL_PNE_9_H                     PowerPortTypeLabel = "P+N+E 9H"
	POWERPORTTYPELABEL__2_PE_4_H                   PowerPortTypeLabel = "2P+E 4H"
	POWERPORTTYPELABEL__2_PE_6_H                   PowerPortTypeLabel = "2P+E 6H"
	POWERPORTTYPELABEL__2_PE_9_H                   PowerPortTypeLabel = "2P+E 9H"
	POWERPORTTYPELABEL__3_PE_4_H                   PowerPortTypeLabel = "3P+E 4H"
	POWERPORTTYPELABEL__3_PE_6_H                   PowerPortTypeLabel = "3P+E 6H"
	POWERPORTTYPELABEL__3_PE_9_H                   PowerPortTypeLabel = "3P+E 9H"
	POWERPORTTYPELABEL__3_PNE_4_H                  PowerPortTypeLabel = "3P+N+E 4H"
	POWERPORTTYPELABEL__3_PNE_6_H                  PowerPortTypeLabel = "3P+N+E 6H"
	POWERPORTTYPELABEL__3_PNE_9_H                  PowerPortTypeLabel = "3P+N+E 9H"
	POWERPORTTYPELABEL_IEC_60906_1                 PowerPortTypeLabel = "IEC 60906-1"
	POWERPORTTYPELABEL__2_PT_10_A__NBR_14136       PowerPortTypeLabel = "2P+T 10A (NBR 14136)"
	POWERPORTTYPELABEL__2_PT_20_A__NBR_14136       PowerPortTypeLabel = "2P+T 20A (NBR 14136)"
	POWERPORTTYPELABEL_NEMA_1_15_P                 PowerPortTypeLabel = "NEMA 1-15P"
	POWERPORTTYPELABEL_NEMA_5_15_P                 PowerPortTypeLabel = "NEMA 5-15P"
	POWERPORTTYPELABEL_NEMA_5_20_P                 PowerPortTypeLabel = "NEMA 5-20P"
	POWERPORTTYPELABEL_NEMA_5_30_P                 PowerPortTypeLabel = "NEMA 5-30P"
	POWERPORTTYPELABEL_NEMA_5_50_P                 PowerPortTypeLabel = "NEMA 5-50P"
	POWERPORTTYPELABEL_NEMA_6_15_P                 PowerPortTypeLabel = "NEMA 6-15P"
	POWERPORTTYPELABEL_NEMA_6_20_P                 PowerPortTypeLabel = "NEMA 6-20P"
	POWERPORTTYPELABEL_NEMA_6_30_P                 PowerPortTypeLabel = "NEMA 6-30P"
	POWERPORTTYPELABEL_NEMA_6_50_P                 PowerPortTypeLabel = "NEMA 6-50P"
	POWERPORTTYPELABEL_NEMA_10_30_P                PowerPortTypeLabel = "NEMA 10-30P"
	POWERPORTTYPELABEL_NEMA_10_50_P                PowerPortTypeLabel = "NEMA 10-50P"
	POWERPORTTYPELABEL_NEMA_14_20_P                PowerPortTypeLabel = "NEMA 14-20P"
	POWERPORTTYPELABEL_NEMA_14_30_P                PowerPortTypeLabel = "NEMA 14-30P"
	POWERPORTTYPELABEL_NEMA_14_50_P                PowerPortTypeLabel = "NEMA 14-50P"
	POWERPORTTYPELABEL_NEMA_14_60_P                PowerPortTypeLabel = "NEMA 14-60P"
	POWERPORTTYPELABEL_NEMA_15_15_P                PowerPortTypeLabel = "NEMA 15-15P"
	POWERPORTTYPELABEL_NEMA_15_20_P                PowerPortTypeLabel = "NEMA 15-20P"
	POWERPORTTYPELABEL_NEMA_15_30_P                PowerPortTypeLabel = "NEMA 15-30P"
	POWERPORTTYPELABEL_NEMA_15_50_P                PowerPortTypeLabel = "NEMA 15-50P"
	POWERPORTTYPELABEL_NEMA_15_60_P                PowerPortTypeLabel = "NEMA 15-60P"
	POWERPORTTYPELABEL_NEMA_L1_15_P                PowerPortTypeLabel = "NEMA L1-15P"
	POWERPORTTYPELABEL_NEMA_L5_15_P                PowerPortTypeLabel = "NEMA L5-15P"
	POWERPORTTYPELABEL_NEMA_L5_20_P                PowerPortTypeLabel = "NEMA L5-20P"
	POWERPORTTYPELABEL_NEMA_L5_30_P                PowerPortTypeLabel = "NEMA L5-30P"
	POWERPORTTYPELABEL_NEMA_L5_50_P                PowerPortTypeLabel = "NEMA L5-50P"
	POWERPORTTYPELABEL_NEMA_L6_15_P                PowerPortTypeLabel = "NEMA L6-15P"
	POWERPORTTYPELABEL_NEMA_L6_20_P                PowerPortTypeLabel = "NEMA L6-20P"
	POWERPORTTYPELABEL_NEMA_L6_30_P                PowerPortTypeLabel = "NEMA L6-30P"
	POWERPORTTYPELABEL_NEMA_L6_50_P                PowerPortTypeLabel = "NEMA L6-50P"
	POWERPORTTYPELABEL_NEMA_L10_30_P               PowerPortTypeLabel = "NEMA L10-30P"
	POWERPORTTYPELABEL_NEMA_L14_20_P               PowerPortTypeLabel = "NEMA L14-20P"
	POWERPORTTYPELABEL_NEMA_L14_30_P               PowerPortTypeLabel = "NEMA L14-30P"
	POWERPORTTYPELABEL_NEMA_L14_50_P               PowerPortTypeLabel = "NEMA L14-50P"
	POWERPORTTYPELABEL_NEMA_L14_60_P               PowerPortTypeLabel = "NEMA L14-60P"
	POWERPORTTYPELABEL_NEMA_L15_20_P               PowerPortTypeLabel = "NEMA L15-20P"
	POWERPORTTYPELABEL_NEMA_L15_30_P               PowerPortTypeLabel = "NEMA L15-30P"
	POWERPORTTYPELABEL_NEMA_L15_50_P               PowerPortTypeLabel = "NEMA L15-50P"
	POWERPORTTYPELABEL_NEMA_L15_60_P               PowerPortTypeLabel = "NEMA L15-60P"
	POWERPORTTYPELABEL_NEMA_L21_20_P               PowerPortTypeLabel = "NEMA L21-20P"
	POWERPORTTYPELABEL_NEMA_L21_30_P               PowerPortTypeLabel = "NEMA L21-30P"
	POWERPORTTYPELABEL_NEMA_L22_20_P               PowerPortTypeLabel = "NEMA L22-20P"
	POWERPORTTYPELABEL_NEMA_L22_30_P               PowerPortTypeLabel = "NEMA L22-30P"
	POWERPORTTYPELABEL_CS6361_C                    PowerPortTypeLabel = "CS6361C"
	POWERPORTTYPELABEL_CS6365_C                    PowerPortTypeLabel = "CS6365C"
	POWERPORTTYPELABEL_CS8165_C                    PowerPortTypeLabel = "CS8165C"
	POWERPORTTYPELABEL_CS8265_C                    PowerPortTypeLabel = "CS8265C"
	POWERPORTTYPELABEL_CS8365_C                    PowerPortTypeLabel = "CS8365C"
	POWERPORTTYPELABEL_CS8465_C                    PowerPortTypeLabel = "CS8465C"
	POWERPORTTYPELABEL_ITA_TYPE_C__CEE_7_16        PowerPortTypeLabel = "ITA Type C (CEE 7/16)"
	POWERPORTTYPELABEL_ITA_TYPE_E__CEE_7_6         PowerPortTypeLabel = "ITA Type E (CEE 7/6)"
	POWERPORTTYPELABEL_ITA_TYPE_F__CEE_7_4         PowerPortTypeLabel = "ITA Type F (CEE 7/4)"
	POWERPORTTYPELABEL_ITA_TYPE_E_F__CEE_7_7       PowerPortTypeLabel = "ITA Type E/F (CEE 7/7)"
	POWERPORTTYPELABEL_ITA_TYPE_G__BS_1363         PowerPortTypeLabel = "ITA Type G (BS 1363)"
	POWERPORTTYPELABEL_ITA_TYPE_H                  PowerPortTypeLabel = "ITA Type H"
	POWERPORTTYPELABEL_ITA_TYPE_I                  PowerPortTypeLabel = "ITA Type I"
	POWERPORTTYPELABEL_ITA_TYPE_J                  PowerPortTypeLabel = "ITA Type J"
	POWERPORTTYPELABEL_ITA_TYPE_K                  PowerPortTypeLabel = "ITA Type K"
	POWERPORTTYPELABEL_ITA_TYPE_L__CEI_23_50       PowerPortTypeLabel = "ITA Type L (CEI 23-50)"
	POWERPORTTYPELABEL_ITA_TYPE_M__BS_546          PowerPortTypeLabel = "ITA Type M (BS 546)"
	POWERPORTTYPELABEL_ITA_TYPE_N                  PowerPortTypeLabel = "ITA Type N"
	POWERPORTTYPELABEL_ITA_TYPE_O                  PowerPortTypeLabel = "ITA Type O"
	POWERPORTTYPELABEL_USB_TYPE_A                  PowerPortTypeLabel = "USB Type A"
	POWERPORTTYPELABEL_USB_TYPE_B                  PowerPortTypeLabel = "USB Type B"
	POWERPORTTYPELABEL_USB_TYPE_C                  PowerPortTypeLabel = "USB Type C"
	POWERPORTTYPELABEL_USB_MINI_A                  PowerPortTypeLabel = "USB Mini A"
	POWERPORTTYPELABEL_USB_MINI_B                  PowerPortTypeLabel = "USB Mini B"
	POWERPORTTYPELABEL_USB_MICRO_A                 PowerPortTypeLabel = "USB Micro A"
	POWERPORTTYPELABEL_USB_MICRO_B                 PowerPortTypeLabel = "USB Micro B"
	POWERPORTTYPELABEL_USB_MICRO_AB                PowerPortTypeLabel = "USB Micro AB"
	POWERPORTTYPELABEL_USB_3_0_TYPE_B              PowerPortTypeLabel = "USB 3.0 Type B"
	POWERPORTTYPELABEL_USB_3_0_MICRO_B             PowerPortTypeLabel = "USB 3.0 Micro B"
	POWERPORTTYPELABEL_MOLEX_MICRO_FIT_1X2         PowerPortTypeLabel = "Molex Micro-Fit 1x2"
	POWERPORTTYPELABEL_MOLEX_MICRO_FIT_2X2         PowerPortTypeLabel = "Molex Micro-Fit 2x2"
	POWERPORTTYPELABEL_MOLEX_MICRO_FIT_2X4         PowerPortTypeLabel = "Molex Micro-Fit 2x4"
	POWERPORTTYPELABEL_DC_TERMINAL                 PowerPortTypeLabel = "DC Terminal"
	POWERPORTTYPELABEL_SAF_D_GRID                  PowerPortTypeLabel = "Saf-D-Grid"
	POWERPORTTYPELABEL_NEUTRIK_POWER_CON__20_A     PowerPortTypeLabel = "Neutrik powerCON (20A)"
	POWERPORTTYPELABEL_NEUTRIK_POWER_CON__32_A     PowerPortTypeLabel = "Neutrik powerCON (32A)"
	POWERPORTTYPELABEL_NEUTRIK_POWER_CON_TRUE1     PowerPortTypeLabel = "Neutrik powerCON TRUE1"
	POWERPORTTYPELABEL_NEUTRIK_POWER_CON_TRUE1_TOP PowerPortTypeLabel = "Neutrik powerCON TRUE1 TOP"
	POWERPORTTYPELABEL_UBIQUITI_SMART_POWER        PowerPortTypeLabel = "Ubiquiti SmartPower"
	POWERPORTTYPELABEL_HARDWIRED                   PowerPortTypeLabel = "Hardwired"
	POWERPORTTYPELABEL_OTHER                       PowerPortTypeLabel = "Other"
)

// All allowed values of PowerPortTypeLabel enum
var AllowedPowerPortTypeLabelEnumValues = []PowerPortTypeLabel{
	"C6",
	"C8",
	"C14",
	"C16",
	"C20",
	"C22",
	"P+N+E 4H",
	"P+N+E 6H",
	"P+N+E 9H",
	"2P+E 4H",
	"2P+E 6H",
	"2P+E 9H",
	"3P+E 4H",
	"3P+E 6H",
	"3P+E 9H",
	"3P+N+E 4H",
	"3P+N+E 6H",
	"3P+N+E 9H",
	"IEC 60906-1",
	"2P+T 10A (NBR 14136)",
	"2P+T 20A (NBR 14136)",
	"NEMA 1-15P",
	"NEMA 5-15P",
	"NEMA 5-20P",
	"NEMA 5-30P",
	"NEMA 5-50P",
	"NEMA 6-15P",
	"NEMA 6-20P",
	"NEMA 6-30P",
	"NEMA 6-50P",
	"NEMA 10-30P",
	"NEMA 10-50P",
	"NEMA 14-20P",
	"NEMA 14-30P",
	"NEMA 14-50P",
	"NEMA 14-60P",
	"NEMA 15-15P",
	"NEMA 15-20P",
	"NEMA 15-30P",
	"NEMA 15-50P",
	"NEMA 15-60P",
	"NEMA L1-15P",
	"NEMA L5-15P",
	"NEMA L5-20P",
	"NEMA L5-30P",
	"NEMA L5-50P",
	"NEMA L6-15P",
	"NEMA L6-20P",
	"NEMA L6-30P",
	"NEMA L6-50P",
	"NEMA L10-30P",
	"NEMA L14-20P",
	"NEMA L14-30P",
	"NEMA L14-50P",
	"NEMA L14-60P",
	"NEMA L15-20P",
	"NEMA L15-30P",
	"NEMA L15-50P",
	"NEMA L15-60P",
	"NEMA L21-20P",
	"NEMA L21-30P",
	"NEMA L22-20P",
	"NEMA L22-30P",
	"CS6361C",
	"CS6365C",
	"CS8165C",
	"CS8265C",
	"CS8365C",
	"CS8465C",
	"ITA Type C (CEE 7/16)",
	"ITA Type E (CEE 7/6)",
	"ITA Type F (CEE 7/4)",
	"ITA Type E/F (CEE 7/7)",
	"ITA Type G (BS 1363)",
	"ITA Type H",
	"ITA Type I",
	"ITA Type J",
	"ITA Type K",
	"ITA Type L (CEI 23-50)",
	"ITA Type M (BS 546)",
	"ITA Type N",
	"ITA Type O",
	"USB Type A",
	"USB Type B",
	"USB Type C",
	"USB Mini A",
	"USB Mini B",
	"USB Micro A",
	"USB Micro B",
	"USB Micro AB",
	"USB 3.0 Type B",
	"USB 3.0 Micro B",
	"Molex Micro-Fit 1x2",
	"Molex Micro-Fit 2x2",
	"Molex Micro-Fit 2x4",
	"DC Terminal",
	"Saf-D-Grid",
	"Neutrik powerCON (20A)",
	"Neutrik powerCON (32A)",
	"Neutrik powerCON TRUE1",
	"Neutrik powerCON TRUE1 TOP",
	"Ubiquiti SmartPower",
	"Hardwired",
	"Other",
}

func (v *PowerPortTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerPortTypeLabel(value)
	for _, existing := range AllowedPowerPortTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerPortTypeLabel", value)
}

// NewPowerPortTypeLabelFromValue returns a pointer to a valid PowerPortTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerPortTypeLabelFromValue(v string) (*PowerPortTypeLabel, error) {
	ev := PowerPortTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerPortTypeLabel: valid values are %v", v, AllowedPowerPortTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerPortTypeLabel) IsValid() bool {
	for _, existing := range AllowedPowerPortTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerPort_type_label value
func (v PowerPortTypeLabel) Ptr() *PowerPortTypeLabel {
	return &v
}

type NullablePowerPortTypeLabel struct {
	value *PowerPortTypeLabel
	isSet bool
}

func (v NullablePowerPortTypeLabel) Get() *PowerPortTypeLabel {
	return v.value
}

func (v *NullablePowerPortTypeLabel) Set(val *PowerPortTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPortTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPortTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPortTypeLabel(val *PowerPortTypeLabel) *NullablePowerPortTypeLabel {
	return &NullablePowerPortTypeLabel{value: val, isSet: true}
}

func (v NullablePowerPortTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPortTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

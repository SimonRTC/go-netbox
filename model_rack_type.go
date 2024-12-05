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

// checks if the RackType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackType{}

// RackType Adds support for custom fields and tags.
type RackType struct {
	Id           int32                  `json:"id"`
	Url          string                 `json:"url"`
	DisplayUrl   string                 `json:"display_url"`
	Display      string                 `json:"display"`
	Manufacturer BriefManufacturer      `json:"manufacturer"`
	Model        string                 `json:"model"`
	Slug         string                 `json:"slug"`
	Description  *string                `json:"description,omitempty"`
	FormFactor   NullableRackFormFactor `json:"form_factor,omitempty"`
	Width        *RackWidth             `json:"width,omitempty"`
	// Height in rack units
	UHeight *int32 `json:"u_height,omitempty"`
	// Starting unit for rack
	StartingUnit *int32 `json:"starting_unit,omitempty"`
	// Units are numbered top-to-bottom
	DescUnits *bool `json:"desc_units,omitempty"`
	// Outer dimension of rack (width)
	OuterWidth NullableInt32 `json:"outer_width,omitempty"`
	// Outer dimension of rack (depth)
	OuterDepth NullableInt32         `json:"outer_depth,omitempty"`
	OuterUnit  NullableRackOuterUnit `json:"outer_unit,omitempty"`
	Weight     NullableFloat64       `json:"weight,omitempty"`
	// Maximum load capacity for the rack
	MaxWeight  NullableInt32                `json:"max_weight,omitempty"`
	WeightUnit NullableDeviceTypeWeightUnit `json:"weight_unit,omitempty"`
	// Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	MountingDepth        NullableInt32          `json:"mounting_depth,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _RackType RackType

// NewRackType instantiates a new RackType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackType(id int32, url string, displayUrl string, display string, manufacturer BriefManufacturer, model string, slug string, created NullableTime, lastUpdated NullableTime) *RackType {
	this := RackType{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewRackTypeWithDefaults instantiates a new RackType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackTypeWithDefaults() *RackType {
	this := RackType{}
	return &this
}

// GetId returns the Id field value
func (o *RackType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RackType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RackType) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *RackType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RackType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RackType) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *RackType) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *RackType) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *RackType) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *RackType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *RackType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *RackType) SetDisplay(v string) {
	o.Display = v
}

// GetManufacturer returns the Manufacturer field value
func (o *RackType) GetManufacturer() BriefManufacturer {
	if o == nil {
		var ret BriefManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *RackType) GetManufacturerOk() (*BriefManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *RackType) SetManufacturer(v BriefManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *RackType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *RackType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *RackType) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *RackType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *RackType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *RackType) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RackType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RackType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RackType) SetDescription(v string) {
	o.Description = &v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetFormFactor() RackFormFactor {
	if o == nil || IsNil(o.FormFactor.Get()) {
		var ret RackFormFactor
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetFormFactorOk() (*RackFormFactor, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *RackType) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullableRackFormFactor and assigns it to the FormFactor field.
func (o *RackType) SetFormFactor(v RackFormFactor) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *RackType) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *RackType) UnsetFormFactor() {
	o.FormFactor.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *RackType) GetWidth() RackWidth {
	if o == nil || IsNil(o.Width) {
		var ret RackWidth
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetWidthOk() (*RackWidth, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *RackType) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given RackWidth and assigns it to the Width field.
func (o *RackType) SetWidth(v RackWidth) {
	o.Width = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *RackType) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *RackType) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *RackType) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetStartingUnit returns the StartingUnit field value if set, zero value otherwise.
func (o *RackType) GetStartingUnit() int32 {
	if o == nil || IsNil(o.StartingUnit) {
		var ret int32
		return ret
	}
	return *o.StartingUnit
}

// GetStartingUnitOk returns a tuple with the StartingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetStartingUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.StartingUnit) {
		return nil, false
	}
	return o.StartingUnit, true
}

// HasStartingUnit returns a boolean if a field has been set.
func (o *RackType) HasStartingUnit() bool {
	if o != nil && !IsNil(o.StartingUnit) {
		return true
	}

	return false
}

// SetStartingUnit gets a reference to the given int32 and assigns it to the StartingUnit field.
func (o *RackType) SetStartingUnit(v int32) {
	o.StartingUnit = &v
}

// GetDescUnits returns the DescUnits field value if set, zero value otherwise.
func (o *RackType) GetDescUnits() bool {
	if o == nil || IsNil(o.DescUnits) {
		var ret bool
		return ret
	}
	return *o.DescUnits
}

// GetDescUnitsOk returns a tuple with the DescUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetDescUnitsOk() (*bool, bool) {
	if o == nil || IsNil(o.DescUnits) {
		return nil, false
	}
	return o.DescUnits, true
}

// HasDescUnits returns a boolean if a field has been set.
func (o *RackType) HasDescUnits() bool {
	if o != nil && !IsNil(o.DescUnits) {
		return true
	}

	return false
}

// SetDescUnits gets a reference to the given bool and assigns it to the DescUnits field.
func (o *RackType) SetDescUnits(v bool) {
	o.DescUnits = &v
}

// GetOuterWidth returns the OuterWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetOuterWidth() int32 {
	if o == nil || IsNil(o.OuterWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterWidth.Get()
}

// GetOuterWidthOk returns a tuple with the OuterWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetOuterWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterWidth.Get(), o.OuterWidth.IsSet()
}

// HasOuterWidth returns a boolean if a field has been set.
func (o *RackType) HasOuterWidth() bool {
	if o != nil && o.OuterWidth.IsSet() {
		return true
	}

	return false
}

// SetOuterWidth gets a reference to the given NullableInt32 and assigns it to the OuterWidth field.
func (o *RackType) SetOuterWidth(v int32) {
	o.OuterWidth.Set(&v)
}

// SetOuterWidthNil sets the value for OuterWidth to be an explicit nil
func (o *RackType) SetOuterWidthNil() {
	o.OuterWidth.Set(nil)
}

// UnsetOuterWidth ensures that no value is present for OuterWidth, not even an explicit nil
func (o *RackType) UnsetOuterWidth() {
	o.OuterWidth.Unset()
}

// GetOuterDepth returns the OuterDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetOuterDepth() int32 {
	if o == nil || IsNil(o.OuterDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.OuterDepth.Get()
}

// GetOuterDepthOk returns a tuple with the OuterDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetOuterDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterDepth.Get(), o.OuterDepth.IsSet()
}

// HasOuterDepth returns a boolean if a field has been set.
func (o *RackType) HasOuterDepth() bool {
	if o != nil && o.OuterDepth.IsSet() {
		return true
	}

	return false
}

// SetOuterDepth gets a reference to the given NullableInt32 and assigns it to the OuterDepth field.
func (o *RackType) SetOuterDepth(v int32) {
	o.OuterDepth.Set(&v)
}

// SetOuterDepthNil sets the value for OuterDepth to be an explicit nil
func (o *RackType) SetOuterDepthNil() {
	o.OuterDepth.Set(nil)
}

// UnsetOuterDepth ensures that no value is present for OuterDepth, not even an explicit nil
func (o *RackType) UnsetOuterDepth() {
	o.OuterDepth.Unset()
}

// GetOuterUnit returns the OuterUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetOuterUnit() RackOuterUnit {
	if o == nil || IsNil(o.OuterUnit.Get()) {
		var ret RackOuterUnit
		return ret
	}
	return *o.OuterUnit.Get()
}

// GetOuterUnitOk returns a tuple with the OuterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetOuterUnitOk() (*RackOuterUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterUnit.Get(), o.OuterUnit.IsSet()
}

// HasOuterUnit returns a boolean if a field has been set.
func (o *RackType) HasOuterUnit() bool {
	if o != nil && o.OuterUnit.IsSet() {
		return true
	}

	return false
}

// SetOuterUnit gets a reference to the given NullableRackOuterUnit and assigns it to the OuterUnit field.
func (o *RackType) SetOuterUnit(v RackOuterUnit) {
	o.OuterUnit.Set(&v)
}

// SetOuterUnitNil sets the value for OuterUnit to be an explicit nil
func (o *RackType) SetOuterUnitNil() {
	o.OuterUnit.Set(nil)
}

// UnsetOuterUnit ensures that no value is present for OuterUnit, not even an explicit nil
func (o *RackType) UnsetOuterUnit() {
	o.OuterUnit.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *RackType) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *RackType) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *RackType) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *RackType) UnsetWeight() {
	o.Weight.Unset()
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetMaxWeight() int32 {
	if o == nil || IsNil(o.MaxWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxWeight.Get()
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetMaxWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxWeight.Get(), o.MaxWeight.IsSet()
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *RackType) HasMaxWeight() bool {
	if o != nil && o.MaxWeight.IsSet() {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given NullableInt32 and assigns it to the MaxWeight field.
func (o *RackType) SetMaxWeight(v int32) {
	o.MaxWeight.Set(&v)
}

// SetMaxWeightNil sets the value for MaxWeight to be an explicit nil
func (o *RackType) SetMaxWeightNil() {
	o.MaxWeight.Set(nil)
}

// UnsetMaxWeight ensures that no value is present for MaxWeight, not even an explicit nil
func (o *RackType) UnsetMaxWeight() {
	o.MaxWeight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetWeightUnit() DeviceTypeWeightUnit {
	if o == nil || IsNil(o.WeightUnit.Get()) {
		var ret DeviceTypeWeightUnit
		return ret
	}
	return *o.WeightUnit.Get()
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetWeightUnitOk() (*DeviceTypeWeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.WeightUnit.Get(), o.WeightUnit.IsSet()
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *RackType) HasWeightUnit() bool {
	if o != nil && o.WeightUnit.IsSet() {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given NullableDeviceTypeWeightUnit and assigns it to the WeightUnit field.
func (o *RackType) SetWeightUnit(v DeviceTypeWeightUnit) {
	o.WeightUnit.Set(&v)
}

// SetWeightUnitNil sets the value for WeightUnit to be an explicit nil
func (o *RackType) SetWeightUnitNil() {
	o.WeightUnit.Set(nil)
}

// UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
func (o *RackType) UnsetWeightUnit() {
	o.WeightUnit.Unset()
}

// GetMountingDepth returns the MountingDepth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackType) GetMountingDepth() int32 {
	if o == nil || IsNil(o.MountingDepth.Get()) {
		var ret int32
		return ret
	}
	return *o.MountingDepth.Get()
}

// GetMountingDepthOk returns a tuple with the MountingDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetMountingDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountingDepth.Get(), o.MountingDepth.IsSet()
}

// HasMountingDepth returns a boolean if a field has been set.
func (o *RackType) HasMountingDepth() bool {
	if o != nil && o.MountingDepth.IsSet() {
		return true
	}

	return false
}

// SetMountingDepth gets a reference to the given NullableInt32 and assigns it to the MountingDepth field.
func (o *RackType) SetMountingDepth(v int32) {
	o.MountingDepth.Set(&v)
}

// SetMountingDepthNil sets the value for MountingDepth to be an explicit nil
func (o *RackType) SetMountingDepthNil() {
	o.MountingDepth.Set(nil)
}

// UnsetMountingDepth ensures that no value is present for MountingDepth, not even an explicit nil
func (o *RackType) UnsetMountingDepth() {
	o.MountingDepth.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *RackType) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *RackType) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *RackType) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RackType) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RackType) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *RackType) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RackType) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RackType) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RackType) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RackType) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *RackType) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RackType) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackType) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *RackType) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o RackType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.FormFactor.IsSet() {
		toSerialize["form_factor"] = o.FormFactor.Get()
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.StartingUnit) {
		toSerialize["starting_unit"] = o.StartingUnit
	}
	if !IsNil(o.DescUnits) {
		toSerialize["desc_units"] = o.DescUnits
	}
	if o.OuterWidth.IsSet() {
		toSerialize["outer_width"] = o.OuterWidth.Get()
	}
	if o.OuterDepth.IsSet() {
		toSerialize["outer_depth"] = o.OuterDepth.Get()
	}
	if o.OuterUnit.IsSet() {
		toSerialize["outer_unit"] = o.OuterUnit.Get()
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.MaxWeight.IsSet() {
		toSerialize["max_weight"] = o.MaxWeight.Get()
	}
	if o.WeightUnit.IsSet() {
		toSerialize["weight_unit"] = o.WeightUnit.Get()
	}
	if o.MountingDepth.IsSet() {
		toSerialize["mounting_depth"] = o.MountingDepth.Get()
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

func (o *RackType) UnmarshalJSON(data []byte) (err error) {
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
		"slug",
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

	varRackType := _RackType{}

	err = json.Unmarshal(data, &varRackType)

	if err != nil {
		return err
	}

	*o = RackType(varRackType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "form_factor")
		delete(additionalProperties, "width")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "starting_unit")
		delete(additionalProperties, "desc_units")
		delete(additionalProperties, "outer_width")
		delete(additionalProperties, "outer_depth")
		delete(additionalProperties, "outer_unit")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "max_weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "mounting_depth")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackType struct {
	value *RackType
	isSet bool
}

func (v NullableRackType) Get() *RackType {
	return v.value
}

func (v *NullableRackType) Set(val *RackType) {
	v.value = val
	v.isSet = true
}

func (v NullableRackType) IsSet() bool {
	return v.isSet
}

func (v *NullableRackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackType(val *RackType) *NullableRackType {
	return &NullableRackType{value: val, isSet: true}
}

func (v NullableRackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

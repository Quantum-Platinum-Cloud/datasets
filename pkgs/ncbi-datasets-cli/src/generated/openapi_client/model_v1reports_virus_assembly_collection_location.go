/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1reportsVirusAssemblyCollectionLocation struct for V1reportsVirusAssemblyCollectionLocation
type V1reportsVirusAssemblyCollectionLocation struct {
	GeographicLocation *string `json:"geographic_location,omitempty"`
	GeographicRegion *string `json:"geographic_region,omitempty"`
}

// NewV1reportsVirusAssemblyCollectionLocation instantiates a new V1reportsVirusAssemblyCollectionLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsVirusAssemblyCollectionLocation() *V1reportsVirusAssemblyCollectionLocation {
	this := V1reportsVirusAssemblyCollectionLocation{}
	return &this
}

// NewV1reportsVirusAssemblyCollectionLocationWithDefaults instantiates a new V1reportsVirusAssemblyCollectionLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsVirusAssemblyCollectionLocationWithDefaults() *V1reportsVirusAssemblyCollectionLocation {
	this := V1reportsVirusAssemblyCollectionLocation{}
	return &this
}

// GetGeographicLocation returns the GeographicLocation field value if set, zero value otherwise.
func (o *V1reportsVirusAssemblyCollectionLocation) GetGeographicLocation() string {
	if o == nil || o.GeographicLocation == nil {
		var ret string
		return ret
	}
	return *o.GeographicLocation
}

// GetGeographicLocationOk returns a tuple with the GeographicLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusAssemblyCollectionLocation) GetGeographicLocationOk() (*string, bool) {
	if o == nil || o.GeographicLocation == nil {
		return nil, false
	}
	return o.GeographicLocation, true
}

// HasGeographicLocation returns a boolean if a field has been set.
func (o *V1reportsVirusAssemblyCollectionLocation) HasGeographicLocation() bool {
	if o != nil && o.GeographicLocation != nil {
		return true
	}

	return false
}

// SetGeographicLocation gets a reference to the given string and assigns it to the GeographicLocation field.
func (o *V1reportsVirusAssemblyCollectionLocation) SetGeographicLocation(v string) {
	o.GeographicLocation = &v
}

// GetGeographicRegion returns the GeographicRegion field value if set, zero value otherwise.
func (o *V1reportsVirusAssemblyCollectionLocation) GetGeographicRegion() string {
	if o == nil || o.GeographicRegion == nil {
		var ret string
		return ret
	}
	return *o.GeographicRegion
}

// GetGeographicRegionOk returns a tuple with the GeographicRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsVirusAssemblyCollectionLocation) GetGeographicRegionOk() (*string, bool) {
	if o == nil || o.GeographicRegion == nil {
		return nil, false
	}
	return o.GeographicRegion, true
}

// HasGeographicRegion returns a boolean if a field has been set.
func (o *V1reportsVirusAssemblyCollectionLocation) HasGeographicRegion() bool {
	if o != nil && o.GeographicRegion != nil {
		return true
	}

	return false
}

// SetGeographicRegion gets a reference to the given string and assigns it to the GeographicRegion field.
func (o *V1reportsVirusAssemblyCollectionLocation) SetGeographicRegion(v string) {
	o.GeographicRegion = &v
}

func (o V1reportsVirusAssemblyCollectionLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeographicLocation != nil  {
		toSerialize["geographic_location"] = o.GeographicLocation
	}
	if o.GeographicRegion != nil  {
		toSerialize["geographic_region"] = o.GeographicRegion
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsVirusAssemblyCollectionLocation struct {
	value *V1reportsVirusAssemblyCollectionLocation
	isSet bool
}

func (v NullableV1reportsVirusAssemblyCollectionLocation) Get() *V1reportsVirusAssemblyCollectionLocation {
	return v.value
}

func (v *NullableV1reportsVirusAssemblyCollectionLocation) Set(val *V1reportsVirusAssemblyCollectionLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsVirusAssemblyCollectionLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsVirusAssemblyCollectionLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsVirusAssemblyCollectionLocation(val *V1reportsVirusAssemblyCollectionLocation) *NullableV1reportsVirusAssemblyCollectionLocation {
	return &NullableV1reportsVirusAssemblyCollectionLocation{value: val, isSet: true}
}

func (v NullableV1reportsVirusAssemblyCollectionLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsVirusAssemblyCollectionLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



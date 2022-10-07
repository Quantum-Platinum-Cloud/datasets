/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 1 API is considred stable and will not be subject to breaking changes. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1OrganismCountByType struct for V1OrganismCountByType
type V1OrganismCountByType struct {
	Type *V1CountType `json:"type,omitempty"`
	Counts *V1OrganismCounts `json:"counts,omitempty"`
}

// NewV1OrganismCountByType instantiates a new V1OrganismCountByType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1OrganismCountByType() *V1OrganismCountByType {
	this := V1OrganismCountByType{}
	var type_ V1CountType = V1COUNTTYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// NewV1OrganismCountByTypeWithDefaults instantiates a new V1OrganismCountByType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1OrganismCountByTypeWithDefaults() *V1OrganismCountByType {
	this := V1OrganismCountByType{}
	var type_ V1CountType = V1COUNTTYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1OrganismCountByType) GetType() V1CountType {
	if o == nil || o.Type == nil {
		var ret V1CountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1OrganismCountByType) GetTypeOk() (*V1CountType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1OrganismCountByType) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V1CountType and assigns it to the Type field.
func (o *V1OrganismCountByType) SetType(v V1CountType) {
	o.Type = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *V1OrganismCountByType) GetCounts() V1OrganismCounts {
	if o == nil || o.Counts == nil {
		var ret V1OrganismCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1OrganismCountByType) GetCountsOk() (*V1OrganismCounts, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *V1OrganismCountByType) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given V1OrganismCounts and assigns it to the Counts field.
func (o *V1OrganismCountByType) SetCounts(v V1OrganismCounts) {
	o.Counts = &v
}

func (o V1OrganismCountByType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil  {
		toSerialize["type"] = o.Type
	}
	if o.Counts != nil  {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableV1OrganismCountByType struct {
	value *V1OrganismCountByType
	isSet bool
}

func (v NullableV1OrganismCountByType) Get() *V1OrganismCountByType {
	return v.value
}

func (v *NullableV1OrganismCountByType) Set(val *V1OrganismCountByType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1OrganismCountByType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1OrganismCountByType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1OrganismCountByType(val *V1OrganismCountByType) *NullableV1OrganismCountByType {
	return &NullableV1OrganismCountByType{value: val, isSet: true}
}

func (v NullableV1OrganismCountByType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1OrganismCountByType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



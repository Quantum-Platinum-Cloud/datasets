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

// V1reportsBioSampleStatus struct for V1reportsBioSampleStatus
type V1reportsBioSampleStatus struct {
	Status *string `json:"status,omitempty"`
	When *string `json:"when,omitempty"`
}

// NewV1reportsBioSampleStatus instantiates a new V1reportsBioSampleStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsBioSampleStatus() *V1reportsBioSampleStatus {
	this := V1reportsBioSampleStatus{}
	return &this
}

// NewV1reportsBioSampleStatusWithDefaults instantiates a new V1reportsBioSampleStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsBioSampleStatusWithDefaults() *V1reportsBioSampleStatus {
	this := V1reportsBioSampleStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1reportsBioSampleStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1reportsBioSampleStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *V1reportsBioSampleStatus) SetStatus(v string) {
	o.Status = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *V1reportsBioSampleStatus) GetWhen() string {
	if o == nil || o.When == nil {
		var ret string
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleStatus) GetWhenOk() (*string, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *V1reportsBioSampleStatus) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given string and assigns it to the When field.
func (o *V1reportsBioSampleStatus) SetWhen(v string) {
	o.When = &v
}

func (o V1reportsBioSampleStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil  {
		toSerialize["status"] = o.Status
	}
	if o.When != nil  {
		toSerialize["when"] = o.When
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsBioSampleStatus struct {
	value *V1reportsBioSampleStatus
	isSet bool
}

func (v NullableV1reportsBioSampleStatus) Get() *V1reportsBioSampleStatus {
	return v.value
}

func (v *NullableV1reportsBioSampleStatus) Set(val *V1reportsBioSampleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsBioSampleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsBioSampleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsBioSampleStatus(val *V1reportsBioSampleStatus) *NullableV1reportsBioSampleStatus {
	return &NullableV1reportsBioSampleStatus{value: val, isSet: true}
}

func (v NullableV1reportsBioSampleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsBioSampleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



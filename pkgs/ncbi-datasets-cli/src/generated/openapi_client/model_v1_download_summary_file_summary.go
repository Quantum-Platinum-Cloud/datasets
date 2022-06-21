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

// V1DownloadSummaryFileSummary struct for V1DownloadSummaryFileSummary
type V1DownloadSummaryFileSummary struct {
	FileCount *int32 `json:"file_count,omitempty"`
	SizeMb *float32 `json:"size_mb,omitempty"`
}

// NewV1DownloadSummaryFileSummary instantiates a new V1DownloadSummaryFileSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DownloadSummaryFileSummary() *V1DownloadSummaryFileSummary {
	this := V1DownloadSummaryFileSummary{}
	return &this
}

// NewV1DownloadSummaryFileSummaryWithDefaults instantiates a new V1DownloadSummaryFileSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DownloadSummaryFileSummaryWithDefaults() *V1DownloadSummaryFileSummary {
	this := V1DownloadSummaryFileSummary{}
	return &this
}

// GetFileCount returns the FileCount field value if set, zero value otherwise.
func (o *V1DownloadSummaryFileSummary) GetFileCount() int32 {
	if o == nil || o.FileCount == nil {
		var ret int32
		return ret
	}
	return *o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DownloadSummaryFileSummary) GetFileCountOk() (*int32, bool) {
	if o == nil || o.FileCount == nil {
		return nil, false
	}
	return o.FileCount, true
}

// HasFileCount returns a boolean if a field has been set.
func (o *V1DownloadSummaryFileSummary) HasFileCount() bool {
	if o != nil && o.FileCount != nil {
		return true
	}

	return false
}

// SetFileCount gets a reference to the given int32 and assigns it to the FileCount field.
func (o *V1DownloadSummaryFileSummary) SetFileCount(v int32) {
	o.FileCount = &v
}

// GetSizeMb returns the SizeMb field value if set, zero value otherwise.
func (o *V1DownloadSummaryFileSummary) GetSizeMb() float32 {
	if o == nil || o.SizeMb == nil {
		var ret float32
		return ret
	}
	return *o.SizeMb
}

// GetSizeMbOk returns a tuple with the SizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DownloadSummaryFileSummary) GetSizeMbOk() (*float32, bool) {
	if o == nil || o.SizeMb == nil {
		return nil, false
	}
	return o.SizeMb, true
}

// HasSizeMb returns a boolean if a field has been set.
func (o *V1DownloadSummaryFileSummary) HasSizeMb() bool {
	if o != nil && o.SizeMb != nil {
		return true
	}

	return false
}

// SetSizeMb gets a reference to the given float32 and assigns it to the SizeMb field.
func (o *V1DownloadSummaryFileSummary) SetSizeMb(v float32) {
	o.SizeMb = &v
}

func (o V1DownloadSummaryFileSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileCount != nil  {
		toSerialize["file_count"] = o.FileCount
	}
	if o.SizeMb != nil  {
		toSerialize["size_mb"] = o.SizeMb
	}
	return json.Marshal(toSerialize)
}

type NullableV1DownloadSummaryFileSummary struct {
	value *V1DownloadSummaryFileSummary
	isSet bool
}

func (v NullableV1DownloadSummaryFileSummary) Get() *V1DownloadSummaryFileSummary {
	return v.value
}

func (v *NullableV1DownloadSummaryFileSummary) Set(val *V1DownloadSummaryFileSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DownloadSummaryFileSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DownloadSummaryFileSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DownloadSummaryFileSummary(val *V1DownloadSummaryFileSummary) *NullableV1DownloadSummaryFileSummary {
	return &NullableV1DownloadSummaryFileSummary{value: val, isSet: true}
}

func (v NullableV1DownloadSummaryFileSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DownloadSummaryFileSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



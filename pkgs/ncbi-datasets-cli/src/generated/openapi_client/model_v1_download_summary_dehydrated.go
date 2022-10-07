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

// V1DownloadSummaryDehydrated struct for V1DownloadSummaryDehydrated
type V1DownloadSummaryDehydrated struct {
	EstimatedFileSizeMb *int32 `json:"estimated_file_size_mb,omitempty"`
	Url *string `json:"url,omitempty"`
	CliDownloadCommandLine *string `json:"cli_download_command_line,omitempty"`
	CliRehydrateCommandLine *string `json:"cli_rehydrate_command_line,omitempty"`
}

// NewV1DownloadSummaryDehydrated instantiates a new V1DownloadSummaryDehydrated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DownloadSummaryDehydrated() *V1DownloadSummaryDehydrated {
	this := V1DownloadSummaryDehydrated{}
	return &this
}

// NewV1DownloadSummaryDehydratedWithDefaults instantiates a new V1DownloadSummaryDehydrated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DownloadSummaryDehydratedWithDefaults() *V1DownloadSummaryDehydrated {
	this := V1DownloadSummaryDehydrated{}
	return &this
}

// GetEstimatedFileSizeMb returns the EstimatedFileSizeMb field value if set, zero value otherwise.
func (o *V1DownloadSummaryDehydrated) GetEstimatedFileSizeMb() int32 {
	if o == nil || o.EstimatedFileSizeMb == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedFileSizeMb
}

// GetEstimatedFileSizeMbOk returns a tuple with the EstimatedFileSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DownloadSummaryDehydrated) GetEstimatedFileSizeMbOk() (*int32, bool) {
	if o == nil || o.EstimatedFileSizeMb == nil {
		return nil, false
	}
	return o.EstimatedFileSizeMb, true
}

// HasEstimatedFileSizeMb returns a boolean if a field has been set.
func (o *V1DownloadSummaryDehydrated) HasEstimatedFileSizeMb() bool {
	if o != nil && o.EstimatedFileSizeMb != nil {
		return true
	}

	return false
}

// SetEstimatedFileSizeMb gets a reference to the given int32 and assigns it to the EstimatedFileSizeMb field.
func (o *V1DownloadSummaryDehydrated) SetEstimatedFileSizeMb(v int32) {
	o.EstimatedFileSizeMb = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *V1DownloadSummaryDehydrated) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DownloadSummaryDehydrated) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *V1DownloadSummaryDehydrated) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *V1DownloadSummaryDehydrated) SetUrl(v string) {
	o.Url = &v
}

// GetCliDownloadCommandLine returns the CliDownloadCommandLine field value if set, zero value otherwise.
func (o *V1DownloadSummaryDehydrated) GetCliDownloadCommandLine() string {
	if o == nil || o.CliDownloadCommandLine == nil {
		var ret string
		return ret
	}
	return *o.CliDownloadCommandLine
}

// GetCliDownloadCommandLineOk returns a tuple with the CliDownloadCommandLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DownloadSummaryDehydrated) GetCliDownloadCommandLineOk() (*string, bool) {
	if o == nil || o.CliDownloadCommandLine == nil {
		return nil, false
	}
	return o.CliDownloadCommandLine, true
}

// HasCliDownloadCommandLine returns a boolean if a field has been set.
func (o *V1DownloadSummaryDehydrated) HasCliDownloadCommandLine() bool {
	if o != nil && o.CliDownloadCommandLine != nil {
		return true
	}

	return false
}

// SetCliDownloadCommandLine gets a reference to the given string and assigns it to the CliDownloadCommandLine field.
func (o *V1DownloadSummaryDehydrated) SetCliDownloadCommandLine(v string) {
	o.CliDownloadCommandLine = &v
}

// GetCliRehydrateCommandLine returns the CliRehydrateCommandLine field value if set, zero value otherwise.
func (o *V1DownloadSummaryDehydrated) GetCliRehydrateCommandLine() string {
	if o == nil || o.CliRehydrateCommandLine == nil {
		var ret string
		return ret
	}
	return *o.CliRehydrateCommandLine
}

// GetCliRehydrateCommandLineOk returns a tuple with the CliRehydrateCommandLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DownloadSummaryDehydrated) GetCliRehydrateCommandLineOk() (*string, bool) {
	if o == nil || o.CliRehydrateCommandLine == nil {
		return nil, false
	}
	return o.CliRehydrateCommandLine, true
}

// HasCliRehydrateCommandLine returns a boolean if a field has been set.
func (o *V1DownloadSummaryDehydrated) HasCliRehydrateCommandLine() bool {
	if o != nil && o.CliRehydrateCommandLine != nil {
		return true
	}

	return false
}

// SetCliRehydrateCommandLine gets a reference to the given string and assigns it to the CliRehydrateCommandLine field.
func (o *V1DownloadSummaryDehydrated) SetCliRehydrateCommandLine(v string) {
	o.CliRehydrateCommandLine = &v
}

func (o V1DownloadSummaryDehydrated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EstimatedFileSizeMb != nil  {
		toSerialize["estimated_file_size_mb"] = o.EstimatedFileSizeMb
	}
	if o.Url != nil  {
		toSerialize["url"] = o.Url
	}
	if o.CliDownloadCommandLine != nil  {
		toSerialize["cli_download_command_line"] = o.CliDownloadCommandLine
	}
	if o.CliRehydrateCommandLine != nil  {
		toSerialize["cli_rehydrate_command_line"] = o.CliRehydrateCommandLine
	}
	return json.Marshal(toSerialize)
}

type NullableV1DownloadSummaryDehydrated struct {
	value *V1DownloadSummaryDehydrated
	isSet bool
}

func (v NullableV1DownloadSummaryDehydrated) Get() *V1DownloadSummaryDehydrated {
	return v.value
}

func (v *NullableV1DownloadSummaryDehydrated) Set(val *V1DownloadSummaryDehydrated) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DownloadSummaryDehydrated) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DownloadSummaryDehydrated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DownloadSummaryDehydrated(val *V1DownloadSummaryDehydrated) *NullableV1DownloadSummaryDehydrated {
	return &NullableV1DownloadSummaryDehydrated{value: val, isSet: true}
}

func (v NullableV1DownloadSummaryDehydrated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DownloadSummaryDehydrated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



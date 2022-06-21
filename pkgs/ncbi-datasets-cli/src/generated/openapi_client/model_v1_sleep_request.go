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

// V1SleepRequest struct for V1SleepRequest
type V1SleepRequest struct {
	SleepMsec *int32 `json:"sleep_msec,omitempty"`
	ErrorRate *float32 `json:"error_rate,omitempty"`
}

// NewV1SleepRequest instantiates a new V1SleepRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SleepRequest() *V1SleepRequest {
	this := V1SleepRequest{}
	return &this
}

// NewV1SleepRequestWithDefaults instantiates a new V1SleepRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SleepRequestWithDefaults() *V1SleepRequest {
	this := V1SleepRequest{}
	return &this
}

// GetSleepMsec returns the SleepMsec field value if set, zero value otherwise.
func (o *V1SleepRequest) GetSleepMsec() int32 {
	if o == nil || o.SleepMsec == nil {
		var ret int32
		return ret
	}
	return *o.SleepMsec
}

// GetSleepMsecOk returns a tuple with the SleepMsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SleepRequest) GetSleepMsecOk() (*int32, bool) {
	if o == nil || o.SleepMsec == nil {
		return nil, false
	}
	return o.SleepMsec, true
}

// HasSleepMsec returns a boolean if a field has been set.
func (o *V1SleepRequest) HasSleepMsec() bool {
	if o != nil && o.SleepMsec != nil {
		return true
	}

	return false
}

// SetSleepMsec gets a reference to the given int32 and assigns it to the SleepMsec field.
func (o *V1SleepRequest) SetSleepMsec(v int32) {
	o.SleepMsec = &v
}

// GetErrorRate returns the ErrorRate field value if set, zero value otherwise.
func (o *V1SleepRequest) GetErrorRate() float32 {
	if o == nil || o.ErrorRate == nil {
		var ret float32
		return ret
	}
	return *o.ErrorRate
}

// GetErrorRateOk returns a tuple with the ErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SleepRequest) GetErrorRateOk() (*float32, bool) {
	if o == nil || o.ErrorRate == nil {
		return nil, false
	}
	return o.ErrorRate, true
}

// HasErrorRate returns a boolean if a field has been set.
func (o *V1SleepRequest) HasErrorRate() bool {
	if o != nil && o.ErrorRate != nil {
		return true
	}

	return false
}

// SetErrorRate gets a reference to the given float32 and assigns it to the ErrorRate field.
func (o *V1SleepRequest) SetErrorRate(v float32) {
	o.ErrorRate = &v
}

func (o V1SleepRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SleepMsec != nil  {
		toSerialize["sleep_msec"] = o.SleepMsec
	}
	if o.ErrorRate != nil  {
		toSerialize["error_rate"] = o.ErrorRate
	}
	return json.Marshal(toSerialize)
}

type NullableV1SleepRequest struct {
	value *V1SleepRequest
	isSet bool
}

func (v NullableV1SleepRequest) Get() *V1SleepRequest {
	return v.value
}

func (v *NullableV1SleepRequest) Set(val *V1SleepRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SleepRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SleepRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SleepRequest(val *V1SleepRequest) *NullableV1SleepRequest {
	return &NullableV1SleepRequest{value: val, isSet: true}
}

func (v NullableV1SleepRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SleepRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



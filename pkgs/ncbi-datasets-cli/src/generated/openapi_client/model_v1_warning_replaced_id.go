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

// V1WarningReplacedId struct for V1WarningReplacedId
type V1WarningReplacedId struct {
	Requested *string `json:"requested,omitempty"`
	Returned *string `json:"returned,omitempty"`
}

// NewV1WarningReplacedId instantiates a new V1WarningReplacedId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1WarningReplacedId() *V1WarningReplacedId {
	this := V1WarningReplacedId{}
	return &this
}

// NewV1WarningReplacedIdWithDefaults instantiates a new V1WarningReplacedId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1WarningReplacedIdWithDefaults() *V1WarningReplacedId {
	this := V1WarningReplacedId{}
	return &this
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *V1WarningReplacedId) GetRequested() string {
	if o == nil || o.Requested == nil {
		var ret string
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1WarningReplacedId) GetRequestedOk() (*string, bool) {
	if o == nil || o.Requested == nil {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *V1WarningReplacedId) HasRequested() bool {
	if o != nil && o.Requested != nil {
		return true
	}

	return false
}

// SetRequested gets a reference to the given string and assigns it to the Requested field.
func (o *V1WarningReplacedId) SetRequested(v string) {
	o.Requested = &v
}

// GetReturned returns the Returned field value if set, zero value otherwise.
func (o *V1WarningReplacedId) GetReturned() string {
	if o == nil || o.Returned == nil {
		var ret string
		return ret
	}
	return *o.Returned
}

// GetReturnedOk returns a tuple with the Returned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1WarningReplacedId) GetReturnedOk() (*string, bool) {
	if o == nil || o.Returned == nil {
		return nil, false
	}
	return o.Returned, true
}

// HasReturned returns a boolean if a field has been set.
func (o *V1WarningReplacedId) HasReturned() bool {
	if o != nil && o.Returned != nil {
		return true
	}

	return false
}

// SetReturned gets a reference to the given string and assigns it to the Returned field.
func (o *V1WarningReplacedId) SetReturned(v string) {
	o.Returned = &v
}

func (o V1WarningReplacedId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Requested != nil  {
		toSerialize["requested"] = o.Requested
	}
	if o.Returned != nil  {
		toSerialize["returned"] = o.Returned
	}
	return json.Marshal(toSerialize)
}

type NullableV1WarningReplacedId struct {
	value *V1WarningReplacedId
	isSet bool
}

func (v NullableV1WarningReplacedId) Get() *V1WarningReplacedId {
	return v.value
}

func (v *NullableV1WarningReplacedId) Set(val *V1WarningReplacedId) {
	v.value = val
	v.isSet = true
}

func (v NullableV1WarningReplacedId) IsSet() bool {
	return v.isSet
}

func (v *NullableV1WarningReplacedId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1WarningReplacedId(val *V1WarningReplacedId) *NullableV1WarningReplacedId {
	return &NullableV1WarningReplacedId{value: val, isSet: true}
}

func (v NullableV1WarningReplacedId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1WarningReplacedId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



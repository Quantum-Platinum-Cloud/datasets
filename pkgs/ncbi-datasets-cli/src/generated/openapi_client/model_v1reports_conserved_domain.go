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

// V1reportsConservedDomain struct for V1reportsConservedDomain
type V1reportsConservedDomain struct {
	Accession *string `json:"accession,omitempty"`
	Name *string `json:"name,omitempty"`
	Range *V1reportsRange `json:"range,omitempty"`
}

// NewV1reportsConservedDomain instantiates a new V1reportsConservedDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsConservedDomain() *V1reportsConservedDomain {
	this := V1reportsConservedDomain{}
	return &this
}

// NewV1reportsConservedDomainWithDefaults instantiates a new V1reportsConservedDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsConservedDomainWithDefaults() *V1reportsConservedDomain {
	this := V1reportsConservedDomain{}
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V1reportsConservedDomain) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsConservedDomain) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V1reportsConservedDomain) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V1reportsConservedDomain) SetAccession(v string) {
	o.Accession = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1reportsConservedDomain) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsConservedDomain) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1reportsConservedDomain) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1reportsConservedDomain) SetName(v string) {
	o.Name = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *V1reportsConservedDomain) GetRange() V1reportsRange {
	if o == nil || o.Range == nil {
		var ret V1reportsRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsConservedDomain) GetRangeOk() (*V1reportsRange, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *V1reportsConservedDomain) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given V1reportsRange and assigns it to the Range field.
func (o *V1reportsConservedDomain) SetRange(v V1reportsRange) {
	o.Range = &v
}

func (o V1reportsConservedDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil  {
		toSerialize["accession"] = o.Accession
	}
	if o.Name != nil  {
		toSerialize["name"] = o.Name
	}
	if o.Range != nil  {
		toSerialize["range"] = o.Range
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsConservedDomain struct {
	value *V1reportsConservedDomain
	isSet bool
}

func (v NullableV1reportsConservedDomain) Get() *V1reportsConservedDomain {
	return v.value
}

func (v *NullableV1reportsConservedDomain) Set(val *V1reportsConservedDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsConservedDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsConservedDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsConservedDomain(val *V1reportsConservedDomain) *NullableV1reportsConservedDomain {
	return &NullableV1reportsConservedDomain{value: val, isSet: true}
}

func (v NullableV1reportsConservedDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsConservedDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



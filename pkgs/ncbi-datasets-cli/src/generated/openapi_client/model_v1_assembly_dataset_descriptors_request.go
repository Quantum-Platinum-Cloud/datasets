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

// V1AssemblyDatasetDescriptorsRequest struct for V1AssemblyDatasetDescriptorsRequest
type V1AssemblyDatasetDescriptorsRequest struct {
	TaxId *int32 `json:"tax_id,omitempty"`
	OrgName *string `json:"org_name,omitempty"`
	TaxName *string `json:"tax_name,omitempty"`
	AssemblyAccession *string `json:"assembly_accession,omitempty"`
	Cutoff *int32 `json:"cutoff,omitempty"`
	Limit *string `json:"limit,omitempty"`
	Filters *V1AssemblyDatasetDescriptorsFilter `json:"filters,omitempty"`
	TaxExactMatch *bool `json:"tax_exact_match,omitempty"`
	ReturnedContent *V1AssemblyDatasetDescriptorsRequestContentType `json:"returned_content,omitempty"`
}

// NewV1AssemblyDatasetDescriptorsRequest instantiates a new V1AssemblyDatasetDescriptorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyDatasetDescriptorsRequest() *V1AssemblyDatasetDescriptorsRequest {
	this := V1AssemblyDatasetDescriptorsRequest{}
	var returnedContent V1AssemblyDatasetDescriptorsRequestContentType = V1ASSEMBLYDATASETDESCRIPTORSREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// NewV1AssemblyDatasetDescriptorsRequestWithDefaults instantiates a new V1AssemblyDatasetDescriptorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyDatasetDescriptorsRequestWithDefaults() *V1AssemblyDatasetDescriptorsRequest {
	this := V1AssemblyDatasetDescriptorsRequest{}
	var returnedContent V1AssemblyDatasetDescriptorsRequestContentType = V1ASSEMBLYDATASETDESCRIPTORSREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxId() int32 {
	if o == nil || o.TaxId == nil {
		var ret int32
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxIdOk() (*int32, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given int32 and assigns it to the TaxId field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetTaxId(v int32) {
	o.TaxId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasOrgName() bool {
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetOrgName(v string) {
	o.OrgName = &v
}

// GetTaxName returns the TaxName field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxName() string {
	if o == nil || o.TaxName == nil {
		var ret string
		return ret
	}
	return *o.TaxName
}

// GetTaxNameOk returns a tuple with the TaxName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxNameOk() (*string, bool) {
	if o == nil || o.TaxName == nil {
		return nil, false
	}
	return o.TaxName, true
}

// HasTaxName returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasTaxName() bool {
	if o != nil && o.TaxName != nil {
		return true
	}

	return false
}

// SetTaxName gets a reference to the given string and assigns it to the TaxName field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetTaxName(v string) {
	o.TaxName = &v
}

// GetAssemblyAccession returns the AssemblyAccession field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetAssemblyAccession() string {
	if o == nil || o.AssemblyAccession == nil {
		var ret string
		return ret
	}
	return *o.AssemblyAccession
}

// GetAssemblyAccessionOk returns a tuple with the AssemblyAccession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetAssemblyAccessionOk() (*string, bool) {
	if o == nil || o.AssemblyAccession == nil {
		return nil, false
	}
	return o.AssemblyAccession, true
}

// HasAssemblyAccession returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasAssemblyAccession() bool {
	if o != nil && o.AssemblyAccession != nil {
		return true
	}

	return false
}

// SetAssemblyAccession gets a reference to the given string and assigns it to the AssemblyAccession field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetAssemblyAccession(v string) {
	o.AssemblyAccession = &v
}

// GetCutoff returns the Cutoff field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetCutoff() int32 {
	if o == nil || o.Cutoff == nil {
		var ret int32
		return ret
	}
	return *o.Cutoff
}

// GetCutoffOk returns a tuple with the Cutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetCutoffOk() (*int32, bool) {
	if o == nil || o.Cutoff == nil {
		return nil, false
	}
	return o.Cutoff, true
}

// HasCutoff returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasCutoff() bool {
	if o != nil && o.Cutoff != nil {
		return true
	}

	return false
}

// SetCutoff gets a reference to the given int32 and assigns it to the Cutoff field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetCutoff(v int32) {
	o.Cutoff = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetLimit() string {
	if o == nil || o.Limit == nil {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetLimitOk() (*string, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetLimit(v string) {
	o.Limit = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetFilters() V1AssemblyDatasetDescriptorsFilter {
	if o == nil || o.Filters == nil {
		var ret V1AssemblyDatasetDescriptorsFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetFiltersOk() (*V1AssemblyDatasetDescriptorsFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given V1AssemblyDatasetDescriptorsFilter and assigns it to the Filters field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetFilters(v V1AssemblyDatasetDescriptorsFilter) {
	o.Filters = &v
}

// GetTaxExactMatch returns the TaxExactMatch field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxExactMatch() bool {
	if o == nil || o.TaxExactMatch == nil {
		var ret bool
		return ret
	}
	return *o.TaxExactMatch
}

// GetTaxExactMatchOk returns a tuple with the TaxExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxExactMatchOk() (*bool, bool) {
	if o == nil || o.TaxExactMatch == nil {
		return nil, false
	}
	return o.TaxExactMatch, true
}

// HasTaxExactMatch returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasTaxExactMatch() bool {
	if o != nil && o.TaxExactMatch != nil {
		return true
	}

	return false
}

// SetTaxExactMatch gets a reference to the given bool and assigns it to the TaxExactMatch field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetTaxExactMatch(v bool) {
	o.TaxExactMatch = &v
}

// GetReturnedContent returns the ReturnedContent field value if set, zero value otherwise.
func (o *V1AssemblyDatasetDescriptorsRequest) GetReturnedContent() V1AssemblyDatasetDescriptorsRequestContentType {
	if o == nil || o.ReturnedContent == nil {
		var ret V1AssemblyDatasetDescriptorsRequestContentType
		return ret
	}
	return *o.ReturnedContent
}

// GetReturnedContentOk returns a tuple with the ReturnedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) GetReturnedContentOk() (*V1AssemblyDatasetDescriptorsRequestContentType, bool) {
	if o == nil || o.ReturnedContent == nil {
		return nil, false
	}
	return o.ReturnedContent, true
}

// HasReturnedContent returns a boolean if a field has been set.
func (o *V1AssemblyDatasetDescriptorsRequest) HasReturnedContent() bool {
	if o != nil && o.ReturnedContent != nil {
		return true
	}

	return false
}

// SetReturnedContent gets a reference to the given V1AssemblyDatasetDescriptorsRequestContentType and assigns it to the ReturnedContent field.
func (o *V1AssemblyDatasetDescriptorsRequest) SetReturnedContent(v V1AssemblyDatasetDescriptorsRequestContentType) {
	o.ReturnedContent = &v
}

func (o V1AssemblyDatasetDescriptorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.OrgName != nil  {
		toSerialize["org_name"] = o.OrgName
	}
	if o.TaxName != nil  {
		toSerialize["tax_name"] = o.TaxName
	}
	if o.AssemblyAccession != nil  {
		toSerialize["assembly_accession"] = o.AssemblyAccession
	}
	if o.Cutoff != nil  {
		toSerialize["cutoff"] = o.Cutoff
	}
	if o.Limit != nil  {
		toSerialize["limit"] = o.Limit
	}
	if o.Filters != nil  {
		toSerialize["filters"] = o.Filters
	}
	if o.TaxExactMatch != nil  {
		toSerialize["tax_exact_match"] = o.TaxExactMatch
	}
	if o.ReturnedContent != nil  {
		toSerialize["returned_content"] = o.ReturnedContent
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyDatasetDescriptorsRequest struct {
	value *V1AssemblyDatasetDescriptorsRequest
	isSet bool
}

func (v NullableV1AssemblyDatasetDescriptorsRequest) Get() *V1AssemblyDatasetDescriptorsRequest {
	return v.value
}

func (v *NullableV1AssemblyDatasetDescriptorsRequest) Set(val *V1AssemblyDatasetDescriptorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetDescriptorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetDescriptorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetDescriptorsRequest(val *V1AssemblyDatasetDescriptorsRequest) *NullableV1AssemblyDatasetDescriptorsRequest {
	return &NullableV1AssemblyDatasetDescriptorsRequest{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetDescriptorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetDescriptorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



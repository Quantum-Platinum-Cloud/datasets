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

// V1reportsBioSampleDescriptor struct for V1reportsBioSampleDescriptor
type V1reportsBioSampleDescriptor struct {
	Accession *string `json:"accession,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	PublicationDate *string `json:"publication_date,omitempty"`
	SubmissionDate *string `json:"submission_date,omitempty"`
	SampleIds *[]V1reportsBioSampleId `json:"sample_ids,omitempty"`
	Description *V1reportsBioSampleDescription `json:"description,omitempty"`
	Owner *V1reportsBioSampleOwner `json:"owner,omitempty"`
	Models *[]string `json:"models,omitempty"`
	Bioprojects *[]V1reportsBioProject `json:"bioprojects,omitempty"`
	Package *string `json:"package,omitempty"`
	Attributes *[]V1reportsBioSampleAttribute `json:"attributes,omitempty"`
	Status *V1reportsBioSampleStatus `json:"status,omitempty"`
}

// NewV1reportsBioSampleDescriptor instantiates a new V1reportsBioSampleDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsBioSampleDescriptor() *V1reportsBioSampleDescriptor {
	this := V1reportsBioSampleDescriptor{}
	return &this
}

// NewV1reportsBioSampleDescriptorWithDefaults instantiates a new V1reportsBioSampleDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsBioSampleDescriptorWithDefaults() *V1reportsBioSampleDescriptor {
	this := V1reportsBioSampleDescriptor{}
	return &this
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V1reportsBioSampleDescriptor) SetAccession(v string) {
	o.Accession = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *V1reportsBioSampleDescriptor) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetPublicationDate returns the PublicationDate field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetPublicationDate() string {
	if o == nil || o.PublicationDate == nil {
		var ret string
		return ret
	}
	return *o.PublicationDate
}

// GetPublicationDateOk returns a tuple with the PublicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetPublicationDateOk() (*string, bool) {
	if o == nil || o.PublicationDate == nil {
		return nil, false
	}
	return o.PublicationDate, true
}

// HasPublicationDate returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasPublicationDate() bool {
	if o != nil && o.PublicationDate != nil {
		return true
	}

	return false
}

// SetPublicationDate gets a reference to the given string and assigns it to the PublicationDate field.
func (o *V1reportsBioSampleDescriptor) SetPublicationDate(v string) {
	o.PublicationDate = &v
}

// GetSubmissionDate returns the SubmissionDate field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetSubmissionDate() string {
	if o == nil || o.SubmissionDate == nil {
		var ret string
		return ret
	}
	return *o.SubmissionDate
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetSubmissionDateOk() (*string, bool) {
	if o == nil || o.SubmissionDate == nil {
		return nil, false
	}
	return o.SubmissionDate, true
}

// HasSubmissionDate returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasSubmissionDate() bool {
	if o != nil && o.SubmissionDate != nil {
		return true
	}

	return false
}

// SetSubmissionDate gets a reference to the given string and assigns it to the SubmissionDate field.
func (o *V1reportsBioSampleDescriptor) SetSubmissionDate(v string) {
	o.SubmissionDate = &v
}

// GetSampleIds returns the SampleIds field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetSampleIds() []V1reportsBioSampleId {
	if o == nil || o.SampleIds == nil {
		var ret []V1reportsBioSampleId
		return ret
	}
	return *o.SampleIds
}

// GetSampleIdsOk returns a tuple with the SampleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetSampleIdsOk() (*[]V1reportsBioSampleId, bool) {
	if o == nil || o.SampleIds == nil {
		return nil, false
	}
	return o.SampleIds, true
}

// HasSampleIds returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasSampleIds() bool {
	if o != nil && o.SampleIds != nil {
		return true
	}

	return false
}

// SetSampleIds gets a reference to the given []V1reportsBioSampleId and assigns it to the SampleIds field.
func (o *V1reportsBioSampleDescriptor) SetSampleIds(v []V1reportsBioSampleId) {
	o.SampleIds = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetDescription() V1reportsBioSampleDescription {
	if o == nil || o.Description == nil {
		var ret V1reportsBioSampleDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetDescriptionOk() (*V1reportsBioSampleDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given V1reportsBioSampleDescription and assigns it to the Description field.
func (o *V1reportsBioSampleDescriptor) SetDescription(v V1reportsBioSampleDescription) {
	o.Description = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetOwner() V1reportsBioSampleOwner {
	if o == nil || o.Owner == nil {
		var ret V1reportsBioSampleOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetOwnerOk() (*V1reportsBioSampleOwner, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given V1reportsBioSampleOwner and assigns it to the Owner field.
func (o *V1reportsBioSampleDescriptor) SetOwner(v V1reportsBioSampleOwner) {
	o.Owner = &v
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetModels() []string {
	if o == nil || o.Models == nil {
		var ret []string
		return ret
	}
	return *o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetModelsOk() (*[]string, bool) {
	if o == nil || o.Models == nil {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasModels() bool {
	if o != nil && o.Models != nil {
		return true
	}

	return false
}

// SetModels gets a reference to the given []string and assigns it to the Models field.
func (o *V1reportsBioSampleDescriptor) SetModels(v []string) {
	o.Models = &v
}

// GetBioprojects returns the Bioprojects field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetBioprojects() []V1reportsBioProject {
	if o == nil || o.Bioprojects == nil {
		var ret []V1reportsBioProject
		return ret
	}
	return *o.Bioprojects
}

// GetBioprojectsOk returns a tuple with the Bioprojects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetBioprojectsOk() (*[]V1reportsBioProject, bool) {
	if o == nil || o.Bioprojects == nil {
		return nil, false
	}
	return o.Bioprojects, true
}

// HasBioprojects returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasBioprojects() bool {
	if o != nil && o.Bioprojects != nil {
		return true
	}

	return false
}

// SetBioprojects gets a reference to the given []V1reportsBioProject and assigns it to the Bioprojects field.
func (o *V1reportsBioSampleDescriptor) SetBioprojects(v []V1reportsBioProject) {
	o.Bioprojects = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *V1reportsBioSampleDescriptor) SetPackage(v string) {
	o.Package = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetAttributes() []V1reportsBioSampleAttribute {
	if o == nil || o.Attributes == nil {
		var ret []V1reportsBioSampleAttribute
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetAttributesOk() (*[]V1reportsBioSampleAttribute, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []V1reportsBioSampleAttribute and assigns it to the Attributes field.
func (o *V1reportsBioSampleDescriptor) SetAttributes(v []V1reportsBioSampleAttribute) {
	o.Attributes = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescriptor) GetStatus() V1reportsBioSampleStatus {
	if o == nil || o.Status == nil {
		var ret V1reportsBioSampleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescriptor) GetStatusOk() (*V1reportsBioSampleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescriptor) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V1reportsBioSampleStatus and assigns it to the Status field.
func (o *V1reportsBioSampleDescriptor) SetStatus(v V1reportsBioSampleStatus) {
	o.Status = &v
}

func (o V1reportsBioSampleDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accession != nil  {
		toSerialize["accession"] = o.Accession
	}
	if o.LastUpdated != nil  {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.PublicationDate != nil  {
		toSerialize["publication_date"] = o.PublicationDate
	}
	if o.SubmissionDate != nil  {
		toSerialize["submission_date"] = o.SubmissionDate
	}
	if o.SampleIds != nil && len(o.GetSampleIds()) > 0  {
		toSerialize["sample_ids"] = o.SampleIds
	}
	if o.Description != nil  {
		toSerialize["description"] = o.Description
	}
	if o.Owner != nil  {
		toSerialize["owner"] = o.Owner
	}
	if o.Models != nil && len(o.GetModels()) > 0  {
		toSerialize["models"] = o.Models
	}
	if o.Bioprojects != nil && len(o.GetBioprojects()) > 0  {
		toSerialize["bioprojects"] = o.Bioprojects
	}
	if o.Package != nil  {
		toSerialize["package"] = o.Package
	}
	if o.Attributes != nil && len(o.GetAttributes()) > 0  {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Status != nil  {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsBioSampleDescriptor struct {
	value *V1reportsBioSampleDescriptor
	isSet bool
}

func (v NullableV1reportsBioSampleDescriptor) Get() *V1reportsBioSampleDescriptor {
	return v.value
}

func (v *NullableV1reportsBioSampleDescriptor) Set(val *V1reportsBioSampleDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsBioSampleDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsBioSampleDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsBioSampleDescriptor(val *V1reportsBioSampleDescriptor) *NullableV1reportsBioSampleDescriptor {
	return &NullableV1reportsBioSampleDescriptor{value: val, isSet: true}
}

func (v NullableV1reportsBioSampleDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsBioSampleDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



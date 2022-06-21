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

// V1Organism struct for V1Organism
type V1Organism struct {
	TaxId *string `json:"tax_id,omitempty"`
	SciName *string `json:"sci_name,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	BlastNode *bool `json:"blast_node,omitempty"`
	Breed *string `json:"breed,omitempty"`
	Cultivar *string `json:"cultivar,omitempty"`
	Ecotype *string `json:"ecotype,omitempty"`
	Isolate *string `json:"isolate,omitempty"`
	Sex *string `json:"sex,omitempty"`
	Strain *string `json:"strain,omitempty"`
	SearchText *[]string `json:"search_text,omitempty"`
	Rank *V1OrganismRankType `json:"rank,omitempty"`
	ParentTaxId *string `json:"parent_tax_id,omitempty"`
	AssemblyCount *string `json:"assembly_count,omitempty"`
	AssemblyCounts *V1OrganismCounts `json:"assembly_counts,omitempty"`
	Counts *[]V1OrganismCountByType `json:"counts,omitempty"`
	Children *[]V1Organism `json:"children,omitempty"`
	Merged *[]V1Organism `json:"merged,omitempty"`
	MergedTaxIds *[]string `json:"merged_tax_ids,omitempty"`
	MinOrd *int32 `json:"min_ord,omitempty"`
	MaxOrd *int32 `json:"max_ord,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Key *string `json:"key,omitempty"`
	Title *string `json:"title,omitempty"`
	Icon *bool `json:"icon,omitempty"`
}

// NewV1Organism instantiates a new V1Organism object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Organism() *V1Organism {
	this := V1Organism{}
	var rank V1OrganismRankType = V1ORGANISMRANKTYPE_NO_RANK
	this.Rank = &rank
	return &this
}

// NewV1OrganismWithDefaults instantiates a new V1Organism object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1OrganismWithDefaults() *V1Organism {
	this := V1Organism{}
	var rank V1OrganismRankType = V1ORGANISMRANKTYPE_NO_RANK
	this.Rank = &rank
	return &this
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1Organism) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1Organism) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *V1Organism) SetTaxId(v string) {
	o.TaxId = &v
}

// GetSciName returns the SciName field value if set, zero value otherwise.
func (o *V1Organism) GetSciName() string {
	if o == nil || o.SciName == nil {
		var ret string
		return ret
	}
	return *o.SciName
}

// GetSciNameOk returns a tuple with the SciName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetSciNameOk() (*string, bool) {
	if o == nil || o.SciName == nil {
		return nil, false
	}
	return o.SciName, true
}

// HasSciName returns a boolean if a field has been set.
func (o *V1Organism) HasSciName() bool {
	if o != nil && o.SciName != nil {
		return true
	}

	return false
}

// SetSciName gets a reference to the given string and assigns it to the SciName field.
func (o *V1Organism) SetSciName(v string) {
	o.SciName = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V1Organism) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V1Organism) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V1Organism) SetCommonName(v string) {
	o.CommonName = &v
}

// GetBlastNode returns the BlastNode field value if set, zero value otherwise.
func (o *V1Organism) GetBlastNode() bool {
	if o == nil || o.BlastNode == nil {
		var ret bool
		return ret
	}
	return *o.BlastNode
}

// GetBlastNodeOk returns a tuple with the BlastNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetBlastNodeOk() (*bool, bool) {
	if o == nil || o.BlastNode == nil {
		return nil, false
	}
	return o.BlastNode, true
}

// HasBlastNode returns a boolean if a field has been set.
func (o *V1Organism) HasBlastNode() bool {
	if o != nil && o.BlastNode != nil {
		return true
	}

	return false
}

// SetBlastNode gets a reference to the given bool and assigns it to the BlastNode field.
func (o *V1Organism) SetBlastNode(v bool) {
	o.BlastNode = &v
}

// GetBreed returns the Breed field value if set, zero value otherwise.
func (o *V1Organism) GetBreed() string {
	if o == nil || o.Breed == nil {
		var ret string
		return ret
	}
	return *o.Breed
}

// GetBreedOk returns a tuple with the Breed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetBreedOk() (*string, bool) {
	if o == nil || o.Breed == nil {
		return nil, false
	}
	return o.Breed, true
}

// HasBreed returns a boolean if a field has been set.
func (o *V1Organism) HasBreed() bool {
	if o != nil && o.Breed != nil {
		return true
	}

	return false
}

// SetBreed gets a reference to the given string and assigns it to the Breed field.
func (o *V1Organism) SetBreed(v string) {
	o.Breed = &v
}

// GetCultivar returns the Cultivar field value if set, zero value otherwise.
func (o *V1Organism) GetCultivar() string {
	if o == nil || o.Cultivar == nil {
		var ret string
		return ret
	}
	return *o.Cultivar
}

// GetCultivarOk returns a tuple with the Cultivar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetCultivarOk() (*string, bool) {
	if o == nil || o.Cultivar == nil {
		return nil, false
	}
	return o.Cultivar, true
}

// HasCultivar returns a boolean if a field has been set.
func (o *V1Organism) HasCultivar() bool {
	if o != nil && o.Cultivar != nil {
		return true
	}

	return false
}

// SetCultivar gets a reference to the given string and assigns it to the Cultivar field.
func (o *V1Organism) SetCultivar(v string) {
	o.Cultivar = &v
}

// GetEcotype returns the Ecotype field value if set, zero value otherwise.
func (o *V1Organism) GetEcotype() string {
	if o == nil || o.Ecotype == nil {
		var ret string
		return ret
	}
	return *o.Ecotype
}

// GetEcotypeOk returns a tuple with the Ecotype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetEcotypeOk() (*string, bool) {
	if o == nil || o.Ecotype == nil {
		return nil, false
	}
	return o.Ecotype, true
}

// HasEcotype returns a boolean if a field has been set.
func (o *V1Organism) HasEcotype() bool {
	if o != nil && o.Ecotype != nil {
		return true
	}

	return false
}

// SetEcotype gets a reference to the given string and assigns it to the Ecotype field.
func (o *V1Organism) SetEcotype(v string) {
	o.Ecotype = &v
}

// GetIsolate returns the Isolate field value if set, zero value otherwise.
func (o *V1Organism) GetIsolate() string {
	if o == nil || o.Isolate == nil {
		var ret string
		return ret
	}
	return *o.Isolate
}

// GetIsolateOk returns a tuple with the Isolate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetIsolateOk() (*string, bool) {
	if o == nil || o.Isolate == nil {
		return nil, false
	}
	return o.Isolate, true
}

// HasIsolate returns a boolean if a field has been set.
func (o *V1Organism) HasIsolate() bool {
	if o != nil && o.Isolate != nil {
		return true
	}

	return false
}

// SetIsolate gets a reference to the given string and assigns it to the Isolate field.
func (o *V1Organism) SetIsolate(v string) {
	o.Isolate = &v
}

// GetSex returns the Sex field value if set, zero value otherwise.
func (o *V1Organism) GetSex() string {
	if o == nil || o.Sex == nil {
		var ret string
		return ret
	}
	return *o.Sex
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetSexOk() (*string, bool) {
	if o == nil || o.Sex == nil {
		return nil, false
	}
	return o.Sex, true
}

// HasSex returns a boolean if a field has been set.
func (o *V1Organism) HasSex() bool {
	if o != nil && o.Sex != nil {
		return true
	}

	return false
}

// SetSex gets a reference to the given string and assigns it to the Sex field.
func (o *V1Organism) SetSex(v string) {
	o.Sex = &v
}

// GetStrain returns the Strain field value if set, zero value otherwise.
func (o *V1Organism) GetStrain() string {
	if o == nil || o.Strain == nil {
		var ret string
		return ret
	}
	return *o.Strain
}

// GetStrainOk returns a tuple with the Strain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetStrainOk() (*string, bool) {
	if o == nil || o.Strain == nil {
		return nil, false
	}
	return o.Strain, true
}

// HasStrain returns a boolean if a field has been set.
func (o *V1Organism) HasStrain() bool {
	if o != nil && o.Strain != nil {
		return true
	}

	return false
}

// SetStrain gets a reference to the given string and assigns it to the Strain field.
func (o *V1Organism) SetStrain(v string) {
	o.Strain = &v
}

// GetSearchText returns the SearchText field value if set, zero value otherwise.
func (o *V1Organism) GetSearchText() []string {
	if o == nil || o.SearchText == nil {
		var ret []string
		return ret
	}
	return *o.SearchText
}

// GetSearchTextOk returns a tuple with the SearchText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetSearchTextOk() (*[]string, bool) {
	if o == nil || o.SearchText == nil {
		return nil, false
	}
	return o.SearchText, true
}

// HasSearchText returns a boolean if a field has been set.
func (o *V1Organism) HasSearchText() bool {
	if o != nil && o.SearchText != nil {
		return true
	}

	return false
}

// SetSearchText gets a reference to the given []string and assigns it to the SearchText field.
func (o *V1Organism) SetSearchText(v []string) {
	o.SearchText = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *V1Organism) GetRank() V1OrganismRankType {
	if o == nil || o.Rank == nil {
		var ret V1OrganismRankType
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetRankOk() (*V1OrganismRankType, bool) {
	if o == nil || o.Rank == nil {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *V1Organism) HasRank() bool {
	if o != nil && o.Rank != nil {
		return true
	}

	return false
}

// SetRank gets a reference to the given V1OrganismRankType and assigns it to the Rank field.
func (o *V1Organism) SetRank(v V1OrganismRankType) {
	o.Rank = &v
}

// GetParentTaxId returns the ParentTaxId field value if set, zero value otherwise.
func (o *V1Organism) GetParentTaxId() string {
	if o == nil || o.ParentTaxId == nil {
		var ret string
		return ret
	}
	return *o.ParentTaxId
}

// GetParentTaxIdOk returns a tuple with the ParentTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetParentTaxIdOk() (*string, bool) {
	if o == nil || o.ParentTaxId == nil {
		return nil, false
	}
	return o.ParentTaxId, true
}

// HasParentTaxId returns a boolean if a field has been set.
func (o *V1Organism) HasParentTaxId() bool {
	if o != nil && o.ParentTaxId != nil {
		return true
	}

	return false
}

// SetParentTaxId gets a reference to the given string and assigns it to the ParentTaxId field.
func (o *V1Organism) SetParentTaxId(v string) {
	o.ParentTaxId = &v
}

// GetAssemblyCount returns the AssemblyCount field value if set, zero value otherwise.
func (o *V1Organism) GetAssemblyCount() string {
	if o == nil || o.AssemblyCount == nil {
		var ret string
		return ret
	}
	return *o.AssemblyCount
}

// GetAssemblyCountOk returns a tuple with the AssemblyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetAssemblyCountOk() (*string, bool) {
	if o == nil || o.AssemblyCount == nil {
		return nil, false
	}
	return o.AssemblyCount, true
}

// HasAssemblyCount returns a boolean if a field has been set.
func (o *V1Organism) HasAssemblyCount() bool {
	if o != nil && o.AssemblyCount != nil {
		return true
	}

	return false
}

// SetAssemblyCount gets a reference to the given string and assigns it to the AssemblyCount field.
func (o *V1Organism) SetAssemblyCount(v string) {
	o.AssemblyCount = &v
}

// GetAssemblyCounts returns the AssemblyCounts field value if set, zero value otherwise.
func (o *V1Organism) GetAssemblyCounts() V1OrganismCounts {
	if o == nil || o.AssemblyCounts == nil {
		var ret V1OrganismCounts
		return ret
	}
	return *o.AssemblyCounts
}

// GetAssemblyCountsOk returns a tuple with the AssemblyCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetAssemblyCountsOk() (*V1OrganismCounts, bool) {
	if o == nil || o.AssemblyCounts == nil {
		return nil, false
	}
	return o.AssemblyCounts, true
}

// HasAssemblyCounts returns a boolean if a field has been set.
func (o *V1Organism) HasAssemblyCounts() bool {
	if o != nil && o.AssemblyCounts != nil {
		return true
	}

	return false
}

// SetAssemblyCounts gets a reference to the given V1OrganismCounts and assigns it to the AssemblyCounts field.
func (o *V1Organism) SetAssemblyCounts(v V1OrganismCounts) {
	o.AssemblyCounts = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *V1Organism) GetCounts() []V1OrganismCountByType {
	if o == nil || o.Counts == nil {
		var ret []V1OrganismCountByType
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetCountsOk() (*[]V1OrganismCountByType, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *V1Organism) HasCounts() bool {
	if o != nil && o.Counts != nil {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []V1OrganismCountByType and assigns it to the Counts field.
func (o *V1Organism) SetCounts(v []V1OrganismCountByType) {
	o.Counts = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *V1Organism) GetChildren() []V1Organism {
	if o == nil || o.Children == nil {
		var ret []V1Organism
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetChildrenOk() (*[]V1Organism, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *V1Organism) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []V1Organism and assigns it to the Children field.
func (o *V1Organism) SetChildren(v []V1Organism) {
	o.Children = &v
}

// GetMerged returns the Merged field value if set, zero value otherwise.
func (o *V1Organism) GetMerged() []V1Organism {
	if o == nil || o.Merged == nil {
		var ret []V1Organism
		return ret
	}
	return *o.Merged
}

// GetMergedOk returns a tuple with the Merged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetMergedOk() (*[]V1Organism, bool) {
	if o == nil || o.Merged == nil {
		return nil, false
	}
	return o.Merged, true
}

// HasMerged returns a boolean if a field has been set.
func (o *V1Organism) HasMerged() bool {
	if o != nil && o.Merged != nil {
		return true
	}

	return false
}

// SetMerged gets a reference to the given []V1Organism and assigns it to the Merged field.
func (o *V1Organism) SetMerged(v []V1Organism) {
	o.Merged = &v
}

// GetMergedTaxIds returns the MergedTaxIds field value if set, zero value otherwise.
func (o *V1Organism) GetMergedTaxIds() []string {
	if o == nil || o.MergedTaxIds == nil {
		var ret []string
		return ret
	}
	return *o.MergedTaxIds
}

// GetMergedTaxIdsOk returns a tuple with the MergedTaxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetMergedTaxIdsOk() (*[]string, bool) {
	if o == nil || o.MergedTaxIds == nil {
		return nil, false
	}
	return o.MergedTaxIds, true
}

// HasMergedTaxIds returns a boolean if a field has been set.
func (o *V1Organism) HasMergedTaxIds() bool {
	if o != nil && o.MergedTaxIds != nil {
		return true
	}

	return false
}

// SetMergedTaxIds gets a reference to the given []string and assigns it to the MergedTaxIds field.
func (o *V1Organism) SetMergedTaxIds(v []string) {
	o.MergedTaxIds = &v
}

// GetMinOrd returns the MinOrd field value if set, zero value otherwise.
func (o *V1Organism) GetMinOrd() int32 {
	if o == nil || o.MinOrd == nil {
		var ret int32
		return ret
	}
	return *o.MinOrd
}

// GetMinOrdOk returns a tuple with the MinOrd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetMinOrdOk() (*int32, bool) {
	if o == nil || o.MinOrd == nil {
		return nil, false
	}
	return o.MinOrd, true
}

// HasMinOrd returns a boolean if a field has been set.
func (o *V1Organism) HasMinOrd() bool {
	if o != nil && o.MinOrd != nil {
		return true
	}

	return false
}

// SetMinOrd gets a reference to the given int32 and assigns it to the MinOrd field.
func (o *V1Organism) SetMinOrd(v int32) {
	o.MinOrd = &v
}

// GetMaxOrd returns the MaxOrd field value if set, zero value otherwise.
func (o *V1Organism) GetMaxOrd() int32 {
	if o == nil || o.MaxOrd == nil {
		var ret int32
		return ret
	}
	return *o.MaxOrd
}

// GetMaxOrdOk returns a tuple with the MaxOrd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetMaxOrdOk() (*int32, bool) {
	if o == nil || o.MaxOrd == nil {
		return nil, false
	}
	return o.MaxOrd, true
}

// HasMaxOrd returns a boolean if a field has been set.
func (o *V1Organism) HasMaxOrd() bool {
	if o != nil && o.MaxOrd != nil {
		return true
	}

	return false
}

// SetMaxOrd gets a reference to the given int32 and assigns it to the MaxOrd field.
func (o *V1Organism) SetMaxOrd(v int32) {
	o.MaxOrd = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *V1Organism) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *V1Organism) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *V1Organism) SetWeight(v int32) {
	o.Weight = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *V1Organism) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *V1Organism) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *V1Organism) SetKey(v string) {
	o.Key = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V1Organism) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V1Organism) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V1Organism) SetTitle(v string) {
	o.Title = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *V1Organism) GetIcon() bool {
	if o == nil || o.Icon == nil {
		var ret bool
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Organism) GetIconOk() (*bool, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *V1Organism) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given bool and assigns it to the Icon field.
func (o *V1Organism) SetIcon(v bool) {
	o.Icon = &v
}

func (o V1Organism) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.SciName != nil  {
		toSerialize["sci_name"] = o.SciName
	}
	if o.CommonName != nil  {
		toSerialize["common_name"] = o.CommonName
	}
	if o.BlastNode != nil  {
		toSerialize["blast_node"] = o.BlastNode
	}
	if o.Breed != nil  {
		toSerialize["breed"] = o.Breed
	}
	if o.Cultivar != nil  {
		toSerialize["cultivar"] = o.Cultivar
	}
	if o.Ecotype != nil  {
		toSerialize["ecotype"] = o.Ecotype
	}
	if o.Isolate != nil  {
		toSerialize["isolate"] = o.Isolate
	}
	if o.Sex != nil  {
		toSerialize["sex"] = o.Sex
	}
	if o.Strain != nil  {
		toSerialize["strain"] = o.Strain
	}
	if o.SearchText != nil && len(o.GetSearchText()) > 0  {
		toSerialize["search_text"] = o.SearchText
	}
	if o.Rank != nil  {
		toSerialize["rank"] = o.Rank
	}
	if o.ParentTaxId != nil  {
		toSerialize["parent_tax_id"] = o.ParentTaxId
	}
	if o.AssemblyCount != nil  {
		toSerialize["assembly_count"] = o.AssemblyCount
	}
	if o.AssemblyCounts != nil  {
		toSerialize["assembly_counts"] = o.AssemblyCounts
	}
	if o.Counts != nil && len(o.GetCounts()) > 0  {
		toSerialize["counts"] = o.Counts
	}
	if o.Children != nil && len(o.GetChildren()) > 0  {
		toSerialize["children"] = o.Children
	}
	if o.Merged != nil && len(o.GetMerged()) > 0  {
		toSerialize["merged"] = o.Merged
	}
	if o.MergedTaxIds != nil && len(o.GetMergedTaxIds()) > 0  {
		toSerialize["merged_tax_ids"] = o.MergedTaxIds
	}
	if o.MinOrd != nil  {
		toSerialize["min_ord"] = o.MinOrd
	}
	if o.MaxOrd != nil  {
		toSerialize["max_ord"] = o.MaxOrd
	}
	if o.Weight != nil  {
		toSerialize["weight"] = o.Weight
	}
	if o.Key != nil  {
		toSerialize["key"] = o.Key
	}
	if o.Title != nil  {
		toSerialize["title"] = o.Title
	}
	if o.Icon != nil  {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullableV1Organism struct {
	value *V1Organism
	isSet bool
}

func (v NullableV1Organism) Get() *V1Organism {
	return v.value
}

func (v *NullableV1Organism) Set(val *V1Organism) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Organism) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Organism) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Organism(val *V1Organism) *NullableV1Organism {
	return &NullableV1Organism{value: val, isSet: true}
}

func (v NullableV1Organism) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Organism) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



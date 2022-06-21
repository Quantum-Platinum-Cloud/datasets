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

// V1GeneDescriptor struct for V1GeneDescriptor
type V1GeneDescriptor struct {
	GeneId *string `json:"gene_id,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	TaxId *string `json:"tax_id,omitempty"`
	Taxname *string `json:"taxname,omitempty"`
	Type *V1GeneDescriptorGeneType `json:"type,omitempty"`
	RnaType *V1GeneDescriptorRnaType `json:"rna_type,omitempty"`
	Orientation *V1Orientation `json:"orientation,omitempty"`
	GenomicRanges *[]V1SeqRangeSet `json:"genomic_ranges,omitempty"`
	ReferenceStandards *[]V1GenomicRegion `json:"reference_standards,omitempty"`
	GenomicRegions *[]V1GenomicRegion `json:"genomic_regions,omitempty"`
	Transcripts *[]V1Transcript `json:"transcripts,omitempty"`
	Proteins *[]V1Protein `json:"proteins,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Chromosome *string `json:"chromosome,omitempty"`
	Chromosomes *[]string `json:"chromosomes,omitempty"`
	NomenclatureAuthority *V1NomenclatureAuthority `json:"nomenclature_authority,omitempty"`
	SwissProtAccessions *[]string `json:"swiss_prot_accessions,omitempty"`
	EnsemblGeneIds *[]string `json:"ensembl_gene_ids,omitempty"`
	OmimIds *[]string `json:"omim_ids,omitempty"`
	Synonyms *[]string `json:"synonyms,omitempty"`
	ReplacedGeneIds *[]string `json:"replaced_gene_ids,omitempty"`
	Annotations *[]V1Annotation `json:"annotations,omitempty"`
}

// NewV1GeneDescriptor instantiates a new V1GeneDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GeneDescriptor() *V1GeneDescriptor {
	this := V1GeneDescriptor{}
	var type_ V1GeneDescriptorGeneType = V1GENEDESCRIPTORGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V1GeneDescriptorRnaType = V1GENEDESCRIPTORRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V1Orientation = V1ORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// NewV1GeneDescriptorWithDefaults instantiates a new V1GeneDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GeneDescriptorWithDefaults() *V1GeneDescriptor {
	this := V1GeneDescriptor{}
	var type_ V1GeneDescriptorGeneType = V1GENEDESCRIPTORGENETYPE_UNKNOWN
	this.Type = &type_
	var rnaType V1GeneDescriptorRnaType = V1GENEDESCRIPTORRNATYPE_RNA_UNKNOWN
	this.RnaType = &rnaType
	var orientation V1Orientation = V1ORIENTATION_NONE
	this.Orientation = &orientation
	return &this
}

// GetGeneId returns the GeneId field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetGeneId() string {
	if o == nil || o.GeneId == nil {
		var ret string
		return ret
	}
	return *o.GeneId
}

// GetGeneIdOk returns a tuple with the GeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetGeneIdOk() (*string, bool) {
	if o == nil || o.GeneId == nil {
		return nil, false
	}
	return o.GeneId, true
}

// HasGeneId returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasGeneId() bool {
	if o != nil && o.GeneId != nil {
		return true
	}

	return false
}

// SetGeneId gets a reference to the given string and assigns it to the GeneId field.
func (o *V1GeneDescriptor) SetGeneId(v string) {
	o.GeneId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *V1GeneDescriptor) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V1GeneDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *V1GeneDescriptor) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTaxname returns the Taxname field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetTaxname() string {
	if o == nil || o.Taxname == nil {
		var ret string
		return ret
	}
	return *o.Taxname
}

// GetTaxnameOk returns a tuple with the Taxname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetTaxnameOk() (*string, bool) {
	if o == nil || o.Taxname == nil {
		return nil, false
	}
	return o.Taxname, true
}

// HasTaxname returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasTaxname() bool {
	if o != nil && o.Taxname != nil {
		return true
	}

	return false
}

// SetTaxname gets a reference to the given string and assigns it to the Taxname field.
func (o *V1GeneDescriptor) SetTaxname(v string) {
	o.Taxname = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetType() V1GeneDescriptorGeneType {
	if o == nil || o.Type == nil {
		var ret V1GeneDescriptorGeneType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetTypeOk() (*V1GeneDescriptorGeneType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V1GeneDescriptorGeneType and assigns it to the Type field.
func (o *V1GeneDescriptor) SetType(v V1GeneDescriptorGeneType) {
	o.Type = &v
}

// GetRnaType returns the RnaType field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetRnaType() V1GeneDescriptorRnaType {
	if o == nil || o.RnaType == nil {
		var ret V1GeneDescriptorRnaType
		return ret
	}
	return *o.RnaType
}

// GetRnaTypeOk returns a tuple with the RnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetRnaTypeOk() (*V1GeneDescriptorRnaType, bool) {
	if o == nil || o.RnaType == nil {
		return nil, false
	}
	return o.RnaType, true
}

// HasRnaType returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasRnaType() bool {
	if o != nil && o.RnaType != nil {
		return true
	}

	return false
}

// SetRnaType gets a reference to the given V1GeneDescriptorRnaType and assigns it to the RnaType field.
func (o *V1GeneDescriptor) SetRnaType(v V1GeneDescriptorRnaType) {
	o.RnaType = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetOrientation() V1Orientation {
	if o == nil || o.Orientation == nil {
		var ret V1Orientation
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetOrientationOk() (*V1Orientation, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given V1Orientation and assigns it to the Orientation field.
func (o *V1GeneDescriptor) SetOrientation(v V1Orientation) {
	o.Orientation = &v
}

// GetGenomicRanges returns the GenomicRanges field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetGenomicRanges() []V1SeqRangeSet {
	if o == nil || o.GenomicRanges == nil {
		var ret []V1SeqRangeSet
		return ret
	}
	return *o.GenomicRanges
}

// GetGenomicRangesOk returns a tuple with the GenomicRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetGenomicRangesOk() (*[]V1SeqRangeSet, bool) {
	if o == nil || o.GenomicRanges == nil {
		return nil, false
	}
	return o.GenomicRanges, true
}

// HasGenomicRanges returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasGenomicRanges() bool {
	if o != nil && o.GenomicRanges != nil {
		return true
	}

	return false
}

// SetGenomicRanges gets a reference to the given []V1SeqRangeSet and assigns it to the GenomicRanges field.
func (o *V1GeneDescriptor) SetGenomicRanges(v []V1SeqRangeSet) {
	o.GenomicRanges = &v
}

// GetReferenceStandards returns the ReferenceStandards field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetReferenceStandards() []V1GenomicRegion {
	if o == nil || o.ReferenceStandards == nil {
		var ret []V1GenomicRegion
		return ret
	}
	return *o.ReferenceStandards
}

// GetReferenceStandardsOk returns a tuple with the ReferenceStandards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetReferenceStandardsOk() (*[]V1GenomicRegion, bool) {
	if o == nil || o.ReferenceStandards == nil {
		return nil, false
	}
	return o.ReferenceStandards, true
}

// HasReferenceStandards returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasReferenceStandards() bool {
	if o != nil && o.ReferenceStandards != nil {
		return true
	}

	return false
}

// SetReferenceStandards gets a reference to the given []V1GenomicRegion and assigns it to the ReferenceStandards field.
func (o *V1GeneDescriptor) SetReferenceStandards(v []V1GenomicRegion) {
	o.ReferenceStandards = &v
}

// GetGenomicRegions returns the GenomicRegions field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetGenomicRegions() []V1GenomicRegion {
	if o == nil || o.GenomicRegions == nil {
		var ret []V1GenomicRegion
		return ret
	}
	return *o.GenomicRegions
}

// GetGenomicRegionsOk returns a tuple with the GenomicRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetGenomicRegionsOk() (*[]V1GenomicRegion, bool) {
	if o == nil || o.GenomicRegions == nil {
		return nil, false
	}
	return o.GenomicRegions, true
}

// HasGenomicRegions returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasGenomicRegions() bool {
	if o != nil && o.GenomicRegions != nil {
		return true
	}

	return false
}

// SetGenomicRegions gets a reference to the given []V1GenomicRegion and assigns it to the GenomicRegions field.
func (o *V1GeneDescriptor) SetGenomicRegions(v []V1GenomicRegion) {
	o.GenomicRegions = &v
}

// GetTranscripts returns the Transcripts field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetTranscripts() []V1Transcript {
	if o == nil || o.Transcripts == nil {
		var ret []V1Transcript
		return ret
	}
	return *o.Transcripts
}

// GetTranscriptsOk returns a tuple with the Transcripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetTranscriptsOk() (*[]V1Transcript, bool) {
	if o == nil || o.Transcripts == nil {
		return nil, false
	}
	return o.Transcripts, true
}

// HasTranscripts returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasTranscripts() bool {
	if o != nil && o.Transcripts != nil {
		return true
	}

	return false
}

// SetTranscripts gets a reference to the given []V1Transcript and assigns it to the Transcripts field.
func (o *V1GeneDescriptor) SetTranscripts(v []V1Transcript) {
	o.Transcripts = &v
}

// GetProteins returns the Proteins field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetProteins() []V1Protein {
	if o == nil || o.Proteins == nil {
		var ret []V1Protein
		return ret
	}
	return *o.Proteins
}

// GetProteinsOk returns a tuple with the Proteins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetProteinsOk() (*[]V1Protein, bool) {
	if o == nil || o.Proteins == nil {
		return nil, false
	}
	return o.Proteins, true
}

// HasProteins returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasProteins() bool {
	if o != nil && o.Proteins != nil {
		return true
	}

	return false
}

// SetProteins gets a reference to the given []V1Protein and assigns it to the Proteins field.
func (o *V1GeneDescriptor) SetProteins(v []V1Protein) {
	o.Proteins = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V1GeneDescriptor) SetCommonName(v string) {
	o.CommonName = &v
}

// GetChromosome returns the Chromosome field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetChromosome() string {
	if o == nil || o.Chromosome == nil {
		var ret string
		return ret
	}
	return *o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetChromosomeOk() (*string, bool) {
	if o == nil || o.Chromosome == nil {
		return nil, false
	}
	return o.Chromosome, true
}

// HasChromosome returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasChromosome() bool {
	if o != nil && o.Chromosome != nil {
		return true
	}

	return false
}

// SetChromosome gets a reference to the given string and assigns it to the Chromosome field.
func (o *V1GeneDescriptor) SetChromosome(v string) {
	o.Chromosome = &v
}

// GetChromosomes returns the Chromosomes field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetChromosomes() []string {
	if o == nil || o.Chromosomes == nil {
		var ret []string
		return ret
	}
	return *o.Chromosomes
}

// GetChromosomesOk returns a tuple with the Chromosomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetChromosomesOk() (*[]string, bool) {
	if o == nil || o.Chromosomes == nil {
		return nil, false
	}
	return o.Chromosomes, true
}

// HasChromosomes returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasChromosomes() bool {
	if o != nil && o.Chromosomes != nil {
		return true
	}

	return false
}

// SetChromosomes gets a reference to the given []string and assigns it to the Chromosomes field.
func (o *V1GeneDescriptor) SetChromosomes(v []string) {
	o.Chromosomes = &v
}

// GetNomenclatureAuthority returns the NomenclatureAuthority field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetNomenclatureAuthority() V1NomenclatureAuthority {
	if o == nil || o.NomenclatureAuthority == nil {
		var ret V1NomenclatureAuthority
		return ret
	}
	return *o.NomenclatureAuthority
}

// GetNomenclatureAuthorityOk returns a tuple with the NomenclatureAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetNomenclatureAuthorityOk() (*V1NomenclatureAuthority, bool) {
	if o == nil || o.NomenclatureAuthority == nil {
		return nil, false
	}
	return o.NomenclatureAuthority, true
}

// HasNomenclatureAuthority returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasNomenclatureAuthority() bool {
	if o != nil && o.NomenclatureAuthority != nil {
		return true
	}

	return false
}

// SetNomenclatureAuthority gets a reference to the given V1NomenclatureAuthority and assigns it to the NomenclatureAuthority field.
func (o *V1GeneDescriptor) SetNomenclatureAuthority(v V1NomenclatureAuthority) {
	o.NomenclatureAuthority = &v
}

// GetSwissProtAccessions returns the SwissProtAccessions field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetSwissProtAccessions() []string {
	if o == nil || o.SwissProtAccessions == nil {
		var ret []string
		return ret
	}
	return *o.SwissProtAccessions
}

// GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetSwissProtAccessionsOk() (*[]string, bool) {
	if o == nil || o.SwissProtAccessions == nil {
		return nil, false
	}
	return o.SwissProtAccessions, true
}

// HasSwissProtAccessions returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasSwissProtAccessions() bool {
	if o != nil && o.SwissProtAccessions != nil {
		return true
	}

	return false
}

// SetSwissProtAccessions gets a reference to the given []string and assigns it to the SwissProtAccessions field.
func (o *V1GeneDescriptor) SetSwissProtAccessions(v []string) {
	o.SwissProtAccessions = &v
}

// GetEnsemblGeneIds returns the EnsemblGeneIds field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetEnsemblGeneIds() []string {
	if o == nil || o.EnsemblGeneIds == nil {
		var ret []string
		return ret
	}
	return *o.EnsemblGeneIds
}

// GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetEnsemblGeneIdsOk() (*[]string, bool) {
	if o == nil || o.EnsemblGeneIds == nil {
		return nil, false
	}
	return o.EnsemblGeneIds, true
}

// HasEnsemblGeneIds returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasEnsemblGeneIds() bool {
	if o != nil && o.EnsemblGeneIds != nil {
		return true
	}

	return false
}

// SetEnsemblGeneIds gets a reference to the given []string and assigns it to the EnsemblGeneIds field.
func (o *V1GeneDescriptor) SetEnsemblGeneIds(v []string) {
	o.EnsemblGeneIds = &v
}

// GetOmimIds returns the OmimIds field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetOmimIds() []string {
	if o == nil || o.OmimIds == nil {
		var ret []string
		return ret
	}
	return *o.OmimIds
}

// GetOmimIdsOk returns a tuple with the OmimIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetOmimIdsOk() (*[]string, bool) {
	if o == nil || o.OmimIds == nil {
		return nil, false
	}
	return o.OmimIds, true
}

// HasOmimIds returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasOmimIds() bool {
	if o != nil && o.OmimIds != nil {
		return true
	}

	return false
}

// SetOmimIds gets a reference to the given []string and assigns it to the OmimIds field.
func (o *V1GeneDescriptor) SetOmimIds(v []string) {
	o.OmimIds = &v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetSynonymsOk() (*[]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *V1GeneDescriptor) SetSynonyms(v []string) {
	o.Synonyms = &v
}

// GetReplacedGeneIds returns the ReplacedGeneIds field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetReplacedGeneIds() []string {
	if o == nil || o.ReplacedGeneIds == nil {
		var ret []string
		return ret
	}
	return *o.ReplacedGeneIds
}

// GetReplacedGeneIdsOk returns a tuple with the ReplacedGeneIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetReplacedGeneIdsOk() (*[]string, bool) {
	if o == nil || o.ReplacedGeneIds == nil {
		return nil, false
	}
	return o.ReplacedGeneIds, true
}

// HasReplacedGeneIds returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasReplacedGeneIds() bool {
	if o != nil && o.ReplacedGeneIds != nil {
		return true
	}

	return false
}

// SetReplacedGeneIds gets a reference to the given []string and assigns it to the ReplacedGeneIds field.
func (o *V1GeneDescriptor) SetReplacedGeneIds(v []string) {
	o.ReplacedGeneIds = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *V1GeneDescriptor) GetAnnotations() []V1Annotation {
	if o == nil || o.Annotations == nil {
		var ret []V1Annotation
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GeneDescriptor) GetAnnotationsOk() (*[]V1Annotation, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *V1GeneDescriptor) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []V1Annotation and assigns it to the Annotations field.
func (o *V1GeneDescriptor) SetAnnotations(v []V1Annotation) {
	o.Annotations = &v
}

func (o V1GeneDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeneId != nil  {
		toSerialize["gene_id"] = o.GeneId
	}
	if o.Symbol != nil  {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Description != nil  {
		toSerialize["description"] = o.Description
	}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.Taxname != nil  {
		toSerialize["taxname"] = o.Taxname
	}
	if o.Type != nil  {
		toSerialize["type"] = o.Type
	}
	if o.RnaType != nil  {
		toSerialize["rna_type"] = o.RnaType
	}
	if o.Orientation != nil  {
		toSerialize["orientation"] = o.Orientation
	}
	if o.GenomicRanges != nil && len(o.GetGenomicRanges()) > 0  {
		toSerialize["genomic_ranges"] = o.GenomicRanges
	}
	if o.ReferenceStandards != nil && len(o.GetReferenceStandards()) > 0  {
		toSerialize["reference_standards"] = o.ReferenceStandards
	}
	if o.GenomicRegions != nil && len(o.GetGenomicRegions()) > 0  {
		toSerialize["genomic_regions"] = o.GenomicRegions
	}
	if o.Transcripts != nil && len(o.GetTranscripts()) > 0  {
		toSerialize["transcripts"] = o.Transcripts
	}
	if o.Proteins != nil && len(o.GetProteins()) > 0  {
		toSerialize["proteins"] = o.Proteins
	}
	if o.CommonName != nil  {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Chromosome != nil  {
		toSerialize["chromosome"] = o.Chromosome
	}
	if o.Chromosomes != nil && len(o.GetChromosomes()) > 0  {
		toSerialize["chromosomes"] = o.Chromosomes
	}
	if o.NomenclatureAuthority != nil  {
		toSerialize["nomenclature_authority"] = o.NomenclatureAuthority
	}
	if o.SwissProtAccessions != nil && len(o.GetSwissProtAccessions()) > 0  {
		toSerialize["swiss_prot_accessions"] = o.SwissProtAccessions
	}
	if o.EnsemblGeneIds != nil && len(o.GetEnsemblGeneIds()) > 0  {
		toSerialize["ensembl_gene_ids"] = o.EnsemblGeneIds
	}
	if o.OmimIds != nil && len(o.GetOmimIds()) > 0  {
		toSerialize["omim_ids"] = o.OmimIds
	}
	if o.Synonyms != nil && len(o.GetSynonyms()) > 0  {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.ReplacedGeneIds != nil && len(o.GetReplacedGeneIds()) > 0  {
		toSerialize["replaced_gene_ids"] = o.ReplacedGeneIds
	}
	if o.Annotations != nil && len(o.GetAnnotations()) > 0  {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableV1GeneDescriptor struct {
	value *V1GeneDescriptor
	isSet bool
}

func (v NullableV1GeneDescriptor) Get() *V1GeneDescriptor {
	return v.value
}

func (v *NullableV1GeneDescriptor) Set(val *V1GeneDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GeneDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GeneDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GeneDescriptor(val *V1GeneDescriptor) *NullableV1GeneDescriptor {
	return &NullableV1GeneDescriptor{value: val, isSet: true}
}

func (v NullableV1GeneDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GeneDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



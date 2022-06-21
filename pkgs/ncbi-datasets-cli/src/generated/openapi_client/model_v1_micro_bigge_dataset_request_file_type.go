/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"fmt"
)

// V1MicroBiggeDatasetRequestFileType the model 'V1MicroBiggeDatasetRequestFileType'
type V1MicroBiggeDatasetRequestFileType string

// List of v1MicroBiggeDatasetRequestFileType
const (
	V1MICROBIGGEDATASETREQUESTFILETYPE_FILE_TYPE_UNSPECIFIED V1MicroBiggeDatasetRequestFileType = "FILE_TYPE_UNSPECIFIED"
	V1MICROBIGGEDATASETREQUESTFILETYPE_ELEMENT_FASTA V1MicroBiggeDatasetRequestFileType = "element_fasta"
	V1MICROBIGGEDATASETREQUESTFILETYPE_ELEMENT_FLANK_FASTA V1MicroBiggeDatasetRequestFileType = "element_flank_fasta"
	V1MICROBIGGEDATASETREQUESTFILETYPE_CONTIG_FASTA V1MicroBiggeDatasetRequestFileType = "contig_fasta"
	V1MICROBIGGEDATASETREQUESTFILETYPE_PROTEIN_FASTA V1MicroBiggeDatasetRequestFileType = "protein_fasta"
)

// All allowed values of V1MicroBiggeDatasetRequestFileType enum
var AllowedV1MicroBiggeDatasetRequestFileTypeEnumValues = []V1MicroBiggeDatasetRequestFileType{
	"FILE_TYPE_UNSPECIFIED",
	"element_fasta",
	"element_flank_fasta",
	"contig_fasta",
	"protein_fasta",
}

func (v *V1MicroBiggeDatasetRequestFileType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1MicroBiggeDatasetRequestFileType(value)
	for _, existing := range AllowedV1MicroBiggeDatasetRequestFileTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1MicroBiggeDatasetRequestFileType", value)
}

// NewV1MicroBiggeDatasetRequestFileTypeFromValue returns a pointer to a valid V1MicroBiggeDatasetRequestFileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1MicroBiggeDatasetRequestFileTypeFromValue(v string) (*V1MicroBiggeDatasetRequestFileType, error) {
	ev := V1MicroBiggeDatasetRequestFileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1MicroBiggeDatasetRequestFileType: valid values are %v", v, AllowedV1MicroBiggeDatasetRequestFileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1MicroBiggeDatasetRequestFileType) IsValid() bool {
	for _, existing := range AllowedV1MicroBiggeDatasetRequestFileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1MicroBiggeDatasetRequestFileType value
func (v V1MicroBiggeDatasetRequestFileType) Ptr() *V1MicroBiggeDatasetRequestFileType {
	return &v
}

type NullableV1MicroBiggeDatasetRequestFileType struct {
	value *V1MicroBiggeDatasetRequestFileType
	isSet bool
}

func (v NullableV1MicroBiggeDatasetRequestFileType) Get() *V1MicroBiggeDatasetRequestFileType {
	return v.value
}

func (v *NullableV1MicroBiggeDatasetRequestFileType) Set(val *V1MicroBiggeDatasetRequestFileType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1MicroBiggeDatasetRequestFileType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1MicroBiggeDatasetRequestFileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1MicroBiggeDatasetRequestFileType(val *V1MicroBiggeDatasetRequestFileType) *NullableV1MicroBiggeDatasetRequestFileType {
	return &NullableV1MicroBiggeDatasetRequestFileType{value: val, isSet: true}
}

func (v NullableV1MicroBiggeDatasetRequestFileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1MicroBiggeDatasetRequestFileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


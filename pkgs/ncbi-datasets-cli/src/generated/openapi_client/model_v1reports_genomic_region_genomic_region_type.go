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

// V1reportsGenomicRegionGenomicRegionType the model 'V1reportsGenomicRegionGenomicRegionType'
type V1reportsGenomicRegionGenomicRegionType string

// List of v1reportsGenomicRegionGenomicRegionType
const (
	V1REPORTSGENOMICREGIONGENOMICREGIONTYPE_UNKNOWN V1reportsGenomicRegionGenomicRegionType = "UNKNOWN"
	V1REPORTSGENOMICREGIONGENOMICREGIONTYPE_REFSEQ_GENE V1reportsGenomicRegionGenomicRegionType = "REFSEQ_GENE"
	V1REPORTSGENOMICREGIONGENOMICREGIONTYPE_PSEUDOGENE V1reportsGenomicRegionGenomicRegionType = "PSEUDOGENE"
	V1REPORTSGENOMICREGIONGENOMICREGIONTYPE_BIOLOGICAL_REGION V1reportsGenomicRegionGenomicRegionType = "BIOLOGICAL_REGION"
	V1REPORTSGENOMICREGIONGENOMICREGIONTYPE_OTHER V1reportsGenomicRegionGenomicRegionType = "OTHER"
)

// All allowed values of V1reportsGenomicRegionGenomicRegionType enum
var AllowedV1reportsGenomicRegionGenomicRegionTypeEnumValues = []V1reportsGenomicRegionGenomicRegionType{
	"UNKNOWN",
	"REFSEQ_GENE",
	"PSEUDOGENE",
	"BIOLOGICAL_REGION",
	"OTHER",
}

func (v *V1reportsGenomicRegionGenomicRegionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1reportsGenomicRegionGenomicRegionType(value)
	for _, existing := range AllowedV1reportsGenomicRegionGenomicRegionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1reportsGenomicRegionGenomicRegionType", value)
}

// NewV1reportsGenomicRegionGenomicRegionTypeFromValue returns a pointer to a valid V1reportsGenomicRegionGenomicRegionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1reportsGenomicRegionGenomicRegionTypeFromValue(v string) (*V1reportsGenomicRegionGenomicRegionType, error) {
	ev := V1reportsGenomicRegionGenomicRegionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1reportsGenomicRegionGenomicRegionType: valid values are %v", v, AllowedV1reportsGenomicRegionGenomicRegionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1reportsGenomicRegionGenomicRegionType) IsValid() bool {
	for _, existing := range AllowedV1reportsGenomicRegionGenomicRegionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1reportsGenomicRegionGenomicRegionType value
func (v V1reportsGenomicRegionGenomicRegionType) Ptr() *V1reportsGenomicRegionGenomicRegionType {
	return &v
}

type NullableV1reportsGenomicRegionGenomicRegionType struct {
	value *V1reportsGenomicRegionGenomicRegionType
	isSet bool
}

func (v NullableV1reportsGenomicRegionGenomicRegionType) Get() *V1reportsGenomicRegionGenomicRegionType {
	return v.value
}

func (v *NullableV1reportsGenomicRegionGenomicRegionType) Set(val *V1reportsGenomicRegionGenomicRegionType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsGenomicRegionGenomicRegionType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsGenomicRegionGenomicRegionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsGenomicRegionGenomicRegionType(val *V1reportsGenomicRegionGenomicRegionType) *NullableV1reportsGenomicRegionGenomicRegionType {
	return &NullableV1reportsGenomicRegionGenomicRegionType{value: val, isSet: true}
}

func (v NullableV1reportsGenomicRegionGenomicRegionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsGenomicRegionGenomicRegionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


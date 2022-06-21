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

// V1AssemblyDatasetDescriptorsFilterAssemblyLevel the model 'V1AssemblyDatasetDescriptorsFilterAssemblyLevel'
type V1AssemblyDatasetDescriptorsFilterAssemblyLevel string

// List of v1AssemblyDatasetDescriptorsFilterAssemblyLevel
const (
	V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYLEVEL_CHROMOSOME V1AssemblyDatasetDescriptorsFilterAssemblyLevel = "chromosome"
	V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYLEVEL_SCAFFOLD V1AssemblyDatasetDescriptorsFilterAssemblyLevel = "scaffold"
	V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYLEVEL_CONTIG V1AssemblyDatasetDescriptorsFilterAssemblyLevel = "contig"
	V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYLEVEL_COMPLETE_GENOME V1AssemblyDatasetDescriptorsFilterAssemblyLevel = "complete_genome"
)

// All allowed values of V1AssemblyDatasetDescriptorsFilterAssemblyLevel enum
var AllowedV1AssemblyDatasetDescriptorsFilterAssemblyLevelEnumValues = []V1AssemblyDatasetDescriptorsFilterAssemblyLevel{
	"chromosome",
	"scaffold",
	"contig",
	"complete_genome",
}

func (v *V1AssemblyDatasetDescriptorsFilterAssemblyLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1AssemblyDatasetDescriptorsFilterAssemblyLevel(value)
	for _, existing := range AllowedV1AssemblyDatasetDescriptorsFilterAssemblyLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1AssemblyDatasetDescriptorsFilterAssemblyLevel", value)
}

// NewV1AssemblyDatasetDescriptorsFilterAssemblyLevelFromValue returns a pointer to a valid V1AssemblyDatasetDescriptorsFilterAssemblyLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1AssemblyDatasetDescriptorsFilterAssemblyLevelFromValue(v string) (*V1AssemblyDatasetDescriptorsFilterAssemblyLevel, error) {
	ev := V1AssemblyDatasetDescriptorsFilterAssemblyLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1AssemblyDatasetDescriptorsFilterAssemblyLevel: valid values are %v", v, AllowedV1AssemblyDatasetDescriptorsFilterAssemblyLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1AssemblyDatasetDescriptorsFilterAssemblyLevel) IsValid() bool {
	for _, existing := range AllowedV1AssemblyDatasetDescriptorsFilterAssemblyLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1AssemblyDatasetDescriptorsFilterAssemblyLevel value
func (v V1AssemblyDatasetDescriptorsFilterAssemblyLevel) Ptr() *V1AssemblyDatasetDescriptorsFilterAssemblyLevel {
	return &v
}

type NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel struct {
	value *V1AssemblyDatasetDescriptorsFilterAssemblyLevel
	isSet bool
}

func (v NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel) Get() *V1AssemblyDatasetDescriptorsFilterAssemblyLevel {
	return v.value
}

func (v *NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel) Set(val *V1AssemblyDatasetDescriptorsFilterAssemblyLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel(val *V1AssemblyDatasetDescriptorsFilterAssemblyLevel) *NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel {
	return &NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetDescriptorsFilterAssemblyLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets version 1 API is considred stable and will not be subject to breaking changes. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"fmt"
)

// V1AssemblyDatasetRequestResolution the model 'V1AssemblyDatasetRequestResolution'
type V1AssemblyDatasetRequestResolution string

// List of v1AssemblyDatasetRequestResolution
const (
	V1ASSEMBLYDATASETREQUESTRESOLUTION_FULLY_HYDRATED V1AssemblyDatasetRequestResolution = "FULLY_HYDRATED"
	V1ASSEMBLYDATASETREQUESTRESOLUTION_DATA_REPORT_ONLY V1AssemblyDatasetRequestResolution = "DATA_REPORT_ONLY"
)

// All allowed values of V1AssemblyDatasetRequestResolution enum
var AllowedV1AssemblyDatasetRequestResolutionEnumValues = []V1AssemblyDatasetRequestResolution{
	"FULLY_HYDRATED",
	"DATA_REPORT_ONLY",
}

func (v *V1AssemblyDatasetRequestResolution) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1AssemblyDatasetRequestResolution(value)
	for _, existing := range AllowedV1AssemblyDatasetRequestResolutionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1AssemblyDatasetRequestResolution", value)
}

// NewV1AssemblyDatasetRequestResolutionFromValue returns a pointer to a valid V1AssemblyDatasetRequestResolution
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1AssemblyDatasetRequestResolutionFromValue(v string) (*V1AssemblyDatasetRequestResolution, error) {
	ev := V1AssemblyDatasetRequestResolution(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1AssemblyDatasetRequestResolution: valid values are %v", v, AllowedV1AssemblyDatasetRequestResolutionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1AssemblyDatasetRequestResolution) IsValid() bool {
	for _, existing := range AllowedV1AssemblyDatasetRequestResolutionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1AssemblyDatasetRequestResolution value
func (v V1AssemblyDatasetRequestResolution) Ptr() *V1AssemblyDatasetRequestResolution {
	return &v
}

type NullableV1AssemblyDatasetRequestResolution struct {
	value *V1AssemblyDatasetRequestResolution
	isSet bool
}

func (v NullableV1AssemblyDatasetRequestResolution) Get() *V1AssemblyDatasetRequestResolution {
	return v.value
}

func (v *NullableV1AssemblyDatasetRequestResolution) Set(val *V1AssemblyDatasetRequestResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyDatasetRequestResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyDatasetRequestResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyDatasetRequestResolution(val *V1AssemblyDatasetRequestResolution) *NullableV1AssemblyDatasetRequestResolution {
	return &NullableV1AssemblyDatasetRequestResolution{value: val, isSet: true}
}

func (v NullableV1AssemblyDatasetRequestResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyDatasetRequestResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


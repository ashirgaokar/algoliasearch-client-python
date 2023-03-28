// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Anchoring Whether the pattern parameter must match the beginning or the end of the query string, or both, or none.
type Anchoring string

// List of anchoring
const (
	ANCHORING_IS          Anchoring = "is"
	ANCHORING_STARTS_WITH Anchoring = "startsWith"
	ANCHORING_ENDS_WITH   Anchoring = "endsWith"
	ANCHORING_CONTAINS    Anchoring = "contains"
)

// All allowed values of Anchoring enum
var AllowedAnchoringEnumValues = []Anchoring{
	"is",
	"startsWith",
	"endsWith",
	"contains",
}

func (v *Anchoring) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Anchoring(value)
	for _, existing := range AllowedAnchoringEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Anchoring", value)
}

// NewAnchoringFromValue returns a pointer to a valid Anchoring
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnchoringFromValue(v string) (*Anchoring, error) {
	ev := Anchoring(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Anchoring: valid values are %v", v, AllowedAnchoringEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Anchoring) IsValid() bool {
	for _, existing := range AllowedAnchoringEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to anchoring value
func (v Anchoring) Ptr() *Anchoring {
	return &v
}

type NullableAnchoring struct {
	value *Anchoring
	isSet bool
}

func (v NullableAnchoring) Get() *Anchoring {
	return v.value
}

func (v *NullableAnchoring) Set(val *Anchoring) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoring) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoring(val *Anchoring) *NullableAnchoring {
	return &NullableAnchoring{value: val, isSet: true}
}

func (v NullableAnchoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

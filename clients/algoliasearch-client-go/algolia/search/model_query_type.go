// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// QueryType Controls if and how query words are interpreted as prefixes.
type QueryType string

// List of queryType
const (
	QUERYTYPE_PREFIX_LAST QueryType = "prefixLast"
	QUERYTYPE_PREFIX_ALL  QueryType = "prefixAll"
	QUERYTYPE_PREFIX_NONE QueryType = "prefixNone"
)

// All allowed values of QueryType enum
var AllowedQueryTypeEnumValues = []QueryType{
	"prefixLast",
	"prefixAll",
	"prefixNone",
}

func (v *QueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryType(value)
	for _, existing := range AllowedQueryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryType", value)
}

// NewQueryTypeFromValue returns a pointer to a valid QueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryTypeFromValue(v string) (*QueryType, error) {
	ev := QueryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryType: valid values are %v", v, AllowedQueryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryType) IsValid() bool {
	for _, existing := range AllowedQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryType value
func (v QueryType) Ptr() *QueryType {
	return &v
}

type NullableQueryType struct {
	value *QueryType
	isSet bool
}

func (v NullableQueryType) Get() *QueryType {
	return v.value
}

func (v *NullableQueryType) Set(val *QueryType) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryType) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryType(val *QueryType) *NullableQueryType {
	return &NullableQueryType{value: val, isSet: true}
}

func (v NullableQueryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

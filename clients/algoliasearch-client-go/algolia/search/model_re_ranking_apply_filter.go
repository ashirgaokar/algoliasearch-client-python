// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// ReRankingApplyFilter - When Dynamic Re-Ranking is enabled, only records that match these filters will be impacted by Dynamic Re-Ranking.
type ReRankingApplyFilter struct {
	ArrayOfMixedSearchFilters *[]MixedSearchFilters
	String                    *string
}

// []MixedSearchFiltersAsReRankingApplyFilter is a convenience function that returns []MixedSearchFilters wrapped in ReRankingApplyFilter
func ArrayOfMixedSearchFiltersAsReRankingApplyFilter(v *[]MixedSearchFilters) ReRankingApplyFilter {
	return ReRankingApplyFilter{
		ArrayOfMixedSearchFilters: v,
	}
}

// stringAsReRankingApplyFilter is a convenience function that returns string wrapped in ReRankingApplyFilter
func StringAsReRankingApplyFilter(v *string) ReRankingApplyFilter {
	return ReRankingApplyFilter{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReRankingApplyFilter) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into ArrayOfMixedSearchFilters
	err = newStrictDecoder(data).Decode(&dst.ArrayOfMixedSearchFilters)
	if err == nil {
		jsonArrayOfMixedSearchFilters, _ := json.Marshal(dst.ArrayOfMixedSearchFilters)
		if string(jsonArrayOfMixedSearchFilters) == "{}" { // empty struct
			dst.ArrayOfMixedSearchFilters = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfMixedSearchFilters = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfMixedSearchFilters = nil
		dst.String = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ReRankingApplyFilter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ReRankingApplyFilter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReRankingApplyFilter) MarshalJSON() ([]byte, error) {
	if src.ArrayOfMixedSearchFilters != nil {
		return json.Marshal(&src.ArrayOfMixedSearchFilters)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReRankingApplyFilter) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfMixedSearchFilters != nil {
		return obj.ArrayOfMixedSearchFilters
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableReRankingApplyFilter struct {
	value *ReRankingApplyFilter
	isSet bool
}

func (v NullableReRankingApplyFilter) Get() *ReRankingApplyFilter {
	return v.value
}

func (v *NullableReRankingApplyFilter) Set(val *ReRankingApplyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReRankingApplyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReRankingApplyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReRankingApplyFilter(val *ReRankingApplyFilter) *NullableReRankingApplyFilter {
	return &NullableReRankingApplyFilter{value: val, isSet: true}
}

func (v NullableReRankingApplyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReRankingApplyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

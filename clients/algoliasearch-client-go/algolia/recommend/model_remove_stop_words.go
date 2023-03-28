// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// RemoveStopWords - Removes stop (common) words from the query before executing it. removeStopWords is used in conjunction with the queryLanguages setting. list: language ISO codes for which ignoring plurals should be enabled. This list will override any values that you may have set in queryLanguages. true: enables the stop word functionality, ensuring that stop words are removed from consideration in a search. The languages supported here are either every language, or those set by queryLanguages. false: disables stop word functionality, allowing stop words to be taken into account in a search.
type RemoveStopWords struct {
	ArrayOfString *[]string
	Bool          *bool
}

// []stringAsRemoveStopWords is a convenience function that returns []string wrapped in RemoveStopWords
func ArrayOfStringAsRemoveStopWords(v *[]string) RemoveStopWords {
	return RemoveStopWords{
		ArrayOfString: v,
	}
}

// boolAsRemoveStopWords is a convenience function that returns bool wrapped in RemoveStopWords
func BoolAsRemoveStopWords(v *bool) RemoveStopWords {
	return RemoveStopWords{
		Bool: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RemoveStopWords) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.Bool = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(RemoveStopWords)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(RemoveStopWords)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RemoveStopWords) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RemoveStopWords) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

type NullableRemoveStopWords struct {
	value *RemoveStopWords
	isSet bool
}

func (v NullableRemoveStopWords) Get() *RemoveStopWords {
	return v.value
}

func (v *NullableRemoveStopWords) Set(val *RemoveStopWords) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveStopWords) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveStopWords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveStopWords(val *RemoveStopWords) *NullableRemoveStopWords {
	return &NullableRemoveStopWords{value: val, isSet: true}
}

func (v NullableRemoveStopWords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveStopWords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

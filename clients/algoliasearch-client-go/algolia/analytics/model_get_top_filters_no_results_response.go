// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetTopFiltersNoResultsResponse struct for GetTopFiltersNoResultsResponse
type GetTopFiltersNoResultsResponse struct {
	// A list of filters without results.
	Values []GetTopFiltersNoResultsValues `json:"values"`
}

// NewGetTopFiltersNoResultsResponse instantiates a new GetTopFiltersNoResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTopFiltersNoResultsResponse(values []GetTopFiltersNoResultsValues) *GetTopFiltersNoResultsResponse {
	this := &GetTopFiltersNoResultsResponse{}
	this.Values = values
	return this
}

// NewGetTopFiltersNoResultsResponseWithDefaults instantiates a new GetTopFiltersNoResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTopFiltersNoResultsResponseWithDefaults() *GetTopFiltersNoResultsResponse {
	this := &GetTopFiltersNoResultsResponse{}
	return this
}

// GetValues returns the Values field value
func (o *GetTopFiltersNoResultsResponse) GetValues() []GetTopFiltersNoResultsValues {
	if o == nil {
		var ret []GetTopFiltersNoResultsValues
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTopFiltersNoResultsResponse) GetValuesOk() ([]GetTopFiltersNoResultsValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *GetTopFiltersNoResultsResponse) SetValues(v []GetTopFiltersNoResultsValues) {
	o.Values = v
}

func (o GetTopFiltersNoResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

func (o GetTopFiltersNoResultsResponse) String() string {
	out := "GetTopFiltersNoResultsResponse {\n"
	out += fmt.Sprintf("  values=%v\n", o.Values)
	out += "}"
	return out
}

type NullableGetTopFiltersNoResultsResponse struct {
	value *GetTopFiltersNoResultsResponse
	isSet bool
}

func (v NullableGetTopFiltersNoResultsResponse) Get() *GetTopFiltersNoResultsResponse {
	return v.value
}

func (v *NullableGetTopFiltersNoResultsResponse) Set(val *GetTopFiltersNoResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTopFiltersNoResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTopFiltersNoResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTopFiltersNoResultsResponse(val *GetTopFiltersNoResultsResponse) *NullableGetTopFiltersNoResultsResponse {
	return &NullableGetTopFiltersNoResultsResponse{value: val, isSet: true}
}

func (v NullableGetTopFiltersNoResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTopFiltersNoResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetSearchesNoClicksResponse struct for GetSearchesNoClicksResponse
type GetSearchesNoClicksResponse struct {
	// A list of searches with no clicks and their count.
	Searches []SearchNoClickEvent `json:"searches"`
}

// NewGetSearchesNoClicksResponse instantiates a new GetSearchesNoClicksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSearchesNoClicksResponse(searches []SearchNoClickEvent) *GetSearchesNoClicksResponse {
	this := &GetSearchesNoClicksResponse{}
	this.Searches = searches
	return this
}

// NewGetSearchesNoClicksResponseWithDefaults instantiates a new GetSearchesNoClicksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSearchesNoClicksResponseWithDefaults() *GetSearchesNoClicksResponse {
	this := &GetSearchesNoClicksResponse{}
	return this
}

// GetSearches returns the Searches field value
func (o *GetSearchesNoClicksResponse) GetSearches() []SearchNoClickEvent {
	if o == nil {
		var ret []SearchNoClickEvent
		return ret
	}

	return o.Searches
}

// GetSearchesOk returns a tuple with the Searches field value
// and a boolean to check if the value has been set.
func (o *GetSearchesNoClicksResponse) GetSearchesOk() ([]SearchNoClickEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Searches, true
}

// SetSearches sets field value
func (o *GetSearchesNoClicksResponse) SetSearches(v []SearchNoClickEvent) {
	o.Searches = v
}

func (o GetSearchesNoClicksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["searches"] = o.Searches
	}
	return json.Marshal(toSerialize)
}

func (o GetSearchesNoClicksResponse) String() string {
	out := "GetSearchesNoClicksResponse {\n"
	out += fmt.Sprintf("  searches=%v\n", o.Searches)
	out += "}"
	return out
}

type NullableGetSearchesNoClicksResponse struct {
	value *GetSearchesNoClicksResponse
	isSet bool
}

func (v NullableGetSearchesNoClicksResponse) Get() *GetSearchesNoClicksResponse {
	return v.value
}

func (v *NullableGetSearchesNoClicksResponse) Set(val *GetSearchesNoClicksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSearchesNoClicksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSearchesNoClicksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSearchesNoClicksResponse(val *GetSearchesNoClicksResponse) *NullableGetSearchesNoClicksResponse {
	return &NullableGetSearchesNoClicksResponse{value: val, isSet: true}
}

func (v NullableGetSearchesNoClicksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSearchesNoClicksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

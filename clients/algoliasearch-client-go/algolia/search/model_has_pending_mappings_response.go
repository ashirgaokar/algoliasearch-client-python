// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// HasPendingMappingsResponse struct for HasPendingMappingsResponse
type HasPendingMappingsResponse struct {
	// If there is any clusters with pending mapping state.
	Pending bool `json:"pending"`
	// Describe cluster pending (migrating, creating, deleting) mapping state.
	Clusters *map[string][]string `json:"clusters,omitempty"`
}

type HasPendingMappingsResponseOption func(f *HasPendingMappingsResponse)

func WithHasPendingMappingsResponseClusters(val map[string][]string) HasPendingMappingsResponseOption {
	return func(f *HasPendingMappingsResponse) {
		f.Clusters = &val
	}
}

// NewHasPendingMappingsResponse instantiates a new HasPendingMappingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasPendingMappingsResponse(pending bool, opts ...HasPendingMappingsResponseOption) *HasPendingMappingsResponse {
	this := &HasPendingMappingsResponse{}
	this.Pending = pending
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewHasPendingMappingsResponseWithDefaults instantiates a new HasPendingMappingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasPendingMappingsResponseWithDefaults() *HasPendingMappingsResponse {
	this := &HasPendingMappingsResponse{}
	return this
}

// GetPending returns the Pending field value
func (o *HasPendingMappingsResponse) GetPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *HasPendingMappingsResponse) GetPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *HasPendingMappingsResponse) SetPending(v bool) {
	o.Pending = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *HasPendingMappingsResponse) GetClusters() map[string][]string {
	if o == nil || o.Clusters == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasPendingMappingsResponse) GetClustersOk() (*map[string][]string, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *HasPendingMappingsResponse) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given map[string][]string and assigns it to the Clusters field.
func (o *HasPendingMappingsResponse) SetClusters(v map[string][]string) {
	o.Clusters = &v
}

func (o HasPendingMappingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["pending"] = o.Pending
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	return json.Marshal(toSerialize)
}

func (o HasPendingMappingsResponse) String() string {
	out := "HasPendingMappingsResponse {\n"
	out += fmt.Sprintf("  pending=%v\n", o.Pending)
	out += fmt.Sprintf("  clusters=%v\n", o.Clusters)
	out += "}"
	return out
}

type NullableHasPendingMappingsResponse struct {
	value *HasPendingMappingsResponse
	isSet bool
}

func (v NullableHasPendingMappingsResponse) Get() *HasPendingMappingsResponse {
	return v.value
}

func (v *NullableHasPendingMappingsResponse) Set(val *HasPendingMappingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHasPendingMappingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHasPendingMappingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasPendingMappingsResponse(val *HasPendingMappingsResponse) *NullableHasPendingMappingsResponse {
	return &NullableHasPendingMappingsResponse{value: val, isSet: true}
}

func (v NullableHasPendingMappingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasPendingMappingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

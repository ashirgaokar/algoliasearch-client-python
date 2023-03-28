// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// AssignUserIdParams Assign userID parameters.
type AssignUserIdParams struct {
	// Name of the cluster.
	Cluster string `json:"cluster"`
}

// NewAssignUserIdParams instantiates a new AssignUserIdParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignUserIdParams(cluster string) *AssignUserIdParams {
	this := &AssignUserIdParams{}
	this.Cluster = cluster
	return this
}

// NewAssignUserIdParamsWithDefaults instantiates a new AssignUserIdParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignUserIdParamsWithDefaults() *AssignUserIdParams {
	this := &AssignUserIdParams{}
	return this
}

// GetCluster returns the Cluster field value
func (o *AssignUserIdParams) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *AssignUserIdParams) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *AssignUserIdParams) SetCluster(v string) {
	o.Cluster = v
}

func (o AssignUserIdParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	return json.Marshal(toSerialize)
}

func (o AssignUserIdParams) String() string {
	out := "AssignUserIdParams {\n"
	out += fmt.Sprintf("  cluster=%v\n", o.Cluster)
	out += "}"
	return out
}

type NullableAssignUserIdParams struct {
	value *AssignUserIdParams
	isSet bool
}

func (v NullableAssignUserIdParams) Get() *AssignUserIdParams {
	return v.value
}

func (v *NullableAssignUserIdParams) Set(val *AssignUserIdParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignUserIdParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignUserIdParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignUserIdParams(val *AssignUserIdParams) *NullableAssignUserIdParams {
	return &NullableAssignUserIdParams{value: val, isSet: true}
}

func (v NullableAssignUserIdParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignUserIdParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

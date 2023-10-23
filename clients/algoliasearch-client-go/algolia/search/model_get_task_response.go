// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// GetTaskResponse struct for GetTaskResponse
type GetTaskResponse struct {
	Status TaskStatus `json:"status"`
}

// NewGetTaskResponse instantiates a new GetTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaskResponse(status TaskStatus) *GetTaskResponse {
	this := &GetTaskResponse{}
	this.Status = status
	return this
}

// NewGetTaskResponseWithDefaults instantiates a new GetTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaskResponseWithDefaults() *GetTaskResponse {
	this := &GetTaskResponse{}
	return this
}

// GetStatus returns the Status field value
func (o *GetTaskResponse) GetStatus() TaskStatus {
	if o == nil {
		var ret TaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetStatusOk() (*TaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTaskResponse) SetStatus(v TaskStatus) {
	o.Status = v
}

func (o GetTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

func (o GetTaskResponse) String() string {
	out := ""
	out += fmt.Sprintf("  status=%v\n", o.Status)
	return fmt.Sprintf("GetTaskResponse {\n%s}", out)
}

type NullableGetTaskResponse struct {
	value *GetTaskResponse
	isSet bool
}

func (v NullableGetTaskResponse) Get() *GetTaskResponse {
	return v.value
}

func (v *NullableGetTaskResponse) Set(val *GetTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaskResponse(val *GetTaskResponse) *NullableGetTaskResponse {
	return &NullableGetTaskResponse{value: val, isSet: true}
}

func (v NullableGetTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

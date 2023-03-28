// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// GetLogsResponse struct for GetLogsResponse
type GetLogsResponse struct {
	Logs []Log `json:"logs"`
}

// NewGetLogsResponse instantiates a new GetLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLogsResponse(logs []Log) *GetLogsResponse {
	this := &GetLogsResponse{}
	this.Logs = logs
	return this
}

// NewGetLogsResponseWithDefaults instantiates a new GetLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLogsResponseWithDefaults() *GetLogsResponse {
	this := &GetLogsResponse{}
	return this
}

// GetLogs returns the Logs field value
func (o *GetLogsResponse) GetLogs() []Log {
	if o == nil {
		var ret []Log
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *GetLogsResponse) GetLogsOk() ([]Log, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *GetLogsResponse) SetLogs(v []Log) {
	o.Logs = v
}

func (o GetLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["logs"] = o.Logs
	}
	return json.Marshal(toSerialize)
}

func (o GetLogsResponse) String() string {
	out := "GetLogsResponse {\n"
	out += fmt.Sprintf("  logs=%v\n", o.Logs)
	out += "}"
	return out
}

type NullableGetLogsResponse struct {
	value *GetLogsResponse
	isSet bool
}

func (v NullableGetLogsResponse) Get() *GetLogsResponse {
	return v.value
}

func (v *NullableGetLogsResponse) Set(val *GetLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogsResponse(val *GetLogsResponse) *NullableGetLogsResponse {
	return &NullableGetLogsResponse{value: val, isSet: true}
}

func (v NullableGetLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

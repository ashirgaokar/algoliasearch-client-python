// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// ListTasksResponse A list of tasks with pagination details.
type ListTasksResponse struct {
	Tasks      []Task     `json:"tasks"`
	Pagination Pagination `json:"pagination"`
}

// NewListTasksResponse instantiates a new ListTasksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTasksResponse(tasks []Task, pagination Pagination) *ListTasksResponse {
	this := &ListTasksResponse{}
	this.Tasks = tasks
	this.Pagination = pagination
	return this
}

// NewListTasksResponseWithDefaults instantiates a new ListTasksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTasksResponseWithDefaults() *ListTasksResponse {
	this := &ListTasksResponse{}
	return this
}

// GetTasks returns the Tasks field value
func (o *ListTasksResponse) GetTasks() []Task {
	if o == nil {
		var ret []Task
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *ListTasksResponse) GetTasksOk() ([]Task, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *ListTasksResponse) SetTasks(v []Task) {
	o.Tasks = v
}

// GetPagination returns the Pagination field value
func (o *ListTasksResponse) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListTasksResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListTasksResponse) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListTasksResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["tasks"] = o.Tasks
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

func (o ListTasksResponse) String() string {
	out := "ListTasksResponse {\n"
	out += fmt.Sprintf("  tasks=%v\n", o.Tasks)
	out += fmt.Sprintf("  pagination=%v\n", o.Pagination)
	out += "}"
	return out
}

type NullableListTasksResponse struct {
	value *ListTasksResponse
	isSet bool
}

func (v NullableListTasksResponse) Get() *ListTasksResponse {
	return v.value
}

func (v *NullableListTasksResponse) Set(val *ListTasksResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTasksResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTasksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTasksResponse(val *ListTasksResponse) *NullableListTasksResponse {
	return &NullableListTasksResponse{value: val, isSet: true}
}

func (v NullableListTasksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTasksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

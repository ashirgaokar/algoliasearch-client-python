// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// SuccessResponse struct for SuccessResponse
type SuccessResponse struct {
	// The status code.
	Status int32 `json:"status"`
	// Message of the response.
	Message string `json:"message"`
}

// NewSuccessResponse instantiates a new SuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResponse(status int32, message string) *SuccessResponse {
	this := &SuccessResponse{}
	this.Status = status
	this.Message = message
	return this
}

// NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResponseWithDefaults() *SuccessResponse {
	this := &SuccessResponse{}
	return this
}

// GetStatus returns the Status field value
func (o *SuccessResponse) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SuccessResponse) SetStatus(v int32) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *SuccessResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SuccessResponse) SetMessage(v string) {
	o.Message = v
}

func (o SuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

func (o SuccessResponse) String() string {
	out := "SuccessResponse {\n"
	out += fmt.Sprintf("  status=%v\n", o.Status)
	out += fmt.Sprintf("  message=%v\n", o.Message)
	out += "}"
	return out
}

type NullableSuccessResponse struct {
	value *SuccessResponse
	isSet bool
}

func (v NullableSuccessResponse) Get() *SuccessResponse {
	return v.value
}

func (v *NullableSuccessResponse) Set(val *SuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResponse(val *SuccessResponse) *NullableSuccessResponse {
	return &NullableSuccessResponse{value: val, isSet: true}
}

func (v NullableSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

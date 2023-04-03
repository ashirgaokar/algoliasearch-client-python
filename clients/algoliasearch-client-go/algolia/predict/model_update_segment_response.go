// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// UpdateSegmentResponse struct for UpdateSegmentResponse
type UpdateSegmentResponse struct {
	// The ID of the segment.
	SegmentID string `json:"segmentID"`
	// The date and time at which the segment was updated (RFC3339).
	UpdatedAt string `json:"updatedAt"`
}

// NewUpdateSegmentResponse instantiates a new UpdateSegmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSegmentResponse(segmentID string, updatedAt string) *UpdateSegmentResponse {
	this := &UpdateSegmentResponse{}
	this.SegmentID = segmentID
	this.UpdatedAt = updatedAt
	return this
}

// NewUpdateSegmentResponseWithDefaults instantiates a new UpdateSegmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSegmentResponseWithDefaults() *UpdateSegmentResponse {
	this := &UpdateSegmentResponse{}
	return this
}

// GetSegmentID returns the SegmentID field value
func (o *UpdateSegmentResponse) GetSegmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SegmentID
}

// GetSegmentIDOk returns a tuple with the SegmentID field value
// and a boolean to check if the value has been set.
func (o *UpdateSegmentResponse) GetSegmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentID, true
}

// SetSegmentID sets field value
func (o *UpdateSegmentResponse) SetSegmentID(v string) {
	o.SegmentID = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UpdateSegmentResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UpdateSegmentResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UpdateSegmentResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o UpdateSegmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["segmentID"] = o.SegmentID
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSegmentResponse) String() string {
	out := "UpdateSegmentResponse {\n"
	out += fmt.Sprintf("  segmentID=%v\n", o.SegmentID)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	out += "}"
	return out
}

type NullableUpdateSegmentResponse struct {
	value *UpdateSegmentResponse
	isSet bool
}

func (v NullableUpdateSegmentResponse) Get() *UpdateSegmentResponse {
	return v.value
}

func (v *NullableUpdateSegmentResponse) Set(val *UpdateSegmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSegmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSegmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSegmentResponse(val *UpdateSegmentResponse) *NullableUpdateSegmentResponse {
	return &NullableUpdateSegmentResponse{value: val, isSet: true}
}

func (v NullableUpdateSegmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSegmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// SourceCreate struct for SourceCreate
type SourceCreate struct {
	Type  SourceType  `json:"type"`
	Name  string      `json:"name"`
	Input SourceInput `json:"input"`
	// The authentication UUID.
	AuthenticationID *string `json:"authenticationID,omitempty"`
}

type SourceCreateOption func(f *SourceCreate)

func WithSourceCreateAuthenticationID(val string) SourceCreateOption {
	return func(f *SourceCreate) {
		f.AuthenticationID = &val
	}
}

// NewSourceCreate instantiates a new SourceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceCreate(type_ SourceType, name string, input SourceInput, opts ...SourceCreateOption) *SourceCreate {
	this := &SourceCreate{}
	this.Type = type_
	this.Name = name
	this.Input = input
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewSourceCreateWithDefaults instantiates a new SourceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceCreateWithDefaults() *SourceCreate {
	this := &SourceCreate{}
	return this
}

// GetType returns the Type field value
func (o *SourceCreate) GetType() SourceType {
	if o == nil {
		var ret SourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SourceCreate) GetTypeOk() (*SourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SourceCreate) SetType(v SourceType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *SourceCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceCreate) SetName(v string) {
	o.Name = v
}

// GetInput returns the Input field value
func (o *SourceCreate) GetInput() SourceInput {
	if o == nil {
		var ret SourceInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *SourceCreate) GetInputOk() (*SourceInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *SourceCreate) SetInput(v SourceInput) {
	o.Input = v
}

// GetAuthenticationID returns the AuthenticationID field value if set, zero value otherwise.
func (o *SourceCreate) GetAuthenticationID() string {
	if o == nil || o.AuthenticationID == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationID
}

// GetAuthenticationIDOk returns a tuple with the AuthenticationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceCreate) GetAuthenticationIDOk() (*string, bool) {
	if o == nil || o.AuthenticationID == nil {
		return nil, false
	}
	return o.AuthenticationID, true
}

// HasAuthenticationID returns a boolean if a field has been set.
func (o *SourceCreate) HasAuthenticationID() bool {
	if o != nil && o.AuthenticationID != nil {
		return true
	}

	return false
}

// SetAuthenticationID gets a reference to the given string and assigns it to the AuthenticationID field.
func (o *SourceCreate) SetAuthenticationID(v string) {
	o.AuthenticationID = &v
}

func (o SourceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["input"] = o.Input
	}
	if o.AuthenticationID != nil {
		toSerialize["authenticationID"] = o.AuthenticationID
	}
	return json.Marshal(toSerialize)
}

func (o SourceCreate) String() string {
	out := "SourceCreate {\n"
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  input=%v\n", o.Input)
	out += fmt.Sprintf("  authenticationID=%v\n", o.AuthenticationID)
	out += "}"
	return out
}

type NullableSourceCreate struct {
	value *SourceCreate
	isSet bool
}

func (v NullableSourceCreate) Get() *SourceCreate {
	return v.value
}

func (v *NullableSourceCreate) Set(val *SourceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceCreate(val *SourceCreate) *NullableSourceCreate {
	return &NullableSourceCreate{value: val, isSet: true}
}

func (v NullableSourceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
